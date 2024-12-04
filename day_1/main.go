package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func load_input(input_path string) ([]int, []int) {
	// Convert input file into two string slices, sorted in ascending order

	var first_column []int
	var second_column []int

	// Load input file as file
	file, err := os.Open(input_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read through file line by line and convert strings to ints
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		left_number, err := strconv.Atoi(scanner.Text()[0:5])
		first_column = append(first_column, left_number)

		right_number, err := strconv.Atoi(scanner.Text()[8:13])
		second_column = append(second_column, right_number)

		if err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Sort columns
	sort.Slice(first_column, func(i, j int) bool {
		return first_column[i] < first_column[j]
	})
	sort.Slice(second_column, func(i, j int) bool {
		return second_column[i] < second_column[j]
	})

	return first_column, second_column
}

func compare_numbers(num1 int, num2 int) int {

	if num1 > num2 {
		return num1 - num2
	} else {
		return num2 - num1
	}
}

func calculate_similarity_score(left_column []int, right_column []int) int {
	// For every value in left_column, check if it exists in right_colum.
	// If it exists, increment the corresponding key in duplicates by 1

	var duplicates = make(map[int]int)
	var sim_score = 0

	for _, i := range left_column {
		for _, num := range right_column {
			if num == i {
				duplicates[i] += 1
			}
		}
	}

	// Calculate the similarity score by multiplying the value times the amount it repeated
	for k, v := range duplicates {
		sim_score += k * v
	}
	return sim_score
}

func main() {

	first_column, second_column := load_input("input.txt")

	if len(first_column) != len(second_column) {
		log.Fatal("The lists are of unequal length")
	}

	sim_score := calculate_similarity_score(first_column, second_column)

	diff := 0
	for i := range len(first_column) {
		diff += compare_numbers(first_column[i], second_column[i])
	}

	fmt.Println("1.1:", diff)
	fmt.Println("1.2:", sim_score)
}
