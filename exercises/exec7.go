package exercises

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func scarpeUrl() {
	file, err := os.Open("./tempFile/urls.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	scarpeUrl()
}
