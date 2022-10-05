package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, _ := os.Open("Hangman.txt")
	fileScanner := bufio.NewScanner(file)
	var array []string
	var temp int
	for fileScanner.Scan() {

		array = append(array, fileScanner.Text())

	}

	file.Close()
	fmt.Println(array)
	fmt.Println(array[0])
	fmt.Println(array[len(array)-1])

	for i := 0; i < len(array); i++ {
		if array[i] == "before" {
			temp, _ = strconv.Atoi(array[i+1])
		}
	}
	fmt.Println(array[temp-1])
}
