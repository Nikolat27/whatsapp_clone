package convert

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StringToObjectId(str string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(str)
}
