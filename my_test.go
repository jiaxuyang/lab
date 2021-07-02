package main

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"

	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

func TestChan(t *testing.T) {
	c := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			i, ok := <-c
			if !ok {
				break
			}
			fmt.Println("A:", i)
			i += 1
			if i > 100 {
				close(c)
				break
			}
			c <- i
		}
	}()

	go func() {
		defer wg.Done()
		c <- 1
		for {
			i, ok := <-c
			if !ok {
				break
			}
			fmt.Println("B:", i)
			i += 1
			if i > 100 {
				close(c)
				break
			}
			c <- i
		}
	}()

	wg.Wait()
}

func TestChan2(t *testing.T) {
	c := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		c <- 1
		for i := range c {
			fmt.Println("A:", i)
			i += 1
			if i > 100 {
				close(c)
				break
			}
			c <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := range c {
			fmt.Println("B:", i)
			i += 1
			if i > 100 {
				close(c)
				break
			}
			c <- i
		}
	}()

	wg.Wait()
}

func TestInterface(t *testing.T) {
	type MyStruct struct{}
	var face interface{}
	// 注释下面一行，ok将变为false
	face = (*MyStruct)(nil)
	v, ok := face.(*MyStruct)
	fmt.Println(v, ok)
}

func TestCreateNilObj(t *testing.T) {
	config := &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}
	conn, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        "test:test@tcp(10.227.19.217:3306)/test",
	}), config)
	if err != nil {
		t.Fatal(err)
	}
	type Treasure struct {
		Id        int64
		Amount    float64
		CreatedAt time.Time
	}
	data := []*Treasure{
		{Amount: 10000.00},
		nil,
	}
	if err := conn.Create(data).Error; err != nil {
		t.Fatal(err)
	}
}
