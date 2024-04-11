package obj

import (
    "time"
)

type Distance struct {
    when time.Time
    val float64
}

type DistanceStack struct {
    head *Distance
    next *Distance

    length int
    distances []*Distance
}

func CreateDistancer() *DistanceStack{
    return &DistanceStack {
        head:nil,
        next:nil,

        length:0,
        distances:nil,
    }
}

func (ds *DistanceStack) Pop() {
    ds.head = ds.next
    ds.next = ds.distances[ds.length - 2]

    ds.length -= 1
    ds.distances = ds.distances[:ds.length]
}

func (ds *DistanceStack) Push(newdistance *Distance) {
    ds.next = ds.head
    ds.head = newdistance
    ds.length += 1
    ds.distances = append(ds.distances, newdistance)
}

func (ds *DistanceStack) Gethead() *Distance {
    return ds.head
}

func (ds *DistanceStack) Getnext() *Distance {
    return ds.next
}
