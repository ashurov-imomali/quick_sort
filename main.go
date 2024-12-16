package main

import (
	"errors"
	"log"
	"math/rand"
	"strconv"
)

func quick(a []int) []int {
	if len(a) < 2 {
		return a
	}

	temp := rand.Int() % len(a)
	r := []int{a[temp]}
	l := []int{}
	for i, v := range a {
		if i == temp {
			continue
		}
		if a[temp] < v {
			r = append(r, v)
			continue
		}
		l = append(l, v)
	}
	return append(quick(l), quick(r)...)
}

func generate(n, max int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = rand.Int() % max
	}
	return a
}

func check(a []int) (bool, error) {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false, errors.New("error at index :" + strconv.Itoa(i))
		}
	}
	return true, nil
}

func main() {
	a := generate(900000, 100)
	log.Println(quick(a))
	log.Println(check(quick(a)))
}
