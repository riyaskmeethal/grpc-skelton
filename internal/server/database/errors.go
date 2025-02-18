package database

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
)

func CheckErrorNoDocument(err error) bool {
	return err == mongo.ErrNoDocuments || err == sql.ErrNoRows
}
