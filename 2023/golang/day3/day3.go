package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var matrix []string

func isNum(ch byte) bool {
    return '0' <= ch && ch <= '9'
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
    }

    file_scanner := bufio.NewScanner(file)

    for file_scanner.Scan() {
       values := file_scanner.Text()
       fmt.Println(values)
       matrix = append(matrix, values)

   }

   var sum int = 0

   for x, row := range matrix {
       var number int = 0
       var isSymbol bool = false
       for y, column := range row {
           if isNum(byte(column)) {
               intNum, _ := strconv.Atoi(string(column))
               number = number*10+intNum

               for dx := -1; dx <= 1; dx++ { // -1,-1 -1,0 -1,1  0,-1 0,0 0,1  1,-1 1,0 1,1
                   for dy := -1; dy <= 1; dy++ {

                       if x-dx >= 0 && y-dy >= 0 && x-dx < len(matrix) && y-dy < len(row) {

                           check := matrix[x-dx][y-dy]
                           if (!isNum(check) && check != '.') {
                               isSymbol = true
                           }
                       }
                   }
               }
           } else if number > 0 { // ended number (you're now in dot or symbol)
               if isSymbol {
                   sum += number
                   isSymbol = false
               }
               number = 0
           }
       }

       if number > 0 { // if number on last collumn is valid it goes here
           if isSymbol {
               sum += number
           }
       }
   }

   fmt.Println("Result:", sum)
}
