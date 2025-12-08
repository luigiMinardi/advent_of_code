package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// This function signature equivalent to the bufio.SplitFunc so that it can be
// used on scanner.Split
func onComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// This loop through a small byte array of the scanner trying to find
	// a comma, if it does it returns that data and the next starting point
	for i, v := range data {
		// v is a rune, ',' means rune 44 (comma)
		if v == ',' {
			return i + 1, data[:i], nil
		}
	}
	// not at the end of the file but still data remaining
	if !atEOF {
		// signal the Scanner to read more data into the slice and try again
		// with a longer slice starting at the same point in the input.
		return 0, nil, nil
	}
	// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
	// but does not trigger an error to be returned from Scan itself.
	return 0, data, bufio.ErrFinalToken
}

func part1() {
	// The ranges are separated by commas (,);
	// each range gives its first ID and last ID separated by a dash (-).
	// invalid IDs are any ID which is made only of some sequence of digits
	// repeated twice (55, 6464, 123123)
	// None of the numbers have leading zeroes; 0101 isn't an ID at all.
	// find all of the invalid IDs
	// What do you get if you add up all of the invalid IDs?
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(onComma)

	var sumInvalidIDs int

	for scanner.Scan() {
		text := scanner.Text()
		var firstID int
		var lastID int
		for i, v := range text {
			if v == '-' {
				firstID, err = strconv.Atoi(text[:i])
				if err != nil {
					fmt.Println(err)
				}
				n := text[i+1:]
				if n[len(n)-1] == '\n' {
					// removing \n from slice
					n = n[:len(n)-1]
				}
				lastID, err = strconv.Atoi(n)
				if err != nil {
					fmt.Println(err)
				}
				break
			}
		}
		//fmt.Println("txt", text)
		//fmt.Println("ids", firstID, lastID)
		for i := firstID; i <= lastID; i++ {
			iStr := strconv.Itoa(i)
			if len(iStr)%2 == 0 {
				half := len(iStr) / 2
				// check if its a repeted digits sequence
				if iStr[:half] == iStr[half:] {
					//fmt.Println("GOTTEM", i)
					sumInvalidIDs += i
				}
				//fmt.Println(i, "even len", half, iStr[:half], iStr[half:])
			} else {
				//fmt.Println(i, "odd len")
			}
		}
	}
	fmt.Println(sumInvalidIDs)
}

// Splits string str every N characters
// Returns slice of substrings and if all substrings are equal
func splitStrInNCharsReturnItAndIfEqual(str string, N int) ([]string, bool) {
	var slice []string
	for charIdx := 0; charIdx < len(str); charIdx += N {
		if charIdx+N > len(str) {
			slice = append(slice, str[charIdx:])
			break
		}
		slice = append(slice, str[charIdx:charIdx+N])
	}
	isEqual := true
	for v := 1; v < len(slice); v++ {
		if slice[v] != slice[v-1] {
			isEqual = false
			break
		}
	}
	return slice, isEqual
}
func part2() {
	// The ranges are separated by commas (,);
	// each range gives its first ID and last ID separated by a dash (-).
	// ID is invalid if it is made only of some sequence of digits repeated
	// at least twice. (12341234, 123123123, 1212121212, 1111111)
	// None of the numbers have leading zeroes; 0101 isn't an ID at all.
	// find all of the invalid IDs
	// What do you get if you add up all of the invalid IDs?
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(onComma)

	var sumInvalidIDs int

	for scanner.Scan() {
		text := scanner.Text()
		var firstID int
		var lastID int
		for i, v := range text {
			if v == '-' {
				firstID, err = strconv.Atoi(text[:i])
				if err != nil {
					fmt.Println(err)
				}
				n := text[i+1:]
				if n[len(n)-1] == '\n' {
					// removing \n from slice
					n = n[:len(n)-1]
				}
				lastID, err = strconv.Atoi(n)
				if err != nil {
					fmt.Println(err)
				}
				break
			}
		}

		for i := firstID; i <= lastID; i++ {
			iStr := strconv.Itoa(i)
			half := len(iStr) / 2
			// start checking half, if fails reduce testingLen until 1
			for j := range half {
				testingLen := half - j
				left := iStr[:testingLen]
				right := iStr[half+j:]
				rightShift := iStr[half+j+1:]
				// fixing odd flooring by ignoring middle index
				if len(left) != len(right) {
					// odd
					if left == rightShift {
						//fmt.Println("GOTTEM", i)
						_, isEqual := splitStrInNCharsReturnItAndIfEqual(iStr, testingLen)
						if isEqual {
							sumInvalidIDs += i
							break
						}
					}
					continue
				}
				// even
				if left == right {
					//fmt.Println("GOTTEM", i)
					_, isEqual := splitStrInNCharsReturnItAndIfEqual(iStr, testingLen)
					if isEqual {
						sumInvalidIDs += i
						break
					}
				}
			}
		}
	}
	fmt.Println(sumInvalidIDs)
}
func main() {
	part2()
}
