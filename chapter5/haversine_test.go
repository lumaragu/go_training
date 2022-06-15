package main

import "testing"

func TestHarvesine(t *testing.T) {
	type args struct {
		point1 Coordinate
		point2 Coordinate
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"AthensAmsterdam",
			args{
				point1: Coordinate{Lat: 37.983972, Lng: 23.727806},
				point2: Coordinate{Lat: 52.366667, Lng: 4.9},
			},
			2163.2310285824487,
		},
		{
			"AthensAmsterdam",
			args{
				point1: Coordinate{Lat: 52.366667, Lng: 4.9},
				point2: Coordinate{Lat: 52.516667, Lng: 13.388889},
			},
			575.2949643958797,
		},
		{
			"AthensAmsterdam",
			args{
				point1: Coordinate{Lat: 52.516667, Lng: 13.388889},
				point2: Coordinate{Lat: 37.983972, Lng: 23.727806},
			},
			1803.1087879059257,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Harvesine(tt.args.point1, tt.args.point2); got != tt.want {
				t.Errorf("Harvesine() = %v, want %v", got, tt.want)
			}
		})
	}
}
