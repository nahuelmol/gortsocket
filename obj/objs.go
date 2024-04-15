package obj 

import (
    "fmt"
    "time"
    "math/rand"
    "errors"
)
//location and coordinates are different entities

type Coordinate struct {
    Xposition int32 
    Yposition int32

    Time time.Time
}

type User struct {
    coordiante Coordinate
    service string //type of service the user looks for
    id_sector int32
    drivers_ammount int32 //drivers available in user's sector
    drivers[]int32 //this should be an slice, a mutable collection
}

type Driver struct {
    id uint32
    coor Coordinate //driver's coordinates
}

func (driver *Driver) Identifier() uint32 {
    return driver.id
}

func (driver *Driver) SetLocation(x,y int32) Coordinate {
    //function that takes the current time and save it into a time variable
    var mycoor Coordinate
    mycoor.Xposition = x
    mycoor.Yposition = y
    mycoor.Time = time.Now()

    driver.id = uint32(rand.Intn(101))
    driver.coor = mycoor

    return mycoor
}

func (driver *Driver) getCoordinate() {
    fmt.Printf("driver id:%d\n", driver.id)
    fmt.Printf("driver x:%d \n", driver.coor.Xposition) 
    fmt.Printf("driver y:%d \n" , driver.coor.Yposition) 
    fmt.Printf("driver time:%d\n", driver.coor.Time.Hour()) 
}

type Node struct {
    data Coordinate
}

func (nn *Node) SeeData() Coordinate {
    fmt.Println(nn.data.Xposition)
    fmt.Println(nn.data.Yposition)
    fmt.Println(nn.data.Time)

    return nn.data
}

type StackLocation struct {
    head *Node
    next *Node

    length int
    nodes []*Node
}

func CreateStack() *StackLocation{
    return &StackLocation {
        head:nil,
        next:nil,

        length:0,
        nodes:nil,
    }
}

func CreateNode(newdata Coordinate) *Node {
    return &Node {
        data:newdata,
    }
}

func (sl *StackLocation) Push(newnode *Node) {
    sl.next = sl.head
    sl.head = newnode

    sl.length += 1
    sl.nodes = append(sl.nodes, newnode)
}

func (sl *StackLocation) Pop() {
    sl.head = sl.next
    sl.length -= 1
    sl.nodes = sl.nodes[:sl.length]
}

func (sl *StackLocation) Topdata() (*Node, error) {
    if sl.head == nil {
        return nil, errors.New("not a head")
    } else {
        return sl.head, nil
    }
}

func (sl *StackLocation) Nextdata() Coordinate {
    return sl.next.data
}

func (sl *StackLocation) Wholedata() []*Node {
    return sl.nodes
}

func (sl *StackLocation) Getlength() int {
    return sl.length
}



