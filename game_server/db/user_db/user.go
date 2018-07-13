package user_db

import (
	"gopkg.in/mgo.v2/bson"
	"crypto/md5"
	"io"
	"time"
)

import (
	"gos_server/config"
	"gos_server/game_server/db"
)


type User struct {
	Id             int32  // 用户id
	Domain         string // 玩家所在分服
	Name           string // 用户名
	Flag           int32  // 状态标记
	Pass           []byte // 密码(MD5 Hash)
	Score          int32  // 分数
	ProtectTimeout int64  // 护盾截止时间
	Mac            string // 玩家MAC地址
	CountryCode    string // 国家代码
	Language       string // 界面语言
	DeviceType     string // 设备类型
	LastSaveTime   int64  // 服务器最后一次刷入数据库的时间
	CreatedAt      int64  // 注册时间
}


const (
	MDCollection   = "gos_users"
	MDCounterName  = "userid_gen"
)



//---------------------------------------------------------- update a user
func Set(user *User) bool {
	ms, c := db.C(MDCollection)
	defer ms.Close()

	info, err := c.Upsert(bson.M{"id": user.Id}, user)
	if err != nil {
		config.LogCenterServer.Errorln(info, err)
		return false
	}

	return true
}

//---------------------------------------------------------- login with name & mac address
func LoginMac(name, mac string) *User {
	ms, c := db.C(MDCollection)
	defer ms.Close()

	user := &User{}
	err := c.Find(bson.M{"name": name, "mac": mac}).One(user)
	if err != nil {
		config.LogCenterServer.Errorln(err, mac)
		return nil
	}

	return user
}

//---------------------------------------------------------- create a new user
func New(name, mac string) *User {
	ms, c := db.C(MDCollection)
	defer ms.Close()

	user := &User{}
	err := c.Find(bson.M{"name": name}).One(user)
	if err != nil {
		user.Id = db.NextVal(MDCounterName)
		user.Name = name
		user.Mac = mac
		user.Domain = config.DomainName
		user.CreatedAt = time.Now().Unix()
		err := c.Insert(user)
		if err != nil {
			config.LogCenterServer.Errorln(err, name, mac)
			return nil
		}
		return user
	}

	return nil
}

//---------------------------------------------------------- query a user by name
func Query(name string) *User {
	ms, c := db.C(MDCollection)
	defer ms.Close()

	user := &User{}
	err := c.Find(bson.M{"name": name}).One(user)
	if err != nil {
		config.LogCenterServer.Errorln(err, name)
		return nil
	}

	return user
}

//---------------------------------------------------------- load a user
func Get(id int32) *User {
	ms, c := db.C(MDCollection)
	defer ms.Close()

	user := &User{}
	err := c.Find(bson.M{"id": id}).One(user)
	if err != nil {
		config.LogCenterServer.Errorln(err, id)
		return nil
	}

	return user
}

//---------------------------------------------------------- load all userss
func GetAll() []User {
	ms, c := db.C(MDCollection)
	defer ms.Close()

	var users []User
	err := c.Find(nil).All(&users)
	if err != nil {
		config.LogCenterServer.Errorln(err)
		return nil
	}

	return users
}

func _md5(str string) []byte {
	configs := config.Get()
	salted := str + configs["salt"]
	h := md5.New()
	io.WriteString(h, salted)
	return h.Sum(nil)
}
