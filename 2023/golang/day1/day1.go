package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
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

func MyIndexNumValue(str string, search []string) int {
    var found_index int = -1
    var value int = -1
    for j := 0; j < len(search); j++ {
        //fmt.Println("*",search[j])
        if found_index == 0 {
            return value 
        } else if search[j] == "1234567890" { // check if you're searching digits (123) or number names (one two three)
            index := strings.IndexAny(str, search[j])
            //fmt.Println(index, found_index, "digits")
            if index != -1 && index < found_index || found_index == -1 && index != -1 {
                found_index = index
                value, _ = strconv.Atoi(string(str[index]))
            }
        } else {
            index := strings.Index(str, search[j])
            //fmt.Println(index, found_index, "numbers")
            if index != -1 && index < found_index || found_index == -1 && index != -1 {
                found_index = index
                value = j // one == arr[1]... so it works
            }
        }
    }
    return value 
}

func MyLastIndexNumValue(str string, search []string) int {
    var found_index int = -1
    var value int = -1
    for j := len(search) -1; j >= 0; j-- {
        //fmt.Println("*",search[j])
        if search[j] == "1234567890" { // check if you're searching digits (123) or number names (one two three)
            index := strings.LastIndexAny(str, search[j])
            //fmt.Println(index, found_index, "digits")
            if index != -1 && index > found_index || found_index == -1 && index != -1 {
                found_index = index
                value, _ = strconv.Atoi(string(str[index]))
            }
        } else {
            index := strings.LastIndex(str, search[j])
            //fmt.Println(index, found_index, "numbers")
            if index != -1 && index > found_index || found_index == -1 && index != -1 {
                found_index = index
                value = j // one == arr[1]... so it works
            }
        }
    }
    return value 
}

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
        var search = []string{"1234567890", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
        
        first_index := fmt.Sprint(MyIndexNumValue(text,search))
        last_index := fmt.Sprint(MyLastIndexNumValue(text,search))

        row_num, _ := strconv.Atoi(first_index+last_index)
        summed_num += row_num 
        fmt.Println(summed_num)
    }

    file.Close()

    fmt.Println("Result:")
    fmt.Println(summed_num)
}
