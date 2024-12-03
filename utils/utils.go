
package utils

import (
    "fmt"
    "math"

    "gonum.org/v1/gonum/geo"
    "gonum.org/v1/gonum/geo/coord"
)

func Takegeographics(geolat, geolon string){
    //this will take geographical coordinates
    lat := geolat * math.PI / 180.0
    lon := geolon * math.PI / 180.0
    
    p := coord.NewGeographic(lat, lon).ToPoint()
    x, y, _ := p.ECEF()

    fmt.Println("x:%.2f, y:%.2f", x, y)
}

func MakeProgressive(coordinates [][]string) {
    new_coords := make([][]float64, len(coordinates))
    for i, row := range coordinates {
        new_coords[i] = make([]float64, len(row))
        for j, element := range row {
            f, err := strconv.ParseFloat(element, 64)
            if err != nil {
                fmt.Println(err)
            } else {
                new_coords[i][j] = f
            }
        }
    }
    fmt.Println(new_coords)

    //deltas i-> lat, j-> lon
    deltas := make([][] floats64, len(newcoordinates) - 1)
    limits := len(new_coords) - 1;
    for i := 0; i <= limits; i++ {
        for j := 0; j <= 2; j++ {
            delta := new_coords[i+1][j] - new_coords[i][j]
            if j == 0 {
                distance := R * delta * math.PI / 180
                deltas[i][j] := distance
            } else {
                lat := new_coords[i][0]
                distance := R * math.Cos(lat) * delta
                deltas[i][j] := distance
            }
        }
    }
}
