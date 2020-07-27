package turf

import (
	"fmt"
	"testing"
)

func TestDistance(t *testing.T) {
	d := Distance(-77.03653, 38.89768, -77.05173, 38.8973)
	if d != 1.317556974720262 {
		t.Errorf("error calculating the distance")
	}
}

func TestPointDistance(t *testing.T) {
	p1 := Point{Lng: -77.03653, Lat: 38.89768}
	p2 := Point{Lng: -77.05173, Lat: 38.8973}
	d := PointDistance(p1, p2)
	if d != 1.317556974720262 {
		t.Errorf("error calculating the distance")
	}
}

func TestBearing(t *testing.T) {
	b := Bearing(-77.03653, 38.89768, -77.05173, 38.8973)
	fmt.Printf("bearing: %v", b)
	if b <= 0 {
		t.Errorf("error calculating the bearing")
	}
}

func TestPointBearing(t *testing.T) {
	p1 := Point{Lng: -77.03653, Lat: 38.89768}
	p2 := Point{Lng: -77.05173, Lat: 38.8973}
	b := PointBearing(p1, p2)
	if b < 0 {
		t.Errorf("error calculating the bearing")
	}
}

func TestMidPoint(t *testing.T) {

	type args struct {
		p1 Point
		p2 Point
	}

	tests := map[string]struct {
		args    args
		wantErr bool
		want    Point
	}{
		"happy path: same lng": {
			args: args{
				p1: Point{Lat: 23.38, Lng: 43.16},
				p2: Point{Lat: 26.38, Lng: 43.16},
			},
			wantErr: false,
			want:    Point{Lat: 24.88, Lng: 43.16},
		},
		"happy path: same lat": {
			args: args{
				p1: Point{Lat: 23.38, Lng: 43.20},
				p2: Point{Lat: 23.38, Lng: 44.20},
			},
			wantErr: false,
			want:    Point{Lat: 23.38079468036304, Lng: 43.699999999999996},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			m := MidPoint(tt.args.p1, tt.args.p2)
			fmt.Printf("midPoint %v", m)
			if tt.want != m {
				t.Errorf("error calculating the midpoint")
				return
			}
		})
	}

}

func TestDestinationPoint(t *testing.T) {
	p := Point{Lat: 23.34, Lng: 43.25}
	d := Destination(p, 10, 230)
	fmt.Printf("destination point: %v", d)

	e := Point{Lat: 23.28223959663299, Lng: 43.175084627817945}

	if e.Lat != d.Lat && e.Lng != d.Lng {
		t.Errorf("error calculating the destination point")
	}

}
