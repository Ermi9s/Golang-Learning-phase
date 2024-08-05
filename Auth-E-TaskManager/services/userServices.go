package services

import (
	"context"
	models "github.com/Ermi9s/Golang-Learning-phase/Auth-E-TaskManager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (DBM *DataBaseManager)GetUser(id int) (models.User , error) {
	filter := bson.D{{Key : "_id" , Value: id}}
	var user models.User

	err := DBM.Users.FindOne(context.TODO() , filter).Decode(&user)
	if err != nil {
		return models.User{} ,err
	}

	return user , nil
}

func (DBM *DataBaseManager) GetUsers(ctx context.Context) ([]models.User, error) {
	cursor, err := DBM.Users.Find(context.TODO(), bson.M{})
  
	if err != nil{
	  return nil, err
	}
	var users []models.User
	for cursor.Next(context.TODO()){
	  user := models.User{}
	  err := cursor.Decode(&user)
  
	  if err != nil{
		return nil, err
	  }
  
	  users = append(users , user)
	}
	return users , nil
  }
  

func (DBM *DataBaseManager)CreateUser(user models.User) (models.User , error) {
	var doc bson.D
	marshaled ,err := bson.Marshal(user)
	if err != nil {
		return models.User{},err
	}
	err = bson.Unmarshal(marshaled , &doc)
	if err != nil {
		return models.User{},err
	}

	_,err = DBM.Users.InsertOne(context.TODO() , doc)
	if err != nil {
		return models.User{},err
	}

	return user , nil
}

  
func (DBM *DataBaseManager)UpdateUser(id primitive.ObjectID , user models.User)(models.User , error) {
	var doc bson.D


	marshaled , err := bson.Marshal(user)
	if err != nil {
		return models.User{},err
	}

	err = bson.Unmarshal(marshaled , &doc)
	if err != nil {
		return models.User{},err
	}

	filter := bson.D{{Key:"_id" , Value : id}}
	update := bson.D{{Key: "$set" , Value: doc}}

	_, err = DBM.Users.UpdateOne(context.TODO() , filter , update)
	if err != nil {
		return models.User{},err
	}
	return user , nil 
}

func (DBM *DataBaseManager)DeleteUser(id primitive.ObjectID) error {
	filter := bson.D{{Key:"_id" , Value: id}}
	_,err := DBM.Users.DeleteOne(context.TODO(),  filter)
	return err
}