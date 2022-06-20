package mongofile

import (
	"context"
	"fmt"

	"github.com/CrowhopTech/shinysorter/backend/pkg/filedb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ListQuestions will return the list of all questions. There are no filter options as this
// list will never be extremely large.
func (mc *mongoConnection) ListQuestions(ctx context.Context) ([]*filedb.Question, error) {
	cursor, err := mc.questionsCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error while running Find: %v", err)
	}

	results := []*filedb.Question{}

	for cursor.Next(ctx) {
		var result filedb.Question
		err := cursor.Decode(&result)
		if err != nil {
			return nil, fmt.Errorf("error while running Decode: %v", err)
		}
		results = append(results, &result)
	}

	return results, nil
}

func (mc *mongoConnection) getNewQuestionID(ctx context.Context) (int64, error) {
	// Find the highest existing ID, then set the ID to one higher
	highestResCursor := mc.questionsCollection.FindOne(ctx, bson.M{}, &options.FindOneOptions{
		Sort: bson.M{
			"_id": -1,
		},
	})

	highestID := int64(0)
	highestRes := filedb.Question{}

	err := highestResCursor.Decode(&highestRes)
	if err != mongo.ErrNoDocuments {
		if err != nil {
			return -1, fmt.Errorf("failed to find highest existing question ID: %v", err)
		}

		highestID = highestRes.ID
	}

	return highestID + 1, nil
}

func (mc *mongoConnection) CreateQuestion(ctx context.Context, q *filedb.Question) (*filedb.Question, error) {
	// TODO: validate question values

	newQuestionID, err := mc.getNewQuestionID(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while getting new question ID: %v", err)
	}

	q.ID = newQuestionID

	res, err := mc.questionsCollection.InsertOne(ctx, *q)
	if err != nil {
		return nil, fmt.Errorf("error while inserting question: %v", err)
	}

	created := mc.questionsCollection.FindOne(ctx, bson.M{"_id": res.InsertedID})
	if created.Err() != nil {
		return nil, fmt.Errorf("error while fetching created question: %v", created.Err())
	}

	createdQuestion := filedb.Question{}
	if err := created.Decode(&createdQuestion); err != nil {
		return nil, fmt.Errorf("error while decoding created question: %v", err)
	}

	return &createdQuestion, nil
}

func (mc *mongoConnection) ModifyQuestion(ctx context.Context, q *filedb.Question) (*filedb.Question, error) {
	trueVal := true

	// TODO: validate that object already exists (we can currently create questions using this call...)

	setParams := bson.M{}

	if len(q.QuestionText) > 0 {
		setParams["questionText"] = q.QuestionText
	}

	// TODO: handle these three parameters better
	if len(q.TagOptions) > 0 {
		setParams["tagOptions"] = q.TagOptions
	}

	if q.RequiresQuestion != nil {
		setParams["requiresQuestion"] = *q.RequiresQuestion
	}

	if q.OrderingID > 0 {
		setParams["orderingID"] = q.OrderingID
	}

	res, err := mc.questionsCollection.UpdateOne(ctx, bson.M{"_id": q.ID}, bson.M{"$set": setParams}, &options.UpdateOptions{
		Upsert: &trueVal,
	})
	if err != nil {
		return nil, fmt.Errorf("error while updating question: %v", err)
	}

	// TODO: better distinguish between "ID didn't exist" and "document matched original"
	if res.ModifiedCount == 0 {
		return nil, fmt.Errorf("no documents updated")
	}

	updated := mc.questionsCollection.FindOne(ctx, bson.M{"_id": q.ID})
	if updated.Err() != nil {
		return nil, fmt.Errorf("error while fetching updated question: %v", updated.Err())
	}

	updatedQuestion := filedb.Question{}
	if err := updated.Decode(&updatedQuestion); err != nil {
		return nil, fmt.Errorf("error while decoding updated question: %v", err)
	}

	return &updatedQuestion, nil
}

func (mc *mongoConnection) DeleteQuestion(ctx context.Context, id int64) error {
	res, err := mc.questionsCollection.DeleteOne(ctx, bson.M{"$_id": id})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no documents deleted")
	}
	return nil
}
