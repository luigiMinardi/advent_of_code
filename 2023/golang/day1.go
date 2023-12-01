package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }

    scanner := bufio.NewScanner(file)
    
    var summed_num int = 0
    
    for scanner.Scan() {
        fmt.Println(scanner.Text())
        var text = scanner.Text()
        var search = "1234567890"
        fmt.Println(strings.IndexAny(scanner.Text(), search))
        fmt.Println(strings.LastIndexAny(scanner.Text(), search))
        
        first_index := string(text[strings.IndexAny(text, search)])
        last_index  := string(text[strings.LastIndexAny(text, search)])

        row_num, _ := strconv.Atoi(first_index+last_index)
        summed_num += row_num 
        fmt.Println(summed_num)
    }

    file.Close()

    fmt.Println("Result:")
    fmt.Println(summed_num)
}
