package main

type RecipeItem struct {
	Ingredient
	Weight_grams float32
	Volume_ml    float32
	//if volume or weight is defined, the weight / volume
	//is the sum of the weight/volume of the quant.
	Quantity float32
}
type Recipe struct {
	RecipeItems []RecipeItem
}
