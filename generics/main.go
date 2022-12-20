package main

import (
	"fmt"
)

type RestStructA struct {
	c int
}

type RestStructB struct {
	d int
}

type MongoStructA struct {
	a int
}

func (a MongoStructA) convertTo() (c RestStructA) {
	ret := RestStructA{a.a + 10}
	return ret
}

type MongoStructB struct {
	b int
}

func (b MongoStructB) convertTo() (d RestStructB) {
	ret := RestStructB{b.b + 100}
	return ret
}

type ConvertableStruct[T any] interface {
	convertTo() T
}

func ConvertTo[M ConvertableStruct[R], R any](in []M) []R {
	restArray := make([]R, 0)

	for _, i := range in {
		restArray = append(restArray, i.convertTo())
	}

	return restArray
}

func main() {
	mongoArrayA := []MongoStructA{{11}, {12}, {13}}
	mongoArrayB := []MongoStructB{{51}, {52}, {53}}

	fmt.Println(ConvertTo[MongoStructA, RestStructA](mongoArrayA))
	fmt.Println(ConvertTo[MongoStructB, RestStructB](mongoArrayB))

}
