package repository

import (
	"context"
	"github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Task_Repository struct {
	Repository
}

func New_Task_Repository(repository Repository) domain.Task_Repository_interface {
	return &Task_Repository{
		Repository: repository,
	}
}

func (repo *Task_Repository)GetTaskDocumentById(id string) (domain.Task , error){
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Task{},err
	}
	var decoded domain.Task
	filter := bson.D{{Key : "_id" , Value: objID}}
	collection := repo.Database.Collection("Task")
	doc := collection.FindOne(context.TODO() , filter)
	
	doc.Decode(&decoded)

	return decoded,doc.Err()
}

func (repo *Task_Repository)GetTaskDocumentByFilter(filter map[string]string)([]domain.Task , error){
	dbfilter := bson.D{}
	for key,val := range filter {
		if key == "creator_id" {
			c_id,_ := primitive.ObjectIDFromHex(val)
			dbfilter = append(dbfilter, bson.E{Key: key , Value: c_id})
			continue
		}
		dbfilter = append(dbfilter, bson.E{Key: key , Value: val})
	}
	

	var result []domain.Task
	collection := repo.Database.Collection("Task")
	cursor, err := collection.Find(context.TODO() , dbfilter)
	if err != nil {
		return nil , err
	}

	for cursor.Next(context.TODO()) {
		var object domain.Task
		err := cursor.Decode(&object)
		if err != nil {
			return nil , err
		}
		result = append(result, object)
	}
	return result , nil
}


func (repo *Task_Repository)UpdateTaskDocumentById(id string , update domain.Task)  error {
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

	collection := repo.Database.Collection("Task")
	
	_,err = collection.UpdateOne(context.TODO() , filter , updater)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Task_Repository)InsertTaskDocument(object domain.Task) (string, error) {
	collection := repo.Database.Collection("Task")
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
	inserted,err := collection.InsertOne(context.TODO() , docModel)
	
	if err != nil {
		return "", err
	}
	id := inserted.InsertedID.(primitive.ObjectID)
	sid := id.Hex()

	return sid,nil
}

func (repo *Task_Repository)DeleteTaskDocument(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.D{{Key : "_id" , Value: objID}}
	collection := repo.Database.Collection("Task")
	_,err = collection.DeleteOne(context.TODO() , filter)
	
	return err
}
