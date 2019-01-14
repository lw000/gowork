package main

import (
	"fmt"
	"sync"
)

type Platform struct {
	Pid   int
	Name  string
	Rooms []*Room
	m     sync.Mutex
}

func NewPlatform(pid int, name string) *Platform {
	return &Platform{
		Pid:  pid,
		Name: name,
	}
}

func (p *Platform) CreateRoom() {
	for i := 0; i < 100; i++ {
		p.Rooms = append(p.Rooms, NewRoom(i+1, fmt.Sprintf("room_%d", i+1)))
	}
}

func (p *Platform) AddRoom(r *Room) {

}

func (p *Platform) RemoveRoom(rid int) {

}

func (p *Platform) DestroyRoom() {

}
