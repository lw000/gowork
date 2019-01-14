// htp_test project main.go
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

// http://192.168.1.73:9090/servers/info/gameinfo/h5
// action=getdata&param=b4b578ff29464e884ed546350202ffcb143463347a6f9d8471e659e0a8659510876bf79fe1a57b59d8eca3ca14e8f55e099aa71f5d9120967d46712e3227e6217e81aaa1356ea2c81893cfa4b0e1135c4ed546350202ffcb143463347a6f9d8471e659e0a8659510876bf79fe1a57b5983f3e983a8d1fc1777aea8b4f6fa339f6dbe006a7584c417190eac4e0d1b749934f004efb95fc7f8fdd7f77672e894fa0ea3a129ddc29d5424ca0a726a8c1f4071fb5bc4aa4f204a13be944c54b44bf4

// http://yll.qmzlcy.com/api/h5
// action=getdata&param=b4b578ff29464e884ed546350202ffcb143463347a6f9d8471e659e0a8659510876bf79fe1a57b59d8eca3ca14e8f55e099aa71f5d9120967d46712e3227e6217e81aaa1356ea2c81893cfa4b0e1135c4ed546350202ffcb143463347a6f9d8471e659e0a8659510876bf79fe1a57b5983f3e983a8d1fc1777aea8b4f6fa339f6dbe006a7584c417190eac4e0d1b749934f004efb95fc7f8fdd7f77672e894fa0ea3a129ddc29d5424ca0a726a8c1f4071fb5bc4aa4f204a13be944c54b44bf4

const (
	HOST = "http://192.168.1.73:9090/servers/info/gameinfo/h5"
	// HOST = "http://47.97.103.250:9090/servers/info/gameinfo/h5"
	DATA = "action=getdata&param=b4b578ff29464e884ed546350202ffcb143463347a6f9d8471e659e0a8659510876bf79fe1a57b59d8eca3ca14e8f55e099aa71f5d9120967d46712e3227e6217e81aaa1356ea2c81893cfa4b0e1135c4ed546350202ffcb143463347a6f9d8471e659e0a8659510876bf79fe1a57b5983f3e983a8d1fc1777aea8b4f6fa339f6dbe006a7584c417190eac4e0d1b749934f004efb95fc7f8fdd7f77672e894fa0ea3a129ddc29d5424ca0a726a8c1f4071fb5bc4aa4f204a13be944c54b44bf4"
	// HOST = "http://192.168.1.73:9090/servers/user/buildunionid"
	// DATA = ""
)

func httpPostForm(cout int) {
	for i := 0; i < cout; i++ {
		go func(i int, w *sync.WaitGroup) {
			w.Add(1)
			defer w.Done()
			for {
				resp, err := http.PostForm(HOST,
					url.Values{"action": {"getdata"},
						"param": {"b4b578ff29464e888d7fd09f9a835850fc062a8caaca4629205a774f01313ccb4ac5b03d7b78da90d8eca3ca14e8f55e099aa71f5d912096c9debf20022c74008f4a4d3d04950aa81893cfa4b0e1135c73ad4cb7e538a57a50bc0db40c0866afbdc42c5b893b2254f04a567c78fdd32e329e576aec99ada146910d9aee34fc7deb68cf24977b501807c0a9d33b506a159fe22b79dffc6de5440d063c50015ff09c1551f28bc51aea"}})
				if err != nil {
					fmt.Println(i, ":", err.Error())
					break
				}
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(i, ":", err.Error())
					break
				}
				if len(body) > 0 {

				}

				fmt.Println(i, resp.StatusCode)
				break
			}
		}(i, &wg)

		time.Sleep(time.Millisecond * 1)
	}
}

func httpPost(cout int) {
	for i := 0; i < cout; i++ {
		go func(i int, w *sync.WaitGroup) {
			w.Add(1)
			defer w.Done()
			for {
				resp, err := http.Post(HOST, "application/x-www-form-urlencoded", strings.NewReader(DATA))
				if err != nil {
					fmt.Println(i, ":", err.Error())
					break
				}
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(i, ":", err.Error())
					break
				}

				if len(body) > 0 {

				}
				fmt.Println(i, resp.StatusCode)
				break
			}
		}(i, &wg)
		time.Sleep(time.Millisecond * 1)
	}
}

func httpGet(cout int) {
	for i := 0; i < cout; i++ {
		go func(i int, w *sync.WaitGroup) {
			w.Add(1)
			defer w.Done()
			for {
				// resp, err := http.Get(HOST + "?" + DATA)
				resp, err := http.Get("http://127.0.0.1:9092/newid/1023")
				if err != nil {
					fmt.Println(i, ":", err.Error())
					break
				}
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(i, ":", err.Error())
					break
				}
				if len(body) > 0 {
					fmt.Println(i, string(body))
				}
				fmt.Println(i, resp.StatusCode)
				break
			}
		}(i, &wg)
		time.Sleep(time.Millisecond * 2)
	}
}

func httpGet1(cout int) {
	for i := 0; i < cout; i++ {
		go func(i int, w *sync.WaitGroup) {
			w.Add(1)
			defer w.Done()
			for {
				url := HOST + "?" + DATA
				request, err := http.NewRequest("GET", url, nil)
				if err != nil {
					fmt.Println(i, ":", err.Error())
					break
				}
				client := &http.Client{}
				resp, err := client.Do(request)
				if err != nil {
					fmt.Println(i, ":", err.Error())
					break
				}

				_, err = io.Copy(os.Stdout, resp.Body)
				if err != nil {
					fmt.Println(i, ":", err.Error())
					break
				}

				fmt.Println(i, resp.StatusCode)
				break
			}
		}(i, &wg)

		time.Sleep(time.Millisecond * 1)
	}
}

func main() {
	start := time.Now().UnixNano()
	httpGet(5000)
	wg.Wait()
	end := time.Now().UnixNano()
	fmt.Println("tms:", (end-start)/(int64)(time.Millisecond*time.Duration(1)))
}
