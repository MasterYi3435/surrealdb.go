package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/MasterYi3435/surrealdb.go"
	"github.com/MasterYi3435/surrealdb.go/pkg/conn/gorilla"
)

var Port = 8000
var LoginName = "root"
var Password = "root"

func main() {
	db, err := surrealdb.New("ws://localhost:"+strconv.FormatUint(uint64(Port), 10)+"/rpc", gorilla.Create())
	if err != nil {
		return
	}

	_, err = db.Signin(&surrealdb.Auth{
		Username: LoginName,
		Password: Password,
	})
	if err != nil {
		return
	}

	_, err = db.Use("researchRadar", "data")
	if err != nil {
		return
	}
	fmt.Println("connected")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			fmt.Println("Sleeping")
			time.Sleep(10 * time.Second)
		}
		wg.Done()
	}()
	wg.Wait()

}
