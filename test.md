```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

func readWomenArticles(ctx context.Context, client *mongo.Client) ([]map[string]interface{}, error) {
	collection := client.Database("clothes_shop").Collection("articles")

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"published_date", -1}})

	cursor, err := collection.Find(ctx, bson.D{{"category", "women"}}, findOptions)
	if err != nil {
		return nil, err
	}

	var articles []map[string]interface{}
	for cursor.Next(context.TODO()) {
		var article map[string]interface{}
		err = cursor.Decode(&article)
		if err != nil {
			log.Println(err)
			continue
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = client.Disconnect(ctx)
	}()

	articles, err := readWomenArticles(ctx, client)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", articles)
}
```