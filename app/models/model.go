package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// TODO 从环境变量引入
const (
	host           = "localhost:27017"
	dbName         = "watchwhat"
	collectionName = "media"
)

// Media model的struct要写如下格式：(不要带Id才能插入)
// 需要带上tag，指定bson（插入数据库的键）和json（转json的键）
type Media struct {
	// ID     bson.ObjectId `bson:"_id" json:"_id"`
	MediaName         string `bson:"media_name" json:"media_name"`
	MediaSnapshot     string `bson:"media_snapshot" json:"media_snapshot"`
	MediaDescription  string `bson:"media_description" json:"media_description"`
	MediaLink         string `bson:"media_link" json:"media_link"`
	PlayedCount       int    `bson:"played_count" json:"played_count"`
	Likes             int    `bson:"likes" json:"likes"`
	UploadTime        int    `bson:"upload_time" json:"upload_time"`
	Uploader          string `bson:"uploader" json:"uploader"`
	UploaderAvatar    string `bson:"uploader_avatar" json:"uploader_avatar"`
	UploaderSpaceLink string `bson:"uploader_space_link" json:"uploader_space_link"`
	CheckTime         int    `bson:"check_time" json:"check_time"`
	Channel           string `bson:"channel" json:"channel"`
	FromWebsite       string `bson:"from_website" json:"from_website"`
	OriginalID        string `bson:"original_id" json:"original_id"`
}

// FindList 根据pagenum，perpage和频道决定查找
func FindList(o int, p int) []Media {
	results := []Media{}
	c := getConnection()
	e := c.Find(nil).Sort("-likes").Skip(o * p).Limit(p).All(&results)
	if e != nil {
		panic(e)
	}
	return results
}

func FindListWithChannel(o int, p int, channel string) []Media {
	results := []Media{}
	c := getConnection()
	e := c.Find(bson.M{"channel": channel}).Sort("-likes").Skip(o * p).Limit(p).All(&results)
	if e != nil {
		panic(e)
	}
	return results
}

// func main() {
// 	connection := getConnection()

// 	// 查
// 	// 查询多条数据，带上错误处理
// 	allResult := make([]Bookstore, 0)
// 	// allResult := make([]interface{}, 0) // 这种可行，但不好

// 	// .Find(bson.M{})或者.Find(nil)
// 	// 如果排序，插入一个.Sort("+name", "-price")
// 	// 如果分页，使用.Skip(<pageNo>).Limit(<perpage>)方法
// 	err := connection.Find(nil).Skip(1).Limit(3).Sort("+name", "-price").All(&allResult)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(allResult)

// 	// 查询单条数据，以下不带错误处理
// 	oneResult := Bookstore{}
// 	connection.Find(bson.M{"name": "oldman"}).One(&oneResult)
// 	fmt.Println(oneResult)

// 	// 增
// 	// 1. 构建结构体，2. 传入结构体的指针
// 	newBook := Bookstore{Name: "leader One", Pages: 44, Price: 55, Url: "http://www.baidu.com", Author: "cuuuuu"}
// 	connection.Insert(newBook)

// 	// 改
// 	// 用id为凭据查找：
// 	// targetId := bson.ObjectIdHex("586daa762d902327803c7428")
// 	// connection.Update(bson.M{"_id": targetId}, bson.M{"$set": bson.M{"title": "title-111122"}})
// 	// 用其他为凭据查找：
// 	connection.Update(bson.M{"name": "leader One"}, bson.M{"$set": bson.M{"price": 60}})
// 	// 更新多条：Update方法改用UpdateAll

// 	// 删
// 	// 根据id删除：
// 	// connection.RemoveId(bson.ObjectIdHex("586daa762d902327803c7428"))
// 	// 删除单条
// 	connection.Remove(bson.M{"name": "leader One"})
// 	// 删除多条：把Remove方法换成RemoveAll
// 	connection.RemoveAll(bson.M{"name": "leader One"})

// 	bs := bson.M{"name": 111, "ss": "sda"}
// 	fmt.Println(bs, bs["name"])

// }

func getConnection() *mgo.Collection {
	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}

	return session.DB(dbName).C(collectionName)
}
