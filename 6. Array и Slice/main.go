package main

import "fmt"

// В цикле спросиь ввод транзакций: -10, 10, 40.5
// Добавлять каждую в массив транзакций
// Вывести в массив
// Вывести сумму баланса в консоль

var transactions = []float64{}

func main() {
	tr := make([]string, 0, 2)
	tr[3] = "10"
	fmt.Println(tr)

	var result float64
	transactions = getUserInp()
	fmt.Println(transactions)

	for _, v := range transactions {
		result += v
	}

	fmt.Println("Ваш баланс составляет: ", result)
	//	temp := transactions
	//	transactions = append(transactions, 100)
	//	fmt.Println(temp)
	//	fmt.Println(transactions)

}

func getUserInp() []float64 {
	for {
		var tr float64
		fmt.Println("Введите вашу транзацикю:")
		fmt.Println("Если хотите завершить ввод, нажмите N")
		fmt.Scanln(&tr)

		if tr == 0 {
			break
		}

		transactions = append(transactions, tr)
	}
	return transactions
}
