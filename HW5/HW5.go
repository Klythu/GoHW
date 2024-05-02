// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	f, err := os.OpenFile("archive.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	count := 1
	for {

		input := ""
		fmt.Scanf("%s\n", &input)
		if input == "exit" {
			break
		}
		timenow := time.Now().UTC()
		dataf := timenow.Format("2006-01-02 15:04:05")
		dataf = strconv.Itoa(count) + " " + dataf + " " + input + "\n"
		count++
		f.WriteString(dataf)

	}
	f.Close()
	f, err = os.OpenFile("archive.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	f.Chmod(0444)
	content, err := os.ReadFile("archive.txt")
	if err != nil {
		fmt.Println("no such file")
		panic(err)
	}
	fmt.Println(string(content))
}
