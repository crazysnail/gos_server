package db

import (
	"os"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

import (
	"gos_server/config"
)

var Gmdb *mgo.Session

const (
	Counters = "gos_counters"
)

type Counter struct {
	Name    string
	NextVal int64
}

func DBService() {
	// dial mongodb
	sess, err := mgo.Dial(config.MongoDBInner)
	if err != nil {
		config.LogCenterServer.Errorln("cannot connect to", config.MongoDBInner, err)
		os.Exit(-1)
	}

	// set default session mode to strong for saving player's data
	sess.SetMode(mgo.Strong, true)
	Gmdb = sess

	config.LogCenterServer.Infoln("DBService connect to %s success!", config.MongoDBInner )
}

//------------------------------------------------ copy connection
// !IMPORTANT!  NEVER FORGET -----> defer ms.Close() <-----
func C(collection string) (*mgo.Session, *mgo.Collection) {
	ms := Gmdb.Copy()
	c := ms.DB(config.MongoDBName).C(collection)
	return ms, c
}

//---------------------------------------------------------- ID GENERATOR
func NextVal(countername string) int32 {
	ms, c := C(Counters)
	defer ms.Close()

	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"nextval": 1}},
		Upsert:    true,
		ReturnNew: true,
	}

	next := &Counter{}
	info, err := c.Find(bson.M{"name": countername}).Apply(change, &next)
	if err != nil {
		config.LogCenterServer.Errorln(info, err)
		return -1
	}

	// round the nextval to 2^31
	return int32(next.NextVal % 2147483648)
}
