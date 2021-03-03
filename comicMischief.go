package main

import "fmt"

func main() {
	var publisher, writer, artist, title string
	publisher, writer, artist, title = "DizzyBooks Publishing Inc.", "Tracey Hatchet", "Jewel Tampson", "Mr. GoToSleep"
	var year, pageNumber int
	year, pageNumber = 1997, 14
	var grade float32
	grade = 6.5

	fmt.Println(title, "written by", writer, "drawn by", artist, "publisheded by", publisher, "written in the year", year, "with", pageNumber, "pages", "gets a grade of", grade)

	publisher, writer, artist, title = "DizzyBooks Publishing Inc.", "Ryan N. Shawn", "Phoebe Paperclips", "Epic Vol. 1"
	year, pageNumber = 2013, 160
	grade = 9.0

	fmt.Println(title, "written by", writer, "drawn by", artist, "publisheded by", publisher, "written in the year", year, "with", pageNumber, "pages", "gets a grade of", grade)
}
