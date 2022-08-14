package dao

import (
	"context"
	"fmt"
	"github.com/miscmo/iknow/proto"
	"github.com/miscmo/iknow/resource"
	"github.com/miscmo/iknow/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var TableMusic *gorm.DB
var TableUserInfo *gorm.DB

var nodeMgoCli *mongo.Client

func InitDB() {
	conf := resource.GetConfig()
	dsn := strings.Replace(conf.MySQL.Dsn, "{addr}", conf.MySQL.Address, 1)
	log.Printf("dsn = %s", dsn)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		errMsg := fmt.Sprintf("连接数据库 %s 失败：%s", dsn, err.Error())
		log.Fatal(errMsg)
	}

	db.DB().SetMaxOpenConns(conf.MySQL.MaxOpenConnCnt)
	db.DB().SetMaxIdleConns(conf.MySQL.MaxIdleConnCnt)

	TableMusic = db.Table(conf.MySQL.TableMusic)
	TableUserInfo = db.Table(conf.MySQL.TableUserInfo)

	log.Printf("初始化数据库成功 %v %v", TableMusic, TableUserInfo)
}

func GetMongoDB() *mongo.Client {
	var err error
	if nodeMgoCli == nil {
		conf := resource.GetConfig()
		clientOptions := options.Client().ApplyURI(conf.MongoDB.URI)

		// 连接到MongoDB
		nodeMgoCli, err = mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal(err)
			panic(err.Error())
		}
		// 检查连接
		err = nodeMgoCli.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal(err)
			panic(err.Error())
		}
	}
	return nodeMgoCli
}

func GetNodeMongoCli() *mongo.Collection {
	mongoCli := GetMongoDB()

	if mongoCli == nil {
		panic("mongoCli is nil")
	}

	db := mongoCli.Database("iknow")

	//选择表 my_collection
	collection := db.Collection("iknow_node")

	return collection
}

//type NoteDiff struct {
//	Version int `json:"version"`
//	Diff string `json:"diff"`
//	Desc string `json:"desc"`
//	CreateTime string `json:"create_time"`
//	UpdateTime string `json:"update_time"`
//}
//
//type Note struct {
//	Id int `json:"id"`
//	Name string `json:"name"`
//	Desc string `json:"desc"`
//	CreateTime string `json:"create_time"`
//	UpdateTime string `json:"update_time"`
//	Content string `json:"content"`
//	Type string	`bson:"type"`	// "text", "audio", "markdown", "video", "pdf"
//	Diff []NoteDiff `json:"diff"`
//}
//
//type Node struct {
//	Id int `json:"id"`
//	Name string 	`json:"name"`
//	Type int	`json:"type"`
//	Notes []int	`json:"notes"`
//	SubNodes []int `json:"sub_nodes"`
//	IsCollapse bool `json:"is_collapse"`
//	Score      int `json:"score"`
//	Permission int `json:"permission"`
//	CreateTime string `json:"create_time"`
//	UpdateTime string `json:"update_time"`
//}

// 插入记录
func InsertNode(node *proto.Node) error {
	nodeCollection := GetNodeMongoCli()

	iResult, err := nodeCollection.InsertOne(context.TODO(), node)
	if err != nil {
		fmt.Print(err)
		return fmt.Errorf("insert mongodb failed, err: %s", err.Error())
	}
	//_id:默认生成一个全局唯一ID
	id := iResult.InsertedID.(primitive.ObjectID)
	fmt.Println("自增ID", id.Hex())

	return nil
}

func InsertManyNode(nodes []*proto.Node) error {
	nodeCollection := GetNodeMongoCli()

	var documents []interface{}
	for _, v := range nodes {
		documents = append(documents, v)
	}

	iResult, err := nodeCollection.InsertMany(context.TODO(), documents)
	if err != nil {
		fmt.Print(err)
		return fmt.Errorf("insert mongodb failed, err: %s", err.Error())
	}
	//_id:默认生成一个全局唯一ID
	//id := iResult.InsertedID.(primitive.ObjectID)
	//fmt.Println("自增ID", id.Hex())
	for _, ret := range iResult.InsertedIDs {
		fmt.Println("insert objectid: ", ret.(primitive.ObjectID).Hex())
	}

	return nil
}

// 修改记录
func UpdateNode(id string, update bson.M) error {
	var (
		nodeCollection = GetNodeMongoCli()
		uResult *mongo.UpdateResult
		err error
	)
	filter := bson.M{"id": id}

	if uResult, err = nodeCollection.UpdateOne(context.TODO(), filter, update); err != nil {
		fmt.Print(err)
		return fmt.Errorf("update node failed, err: %s", err.Error())
	}

	fmt.Print("update count: ", uResult.ModifiedCount)


	return nil
}

// 查找记录
func SearchNode(id int64, page, pageSize int32) ([]*proto.Node, error) {
	var (
		nodeCollection = GetNodeMongoCli()
		err error
		cursor *mongo.Cursor
		result []*proto.Node
	)

	limit, offset := utils.GetLimitOffset(page, pageSize)

	filter := bson.M{}

	if id != 0 {
		filter["id"] = id
	}

	if cursor, err = nodeCollection.Find(context.TODO(), filter,
		options.Find().SetSkip(int64(offset)),
		options.Find().SetLimit(int64(limit)),
	); err != nil {
		fmt.Print(err)
		return result, fmt.Errorf("search node failed, err: %s", err.Error())
	}

	defer func() {
		if err := cursor.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	if err = cursor.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
	}


	return result, nil
}
