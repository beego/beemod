/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/7/30 11:23
 */
package main

import (
	"github.com/beego/beemod"
	"github.com/beego/beemod/pkg/database/mongo"
	"github.com/globalsign/mgo/bson"
)

const (
	db         = "Users"
	collection = "UserModel"
)

const config = `
[beego.mongo.dev]
    URL   = "127.0.0.1:27017"
	debug = true
    source = "admin"
    user   = "user"
    password   = "123456"
`

type User struct {
	Id   bson.ObjectId `bson:"_id" json:"id"`
	Name string        `bson:"name" json:"name"`
}

func main() {
	var (
		err           error
		client        *mongo.Client
		user, newUser User
		users         []User
		id            string
	)

	err = beemod.Register(
		mongo.DefaultBuild,
	).SetCfg([]byte(config), "toml").Run()

	if err != nil {
		panic("register err:" + err.Error())
	}
	client = mongo.Invoker("dev")

	//Insert
	user.Name = "user123"
	user.Id = bson.NewObjectId()
	err = client.Insert(db, collection, user)
	if err != nil {
		panic("Insert err:" + err.Error())
	}

	//FindAll
	err = client.FindAll(db, collection, nil, nil, &users)
	if err != nil {
		panic("FindAll err:" + err.Error())
	}

	//FindOne
	id = "1"
	err = client.FindOne(db, collection, bson.M{"_id": bson.ObjectIdHex(id)}, nil, &user)
	if err != nil {
		panic("FindOne err:" + err.Error())
	}

	//Update
	newUser.Name = "Update Name"
	err = client.Update(db, collection, bson.M{"_id": user.Id}, newUser)
	if err != nil {
		panic("Update err:" + err.Error())
	}

	//Remove
	id = "2"
	err = client.Remove(db, collection, bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		panic("Remove err:" + err.Error())
	}
}
