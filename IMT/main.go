package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	fmt.Println("__ Калькулятор индекса массы тела __")
	for {

		uHeight, uWeight := getUserInput()
		IMT, err := calcIMT(uWeight, uHeight)
		if err != nil {
			//fmt.Println("Заданы неккоректные параметры")
			//continue
			panic("Заданы неккоретные параметры для расчёта")
		}
		outRes(IMT)
		isRepCalc := checkRepCalc()

		if !isRepCalc {
			break
		}

	}
}

func outRes(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела составляет: %.02f\n", IMT)
	fmt.Println(result)

	switch {
	case IMT < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case IMT < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case IMT < 25:
		fmt.Println("У вас нормальный вес")
	case IMT < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
}

func calcIMT(usW, usH float64) (float64, error) {
	if usW <= 0 || usH <= 0 {
		return 0, errors.New("NO_PARAMS_ERR")
	}
	const IMTPow = 2
	IMT := usW / math.Pow(usH/100, IMTPow)
	return IMT, nil
}

func getUserInput() (float64, float64) {
	var uHeight float64
	var uWeight float64
	fmt.Println("Введите свой рост в сантиметрах")
	fmt.Scanln(&uHeight)
	fmt.Println("Введите свой вес в килограммах")
	fmt.Scanln(&uWeight)
	return uHeight, uWeight
}

func checkRepCalc() bool {
	var uBtn int
	fmt.Print("1) Калькулятор IMT\n2) Выход из программы\nВыберите 1 или 2: ")
	fmt.Scanln(&uBtn)

	if uBtn == 2 {
		return false
	}
	return true

}
