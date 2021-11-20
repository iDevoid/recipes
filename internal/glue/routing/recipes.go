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
			Method:  http.MethodPost,
			Path:    "/",
			Handler: handler.CreateRecipe,
		},

		{
			Method:  http.MethodGet,
			Path:    "/:id",
			Handler: handler.GetRecipe,
		},
		{
			Method:  http.MethodGet,
			Path:    "/",
			Handler: handler.GetAllRecipes,
		},

		{
			Method:  http.MethodPatch,
			Path:    "/:id",
			Handler: handler.UpdateRecipe,
		},

		{
			Method:  http.MethodDelete,
			Path:    "/:id",
			Handler: handler.DeleteRecipe,
		},
	}
}
