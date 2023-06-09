package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.27

import (
	"context"
	DBModel "gql-go/db/model"
	"gql-go/graph/model"
)

// CreateSketch is the resolver for the createSketch field.
func (r *mutationResolver) CreateSketch(ctx context.Context, input model.CreateSketchInput) (*model.Sketch, error) {
	tokenUser, err := GetTokenUser(ctx)
	if err != nil {
		panic(err)
	}

	sketch := DBModel.Sketch{
		UserID:  tokenUser.ID,
		Title:   input.Title,
		Content: input.Content,
	}

	result := r.DB.Create(&sketch)

	return &model.Sketch{
		ID:        sketch.ID,
		UserID:    sketch.UserID,
		Title:     sketch.Title,
		Content:   sketch.Content,
		CreatedAt: sketch.CreatedAt,
		UpdatedAt: sketch.UpdatedAt,
	}, result.Error
}

// UpdateSketch is the resolver for the updateSketch field.
func (r *mutationResolver) UpdateSketch(ctx context.Context, input *model.UpdateSketchInput) (*model.Sketch, error) {
	tokenUser, err := GetTokenUser(ctx)
	if err != nil {
		return nil, err
	}

	var sketch DBModel.Sketch
	result := r.DB.Where("id = ? AND user_id = ?", input.ID, tokenUser.ID).First(&sketch)

	if result.Error != nil {
		return nil, result.Error
	}

	sketch.Title = input.Title
	sketch.Content = input.Content

	result = r.DB.Save(&sketch)

	return &model.Sketch{
		ID:        sketch.ID,
		UserID:    sketch.UserID,
		Title:     sketch.Title,
		Content:   sketch.Content,
		CreatedAt: sketch.CreatedAt,
		UpdatedAt: sketch.UpdatedAt,
	}, result.Error
}

// DeleteSketch is the resolver for the deleteSketch field.
func (r *mutationResolver) DeleteSketch(ctx context.Context, id int) (bool, error) {
	tokenUser, err := GetTokenUser(ctx)
	if err != nil {
		return false, err
	}

	var sketch DBModel.Sketch
	result := r.DB.Where("id = ? AND user_id = ?", id, tokenUser.ID).First(&sketch)
	if result.Error != nil {
		return false, result.Error
	}

	result = r.DB.Delete(&sketch)

	return result.Error == nil, result.Error
}

// Sketches is the resolver for the sketches field.
func (r *queryResolver) Sketches(ctx context.Context) ([]*model.Sketch, error) {
	tokenUser, err := GetTokenUser(ctx)
	if err != nil {
		return nil, err
	}

	var sketches []*DBModel.Sketch
	result := r.DB.Where("user_id = ?", tokenUser.ID).Find(&sketches)

	var sketchModels []*model.Sketch
	for _, sketch := range sketches {
		sketchModels = append(sketchModels, &model.Sketch{
			ID:        sketch.ID,
			UserID:    sketch.UserID,
			Title:     sketch.Title,
			Content:   sketch.Content,
			CreatedAt: sketch.CreatedAt,
			UpdatedAt: sketch.UpdatedAt,
		})
	}

	return sketchModels, result.Error
}
