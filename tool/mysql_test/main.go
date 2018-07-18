package main

import (
	_ "os"
	_ "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

import (
	"fmt"
	_ "gos_server/config"
	"gos_server/center_server/db"
	"gos_server/center_server/db/user"
)

func main() {

	gdb := db.DBService()

	user.Create(gdb)

	user1 := user.New(gdb,"test","dddddddddd")
	if user1 != nil {
		fmt.Println("user1",user1)
	}

	user2 := user.QueryById(gdb,user1.Id)
	if user2 != nil {
		fmt.Println("user2",user2)
	}

	user3 := user.QueryByName(gdb,user1.Name)
	if user2 != nil {
		fmt.Println("user3",user3)
	}

	user4 := user.New(gdb,"test2","dddddddddd")
	if user4 != nil {
		fmt.Println("user4",user4)
	}

	user1.Name = "kkkkkk"
	user.Update(gdb,*user1)

	var users []user.User	
	//users:= []user.User

	users = user.GetAll(gdb)

	for i := 0 ; i<len(users); i++ {
		fmt.Println(users[i])
		//user.Delete(gdb,users[i].Id)
	}
	
	//fmt.Println("test delele result!")
	//user5 := user.QueryByName(gdb,"kkkkkk")
	//if user5 != nil {
	//	fmt.Println("user5",user5)
	//}
}
