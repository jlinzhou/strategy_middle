package main

import (
	"fmt"
	"log"
	"strategy_middle/models"
)

func main() {
	c := models.ConnecToDB("student2")

	students := make([]models.Student, 20)
	err := c.Find(nil).All(&students)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(students[0])
}
