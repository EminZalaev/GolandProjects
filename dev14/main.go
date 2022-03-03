//Разработать программу, которая в рантайме способна определить тип переменной:
//int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"reflect"
)

//Пакет Reflect реализует отражение во время выполнения, позволяя программе
//манипулировать объектами произвольного типа. Типичное использование состоит в том, чтобы взять
//значение со статическим интерфейсом типа {}
//и извлечь информацию о его динамическом типе, вызвав TypeOf, который возвращает Type.

func main() {
	valueStr := "string"
	valueInt := 10
	valueBool := true
	valueCh := make(chan int)

	fmt.Println(ReflectDeterminateType(valueStr))
	fmt.Println(ReflectDeterminateType(valueInt))
	fmt.Println(ReflectDeterminateType(valueBool))
	fmt.Println(ReflectDeterminateType(valueCh))
}

func ReflectDeterminateType(value interface{}) reflect.Type {
	return reflect.TypeOf(value)
}
