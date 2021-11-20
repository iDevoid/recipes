package recipes

import (
	"context"

	"github.com/iDevoid/recipes/internal/constant/model"
)

func (s *service) UpdateRecipe(ctx context.Context, data *model.Recipe) error {
	s.transaction.Begin(&ctx)
	return s.recipesPersist.UpdateRecipe(ctx, data)
}
