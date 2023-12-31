package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func numOfInts(line string) int {
	count := 0
	for x := 0; x < len(line); x++ {
		if unicode.IsDigit(rune(line[x])) {
			count++
		}
	}
	return count
}

func readLine(path string) []string {
	file, err := os.Open(path)
	array := []string{}
	if err != nil {
		log.Fatal(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			array = append(array, scanner.Text())
		}
	}
	return array
}

func main() {
	lines := readLine("./file.txt")
	for x := 0; x < len(lines); x++ {
		numRead := ""
		switch numOfInts(lines[x]) {
		case 0:
			fmt.Println("No integers in this line:" + lines[x])

		case 1:
			for y := 0; y < len(lines[x]); y++ {
				if unicode.IsDigit(rune(lines[x][y])) {
					numRead = string(lines[x][y]) + string(lines[x][y])
					break
				}
			}

		default:
			count := 0
			for y := 0; y < len(lines[x]); y++ {
				if unicode.IsDigit(rune(lines[x][y])) {
					switch count {
					case 0:
						numRead = string(lines[x][y])
						count++
					case 1:
						numRead = numRead + string(lines[x][y])
						count++
					case 2:
						numRead = string(numRead[0]) + string(lines[x][y])
					}
				}
			}
		}
		fmt.Println(numRead)
	}
}
