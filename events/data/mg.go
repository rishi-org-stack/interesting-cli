package data

import (
	"context"
	"fmt"
	"log"

	"time" 

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Event struct{
	ID primitive.ObjectID `bson:"_id,omitempty"`
	UID int32 `bson:"uid,omitempty"`
	Name string `bson:"name,omitempty"`
	Description string `bson:"description"`
	Date string`bson:"date,omitempty"`
}

//Todo:check validity of given string as dateif yes then return 1 else 0
// func valid(s string) int {
// 	return 1
// }
//Construct base class for all query
func Construct(name string,des string, date string,id int32)Event {
	var e Event
		e.Name= name
		e.Description=des
		e.UID = id
		e.Date=date
	return e
}

func connection(d string, c string) (*mongo.Collection, *mongo.Client, error) {
	var client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	var col = client.Database(d).Collection(c)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err != nil {
		log.Fatal(err)
		return col, client, err
	}
	client.Connect(ctx)
	return col, client, nil
}

func (E*Event) GetAll(d string, s string) []bson.M {
	data := make([]bson.M, 0)
	var c *mongo.Collection
	var cur *mongo.Cursor
	var e error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	c, _, e = connection(d, s)
	if e != nil {
		log.Fatal(e)
	}
	cur, e = c.Find(ctx, bson.M{})
	if e != nil {
		log.Fatal(e)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var epi bson.M
		if e = cur.Decode(&epi); e != nil {
			log.Fatal(e)
		}
		data = append(data, epi)
	}
	return data
}

func (S *Event) Getone(d string, c string) bson.M {
	var data = S.GetAll(d, c)
	var s bson.M
	for _, studs := range data {

		if studs["name"].(string) == S.Name {
			s = studs
		} else {
			// fmt.Println("you are not  present in ur databse")
		}
	}
	return s
}

func (S *Event)Insert(d string, c string){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col,_,err:=connection(d,c)
	if err!=nil{
		log.Fatal(err)
	}
	i,e:=col.InsertOne(ctx,bson.M{"id":S.UID,"name":S.Name,"des":S.Description,"date":S.Date})
	if e!=nil{
		log.Fatal(e)
	}
	fmt.Println(i)
}

func (S*Event)Update(d,c string){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var col *mongo.Collection
	var err error
	col,_,err=connection(d,c)
	if err!=nil{
		log.Fatal(err)
	}
	result,err:=col.UpdateOne(ctx,bson.M{"id":S.UID},bson.M{"$set":bson.M{"name":S.Name,"date":S.Date}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func (S *Event)Delete(d,c string){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col,_,err:=connection(d,c)
	if err != nil {
		log.Fatal(err)
	}
	result,err:=col.DeleteOne(ctx,bson.M{"id":S.UID})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}