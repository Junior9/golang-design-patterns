package main

import (
	"fmt"
	"sync"
	"time"
)

type DataBase struct {
}

func (DataBase) createSingleConnection() {
	fmt.Println("Creating a single db...")
	time.Sleep(2 * time.Second)
	fmt.Println("Creating DB Done")
}

var db *DataBase
var lock sync.Mutex

func getDataBaseInstance() *DataBase {
	lock.Lock()
	defer lock.Unlock()

	if db == nil {
		fmt.Println("Creating database...")
		db = &DataBase{}
		db.createSingleConnection()

	} else {
		fmt.Println("DataBase already created...")
	}

	return db
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDataBaseInstance()
		}()
	}
	wg.Wait()
}
