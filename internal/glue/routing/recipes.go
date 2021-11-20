package routing

import (
	"net/http"

	"github.com/iDevoid/recipes/internal/handler/rest"
	"github.com/iDevoid/recipes/platform/routers"
)

// RecipesRouting returns the list of routers for domain recipes
func RecipesRouting(handler rest.RecipesHandler) []routers.Router {
	return []routers.Router{
		{
			Method:  http.MethodGet,
			Path:    "/test",
			Handler: handler.Test,
		},

		{
			Method:  http.MethodPost,
			Path:    "/recipes",
			Handler: handler.CreateRecipe,
		},

		{
			Method:  http.MethodGet,
			Path:    "/recipes/:id",
			Handler: handler.GetRecipe,
		},
		{
			Method:  http.MethodGet,
			Path:    "/recipes",
			Handler: handler.GetAllRecipes,
		},

		{
			Method:  http.MethodPatch,
			Path:    "/recipes/:id",
			Handler: handler.UpdateRecipe,
		},

		{
			Method:  http.MethodDelete,
			Path:    "/recipes/:id",
			Handler: handler.DeleteRecipe,
		},
	}
}
