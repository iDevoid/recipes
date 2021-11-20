package recipes

import (
	"context"

	"github.com/iDevoid/recipes/internal/constant/model"
)

func (s *service) GetAllRecipes(ctx context.Context) ([]model.Recipe, error) {
	return s.recipesPersist.SelectAllRecipes(ctx)
}

func (s *service) GetRecipe(ctx context.Context, id int64) (model.Recipe, error) {
	return s.recipesPersist.SelectRecipe(ctx, id)
}
