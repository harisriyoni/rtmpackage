package rtm_package

import "go.mongodb.org/mongo-driver/bson/primitive"

type Job struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Job_title string             `bson:"job_title,omitempty" json:"job_title,omitempty"`
	Deskripsi string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Deadline  string             `bson:"deadline,omitempty" json:"deadline,omitempty"`
	Priority  string             `bson:"priority,omitempty" json:"priority,omitempty"`
}
