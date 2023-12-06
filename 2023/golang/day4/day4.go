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

    var cards map[int][]int = make(map[int][]int)
    var total_cards int = 0
    var game int = 1

    for file_scanner.Scan() {
        var hasValid bool = false
        var roundPoints int = 1
        var matches int = 0

        value := strings.Split(strings.TrimSpace(strings.Split(file_scanner.Text(), ": ")[1]), " | ")

        win_nums := strings.Split(strings.ReplaceAll(value[0], "  ", " "), " ")
        my_nums := strings.Split(strings.ReplaceAll(value[1], "  ", " "), " ")

        for _, num := range win_nums {
            for _, mnum := range my_nums {
                if (num == mnum) {
                    matches+= 1
                    for range cards[game] {
                        cards[game+matches] = append(cards[game+matches], game+matches)
                    }
                    if len(cards[game]) > 0 {
                    }
                    if(!hasValid) {
                        hasValid = true
                    } else {
                        roundPoints *= 2 
                    }
                }
            }
        }

        if hasValid {
            part1_answr += roundPoints
        } else {
            cards[game] = append(cards[game], game)
        }

        for match := 1; match <= matches; match++ {
            if hasValid {
                cards[game] = append(cards[game], game)
                hasValid = false
            }
            cards[game+match] = append(cards[game+match], game+match)
        }

        game++
    }

    for _, val := range cards {
        total_cards += len(val)
    }

    fmt.Println("p1",part1_answr)
    fmt.Println("p2",total_cards)
}
