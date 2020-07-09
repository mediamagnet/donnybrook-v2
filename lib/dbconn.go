package lib

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Players should have comment
type Players struct {
	Name      string    `bson:"Name"`
	GuildID   string    `bson:"GuildID"`
	PlayerID  string    `bson:"PlayerID"`
	ChannelID string    `bson:"ChannelID"`
	RaceID    string    `bson:"RaceID"`
	Done      bool      `bson:"Done"`
	Ready     bool      `bson:"Ready"`
	JoinTime  time.Time `bson:"Join Time,omitempty"`
	DoneTime  time.Time `bson:"Done Time,omitempty"`
	TotalTime time.Time `bson:"Total Time,omitempty"`
}

// Races should have comment
type Races struct {
	RaceID         string    `bson:"RaceID"`
	GuildID        string    `bson:"GuildID"`
	ChannelID      string    `bson:"ChannelID"`
	Game           string    `bson:"Game"`
	Category       string    `bson:"Category"`
	StartTime      time.Time `bson:"Start Time"`
	Started        bool      `bson:"Started"`
	PlayersEntered int       `bson:"Players Entered"`
	PlayersReady   int       `bson:"Players Ready"`
	PlayersDone    int       `bson:"Players Done"`
}

// Settings should have comment
type Settings struct {
	GuildID string `bson:"GuildID"`
	Volume  int    `bson:"Volume"`
}

var err error

// GetClient gets client
func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// mongoDB connection stuff

// MonPlayer inserts a player
func MonPlayer(dbase string, collect string, players Players) {
	// Connecting to mongoDB
	client := GetClient()
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	}
	collection := client.Database(dbase).Collection(collect)

	insertResult, err := collection.InsertOne(context.TODO(), players)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted:", insertResult.InsertedID)
}

// MonRace inserts a race
func MonRace(dbase string, collect string, races Races) {
	// Connecting to mongoDB
	client := GetClient()
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	}
	collection := client.Database(dbase).Collection(collect)
	insertResult, err := collection.InsertOne(context.TODO(), races)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted:", insertResult.InsertedID)
}

// MonSettings inserts guild settings
func MonSettings(dbase string, collect string, settings Settings) {
	// Connecting to mongoDB
	client := GetClient()
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	}
	collection := client.Database(dbase).Collection(collect)
	insertResult, err := collection.InsertOne(context.TODO(), settings)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted:", insertResult.InsertedID)
}

// MonReturnAllPlayers returns all players
func MonReturnAllPlayers(client *mongo.Client, filter bson.M) []*Players {

	var players []*Players
	collection := client.Database("donnybrook").Collection("players")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Could not find document ", err)
	}
	for cur.Next(context.TODO()) {
		var player Players
		err = cur.Decode(&player)
		if err != nil {
			log.Fatal("Decode Error ", err)
		}
		players = append(players, &player)
	}
	return players
}

// MonReturnAllRaces returns all races
func MonReturnAllRaces(client *mongo.Client, filter bson.M) []*Races {

	var races []*Races
	collection := client.Database("donnybrook").Collection("races")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Could not find document ", err)
	}
	for cur.Next(context.TODO()) {
		var race Races
		err = cur.Decode(&race)
		if err != nil {
			log.Fatal("Decode Error ", err)
		}
		races = append(races, &race)
	}
	return races
}

// MonReturnAllSettings returns all settings
func MonReturnAllSettings(client *mongo.Client, filter bson.M) []*Settings {

	var settings []*Settings
	collection := client.Database("donnybrook").Collection("settings")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Could not find document ", err)
	}
	for cur.Next(context.TODO()) {
		var setting Settings
		err = cur.Decode(&setting)
		if err != nil {
			log.Fatal("Decode Error ", err)
		}
		settings = append(settings, &setting)
	}
	return settings
}

// MonUpdatePlayer updates the object
func MonUpdatePlayer(client *mongo.Client, updatedData bson.M, filter bson.M) int64 {
	collection := client.Database("donnybrook").Collection("players")
	update := bson.D{{Key: "$set", Value: updatedData}}
	updatedResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal("Error updating player", err)
	}
	return updatedResult.ModifiedCount
}

// MonUpdateRace updates the object
func MonUpdateRace(client *mongo.Client, updatedData bson.M, filter bson.M) int64 {
	collection := client.Database("donnybrook").Collection("races")
	update := bson.D{{Key: "$set", Value: updatedData}}
	updatedResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal("Error updating race,", err)
	}
	return updatedResult.ModifiedCount
}

// MonUpdateSettings updates the object
func MonUpdateSettings(client *mongo.Client, updatedData bson.M, filter bson.M) int64 {
	collection := client.Database("donnybrook").Collection("settings")
	update := bson.D{{Key: "$set", Value: updatedData}}
	updatedResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal("Error updating player", err)
	}
	return updatedResult.ModifiedCount
}

// MonDeletePlayer removes a player from collection
func MonDeletePlayer(client *mongo.Client, filter bson.M) int64 {
	collection := client.Database("donnybrook").Collection("players")
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error deleting player", err)
	}
	return deleteResult.DeletedCount
}

// MonDeleteRace removes a race from collection
func MonDeleteRace(client *mongo.Client, filter bson.M) int64 {
	collection := client.Database("donnybrook").Collection("race")
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error deleting player", err)
	}
	return deleteResult.DeletedCount
}
