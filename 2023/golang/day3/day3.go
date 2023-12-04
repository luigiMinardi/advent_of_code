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
   var sum_ratios int = 0

   var gears map[[2]int][]int = make(map[[2]int][]int)

   for x, row := range matrix {
       var number int = 0
       var isSymbol bool = false
       var stars map[int]int = make(map[int]int)
       for y, column := range row {
           if isNum(byte(column)) {
               intNum, _ := strconv.Atoi(string(column))
               number = number*10+intNum

               for dx := -1; dx <= 1; dx++ { // -1,-1 -1,0 -1,1  0,-1 0,0 0,1  1,-1 1,0 1,1
                   for dy := -1; dy <= 1; dy++ {
                       chx, chy := x-dx, y-dy
                       if chx >= 0 && chy >= 0 && chx < len(matrix) && chy < len(row) {

                           check := matrix[chx][chy]
                           if (!isNum(check) && check != '.') {
                               isSymbol = true
                           }

                           if check == '*' {
                               stars[chx] = chy
                           }
                       }
                   }
               }
           } else if number > 0 { // ended number (you're now in dot or symbol)
               for key, val := range stars {
                   gears[[2]int{key,val}] = append(gears[[2]int{key, val}], number)
               }
               if isSymbol {
                   sum += number
                   isSymbol = false
               }
               number = 0
               stars = make(map[int]int)
           }
       }

       if number > 0 { // if number on last collumn is valid it goes here
           for key, val := range stars {
               gears[[2]int{key,val}] = append(gears[[2]int{key, val}], number)
           }
           if isSymbol {
               sum += number
           }
       }
   }

   for key := range gears {
       if len(gears[key]) == 2 {
           ratio := gears[key][0] * gears[key][1]
           sum_ratios += ratio
       }
   }
   fmt.Println("Result:", sum)
   fmt.Printf("sum_ratios: %v\n", sum_ratios)
}
