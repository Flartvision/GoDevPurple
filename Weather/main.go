package main

import (
	"flag"
	"fmt"
	"weather/geo"
)

func main() {
	fmt.Println("Флаги")
	city := flag.String("city", "", "Город")
	//	format := flag.Int("format", 1, "Выбор формата вывода")
	flag.Parse()
	fmt.Println(*city)

	geoData, err := geo.GetLocation(*city)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(geoData)

}
