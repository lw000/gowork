// leveldb_test project main.go
package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

func LevelDBTest() {
	ldb, err := leveldb.OpenFile("./db", nil)
	if err != nil {

	}
	defer ldb.Close()

	if err := ldb.Put([]byte("key"), []byte("levi"), nil); err == nil {
		data, err := ldb.Get([]byte("key"), nil)
		if err != nil {

		}
		fmt.Println("data", string(data))
	}
}

func main() {
	LevelDBTest()
}
