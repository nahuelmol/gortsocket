
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
