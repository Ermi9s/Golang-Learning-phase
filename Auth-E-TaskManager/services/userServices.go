package services

import (
	"context"

	models "github.com/Ermi9s/Golang-Learning-phase/Auth-E-TaskManager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func (DBM *DataBaseManager)GetUser(id primitive.ObjectID) (models.User , error) {
	filter := bson.D{{Key:"_id" , Value: id}}
	var user models.User
	
	err := DBM.Users.FindOne(context.TODO(),filter).Decode(&user)
	if err != nil {
		return models.User{},err
	}

	return user,nil
}

func (DBM *DataBaseManager) GetUsers() ([]models.User, error) {
	cursor, err := DBM.Users.Find(context.TODO(), bson.D{})
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
	hasshedPasskey,err := bcrypt.GenerateFromPassword([]byte(user.Password) , bcrypt.DefaultCost); 
	if err != nil {
		return models.User{},err
	}
	user.Password = string(hasshedPasskey)
	user.ID = primitive.NewObjectID()
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

func (DBM *DataBaseManager)MakeAdmin(id primitive.ObjectID)(models.User , error) {
	var user models.User
	filter := bson.D{{Key: "_id"  , Value: id}}
	err := DBM.Users.FindOne(context.TODO() , filter).Decode(&user)
	if err != nil {
		return models.User{},err
	}
	user.Is_admin = true
	promoted_user , err := DBM.UpdateUser(user.ID , user)
	if err != nil {
		return models.User{} , err
	}
	return promoted_user,nil
}

func (DBM *DataBaseManager)LogIn(user models.LogIN)(models.LoggedInUser , error) {
	var iuser models.User
	filter := bson.D{{Key : "username" , Value: user.UserName} , {Key : "email" , Value: user.Email}}
	err := DBM.Users.FindOne(context.TODO() , filter).Decode(&iuser)
	if err != nil {
		return models.LoggedInUser{} , err
	}
	err = bcrypt.CompareHashAndPassword([]byte(iuser.Password) , []byte(user.Password))
	if err != nil {
		return models.LoggedInUser{} , err
	}

	var logged_user models.LoggedInUser
	err = DBM.Users.FindOne(context.TODO() , filter).Decode(&logged_user)
	if err != nil {
		return models.LoggedInUser{} , err
	}

	return logged_user,nil
}