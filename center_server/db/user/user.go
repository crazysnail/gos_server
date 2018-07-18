package user

import (
	_ "crypto/md5"
	_ "io"
	"time"
	"github.com/jinzhu/gorm"
)

import (
	_ "gos_server/config"
	_ "gos_server/center_server/db"
)

type User struct {
	Id	           int32 `gorm:"primary_key"` // 用户id
	ServerId       int32  // serverid
	Name           string `gorm:"type:varchar(32);not null;"`// 用户名
	Pass           string `gorm:"type:varchar(64);not null;"`// 密码
	Token          int32  // 代币数
    Ip        	   string `gorm:"type:varchar(20);not null;index:ip_idx"`//ip
	Mac            string `gorm:"type:varchar(64);not null;"`// 玩家MAC地址
	CountryCode    string `gorm:"type:varchar(16);not null;"`// 国家代码
	Language       string `gorm:"type:varchar(32);not null;"`// 界面语言
	DeviceType     string `gorm:"type:varchar(32);not null;"`// 设备类型
	LastSaveAt	   time.Time  // 服务器最后一次刷入数据库的时间
	CreatedAt      time.Time  // 注册时间
	IsNew		   int16	  // 是否是新号
}


func Create(gdb *gorm.DB){
	if gdb == nil {
		return
	}
	//if !gdb.HasTable(&User{}) {
	//	err := gdb.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&User{}).Error;
	//	if  err != nil {
	//		panic(err)
	//	}
	//}else{
	gdb.AutoMigrate(&User{})
	//}
}

func New(gdb *gorm.DB, name string, mac string) *User {
	if gdb == nil {
		return nil
	}
	pUser := QueryByName(gdb,name)
	if pUser == nil {
		id := GenId(gdb)
		if id != -1 && id <1999999999 {
			user := User{Id:id, Name: name, Mac: mac, CreatedAt: time.Now(), IsNew: 1}
			//db.GDB.NewRecord(user) 	
			gdb.Create(&user)
			return &user
		}
	}
	return nil
}

func GenId(gdb *gorm.DB ) int32 {
	if gdb == nil {
		return -1
	}
	var count int32 = 0	
	err := gdb.Model(&User{}).Count(&count).Error
	if err != nil {
		panic(err)
		//return -1
	}
	return count+1000000000
}

func QueryById(gdb *gorm.DB,  id int32 ) *User {
	if gdb == nil {
		return nil
	}
	var user = User{}
	err := gdb.Model(&User{}).Where(&User{Id: id }).First(&user).Error
	if err != nil {
		return nil
	}
	return &user
}

func QueryByName(gdb *gorm.DB,  name string ) *User {
	if gdb == nil {
		return nil
	}
	var user = User{}
	err := gdb.Model(&User{}).Where(&User{ Name: name }).First(&user).Error
	if err != nil {
		return nil
	}
	return &user
}

func Update(gdb *gorm.DB,  user User ){
	if gdb == nil {
		return
	}
	err := gdb.Model(&User{}).Update(&user).Error
	if err != nil {
		panic(err)
	}
}

func Delete(gdb *gorm.DB,  id int32 ){
	if gdb == nil {
		return
	}
	var user = User{}
	err := gdb.Model(&User{}).Where(&User{Id: id}).Delete(&user).Error
	if err != nil {
		panic(err)
	}
}

func GetAll(gdb *gorm.DB ) []User {
	if gdb == nil {
		return nil
	}
	var users []User	
	err := gdb.Model(&User{}).Find(&users).Error
	if err != nil {
		panic(err)
	}
	return users
}

//
//func _md5(str string) []byte {
//	configs := config.Get()
//	salted := str + configs["salt"]
//	h := md5.New()
//	io.WriteString(h, salted)
//	return h.Sum(nil)
//}
