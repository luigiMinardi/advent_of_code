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
    file, _ := os.Open("input.txt")

    file_scanner := bufio.NewScanner(file)

    type SourceToDestMap struct {
        dest_range int
        source_range int
        range_len int
    }

    type PlantProcess int
    const ( 
        soil PlantProcess = iota
        fertilizer
		water
		light
		temperature
		humidity
		location
    )

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
    var plant_data map[int]map[PlantProcess]int = make(map[int]map[PlantProcess]int)

    var p1 int = int(^uint(0) >> 1) // max int

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
    fmt.Println(plant_data, searching)
    fmt.Println(maps)
    for _, seed := range seeds {
        plant_data[seed] = make(map[PlantProcess]int)
        searching = seed
        for curr_cat := 0; curr_cat < len(maps); curr_cat++ {
            var isMapped bool = false
            for _, value := range maps[Categories(curr_cat)] {
                if isOnRange(searching, value.source_range, value.range_len) {
                    curr_process := value.dest_range+searching-value.source_range
                    plant_data[seed][PlantProcess(curr_cat)] = curr_process
                    isMapped = true
                    break
                }
            }
            if !isMapped { 
                plant_data[seed][PlantProcess(curr_cat)] = searching
            }
            searching = plant_data[seed][PlantProcess(curr_cat)]
        }
        if plant_data[seed][6] < p1 {
            fmt.Println("------------------------", plant_data[seed][6])
            p1 = plant_data[seed][6]
        }
    }
    fmt.Println(p1)
}
