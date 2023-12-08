package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isDigit(ch byte) bool {
    return ch >= '0' && ch <= '9'
}

func isOnRange(value int, start int, rangeN int) bool {
    return value == start || value > start && value <= start+rangeN
}

func main() {
    file, _ := os.Open("example.txt")

    file_scanner := bufio.NewScanner(file)

    type SourceToDestMap struct {
        dest_range int
        source_range int
        range_len int
    }

    type PlantProcess struct {
        soil, fertilizer, water, light, temperature, humidity, location int
    }

    type Categories int
    const (
        seed_to_soil Categories = iota
        soil_to_fetilizer
        fertilizer_to_water
        water_to_light
        light_to_temperature
        temperature_to_humidity
        humidity_to_location
    )

    var maps map[Categories][]SourceToDestMap = make(map[Categories][]SourceToDestMap)
    var plant_data map[int]PlantProcess = make(map[int]PlantProcess)

    var seeds []int
    var line int = 0
    var cat Categories = -1
    for file_scanner.Scan() {
        if line == 0 {
            seeds_str := strings.Split(strings.Split(file_scanner.Text(), ": ")[1]," ")
            for seed := range seeds_str {
                new_seed, _ := strconv.Atoi(seeds_str[seed])
                seeds = append(seeds, new_seed)
            }
            fmt.Println(seeds)
        } else if file_scanner.Text() != ""{
            if !isDigit(file_scanner.Text()[0]) {
                cat++
            } else {
                values_str := strings.Split(file_scanner.Text(), " ")
                var values [3]int
                for value := range values {
                    new_value, _ := strconv.Atoi(values_str[value])
                    values[value] = new_value
                }
                maps[cat] = append(maps[cat], SourceToDestMap{values[0], values[1], values[2]})
            }
        }
        line++
    }

    var searching = 0
    var curr_cat Categories = 0 
    fmt.Println(plant_data, searching, curr_cat)
    fmt.Println(maps)
    for _, seed := range seeds {
        fmt.Println(seed, isOnRange(seed, maps[0][1].source_range, maps[0][1].range_len))
    }
}
