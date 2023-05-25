package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-test/common"
)

func TestCompare(t *testing.T) {
	res := common.CompareNumber(1, 2)
	fmt.Println(res)
}

func TestSampleFuncModifier(t *testing.T) {
	defer func() {
		fmt.Println("All done!")
	}()
	go func() {

		time.Sleep(2 * time.Second)
		fmt.Println("After 2 seconds")
	}()
	fmt.Println("Hello, World!")
	time.Sleep(3 * time.Second)
	fmt.Println("Bye, World!")
}

func TestSampleComposite(t *testing.T) {
	ch := make(chan int)

	go func() {
		fmt.Println("ready to send number to channel")
		time.Sleep(2 * time.Second)
		ch <- 3
	}()

	fmt.Println("wait some thing from channel")
	res := <-ch
	fmt.Println(res)
	fmt.Println("end")
}

func TestSampleComposite2(t *testing.T) {

	var myMap map[string]interface{} = make(map[string]interface{})

	myMap["name"] = "John"
	myMap["age"] = 30
	myMap["married"] = true

	fmt.Println(myMap)

	if value, ok := myMap["noExist"]; !ok {
		fmt.Println(value)
	}
}

func TestSampleControlFlow(t *testing.T) {

	var myMap map[string]interface{} = make(map[string]interface{})

	myMap["name"] = "John"
	myMap["age"] = 30
	myMap["married"] = true

	for key, value := range myMap {
		fmt.Println(key, value)
	}

	var mySlice []int = []int{1, 2, 3, 4, 5}
	for index, value := range mySlice {
		fmt.Println(index, value)
	}
}

func TestSampleControlFlow2(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Hello, World!"
	}()

	select {
	case num := <-ch1:
		fmt.Println("Received from ch1:", num)
	case msg := <-ch2:
		fmt.Println("Received from ch2:", msg)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout occurred!")
	}
}
