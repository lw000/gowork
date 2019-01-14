// pb_test project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"time"

	"pb_test/test"
	"pb_test/ty"

	"github.com/golang/protobuf/proto"
)

func Test() {
	p1 := &test.Person{
		Id:   1,
		Name: "小米",
		Phones: []*test.Phone{
			{test.PhoneType_HOME, "13632767233"},
			{test.PhoneType_WORK, "13632767233"},
		},
	}

	p2 := &test.Person{
		Id:   1,
		Name: "华为",
		Phones: []*test.Phone{
			{test.PhoneType_HOME, "13632767233"},
			{test.PhoneType_WORK, "13632767233"},
		},
	}

	book := &test.ContactBook{}
	book.Persons = append(book.Persons, p1)
	book.Persons = append(book.Persons, p2)

	data, err := proto.Marshal(book)
	if err != nil {

	}

	ioutil.WriteFile("./log.txt", data, os.ModePerm)

	var newbook test.ContactBook
	if err := proto.Unmarshal(data, &newbook); err == nil {
		// fmt.Println(newbook)
		for k, v := range newbook.Persons {
			fmt.Println(k, v.GetName(), v.GetId())
			for k1, v1 := range v.GetPhones() {
				fmt.Println(k1, v1.GetNumber())
			}
		}
	}
}

func Test1() {
	r := &ty.ReqEcho{
		Id:   1,
		Data: "11111111111111",
	}

	if r != nil {

	}

	data, err := proto.Marshal(r)
	if err != nil {

	}
	fmt.Println(data)

	var p1 ty.ReqEcho

	if err := proto.Unmarshal(data, &p1); err == nil {
		fmt.Println(p1)
	}
}

const (
	UINT32_MAX = ^uint32(0)
)

type Struct1 struct {
	mid int
	sid int
	cid int
}

func main() {
	Test()
	Test1()
	fmt.Printf("%d\n", UINT32_MAX)

	var s map[Struct1]string
	s = make(map[Struct1]string)
	s[Struct1{0, 0, 0}] = "111111111"
	s[Struct1{1, 0, 0}] = "222222222"
	s[Struct1{2, 0, 0}] = "222222222"
	s[Struct1{3, 0, 0}] = "333333333"

	for k, v := range s {
		fmt.Println(k, v)
		k1 := Struct1{2, 0, 0}
		if k == k1 {
			fmt.Println("k = k1", k)
			delete(s, k1)
		}
	}
	s1 := Struct1{
		mid: 0,
		sid: 0,
		cid: 0,
	}

	s2 := Struct1{
		mid: 0,
		sid: 0,
		cid: 0,
	}

	if s1 == s2 {
		fmt.Println("s1 = s2")
	}

	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt, os.Kill)

	quitcount := 2
	quit := make(chan bool, quitcount)

	go func() {
		defer func() {
			log.Println("go 1 exit")
		}()

		for {
			select {
			case <-quit:
				return
			default:
				log.Printf("[go 1] %s", time.Now().Format("2006-01-02 15:04:05"))
				time.Sleep(time.Millisecond * time.Duration(16))
			}
		}
	}()

	go func() {
		defer func() {
			log.Println("go 2 exit")
		}()

		for {
			select {
			case <-quit:
				return
			default:
				log.Printf("[go 2] %s", time.Now().Format("2006-01-02 15:04:05"))
				time.Sleep(time.Millisecond * time.Duration(16))
			}
		}
	}()

	<-exit

	log.Println("Interrupt")

	for i := 0; i < 2; i++ {
		quit <- true
	}

	time.Sleep(time.Second * 1)

	log.Println("main exit")
}
