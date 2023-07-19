package rtm_package

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDataJob(db *mongo.Database, job_title, deskripsi, deadline string, priority string) (InsertedID interface{}) {
	var datajob Job
	datajob.Job_title = job_title
	datajob.Deskripsi = deskripsi
	datajob.Deadline = deadline
	datajob.Priority = priority
	return InsertOneDoc(db, "data_job", datajob)
}

func GetDataJob(priority string, db *mongo.Database, col string) (data Job) {
	user := db.Collection(col)
	filter := bson.M{"priority": priority}
	err := user.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getbypriority: %v\n", err)
	}
	return data
}

func DeleteDataJob(priority string, db *mongo.Database, col string) (data Job) {
	user := db.Collection(col)
	filter := bson.M{"priority": priority}
	err, _ := user.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDataJob : %v\n", err)
	}
	fmt.Println("Succes Delete data")
	return data
}
