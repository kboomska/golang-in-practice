package main

import "fmt"

func main() {
	// Go статически типизирован.
	// Переменные явно объявляются, типы заранее известны компилятору.

	// var объявляет переменную, = инициализирует конкретным значением:

	var b bool = true
	fmt.Println(b) // true

	var s string = "hello"
	fmt.Println(s) // hello

	var i int = 42
	fmt.Println(i) // 42

	var f float64 = 12.34
	fmt.Println(f) // 12.34

	// Можно объявить сразу несколько:
	var one, two int = 1, 2
	fmt.Println(one, two) // 1 2

	// Если инициализировать переменную при объявлении,
	// тип можно и не указывать. Go сам догадается:
	var sunny = true
	fmt.Printf("%v is %T\n", sunny, sunny) // true is bool

	// fmt.Printf() подставляет переменные в строку по шаблону. %v — формат по
	// умолчанию, %T — тип переменной.

	// Если не инициализировать переменную при объявлении, она получит нулевое
	// значение (zero value). У каждого типа оно свое: int — 0, string — "",
	// bool — false.

	var num int
	var str string
	var ok bool

	fmt.Printf("%#v %#v %#v\n", num, str, ok) // 0 "" false

	// Оператор := объявляет и инициализирует переменную. var и тип можно
	// не указывать:

	food := "apple"   // var food string = "apple"
	fmt.Println(food) // apple

	// %v печатает представление значения «для человека», а %#v — синтаксически
	// корректное (с точки зрения Go) представление:
	fmt.Printf("%v is %T\n", food, food)  // apple is string
	fmt.Printf("%#v is %T\n", food, food) // "apple" is string
}
