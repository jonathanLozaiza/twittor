package db

import (
	"context"
	"time"
	"twitter/src/models"

	"go.mongodb.org/mongo-driver/bson"
)

func CheckUser(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("users")

	condition := bson.M{"email": email}
	var results models.Usuario
	err := col.FindOne(ctx, condition).Decode(&results)
	ID := results.ID.Hex()
	if err != nil {
		return results, false, ID
	}
	return results, true, ID
}
