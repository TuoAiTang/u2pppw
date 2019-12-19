package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type User struct {
	Age int
	Name string
	Alive bool
}

func Sprint(x interface{}) string {
	type stringer interface {
		String() string
	}

	switch x := x.(type) {
	case int:
		fmt.Println("int")
		return strconv.Itoa(x)
	case bool:
		fmt.Println("bool")
		if x {
			return "true"
		}else{
			return "false"
		}
	case stringer :
		fmt.Println("stringer")
		return x.String()
	case string:
		fmt.Println("string")
		return x
	}

	v := reflect.ValueOf(x)
	var s string
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			println(reflect.TypeOf(v.Field(i).String()))
			println(v.Type().Field(i).Name)
			s += v.Type().Field(i).Name + Sprint(v.Field(i).Interface()) + "/"
		}
	}else {
		panic("不是结构体")
		return "???"
	}
	return s
}

func main() {
	user := User{
		Age:  12,
		Name:  "rick",
		Alive: true,
	}

	println(Sprint(user))
}

