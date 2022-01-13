package graph

//from sql server
type sqlimage struct {
	foodImageUrl string `db:"foodImageUrl"`
	mediumImageUrl string `db:"mediumImageUrl"`
	recipeCost string `db:"recipeCost"`
	recipeId string `db:"recipeId"`
	recipeMaterial string `db:"recipeMaterial"`
	recipeTitle string `db:"recipeTitle"`
	recipeUrl string `db:"recipeUrl"`
	smallImageUrl string `db:"smallImageUrl"`
}