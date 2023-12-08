package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"

	"github.com/felipefbs/graphql/graph/model"
)

// Courses is the resolver for the courses field.
func (r *categoryResolver) Courses(ctx context.Context, obj *model.Category) ([]*model.Course, error) {
	modelList, err := r.CourseRepository.FindAllByCategoryID(obj.ID)
	if err != nil {
		return nil, err
	}

	response := make([]*model.Course, len(modelList))

	for k, v := range modelList {
		response[k] = &model.Course{
			ID:          v.ID,
			Name:        &v.Name,
			Description: &v.Description,
			Category: &model.Category{
				ID: v.CategoryID,
			},
		}
	}

	return response, nil
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	categoryModel, err := r.CategoryRepository.Create(input.Name, *input.Description)
	if err != nil {
		return nil, err
	}

	return &model.Category{
		ID:          categoryModel.ID,
		Name:        categoryModel.Name,
		Description: &categoryModel.Description,
	}, nil
}

// CreateCourse is the resolver for the createCourse field.
func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	courseModel, err := r.CourseRepository.Create(input.Name, *input.Description, input.CategoryID)
	if err != nil {
		return nil, err
	}

	return &model.Course{
		ID:          courseModel.ID,
		Name:        &courseModel.Name,
		Description: &courseModel.Description,
		Category: &model.Category{
			ID: courseModel.CategoryID,
		},
	}, nil
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	modelList, err := r.CategoryRepository.FindAll()
	if err != nil {
		return nil, err
	}

	response := make([]*model.Category, len(modelList))

	for k, v := range modelList {
		response[k] = &model.Category{
			ID:          v.ID,
			Name:        v.Name,
			Description: &v.Description,
		}
	}

	return response, nil
}

// Courses is the resolver for the courses field.
func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	modelList, err := r.CourseRepository.FindAll()
	if err != nil {
		return nil, err
	}

	response := make([]*model.Course, len(modelList))

	for k, v := range modelList {
		response[k] = &model.Course{
			ID:          v.ID,
			Name:        &v.Name,
			Description: &v.Description,
			Category: &model.Category{
				ID: v.CategoryID,
			},
		}
	}

	return response, nil
}

// Category returns CategoryResolver implementation.
func (r *Resolver) Category() CategoryResolver { return &categoryResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type (
	categoryResolver struct{ *Resolver }
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)
