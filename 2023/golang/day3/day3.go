package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var matrix [][]string

type VType uint8

const (
    number VType = iota
    dot
    symbol
)

type Value struct {
    value   string
    vtype   VType
    row     int 
    column  int
    isValid bool
}

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println(err)
    }

    file_scanner := bufio.NewScanner(file)

    var values_list [][]Value
    var sum_num = 0

    var rw = 0
    for file_scanner.Scan() {
       values := strings.Split(file_scanner.Text(), "")
       fmt.Println(values)
       matrix = append(matrix, values)

       var value_row []Value
       for cl := range values {
           if strings.Contains(values[cl], ".") {
               value_row = append(value_row, Value{values[cl],dot,rw,cl,false})

           } else if strings.ContainsAny(values[cl], "1234567890") {
               value_row = append(value_row, Value{values[cl],number,rw,cl,false})

           } else {
               value_row = append(value_row, Value{values[cl],symbol,rw,cl,false})
           }
       }
       values_list = append(values_list, value_row)
       rw++
    }

    for i, r := range values_list {
        for col := range r {
            current := values_list[i][col]

            if current.vtype == symbol {
                left_of_curr := &values_list[i][current.column-1]
                right_of_curr := &values_list[i][current.column+1]
                top_of_curr := &values_list[current.row-1][col]
                bottom_of_curr := &values_list[current.row+1][col]
                top_left_of_curr := &values_list[current.row-1][current.column-1]
                top_right_of_curr := &values_list[current.row-1][current.column+1]
                bottom_left_of_curr := &values_list[current.row+1][current.column-1]
                bottom_right_of_curr := &values_list[current.row+1][current.column+1]

                fmt.Println(current, "symb")

                if left_of_curr.vtype == number {
                    if !left_of_curr.isValid {
                        for c := left_of_curr.column; c >= 0; c-- {
                            if values_list[left_of_curr.row][c].vtype == number {
                                values_list[left_of_curr.row][c].isValid = true
                                fmt.Println(values_list[left_of_curr.row][c], "----------------")
                            } else {
                                break
                            }
                        }
                    }
                    fmt.Println(left_of_curr, "found left")
                }

                if right_of_curr.vtype == number {
                    if !right_of_curr.isValid {
                        for c := right_of_curr.column; c <= len(r); c++ {
                            if values_list[right_of_curr.row][c].vtype == number {
                                values_list[right_of_curr.row][c].isValid = true
                                fmt.Println(values_list[right_of_curr.row][c], "----------------")
                            } else {
                                break
                            }
                        }
                    }

                    fmt.Println(right_of_curr, "found right")
                }

                if top_of_curr.vtype == number {
                    top_of_curr.isValid = true
                    fmt.Println(top_of_curr, "found top")
                }

                if bottom_of_curr.vtype == number {
                    bottom_of_curr.isValid = true
                    fmt.Println(bottom_of_curr, "found bottom")
                }

                if top_left_of_curr.vtype == number {
                    if !top_left_of_curr.isValid {
                        for c := top_left_of_curr.column; c >= 0; c-- {
                            if values_list[top_left_of_curr.row][c].vtype == number {
                                values_list[top_left_of_curr.row][c].isValid = true
                                fmt.Println(values_list[top_left_of_curr.row][c], "----------------")
                            } else {
                                break
                            }
                        }
                    }

                    fmt.Println(top_left_of_curr, "found top left")
                }

                if top_right_of_curr.vtype == number {
                    if !top_right_of_curr.isValid {
                        for c := top_right_of_curr.column; c <= len(r); c++ {
                            if values_list[top_right_of_curr.row][c].vtype == number {
                                values_list[top_right_of_curr.row][c].isValid = true
                                fmt.Println(values_list[top_right_of_curr.row][c], "----------------")
                            } else {
                                break
                            }
                        }
                    }


                    fmt.Println(top_right_of_curr, "found top right")
                }

                if bottom_left_of_curr.vtype == number {
                    if !bottom_left_of_curr.isValid {
                        for c := bottom_left_of_curr.column; c >= 0; c-- {
                            if values_list[bottom_left_of_curr.row][c].vtype == number {
                                values_list[bottom_left_of_curr.row][c].isValid = true
                                fmt.Println(values_list[bottom_left_of_curr.row][c], "----------------")
                            } else {
                                break
                            }
                        }
                    }

                    fmt.Println(bottom_left_of_curr, "found bottom left")
                }

                if bottom_right_of_curr.vtype == number {
                    if !bottom_right_of_curr.isValid {
                        for c := bottom_right_of_curr.column; c <= len(r); c++ {
                            if values_list[bottom_right_of_curr.row][c].vtype == number {
                                fmt.Println(values_list[bottom_right_of_curr.row][c], "----------------")
                                values_list[bottom_right_of_curr.row][c].isValid = true
                            } else {
                                break
                            }
                        }
                    }

                    fmt.Println(bottom_right_of_curr, "found bottom right")
                }
            }
        }
    }

    fmt.Println(values_list)

    for i, r := range values_list {
        for el := 0; el < len(r); el++ {
            if values_list[i][el].isValid {
                number := ""
                for iel := el; iel < len(r); iel++ {
                    fmt.Println(values_list[i][iel]) 
                    if values_list[i][iel].isValid {
                        number += values_list[i][iel].value
                        el++
                    } else {
                        num, _ := strconv.Atoi(number)
                        sum_num += num
                        break
                    }
                    fmt.Println(number, "num")
                }
            }
            fmt.Println(sum_num)
        }
    }
}
