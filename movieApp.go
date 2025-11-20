package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("ENTER YOUR NAME: ")
	userName, _ := reader.ReadString('\n')
	userName = strings.TrimSpace(userName)

	fmt.Println("HOW MANY MOVIES DO YOU WANT TO ADD TO YOUR COLLECTION: ")
	movieCount, _ := reader.ReadString('\n')
	movieCount = strings.TrimSpace(movieCount)
	count, _ := strconv.Atoi(movieCount)

	movies := []struct {
		Title  string
		Year   int
		Rating int
	}{}

	for i := 1; i <= count; i++ {
		fmt.Printf("\nMOVIE %d: ", i)

		//Getting the title of the movie
		fmt.Println("TITLE: ")
		title, _ := reader.ReadString('\n')
		title = strings.TrimSpace(title)

		//Getting the year of the movie
		fmt.Println("YEAR: ")
		yearStr, _ := reader.ReadString('\n')
		year, _ := strconv.Atoi(strings.TrimSpace(yearStr))

		fmt.Println("RATING (1-10): ")
		ratingStr, _ := reader.ReadString('\n')
		rating, _ := strconv.Atoi(strings.TrimSpace(ratingStr))

		movies = append(movies, struct{
			Title string; 
			Year int; 
			Rating int}{
				title,
				year,
				rating,
			})
	}
	fmt.Println("\nYOUR MOVIE COLLECTION:")
	for i, movie := range movies {
		fmt.Printf("%s (%d) - Rating: %d/10\n", movie.Title, movie.Year, movie.Rating)
		if movie.Rating >= 8 {
			fmt.Println("⭐ Highly Recommended!")
		}else if movie.Rating >= 5 {
			fmt.Println("👍 Worth Watching.")
		}else {
			fmt.Println("👎 Maybe Skip This One.")
		}
		fmt.Printf("TOTAL MOVIES: %d", i)
		fmt.Println("AVERAGE RATINGS: ")

	}


}
