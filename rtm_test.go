package rtm_package

import (
	"fmt"
	"os"
	"testing"

	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "HelpdeskData",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func TestInsertData(t *testing.T) {
	Id := "112233"
	Job_title := "Manager"
	Deadline := "01/10/2023"
	Priority := "Medium"
	hasil := InsertDataUser(MongoConn, Id, Job_title, Deadline, Priority)
	fmt.Println(hasil)

}
func TestInsertDataJob(t *testing.T) {
	dbname := "Job"
	user := UserSurat{
		ID:      primitive.NewObjectID(),
		Nama:    "Haris Riyoni",
		Email:   "harisriyoni@gmail.com",
		Telepon: "081234567890",
	}
	insertedID := InsertUser(dbname, user)
	if insertedID == nil {
		t.Error("Failed to insert user")
	}
}

func TestGetDatauser(t *testing.T) {
	id := "112233"
	hasil := GetDataJob(id, MongoConn, "data_user")
	fmt.Println(hasil)

}

func TestGetDatabyphone(t *testing.T) {
	Handphone := "0000000"
	hasil := GetDataUserFromPhone(Handphone, MongoConn, "data_user")
	fmt.Println(hasil)

}

func TestDeleteData(t *testing.T) {
	Handphone := "0000000"
	hasil := DeleteData(Handphone, MongoConn, "data_user")
	fmt.Println(hasil)

}
