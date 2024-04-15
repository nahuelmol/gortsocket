package obj

import ( 
    "math"
)

func abs(val_1, val_2 int32) float64 {
    if val_1 > val_2 {
        result := float64(val_1) - float64(val_2)
        return result
    } else {
        result := float64(val_2) - float64(val_1)
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
    xdis := abs(user.Xposition, driver.Xposition)
    ydis := abs(user.Yposition, driver.Yposition)
    distance := eulerian(xdis, ydis)
    return distance
}
