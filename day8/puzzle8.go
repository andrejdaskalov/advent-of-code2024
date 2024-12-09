package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFromFile(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("file not present.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	antennaMap := make([][]string, 0, 50)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		antennaMap = append(antennaMap, split)
	}
	return antennaMap
}

type Location struct {
	x int
	y int
}
type LocationMap map[string][]Location

type LocationSet map[Location]bool

func CreateLocationMap(input [][]string) (LocationMap, int) {
	locationMap := make(LocationMap)
	for i, row := range input {
		for j, e := range row {
			if e == "." {
				continue
			}
			slc, ok := locationMap[e]
			if !ok {
				slc = make([]Location, 0, 50)
				locationMap[e] = slc
			}
			locationMap[e] = append(locationMap[e], Location{i, j})
		}
	}
	return locationMap, len(input)
}

// part 1
func FindAntinodesLocation(a Location, b Location) (Location, Location) {
	a1x := a.x + (a.x - b.x)
	a1y := a.y + (a.y - b.y)
	a1 := Location{a1x, a1y}

	a2x := b.x + (b.x - a.x)
	a2y := b.y + (b.y - a.y)
	a2 := Location{a2x, a2y}

	return a1, a2
}

// part 2
func FindAllAntinodeLocations(a Location, b Location, size int) []Location {
	locations := make([]Location, 0, 50)
	diff1 := Location{a.x - b.x, a.y - b.y}
	diff2 := Location{b.x - a.x, b.y - a.y}

	tmp := Location{a.x, a.y}
	for tmp.x >= 0 && tmp.x < size && tmp.y >= 0 && tmp.y < size {
		locations = append(locations, tmp)
		tmp = Location{tmp.x + diff1.x, tmp.y + diff1.y}
	}
	tmp = Location{b.x, b.y}
	for tmp.x >= 0 && tmp.x < size && tmp.y >= 0 && tmp.y < size {
		locations = append(locations, tmp)
		tmp = Location{tmp.x + diff2.x, tmp.y + diff2.y}
	}
	return locations
}

func GetAntinodeSet(locationMap LocationMap, size int) LocationSet {
	set := make(LocationSet)
	for _, slc := range locationMap {
		for _, a := range slc {
			for _, b := range slc {
				if a == b {
					continue
				}
				antinodeLocations := FindAllAntinodeLocations(a, b, size)
				for _, loc := range antinodeLocations {
					set[loc] = true
				}
				// a1, a2 := FindAntinodesLocation(a, b)
				// if a1.x >= 0 && a1.x < size && a1.y >= 0 && a1.y < size {
				// 	set[a1] = true
				// }
				// if a2.x >= 0 && a2.x < size && a2.y >= 0 && a2.y < size {
				// 	set[a2] = true
				// }

			}
		}
	}
	return set
}
func main() {
	input := ReadFromFile("day8-input.txt")
	mapLocations, size := CreateLocationMap(input)
	fmt.Println(mapLocations["M"])
	set := GetAntinodeSet(mapLocations, size)
	fmt.Println(len(set))

}
