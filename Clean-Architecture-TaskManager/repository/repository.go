package repository

import (
	"context"
	"github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (repo *Repository)GetDocumentById(collection_name string ,  id string) (domain.Model , error){
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil , err
	}

	var decoded domain.Model // an interface both user and task are implementing 

	filter := bson.D{{Key : "_id" , Value: objID}}
	collection := repo.Database.Collection(collection_name)
	doc := collection.FindOne(context.TODO() , filter)
	doc.Decode(&decoded)

	return decoded,doc.Err()
}

func (repo *Repository)GetDocumentByFilter(collection_name string , filter map[string]string)([]domain.Model , error){
	dbfilter := bson.D{{}}
	for key,val := range filter {
		dbfilter = append(dbfilter, bson.E{Key: key , Value: val})
	}
	var result []domain.Model
	collection := repo.Database.Collection(collection_name)
	cursor, err := collection.Find(context.TODO() , filter)
	if err != nil {
		return nil , err
	}

	for cursor.Next(context.TODO()) {
		var object domain.Model
		err := cursor.Decode(&object)
		if err != nil {
			return nil , err
		}
		result = append(result, object)
	}
	return result , nil
}


func (repo *Repository)UpdateDocumentById(collection_name string ,  id string , update domain.Model)  error {
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

	collection := repo.Database.Collection(collection_name)
	
	_,err = collection.UpdateOne(context.TODO() , filter , updater)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository)InsertDocument(collection_name string , object domain.Model) error {
	collection := repo.Database.Collection(collection_name)
	var docModel bson.D

	byteModel,err := bson.Marshal(object)
	if err != nil {
		return err
	}
	err = bson.Unmarshal(byteModel , &docModel)
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(context.TODO() , docModel)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository)DeleteDocument(collection_name string , id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.D{{Key : "_id" , Value: objID}}
	collection := repo.Database.Collection(collection_name)
	_,err = collection.DeleteOne(context.TODO() , filter)
	
	return err
}
