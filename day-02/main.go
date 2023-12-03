package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const (
	redCubes   = 12
	greenCubes = 13
	blueCubes  = 14

	outerPattern = "Game (?P<gameNum>[[:digit:]]{1,}): (?P<sets>.*)"
	colorPattern = "(?P<count>[[:digit:]]{1,}) (?P<color>.*)"
	file         = "02_input.txt"
)

// Game N: x green, y blue, z red; ...
// keep a total of each color per game. Once a color exceeds the set limit, that game is not possible.
// If the game is possible, add N to total
type reg map[string]*regexp.Regexp

// https://stackoverflow.com/questions/44406077
type re *regexp.Regexp

func patternToMap(re *regexp.Regexp, b []byte) map[string][]byte {
	matches := re.FindAllSubmatch(b, -1)
	result := make(map[string][]byte)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = matches[0][i]
		}
	}
	return result
}

func main() {

	patterns := make(reg)
	patterns["outer"] = regexp.MustCompile(outerPattern)
	patterns["color"] = regexp.MustCompile(colorPattern)

	data, err := os.ReadFile(file)
	p(err)

	// https://stackoverflow.com/questions/25691879
	lines := bytes.Split(data, []byte("\n"))
	// Sum of gameNums that satisfy conditions
	var gameNumTotal, gameSetTotal int
	var impossible bool
	// Each 'line' is a game
	for _, line := range lines {
		// TODO - figure out how to name matches in output
		// https://stackoverflow.com/questions/20750843/
		gameMatches := patternToMap(patterns["outer"], line)

		gameNum, err := strconv.Atoi(string(gameMatches["gameNum"]))
		p(err)
		sets := gameMatches["sets"]

		impossible = false
		// Decision is by set, not by game
		var setMax = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, set := range bytes.Split(sets, []byte(";")) {
			m := make(map[string]int)
			_ = m
			for _, color := range bytes.Split(set, []byte(",")) {
				colorMap := patternToMap(patterns["color"], color)
				count, err := strconv.Atoi(string(colorMap["count"]))
				p(err)
				c := string(colorMap["color"])
				m[c] = count
				if m[c] > setMax[c] {
					setMax[c] = m[c]
				}
			}
			if m["red"] > redCubes ||
				m["green"] > greenCubes ||
				m["blue"] > blueCubes {
				impossible = true
			}
		}

		if !impossible {
			gameNumTotal += gameNum
		}

		gameSetTotal += mapProduct(setMax)

		/*
			for _, color := range bytes.Split(sets, []byte(",")) {
				colorMap := patternToMap(patterns["color"], color)
				count, err := strconv.Atoi(string(colorMap["count"]))
				p(err)
				colorCounts[string(colorMap["color"])] += count
			}
		*/

		/*
			for _, set := range bytes.Split(sets, []byte(";")) {
				for _, color := range bytes.Split(set, []byte(",")) {
					v := patterns["color"].FindAllSubmatch(color, -1)
					rv := make(map[string][]byte)
					for j, n := range patterns["color"].SubexpNames() {
						if j != 0 && n != "" {
							rv[n] = v[0][j]
						}
					}
					c, err := strconv.Atoi(string(rv["count"]))
					p(err)
					colorCounts[string(rv["color"])] += c
				}
			}
		*/

		/*
			if colorCounts["red"] <= redCubes &&
				colorCounts["green"] <= greenCubes &&
				colorCounts["blue"] <= blueCubes {
				gameNumTotal += gameNum
				fmt.Printf("%d ", gameNum)
			}
		*/
	}
	fmt.Println(gameSetTotal)
}

func p(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func mapProduct(m map[string]int) int {
	var total int = 1
	if len(m) == 0 {
		return 0
	}
	for _, v := range m {
		total *= v
	}
	return total
}
