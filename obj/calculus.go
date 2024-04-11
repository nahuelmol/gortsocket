package obj

import ( 
    "fmt"
)
func abs(val_1, val_2 float64) float64 {
    var result float64
    if val_1 > val_2 {
        result = val_1 - val_2
        return result
    } else {
        result = val_2 - val_1
        return result
    }
}

func DistanceXY() float64 {
    //take the last location in user Stack
    //take the last location in dirver Stack
    //xdis := abs(userDistance.x, driverDistance.x)
    //ydis := abs(userDistance.y, driverDistance.y)
    var distance float64 = 11
    //distance := sqrt(pow(xdis) + pow(ydis))
    fmt.Println(distance)
    return distance
}
