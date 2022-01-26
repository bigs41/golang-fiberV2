package main

import (
	"github.com/gofiber/fiber/v2"

	"log"

	"exv2/utils/collection"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		// type Student struct {
		// 	Name  string
		// 	Marks []int
		// }

		// students := []Student{
		// 	{Name: "Hugo", Marks: []int{91, 88, 76, 93}},
		// 	{Name: "Rick", Marks: []int{70, 73, 66, 90}},
		// 	{Name: "Michael", Marks: []int{73, 80, 75, 88}},
		// 	{Name: "Fadi", Marks: []int{82, 75, 66, 84}},
		// 	{Name: "Peter", Marks: []int{67, 78, 70, 82}},
		// }

		// var approvedStudents []Student
		// From(students).
		// 	WhereT(
		// 		func(student Student) bool {
		// 			return student.Name == "Peter"
		// 		},
		// 	).
		// 	ToSlice(&approvedStudents)

		inputData4 := collection.CollSets{
			{
				"word": "hello man1!",
				"code": "321234",
			},
			{
				"word": "hello man2!",
				"code": "321234",
			},
			{
				"word": "hello man!",
				"code": "123456",
			},
			{
				"word": "hello girl!",
				"code": "789102",
			},
		}

		log.Print(inputData4)

		return c.JSON(inputData4.WhereStr("word", "hello man!").Groupby("code").OrderBy().Get())
	})

	app.Listen(":3000")
}
