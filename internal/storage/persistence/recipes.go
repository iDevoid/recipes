package persistence

//go:generate mockgen -destination=../../../mocks/recipes/persistence_mock.go -package=recipes_mock -source=recipes.go

import (
	"context"

	"github.com/iDevoid/cptx"
	"github.com/iDevoid/recipes/internal/constant/model"
	"github.com/iDevoid/recipes/internal/constant/query"
	"github.com/jmoiron/sqlx"
)

type RecipesPersist interface {
	InsertRecipes(ctx context.Context, recipe *model.Recipe) error
	UpdateRecipe(ctx context.Context, recipe *model.Recipe) error
	DeleteRecipe(ctx context.Context, id int64) error
	SelectAllRecipes(ctx context.Context) (res []model.Recipe, err error)
	SelectRecipe(ctx context.Context, id int64) (res model.Recipe, err error)
}

type recipesPersist struct {
	db cptx.Database
}

func RecipesInit(db cptx.Database) RecipesPersist {
	return &recipesPersist{
		db,
	}
}

func (rp *recipesPersist) InsertRecipes(ctx context.Context, recipe *model.Recipe) error {
	params := map[string]interface{}{
		"title":       recipe.Title,
		"making_time": recipe.MakingTime,
		"serves":      recipe.Serves,
		"ingredients": recipe.Ingredients,
		"cost":        recipe.Cost,
		"created_at":  recipe.CreateAt,
		"updated_at":  recipe.UpdateAt,
	}

	return rp.db.Main().QueryRowMustTx(ctx, query.RecipeInsert, params, &recipe.ID)
}

func (rp *recipesPersist) UpdateRecipe(ctx context.Context, recipe *model.Recipe) error {
	params := map[string]interface{}{
		"title":       recipe.Title,
		"making_time": recipe.MakingTime,
		"serves":      recipe.Serves,
		"ingredients": recipe.Ingredients,
		"cost":        recipe.Cost,
	}

	return rp.db.Main().QueryRowMustTx(ctx, query.RecipeUpdate, params)
}

func (rp *recipesPersist) DeleteRecipe(ctx context.Context, id int64) error {
	params := map[string]interface{}{
		"id": id,
	}
	return rp.db.Main().QueryRowMustTx(ctx, query.RecipeDelete, params)
}

func (rp *recipesPersist) SelectAllRecipes(ctx context.Context) (res []model.Recipe, err error) {
	err = rp.db.Replica().SelectContext(ctx, &res, query.RecipesSelectAll)
	return
}

func (rp *recipesPersist) SelectRecipe(ctx context.Context, id int64) (res model.Recipe, err error) {
	params := map[string]interface{}{
		"id": id,
	}
	query, args, _ := sqlx.Named(query.RecipeSelect, params)
	query = rp.db.Replica().Rebind(query)
	err = rp.db.Replica().GetContext(ctx, &res, query, args...)
	return
}
