package main

import (
	"errors"
	"fmt"
	"log"
)

// #include <stdio.h>

// int add(int i, int j){
// 	return i +j;
// }

func main() {
	var data = 10
	var err = errors.New("hello world")
	q, err := a()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(q)
	}
	fmt.Println(data, err)
}
func a() (int, error) {
	return 1, nil
}
