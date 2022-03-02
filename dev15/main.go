//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
//Приведите корректный пример реализации.
//
//var justString string
//func someFunc() {
//  v := createHugeString(1 << 10)
//  justString = v[:100]
//}
//
//func main() {
//  someFunc()
//}

//1)Напрягает то, что переменная justString глобальная, возможны проблемы
//со сборщиком мусора
//2)К нам прилетает строка 2^10, а нам нужно только 100, поэтому копируем срез
//3) по хорошему строку ставим в []rune из за проблем с uft

//package main

//var justString string
//func someFunc() {
//	v := createHugeString(1 << 10)
//	c := make([]byte, 0, len(v))
//	copy(c, v)
//	justString = string(c)
//}
//
//func main() {
//	someFunc()
//}
