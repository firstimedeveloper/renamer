package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := "01 - German - I.mp3"

	newName, err := match(fileName)
	if err != nil {
		fmt.Println("no match")
		os.Exit(0)
	}
	fmt.Println(newName)
}

//match returns the new filename or an error if the file name
// didn't match the pattern
func match(fileName string) (string, error) {
	//"01", "German", "I "

	pieces := strings.Split(fileName, ".")
	ext := pieces[len(pieces)-1]
	tmp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(tmp, " - ")
	name := strings.Join(pieces[1:], " ")
	number := pieces[0]
	// number, err := strconv.Atoi(pieces[0])
	// if err != nil {
	// 	return "", fmt.Errorf("%s didn't match our pattern", fileName)
	// }
	return fmt.Sprintf("%s - %s.%s", name, number, ext), nil
}
