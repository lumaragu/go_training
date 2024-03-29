package haversine

import "math"

const earthRadius = 6371

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

	d := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLng/2), 2)
	distance := 2 * math.Atan2(math.Sqrt(d), math.Sqrt(1-d)) * earthRadius

	return distance
}
