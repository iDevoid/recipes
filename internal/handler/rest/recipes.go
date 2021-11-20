package rest

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iDevoid/recipes/internal/constant/model"
	"github.com/iDevoid/recipes/internal/constant/state"
	"github.com/iDevoid/recipes/internal/module/recipes"
)

//go:generate mockgen -destination=../../../mocks/recipes/handler_mock.go -package=recipes_mock -source=recipes.go

type RecipesHandler interface {
	Test(c *fiber.Ctx) error
	CreateRecipe(ctx *fiber.Ctx) error
	GetAllRecipes(ctx *fiber.Ctx) error
	GetRecipe(ctx *fiber.Ctx) error
	UpdateRecipe(ctx *fiber.Ctx) error
	DeleteRecipe(ctx *fiber.Ctx) error
}

type recipesHandler struct {
	recipesCase recipes.Usecase
}

// RecipesInit is to initialize the rest handler for domain recipes
func RecipesInit(recipesCase recipes.Usecase) RecipesHandler {
	return &recipesHandler{
		recipesCase,
	}
}

// Test is handler testing
func (uh *recipesHandler) Test(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World!")
}

func (h *recipesHandler) CreateRecipe(ctx *fiber.Ctx) error {
	ctx.Context().SetContentType("application/json; charset=UTF-8")
	failedResp := model.FailedResp{
		Message:  state.ErrorCreation,
		Required: state.Required,
	}
	rawFail, _ := json.Marshal(failedResp)

	var body model.Recipe
	err := json.Unmarshal(ctx.Body(), &body)
	if err != nil || !body.Valid() {
		return ctx.Send(rawFail)
	}

	err = h.recipesCase.CreateRecipe(ctx.Context(), &body)
	if err != nil {
		return ctx.Send(rawFail)
	}

	resp := model.SuccessResp{
		Message: state.SuccessCreation,
		Recipes: []model.Recipe{
			body,
		},
	}

	raw, _ := json.Marshal(resp)
	return ctx.Send(raw)
}

func (h *recipesHandler) GetAllRecipes(ctx *fiber.Ctx) error {
	data, err := h.recipesCase.GetAllRecipes(ctx.Context())
	if err != nil {
		return err
	}

	resp := model.SuccessResp{
		Recipes: data,
	}
	raw, _ := json.Marshal(resp)
	ctx.Context().SetContentType("application/json; charset=UTF-8")
	return ctx.Send(raw)
}

func (h *recipesHandler) GetRecipe(ctx *fiber.Ctx) error {
	ctx.Context().SetContentType("application/json; charset=UTF-8")
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return err
	}
	data, err := h.recipesCase.GetRecipe(ctx.Context(), int64(id))
	if err != nil {
		return err
	}

	resp := model.SuccessResp{
		Message: "Recipe details by id",
		Recipes: []model.Recipe{
			data,
		},
	}
	raw, _ := json.Marshal(resp)
	return ctx.Send(raw)
}

func (h *recipesHandler) UpdateRecipe(ctx *fiber.Ctx) error {
	ctx.Context().SetContentType("application/json; charset=UTF-8")
	failedResp := model.FailedResp{
		Message:  state.ErrorCreation,
		Required: state.Required,
	}
	rawFail, _ := json.Marshal(failedResp)

	var body model.Recipe
	err := json.Unmarshal(ctx.Body(), &body)
	if err != nil || !body.Valid() {
		return ctx.Send(rawFail)
	}

	err = h.recipesCase.UpdateRecipe(ctx.Context(), &body)
	if err != nil {
		return ctx.Send(rawFail)
	}

	resp := model.SuccessResp{
		Message: "Recipe successfully updated!",
		Recipes: []model.Recipe{
			body,
		},
	}

	raw, _ := json.Marshal(resp)
	return ctx.Send(raw)
}

func (h *recipesHandler) DeleteRecipe(ctx *fiber.Ctx) error {
	ctx.Context().SetContentType("application/json; charset=UTF-8")
	resp := map[string]string{
		"message": "No recipe found",
	}
	raw, _ := json.Marshal(resp)

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Send(raw)
	}
	err = h.recipesCase.RemoveRecipe(ctx.Context(), int64(id))
	if err != nil {
		return ctx.Send(raw)
	}

	resp["message"] = "Recipe successfully removed!"
	raw, _ = json.Marshal(resp)
	return ctx.Send(raw)
}
