package obj

import (
    "fmt"
    "time"
)

type ServiceBuffer struct {
    head *Service
}

func CreateServiceBuffer() *ServiceBuffer {
    services = []Service
    services = append[services, nil]
    return &ServiceBuffer {
        services,
    }
}

func (sb *ServiceBuffer) Set(service *Service) {
    ss.head = service
}

func (sb *ServiceBuffer) Del(service *Service) {
    ss.head = nil
}

func UserCurrentService() func(action string, service *Service) {
    bufferService := CreateServiceBuffer()
    return func(action string, service *Service) error {
        switch action {
        case "set":
            bufferService.Set(service)
            fmt.Println("setting")
            return nil
        case "del":
            bufferService.Del()
            fmt.Println("deleting")
            return nil
        default:
            fmt.Println("unrecognized command")
        }
    }
}

type Request struct {
    time time.Time

    userid uint32
    driverid uint32
}

type Service struct {
    duration time.Time

    initial time.Time
    finish time.Time
}

func (s Service) Start() *Service {
    s.initial:= time.Now()
    s.finish := nil
    s.duration = nil

    services := UserCurrentService()
    services("set", s)

    return &s
}

func (s Service) Finish(service Service) {
    services := UserCurrentService()
    services("del", service)
}


func CreateRequest() *Request {
    //the user starts the request
    time := time.Now()
    return &Request {
        time,
        userid,
        driverid,//the driver's that the user selected
    }
}

func (r Request) Send() {
    //send the request to the driver
    //wait for the driver to accept or not
}
