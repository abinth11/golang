package movieHelpers

import (
	"fmt"
	"log"

	dbConfig "github.com/abinth11/gowithmongodb/config"
	movies "github.com/abinth11/gowithmongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var movieCollection *mongo.Collection = dbConfig.Database.Collection("movies")

func InsertOne(movie movies.Movie) {
	response, err := movieCollection.InsertOne(dbConfig.Ctx, movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}

func UpdateOne(movieId string, updateData movies.Movie) (bool, error) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	fmt.Println(id)
	filter := bson.M{"_id": movieId}
	update := bson.M{
		"$set": updateData,
	}

	result, err := movieCollection.UpdateOne(dbConfig.Ctx, filter, update)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	if result.MatchedCount == 0 {
		return false, nil
	}

	fmt.Println(result)

	return true, nil
}

func DeleteOne(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	response, err := movieCollection.DeleteOne(dbConfig.Ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}

func DeleteAll() {
	response, err := movieCollection.DeleteMany(dbConfig.Ctx, bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}

func FindOne(movieId string) (movies.Movie, error) {
	var movie movies.Movie
	filter := bson.M{"_id": movieId}

	err := movieCollection.FindOne(dbConfig.Ctx, filter).Decode(&movie)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return movies.Movie{}, nil
		}
		return movies.Movie{}, err
	}

	fmt.Println(movie)
	return movie, nil
}

func FindAll() []primitive.M {
	cursor, err := movieCollection.Find(dbConfig.Ctx, bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(dbConfig.Ctx) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(dbConfig.Ctx)
	return movies
}
