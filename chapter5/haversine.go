package src

import (
	"fmt"
	"math"
)

const (
	earthRaidus = 6371
)

type Coordinate struct {
	Lat float64
	Lng float64
}

func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

func Harvesine(point1, point2 Coordinate) float64 {
	lat1 := degreesToRadians(point1.Lat)
	lng1 := degreesToRadians(point1.Lng)
	lat2 := degreesToRadians(point2.Lat)
	lng2 := degreesToRadians(point2.Lng)

	diffLat := lat2 - lat1
	diffLng := lng2 - lng1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLng/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return c * earthRaidus
}

func main() {
	Athens := Coordinate{Lat: 37.983972, Lng: 23.727806}
	Amsterdam := Coordinate{Lat: 52.366667, Lng: 4.9}
	Berlin := Coordinate{Lat: 52.516667, Lng: 13.388889}
	fmt.Println(Harvesine(Athens, Amsterdam))
	fmt.Println(Harvesine(Amsterdam, Berlin))
	fmt.Println(Harvesine(Berlin, Athens))
}
