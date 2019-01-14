// test project main.go
package main

import (
	"bytes"
	"container/list"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/url"
	"os"
	"sort"
	"strings"
	"test/worker"
	"time"

	// "github.com/golang/groupcache"
	"tuyue_common/IdWorker"

	"github.com/json-iterator/go"
	"github.com/zheng-ji/goSnowFlake"

	"github.com/emirpasic/gods/lists/arraylist"
	lls "github.com/emirpasic/gods/stacks/linkedliststack"
	// "gopkg.in/gomail.v2"
)

type Hanlder struct {
	f   func()
	cid uint32
}

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

type SalaryCalculator interface {
	CalculateSalary() int
}

// 普通挖掘机员工
type Contract struct {
	empId    int
	basicpay int
}

// 有蓝翔技校证的员工
type Permanent struct {
	empId    int
	basicpay int
	jj       int
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.jj
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}

	fmt.Printf("总开支: $%d", expense)
}

func productor(channel chan<- string) {
	for {
		channel <- fmt.Sprintf("%v", rand.Float64())
		time.Sleep(time.Second * time.Duration(1))
	}
}

func customer(channel <-chan string) {
	for {
		message := <-channel // 此处会阻塞, 如果信道中没有数据的话
		fmt.Println(message)
	}
}

func ListTest() {
	l := list.New()
	l.PushBack("11111111111")
	l.PushBack("22222222222")
	l.PushBack("33333333333")
	l.PushBack("44444444444")

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func readXml() {
	file, err := os.Open("./configs/xml.xml")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	v := Recurlyservers{}

	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(v)
}

func idWorkerTest() {
	{
		tyIdWorker.SharedIdworkerInstance().InitIdWorker(1000)
		go func() {
			i := 0
			for i < 1000 {
				fmt.Println(tyIdWorker.SharedIdworkerInstance().NextId())
				time.Sleep(time.Duration(10) * time.Millisecond)
			}
		}()
	}

	{
		iw, err := goSnowFlake.NewIdWorker(1000)
		if err != nil {
			return
		}

		go func() {
			i := 0
			for i < 1000 {
				id, err := iw.NextId()
				if err != nil {
					continue
				}
				fmt.Println(id)
				time.Sleep(time.Duration(1) * time.Millisecond)
			}
		}()
	}
}

func GodsTest() {
	{
		list := arraylist.New()
		list.Add(1)
		list.Add("1111")
		fmt.Println(list.String())
	}
	{
		stack := lls.New()
		stack.Push(1)
		stack.Push(2)
		stack.Push("11111111111")
		fmt.Println(stack.Peek())
		fmt.Println(stack.Values())
		js, err := stack.ToJSON()
		fmt.Println(string(js), err)
	}

}

func main() {

	// {
	// 	s := ListServer()
	// 	s.Add("1111111111111111")
	// 	s.Add("1111111111111111")
	// 	s.Add("1111111111111111")
	// 	s.Add("1111111111111111")
	// 	s.Add("1111111111111111")
	// }

	// {
	// 	idWorkerTest()
	// }

	{
		ss := list.New()
		var mapHandlers map[string]*list.List
		mapHandlers = make(map[string]*list.List)
		l := list.New()
		l.PushFront(Hanlder{cid: 1, f: func() {}})
		l.PushFront(Hanlder{cid: 2, f: func() {}})
		l.PushFront(Hanlder{cid: 3, f: func() {}})
		l.PushFront(Hanlder{cid: 4, f: func() {}})
		l.PushFront(Hanlder{cid: 5, f: func() {}})
		mapHandlers["1"] = l
		ls := mapHandlers["1"]

		for v := ls.Front(); v != nil; v = v.Next() {
			fmt.Println(v)
		}
		fmt.Println(ss)
	}

	{
		GodsTest()
	}

	{
		readXml()
	}

	{
		ListTest()
	}

	{
		pemp1 := Permanent{
			1, 3000, 3000,
		}
		pemp2 := Permanent{
			2, 3000, 3000,
		}
		cemp1 := Contract{
			3, 3000,
		}
		cemp2 := Contract{
			4, 3000,
		}

		employees := []SalaryCalculator{pemp1, pemp2, cemp1, cemp2}
		totalExpense(employees)
	}

	{
		var buffer bytes.Buffer
		for i := 0; i < 10; i++ {
			buffer.WriteString("hello")
			buffer.WriteString("\r\n")
		}
		fmt.Println(buffer.String())
	}

	{
		var ar []int
		ar = append(ar, 1)
		ar = append(ar, 1)
		ar = append(ar, 1)
		ar = append(ar, 1)

		for i := 0; i < len(ar); i++ {
			fmt.Printf("%d\n", ar[i])
		}

		for k, v := range ar {
			fmt.Printf("%d: %d\n", k, v)
		}
	}

	{
		var ProgramingLanguage = map[string]int{
			"Java":              0,
			"C":                 1,
			"C++":               2,
			"Python":            3,
			"C#":                4,
			"PHP":               5,
			"JavaScript":        6,
			"Visual Basic.NET":  7,
			"Perl":              8,
			"Assembly language": 9,
			"Ruby":              10,
		}
		var SortString []string
		for k := range ProgramingLanguage {
			SortString = append(SortString, k)
		}
		sort.Strings(SortString) //会根据字母的顺序进行排序。

		for _, k := range SortString {
			fmt.Println("Key:", k, "Value:", ProgramingLanguage[k])
		}
	}

	{
		var m map[string]string
		m = make(map[string]string)
		m["name"] = "levi"
		m["age"] = "30"
		m["address"] = "深圳市"
		m["sex"] = "男"
		fmt.Printf("%+v", m)
		delete(m, "sex")
		fmt.Println(m)

		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		if data, err := json.Marshal(m); err == nil {
			fmt.Println(string(data))
		}
	}

	worker.BuildWorker().Create().Start()

	ssss := "11111,1111,11111111,111111111,11,111111,1111,11111111,1111,11111"
	ars := strings.Split(ssss, ",")
	fmt.Println(ars)

	fmt.Println(strings.ToUpper("aaaaaaaaaaaaaaaaaaa"))
	fmt.Println(strings.ToLower("SDFDSFWETFGDSFG"))

	{
		vals, err := url.ParseQuery("a=1111&b=2222&c=3333")
		if err != nil {

		}
		for k, v := range vals {
			for k1, v1 := range v {
				fmt.Println(k, k1, v1)
			}
		}
		fmt.Println(vals)
	}

	for {
		select {
		case <-time.After(time.Second * time.Duration(1)):
			fmt.Println("time.After", time.Now().Format("2006-01-02 15:04:05"))
			// default:
			// 	fmt.Println("default", time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}
