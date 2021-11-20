package recipes

import "context"

func (s *service) RemoveRecipe(ctx context.Context, id int64) error {
	return s.recipesPersist.DeleteRecipe(ctx, id)
}
