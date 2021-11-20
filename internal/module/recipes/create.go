package recipes

import (
	"context"
	"time"

	"github.com/iDevoid/recipes/internal/constant/model"
)

func (s *service) CreateRecipe(ctx context.Context, data *model.Recipe) error {
	now := time.Now()
	data.CreateAt = now
	data.UpdateAt = now

	s.transaction.Begin(&ctx)
	return s.recipesPersist.InsertRecipes(ctx, data)
}
