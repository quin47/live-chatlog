package record

import (
	"context"
	"github.com/bigemon/gobilibili"
	"github.com/quin47/live-chatlog/conf"
	"github.com/quin47/live-chatlog/model"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"strconv"
	"time"
)


var chat = conf.Mongoclient.Database("chat")
var collection = chat.Collection("messages")


func Record(roomId int) {
	bili := gobilibili.NewBiliBiliClient()
	bili.RegHandleFunc(gobilibili.CmdDanmuMsg, func(c *gobilibili.Context) bool {
		dinfo := c.GetDanmuInfo()
		//log.Printf("[%d]%d 说: %s\r\n", c.RoomID, dinfo.UID, dinfo.Text)
		message := model.Message{Source: "bilibili", UID:strconv.Itoa(dinfo.UID),Time:time.Now(),Body: dinfo.Text, RoomId: strconv.Itoa(c.RoomID), MsgType: "msg"}
		b, _ := bson.Marshal(message)
		_, e := collection.InsertOne(context.Background(), b)
		if e != nil {
			log.Println(" 保存数据失败",e)
		}
		return true
	})
	bili.ConnectServer(roomId)


}
