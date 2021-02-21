package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection is the struct for mongo collection
type Collection struct {
	client         *mongo.Client
	databaseName   string
	collectionName string
	collection     *mongo.Collection
}

// SimpleClient generate a simple client
func SimpleClient(uri string) *mongo.Client {
	// mongodb://[username:password@]host1[:port1][,host2[:port2],...[,hostN[:portN]]][/[database][?options]]
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic("failed to connect mongo.")
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		panic("failed for check connect.")
	}

	return client
}

// NewCollection create a collection
func NewCollection(uri string, dbName string, collectionName string) *Collection {
	client := SimpleClient(uri)

	return &Collection{
		client:         client,
		databaseName:   dbName,
		collectionName: collectionName,
		collection:     client.Database(dbName).Collection(collectionName),
	}
}

// InsertOne insert a item. d should be a struct, bson...
func (c *Collection) InsertOne(d interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return c.collection.InsertOne(ctx, d)
}

// todo
// func (c *Collection) InsertMany(ds []interface{}) {
// 	students := []interface{}{s2, s3}
// 	insertManyResult, err := collection.InsertMany(context.TODO(), students)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
// }

// todo
// UpdateOneSet update a item
// d should be a map, struct or bson
// func (c *Collection) UpdateOneSet(filter interface{}, d interface{}) (*mongo.UpdateResult, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	return c.collection.UpdateOne(ctx, filter, bson.D{{"$set", d}})
// }

// FindOneFill query a item from db.
// filter should be a bson.M or bson.D
// s should be a struct, map
func (c *Collection) FindOneFill(filter interface{}, s interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := c.collection.FindOne(ctx, filter).Decode(s)
	return err
}

// todo
// func (c *Collection) Find() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()
// 	cur, err := c.collection.Find(ctx, bson.D{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer cur.Close(ctx)
// 	for cur.Next(ctx) {
// 		var result bson.M
// 		err := cur.Decode(&result)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// do something with result....
// 	}
// 	if err := cur.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// }

// DeleteOne delete a item from db
// filter should be a bson. example filter = bson.D{{"name": "yourname"}}
func (c *Collection) DeleteOne(filter interface{}) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return c.collection.DeleteOne(ctx, filter)
}

// DeleteMany delete items from db
// filter should be a bson. example filter = bson.D{{"name": "yourname"}}
func (c *Collection) DeleteMany(filter interface{}) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return c.collection.DeleteMany(ctx, filter)
}
