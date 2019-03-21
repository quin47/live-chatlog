package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Message struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UID string `json:"uid" bson:"uid" `
	Time time.Time `json:"time" bson:"time"`
	Body    string            `json:"body" bson:"body"`
	Source  string            `json:"source" bson:"source"`
	MsgType string            ` json:"msg_type" bson:"msg_type"`
	RoomId string `json:"room_id" bson:"room_id"`
}
