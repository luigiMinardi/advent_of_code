package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    file, _ := os.Open("input.txt")
    
    file_scanner := bufio.NewScanner(file)

    var part1_answr int

    for file_scanner.Scan() {
        var hasValid bool = false
        var roundPoints int = 1

        value := strings.Split(strings.TrimSpace(strings.Split(file_scanner.Text(), ": ")[1]), " | ")

        win_nums := strings.Split(strings.ReplaceAll(value[0], "  ", " "), " ")
        my_nums := strings.Split(strings.ReplaceAll(value[1], "  ", " "), " ")
        fmt.Println(win_nums)
        fmt.Println(my_nums)

        for _, num := range win_nums {
            for _, mnum := range my_nums {
                if (num == mnum) {
                    if(!hasValid) {
                        hasValid = true
                    } else {
                        roundPoints *= 2 
                    }
                }
            }
            fmt.Println(num, hasValid, roundPoints)
        }

        if hasValid {
            part1_answr += roundPoints
            fmt.Println(roundPoints, part1_answr)
        }

        fmt.Println(part1_answr)
    }
}
