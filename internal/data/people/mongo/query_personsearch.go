package mongo

import (
	"context"
	"fmt"
	"log"

	. "github.com/Goldwin/ies-pik-cms/pkg/common/queries"
	"github.com/Goldwin/ies-pik-cms/pkg/people/dto"
	"github.com/Goldwin/ies-pik-cms/pkg/people/queries"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type searchPersonImpl struct {
	ctx context.Context
	db  *mongo.Database
}

// Execute implements queries.SearchPerson.
func (s *searchPersonImpl) Execute(query queries.SearchPersonFilter) (queries.SearchPersonResult, QueryErrorDetail) {
	opts := options.Find().SetSort(bson.D{{Key: "_id", Value: 1}}).SetLimit(int64(query.Limit))
	regexOp := "$regex"
	cursor, err := s.db.Collection(personCollectionName).Find(s.ctx,
		bson.M{
			"_id": bson.M{"$gt": query.LastID},
			"$or": []interface{}{
				bson.M{"firstName": bson.M{regexOp: fmt.Sprintf("^%s", query.NamePrefix)}},
				bson.M{"middleName": bson.M{regexOp: fmt.Sprintf("^%s", query.NamePrefix)}},
				bson.M{"lastName": bson.M{regexOp: fmt.Sprintf("^%s", query.NamePrefix)}},
			},
		},
		opts,
	)
	if err != nil {
		log.Default().Printf("Failed to connect to database: %s", err.Error())
		return queries.SearchPersonResult{}, QueryErrorDetail{
			Code:    500,
			Message: "Failed to connect to database",
		}
	}
	defer cursor.Close(s.ctx)
	var result = make([]dto.Person, 0)

	for cursor.Next(s.ctx) {
		var person PersonModel
		if err := cursor.Decode(&person); err != nil {
			return queries.SearchPersonResult{}, QueryErrorDetail{
				Code:    500,
				Message: "Failed to Decode Person Information",
			}
		}
		result = append(result, toPersonDTO(person))
	}

	return queries.SearchPersonResult{
		Data: result,
	}, NoQueryError
}

func SearchPerson(ctx context.Context, db *mongo.Database) queries.SearchPerson {
	return &searchPersonImpl{
		ctx: ctx,
		db:  db,
	}
}
