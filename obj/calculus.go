package obj

import ( 
    "fmt"
    "math"
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

func eulerian(xdis, ydis float64) float64 {
    radical := math.Pow(xdis, 2) + math.Pow(ydis, 2)
    result  := math.Sqrt(radical)
    return result

}

func CalculateDistance(driver Coordinate, user Coordinate) float64 {
    //take the last location in user Stack
    //take the last location in dirver Stack
    xdis := abs(user.x, driver.xposition)
    ydis := abs(user.y, driver.yposition)
    distance := eulerian(xdis, ydis)
    return distance
}
