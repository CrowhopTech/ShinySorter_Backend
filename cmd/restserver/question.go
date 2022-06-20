package main

import (
	"fmt"

	"github.com/CrowhopTech/shinysorter/backend/pkg/filedb"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/models"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func translateDBQuestionToREST(question *filedb.Question) *models.Question {
	return &models.Question{
		OrderingID:   question.OrderingID,
		QuestionID:   question.ID,
		QuestionText: question.QuestionText,
		TagOptions:   tagOptionArrayToSwagger(question.TagOptions),
	}
}

//ListQuestions lists all registered questions
func ListQuestions(params operations.ListQuestionsParams) middleware.Responder {
	requestCtx := rootCtx

	results, err := imageMetadataConnection.ListQuestions(requestCtx)
	if err != nil {
		return operations.NewListQuestionsInternalServerError().WithPayload(fmt.Sprintf("failed to list questions: %v", err))
	}

	output := []*models.Question{}

	for _, question := range results {
		converted := translateDBQuestionToREST(question)
		output = append(output, converted)
	}

	return operations.NewListQuestionsOK().WithPayload(output)
}

func CreateQuestion(params operations.CreateQuestionParams) middleware.Responder {
	requestCtx := rootCtx

	createdQuestion, err := imageMetadataConnection.CreateQuestion(requestCtx, &filedb.Question{
		ID:           params.NewQuestion.QuestionID,
		OrderingID:   params.NewQuestion.OrderingID,
		QuestionText: params.NewQuestion.QuestionText,
		TagOptions:   tagOptionArrayTofiledb(params.NewQuestion.TagOptions),
	})
	if err != nil {
		return operations.NewCreateQuestionInternalServerError().WithPayload(fmt.Sprintf("failed to insert question: %v", err))
	}

	output := translateDBQuestionToREST(createdQuestion)

	return operations.NewCreateQuestionCreated().WithPayload(output)
}

func PatchQuestionByID(params operations.PatchQuestionByIDParams) middleware.Responder {
	requestCtx := rootCtx

	question := filedb.Question{
		ID: params.ID,
	}

	if len(params.Patch.QuestionText) > 0 {
		question.QuestionText = params.Patch.QuestionText
	}

	// TODO: handle these three parameters better
	if len(params.Patch.TagOptions) > 0 {
		question.TagOptions = tagOptionArrayTofiledb(params.Patch.TagOptions)
	}

	if params.Patch.OrderingID > 0 {
		question.OrderingID = params.Patch.OrderingID
	}

	newQuestion, err := imageMetadataConnection.ModifyQuestion(requestCtx, &question)
	if err != nil {
		return operations.NewPatchQuestionByIDInternalServerError().WithPayload(fmt.Sprintf("failed to modify question entry %d: %v", params.ID, err))
	}

	output := translateDBQuestionToREST(newQuestion)

	return operations.NewPatchQuestionByIDOK().WithPayload(output)
}
