package main

import (
	"log"
	"os"
	"text/template"
)

type restaurant struct {
	Name string
	Menu menu
}

type menu struct {
	Breakfast []menuItem
	Lunch     []menuItem
	Dinner    []menuItem
}

type menuItem struct {
	ItemName      string
	Price         string
	NutritionData nutritionData
}

type nutritionData struct {
	Calories      int
	Protien       int
	Fat           int
	Carbohydrates int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	m := menu{
		Breakfast: []menuItem{
			menuItem{
				ItemName: "Pancakes",
				Price:    "7.99",
				NutritionData: nutritionData{
					Calories:      1000,
					Protien:       27,
					Fat:           8,
					Carbohydrates: 57,
				},
			},
			menuItem{
				ItemName: "French Toast",
				Price:    "7.99",
				NutritionData: nutritionData{
					Calories:      1005,
					Protien:       34,
					Fat:           9,
					Carbohydrates: 62,
				},
			},
		},
		Lunch: []menuItem{
			menuItem{
				ItemName: "Hamburger",
				Price:    "7.99",
				NutritionData: nutritionData{
					Calories:      1200,
					Protien:       47,
					Fat:           12,
					Carbohydrates: 37,
				},
			},
			menuItem{
				ItemName: "Salad",
				Price:    "9.99",
				NutritionData: nutritionData{
					Calories:      700,
					Protien:       47,
					Fat:           12,
					Carbohydrates: 37,
				},
			},
		},
		Dinner: []menuItem{
			menuItem{
				ItemName: "Salad",
				Price:    "9.99",
				NutritionData: nutritionData{
					Calories:      700,
					Protien:       47,
					Fat:           12,
					Carbohydrates: 37,
				},
			},
			menuItem{
				ItemName: "Steak",
				Price:    "17.99",
				NutritionData: nutritionData{
					Calories:      1200,
					Protien:       47,
					Fat:           12,
					Carbohydrates: 0,
				},
			},
		},
	}
	restaurants := []restaurant{
		restaurant{
			Name: "Keg and Barrel",
			Menu: m,
		},
		restaurant{
			Name: "Trevors Burger Place",
			Menu: m,
		},
	}

	err := tpl.Execute(os.Stdout, restaurants)
	if err != nil {
		log.Fatalln(err)
	}
}
