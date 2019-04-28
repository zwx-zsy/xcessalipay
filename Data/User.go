package Data

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const CollectionName_User string = "wct_user"

type User struct {
	Id             bson.ObjectId `bson:"_id,omitempty"`
	WxOpenId       string        `bson:"WxOpenId"`
	UserCode       string        `bson:"UserCode"`
	PersonId       string        `bson:"PersonId"`
	CreateDateTime time.Time     `bson:"CreateDateTime"`
}

func Users() *mgo.Collection {
	return DB.C(CollectionName_User)
}
