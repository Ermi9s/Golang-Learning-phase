package repository

import (
	"context"
	"errors"

	databasedomain "github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/database/databaseDomain"
	"github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User_Repository struct {
	Repository
	Collection databasedomain.Collection
}

func New_User_Repository(repository Repository , collection databasedomain.Collection) domain.User_Repository_interface {
	return &User_Repository{
		Repository: repository,
		Collection: collection,
	}
}

func (repo *User_Repository)GetUserDocumentById(id string) (domain.User , error){
	objID,err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{},err
	}
	var decoded domain.User
	filter := bson.D{{Key : "_id" , Value: objID}}
	doc := repo.Collection.FindOne(context.TODO() , filter)
	
	doc.Decode(&decoded)

	return decoded,nil
}

func (repo *User_Repository)GetUserDocumentByFilter(filter map[string]string)([]domain.User , error){
	dbfilter := bson.D{{}}
	for key,val := range filter {
		dbfilter = append(dbfilter, bson.E{Key: key , Value: val})
	}
	var result []domain.User
	cursor, err := repo.Collection.Find(context.TODO() , filter)
	if err != nil {
		return nil , err
	}

	for cursor.Next(context.TODO()) {
		var object domain.User
		err := cursor.Decode(&object)
		if err != nil {
			return nil , err
		}
		result = append(result, object)
	}
	if len(result) == 0 {
		return nil , errors.New("no user found with the filter")
	}
	return result , nil
}


func (repo *User_Repository)UpdateUserDocumentById(id string , update domain.User)  error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	byteModel, err := bson.Marshal(update)
	if err != nil {
		return err
	}
	var docModel bson.D
	err = bson.Unmarshal(byteModel , &docModel)
	if err != nil {
		return err
	}
	filter := bson.D{{Key : "_id" , Value: objID}}
	updater := bson.D{{Key: "$set" , Value: docModel}}
	
	_,err = repo.Collection.UpdateOne(context.TODO() , filter , updater)
	if err != nil {
		return err
	}
	return nil
}

func (repo *User_Repository)InsertUserDocument(object domain.User) (string, error) {
	var docModel bson.D
	var byteModel []byte
	
	byteModel,err := bson.Marshal(object)
	if err != nil {
		return "" , err
	}
	err = bson.Unmarshal(byteModel , &docModel)
	if err != nil {
		return "", err
	}
	inserted,err := repo.Collection.InsertOne(context.TODO() , docModel)
	
	if err != nil {
		return "", err
	}
	id := inserted.(primitive.ObjectID)
	sid := id.Hex()

	return sid,nil
}

func (repo *User_Repository)DeleteUserDocument(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.D{{Key : "_id" , Value: objID}}

	_,err = repo.Collection.DeleteOne(context.TODO() , filter)
	
	return err
}
