package main

import (
	"fmt"

	"github.com/zhhink/common/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type user struct {
	Name string
	Age  int
}

var (
	uri string
)

func init() {
	uri = ""
}

func main() {
	col := mongo.NewCollection(uri, "uitest", "testcol")

	u := user{
		Name: "jdd0d",
		Age:  20,
	}

	rs, _ := col.InsertOne(u)
	fmt.Printf("insert: %v", rs.InsertedID)
	println(rs.InsertedID)

	filter := bson.D{{"_id", rs.InsertedID}}
	var userFind user
	err := col.FindOneFill(filter, &userFind)
	if err != nil {
		fmt.Printf("find one struct err: %+v", err)
		panic("find failed.")
	}

	fmt.Printf("find one struct: %+v\n", userFind)

	userMap := make(map[string]interface{})
	err = col.FindOneFill(filter, &userMap)
	if err != nil {
		fmt.Printf("find one struct err: %+v", err)
		panic("find failed.")
	}

	fmt.Printf("find one struct: %+v\n", userMap)

	uModify := user{
		Name: "jdd0d",
		Age:  40,
	}

	uModifyMap := map[string]interface{}{
		"email": "test@163.com",
	}

	// _, err = col.UpdateOneSet(filter, uModify)
	// if err != nil {
	// 	fmt.Printf("update one err: %v", err)
	// }

	// _, err = col.UpdateOneSet(filter, uModifyMap)

	// _, err = col.DeleteMany(bson.D{})
	// _, err = col.DeleteOne(filter)
}
