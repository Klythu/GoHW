// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

func main() {
	count := 1
	strinf := ""
	for {

		input := ""
		fmt.Scanf("%s\n", &input)
		if input == "exit" {
			break
		}
		timenow := time.Now().UTC()
		dataf := timenow.Format("2006-01-02 15:04:05")
		strinf = strinf + strconv.Itoa(count) + " " + dataf + " " + input + "\n"
		count++
	}
	ioutil.WriteFile("archive.txt", []byte(strinf), 0644)

	information, err := ioutil.ReadFile("archive.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", information)
}
