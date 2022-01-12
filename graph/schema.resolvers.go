package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"database/sql"
	_ "github.com/lib/pq"

	"github.com/mercy34mercy/gqlgen-todos/graph/generated"
	"github.com/mercy34mercy/gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	
	DATABASE_URL := "postgres://aiddbjmxylnjxm:d0d3756986638bd8f399a370d3aace891221c4bffb21c4543c3e65f2d25e7cd0@ec2-54-225-187-177.compute-1.amazonaws.com:5432/der1lubvsuh4mb"
	var a sqlimage
	//sql Open
	db, err := sql.Open("postgres", DATABASE_URL)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

			//sqlにrequest送信
	if result, err := db.Query("SELECT * FROM BANMESHI ORDER BY random() LIMIT 1;"); err != nil {
			fmt.Printf("Error get database table: %q", err)
			//return c.String(http.StatusOK, "Error get database table")
	} else {
		for result.Next() {
			fmt.Printf("%s",result)
			result.Scan(&a.foodImageUrl,&a.mediumImageUrl,&a.recipeCost,&a.recipeId,&a.recipeMaterial,&a.recipeTitle,&a.smallImageUrl)
			fmt.Printf("foodImageUrl: %s, mediumImageUrl: %s, recipeCost:%, recipeId:%s,recipeMaterial:%s, recipeTitle:%s, smallImageUrl:%s\n", string(a.foodImageUrl), a.mediumImageUrl,a.recipeCost,a.recipeId,a.recipeMaterial,a.recipeTitle,a.smallImageUrl)
		}
	}

	return []*model.Todo{
		{
		FoodImageURL:a.foodImageUrl,
		MediumImageURL:a.mediumImageUrl,
		RecipeCost:a.recipeCost,
		RecipeID: a.recipeId, 
		RecipeMaterial: a.recipeMaterial,
		RecipeTitle: a.recipeTitle,
		SmallImageURL: a.smallImageUrl,},
	}, nil

}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
