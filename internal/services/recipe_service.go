package services

type Recipe struct {
	Title       string
	Ingredients []string
	Instructions string
}

func SaveRecipe(recipe Recipe) error {
	// Placeholder for saving recipe
	return nil
}