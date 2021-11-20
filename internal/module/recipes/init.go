package recipes

import (
	"context"

	"github.com/iDevoid/cptx"
	"github.com/iDevoid/recipes/internal/constant/model"
	"github.com/iDevoid/recipes/internal/storage/persistence"
)

//go:generate mockgen -destination=../../../mocks/recipes/usecase_mock.go -package=recipes_mock -source=init.go

type Usecase interface {
	CreateRecipe(ctx context.Context, data *model.Recipe) error
	RemoveRecipe(ctx context.Context, id int64) error
	GetAllRecipes(ctx context.Context) ([]model.Recipe, error)
	GetRecipe(ctx context.Context, id int64) (model.Recipe, error)
	UpdateRecipe(ctx context.Context, data *model.Recipe) error
}

type service struct {
	transaction    cptx.Transaction
	recipesPersist persistence.RecipesPersist
}

func Initialize(transaction cptx.Transaction, recipesPersist persistence.RecipesPersist) Usecase {
	return &service{
		transaction,
		recipesPersist,
	}
}
