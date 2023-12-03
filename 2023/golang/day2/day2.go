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

    file_scanner := bufio.NewScanner(file)
    
    var current_game_id = 1
    var part1_answer = 0
    var part2_answer = 0

    for file_scanner.Scan() {
        fmt.Println(file_scanner.Text())
        var rgbmap map[string]int = make(map[string]int)

        all_rounds := strings.Split(strings.Split(file_scanner.Text(), ": ")[1], "; ") // ["1 green, 3 red, 6 blue", "3 green, 6 red", "9 blue"]
        for i := range all_rounds {
            round_cubes := strings.Split(all_rounds[i], ", ") // ["1 green", "3 red", "6 blue"]
            for j := range round_cubes {
                cube_values := strings.Split(round_cubes[j], " ") // ["123", "red/green/blue"] 
                v, _ := strconv.Atoi(cube_values[0])

                if rgbmap[cube_values[1]] < v {
                    rgbmap[cube_values[1]] = v
                }
                fmt.Printf("rounds: %v\n", rgbmap)
            }
        }
        part2_answer += rgbmap["red"]*rgbmap["green"]*rgbmap["blue"]

        if rgbmap["red"] <= 12 && rgbmap["green"] <= 13 && rgbmap["blue"] <= 14 {
            part1_answer += current_game_id
        }
        current_game_id++
    }

    fmt.Printf("part1_answer: %v\n", part1_answer)
    fmt.Printf("part2_answer: %v\n", part2_answer)
}
