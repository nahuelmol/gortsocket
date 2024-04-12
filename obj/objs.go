package obj 

import (
    "fmt"
    "time"
    "math/rand"
    "errors"
)
//location and coordinates are different entities

type Coordinate struct {
    xposition int32 
    yposition int32

    pTime time.Time
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
    mycoor.xposition = x
    mycoor.yposition = y
    mycoor.pTime = time.Now()

    driver.id = uint32(rand.Intn(101))
    driver.coor = mycoor

    return mycoor
}

func (driver *Driver) getCoordinate() {
    fmt.Printf("driver id:%d\n", driver.id)
    fmt.Printf("driver x:%d \n", driver.coor.xposition) 
    fmt.Printf("driver y:%d \n" , driver.coor.yposition) 
    fmt.Printf("driver time:%d\n", driver.coor.pTime.Hour()) 
}

type Node struct {
    data Coordinate
}

func (nn *Node) SeeData() Coordinate {
    fmt.Println(nn.data.xposition)
    fmt.Println(nn.data.yposition)
    fmt.Println(nn.data.pTime)

    return nn.data
}

type StackLocation struct {
    head *Node
    next *Node

    length int
    nodes []*Node
}

func  CreateStack() *StackLocation{

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

func (sl *StackLocation) Topdata() (Coordinate, error) {
    nilco := new(Coordinate)
    if sl.head == nil {
        return *nilco, errors.New("not a head")
    } else {
        return sl.head.data, nil
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



