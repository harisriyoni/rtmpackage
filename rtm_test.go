package rtm_package

import (
	"fmt"
	"os"
	"testing"

	"github.com/aiteung/atdb"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "Job",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func TestInsertData(t *testing.T) {
	Id := "112233"
	Job_title := "Manager"
	Deadline := "01/10/2023"
	Priority := "Medium"
	hasil := InsertDataJob(MongoConn, Id, Job_title, Deadline, Priority)
	fmt.Println(hasil)

}

func TestGetDatauser(t *testing.T) {
	id := "112233"
	hasil := GetDataJob(id, MongoConn, "data_user")
	fmt.Println(hasil)

}

func TestDeleteData(t *testing.T) {
	Priority := "Medium"
	hasil := DeleteDataJob(Priority, MongoConn, "data_user")
	fmt.Println(hasil)

}
