package query

const (
	RecipeInsert = `
		INSERT INTO recipes (
			title,
			making_time,
			serves,
			ingredients,
			cost,
			created_at,
			updated_at
		) VALUES (	
			:title,
			:making_time,
			:serves,
			:ingredients,
			:cost,
			:created_at,
			:updated_at	
		) RETURNING id;
	`

	RecipesSelectAll = `
		SELECT
			id,
			title,
			making_time,
			serves,
			ingredients,
			cost,
			created_at,
			updated_at
		FROM
			recipes;
	`

	RecipeSelect = `
		SELECT
			id,
			title,
			making_time,
			serves,
			ingredients,
			cost,
			created_at,
			updated_at
		FROM
			recipes
		WHERE
			id = :id;
	`

	RecipeUpdate = `
		UPDATE recipes
		SET
			title = :title,
			making_time = :making_time,
			serves = :serves,
			ingredients = :ingredients,
			cost = :cost
		WHERE
			id = :id;
	`

	RecipeDelete = `
		DELETE FROM
			recipes
		WHERE
			id = :id;
	`
)
