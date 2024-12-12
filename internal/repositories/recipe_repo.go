package repositories

type Recipe struct {
	ID          int
	Title       string
	Ingredients string
	Instructions string
}

func SaveRecipeToDB(recipe Recipe) error {
	// Placeholder for saving to database
	return nil
}
