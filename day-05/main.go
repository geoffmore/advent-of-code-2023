package main

import (
	"bytes"
	"fmt"
	"github.com/geoffmore/advent-of-code-2023/aoclib"
	"regexp"
	"strconv"
	"strings"
)

const pattern = "CHANGEME"

const file = "input.txt"

type problemMap []problem

// [3]int{soil=50, seed=98, 2}
// expands to soil={50, 51} seed={98, 99}
// problem is [3]int{dest, src, range}
type problem [3]int

type Foo struct {
	seeds              []int
	seedSoilMap        problemMap
	soilFertilizerMap  problemMap
	fertilizerWaterMap problemMap
	waterLightMap      problemMap
	lightTempMap       problemMap
	tempHumidMap       problemMap
	humidLocMap        problemMap
}

// This will fill up memory very quickly
func (p problemMap) genMap() map[int]int {
	// map[src]dst
	m := make(map[int]int)

	for _, prob := range p {
		for i := 0; i < prob[2]; i++ {
			m[prob[1]+i] = prob[0] + i
		}
	}
	return m
}

func (pm problemMap) mapFunc(i int) int {
	for _, p := range pm {
		dst, src, r := p[0], p[1], p[2]-1
		if i >= src && i <= src+r {
			// i is within src + r, so just add diff of dst and src
			i = i + (dst - src)
			break
		}
	}
	return i
}

func aToProblem(b []byte) problem {
	var p problem
	nums := bytes.Split(b, []byte(" "))
	if len(nums) == 3 {
		for i, num := range nums {
			n, err := aoclib.B2i(num)
			aoclib.PanicIf(err)
			p[i] = n
		}
	}
	return p
}

func unmarshalField(b []byte) problemMap {
	var p problemMap
	lines := bytes.Split(b, []byte("\n"))
	for i := 1; i < len(lines); i++ {
		p = append(p, aToProblem(lines[i]))
	}
	return p
}

func unmarshal(data []byte, f *Foo) error {
	b := bytes.Split(data, []byte("\n\n"))
	for i := range b {
		if i == 0 {
			s := strings.Split(string(b[i]), " ")
			for j := 1; j < len(s); j++ {
				n, err := strconv.Atoi(s[j])
				aoclib.PanicIf(err)
				f.seeds = append(f.seeds, n)
			}
		} else if i == 1 {
			f.seedSoilMap = unmarshalField(b[i])
		} else if i == 2 {
			f.soilFertilizerMap = unmarshalField(b[i])
		} else if i == 3 {
			f.fertilizerWaterMap = unmarshalField(b[i])
		} else if i == 4 {
			f.waterLightMap = unmarshalField(b[i])
		} else if i == 5 {
			f.lightTempMap = unmarshalField(b[i])
		} else if i == 6 {
			f.tempHumidMap = unmarshalField(b[i])
		} else if i == 7 {
			f.humidLocMap = unmarshalField(b[i])
		}
	}
	return nil
}
func main() {
	data, err := aoclib.Scan(file)
	aoclib.PanicIf(err)
	fmt.Println("Problem A: ", ProblemA(data))
	fmt.Println("Problem B: ", ProblemB(data))
}

func (f Foo) getLocFromSeed(i int) int {
	v := i
	v = f.seedSoilMap.mapFunc(v)
	v = f.soilFertilizerMap.mapFunc(v)
	v = f.fertilizerWaterMap.mapFunc(v)
	v = f.waterLightMap.mapFunc(v)
	v = f.lightTempMap.mapFunc(v)
	v = f.tempHumidMap.mapFunc(v)
	v = f.humidLocMap.mapFunc(v)
	return v
}

func ProblemA(data []byte) int {
	var f Foo
	var locs []int

	// unmarshal data into f
	unmarshal(data, &f)
	// generate mappings into b
	for _, seed := range f.seeds {
		locs = append(locs, f.getLocFromSeed(seed))
	}

	i := locs[0] // Assumes loc has len >=1
	for _, loc := range locs {
		if i > loc {
			i = loc
		}
	}
	return i
}

func ProblemB(data []byte) int {
	var total int
	re := regexp.MustCompile(pattern)

	for _, line := range data {
		_, _ = line, re
	}
	return total
}
