//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
//
//Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

package main

import "fmt"

func main() {
	temperatureArray := []float64{-39.4, -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	res := groupBy(temperatureArray, steps(10))
	fmt.Println(res)
}

func steps(step int) func(float64) int {
	return func(v float64) int {
		return int(v) / step * step // находим ключи с шагом 10
	}
}

func groupBy(slice []float64, step func(float64) int) map[int][]float64 {
	res := make(map[int][]float64)

	for _, v := range slice {
		key := step(v)
		res[key] = append(res[key], v)
	}
	return res
}
