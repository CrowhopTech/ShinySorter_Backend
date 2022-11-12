package main

import (
	"fmt"

	"github.com/CrowhopTech/shinysorter/backend/pkg/filedb"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/models"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/server/restapi/operations/questions"
	"github.com/go-openapi/runtime/middleware"
)

func translateDBQuestionToREST(question *filedb.Question) *models.QuestionEntry {
	return &models.QuestionEntry{
		OrderingID:        &question.OrderingID,
		QuestionID:        &question.ID,
		QuestionText:      &question.QuestionText,
		TagOptions:        tagOptionArrayToSwagger(question.TagOptions),
		MutuallyExclusive: question.MutuallyExclusive,
	}
}

//ListQuestions lists all registered questions
func ListQuestions(params questions.ListQuestionsParams) middleware.Responder {
	requestCtx := rootCtx

	results, err := imageMetadataConnection.ListQuestions(requestCtx)
	if err != nil {
		return questions.NewListQuestionsInternalServerError().WithPayload(fmt.Sprintf("failed to list questions: %v", err))
	}

	output := []*models.QuestionEntry{}

	for _, question := range results {
		converted := translateDBQuestionToREST(question)
		output = append(output, converted)
	}

	return questions.NewListQuestionsOK().WithPayload(output)
}

func CreateQuestion(params questions.CreateQuestionParams) middleware.Responder {
	requestCtx := rootCtx

	createdQuestion, err := imageMetadataConnection.CreateQuestion(requestCtx, &filedb.Question{
		// TODO: determine new ID? Or will it auto-increment?
		OrderingID:   *params.NewQuestion.OrderingID,
		QuestionText: *params.NewQuestion.QuestionText,
		TagOptions:   tagOptionArrayTofiledb(params.NewQuestion.TagOptions),
	})
	if err != nil {
		return questions.NewCreateQuestionInternalServerError().WithPayload(fmt.Sprintf("failed to insert question: %v", err))
	}

	output := translateDBQuestionToREST(createdQuestion)

	return questions.NewCreateQuestionCreated().WithPayload(output)
}

func PatchQuestionByID(params questions.PatchQuestionByIDParams) middleware.Responder {
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

	if params.Patch.MutuallyExclusive != "" {
		me := false
		if params.Patch.MutuallyExclusive == "true" {
			me = true
		}
		question.MutuallyExclusive = &me
	}

	newQuestion, err := imageMetadataConnection.ModifyQuestion(requestCtx, &question)
	if err != nil {
		return questions.NewPatchQuestionByIDInternalServerError().WithPayload(fmt.Sprintf("failed to modify question entry %d: %v", params.ID, err))
	}

	output := translateDBQuestionToREST(newQuestion)

	return questions.NewPatchQuestionByIDOK().WithPayload(output)
}

func DeleteQuestion(params questions.DeleteQuestionParams) middleware.Responder {
	requestCtx := rootCtx

	err := imageMetadataConnection.DeleteQuestion(requestCtx, params.ID)
	if err != nil {
		return questions.NewDeleteQuestionInternalServerError().WithPayload(fmt.Sprintf("failed to delete question: %v", err))
	}

	return questions.NewDeleteQuestionOK()
}

func ReorderQuestions(params questions.ReorderQuestionsParams) middleware.Responder {
	requestCtx := rootCtx

	err := imageMetadataConnection.ReorderQuestions(requestCtx, params.NewOrder)
	if err != nil {
		return questions.NewReorderQuestionsInternalServerError().WithPayload(fmt.Sprintf("failed to reorder questions: %v", err))
	}

	return questions.NewReorderQuestionsOK()
}
