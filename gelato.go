package main

import "fmt"

func main() {
	var Milk, Cream, Sugar, Dex, LMP, Base50 float32
	fmt.Println("Привет, этот бот поможет Вам сбалансировать рецепт ващего Gelato")
	fmt.Println("Давайте приступим")

	fmt.Println("Ввведите колличество молока")
	fmt.Scan(&Milk)
	fmt.Println("Ввведите колличество сливок")
	fmt.Scan(&Cream)
	fmt.Println("Ввведите колличество сахара")
	fmt.Scan(&Sugar)
	fmt.Println("Введите колличество декстрозы")
	fmt.Scan(&Dex)
	fmt.Println("Введите колличество СОМ")
	fmt.Scan(&LMP)
	fmt.Println("Введите колличество Базы 50")
	fmt.Scan(&Base50)

	var TotalW float32 //Рассчитывает общий вес нашей базы
	TotalW = Milk + Cream + Sugar + Dex + LMP + Base50
	fmt.Println("Общий вес вашей смеси =", TotalW, "г.")

	var MilkFat float32 //Запрашивает жирность молока
	fmt.Println("Введите жирность вашего молока")
	fmt.Scan(&MilkFat)

	var IFM float32 //Рассчитывает промежуточную жирность молока
	IFM = Milk * MilkFat / 100

	var CreamFat float32 //Запрашивает жирность сливок
	fmt.Println("Введите жирность ваших сливок")
	fmt.Scan(&CreamFat)

	var IFC float32 //Рассчитывает промежуточную жирность сливок
	IFC = Cream * CreamFat / 100

	var FatLMP, IFLMP float32 //Рассчитывает промежуточную жирность сливок
	FatLMP = 1.5
	IFLMP = LMP * FatLMP / 100

	var TotalFat float32 //Рассчитывает жирность нашей базы
	TotalFat = (IFM + IFC + IFLMP) / TotalW * 100

	var MilkProt, CarbMilk float32 //Вводим переменные белок молока и углеводы молока

	fmt.Println("Какое содержание белка в вашем молоке?")
	fmt.Scan(&MilkProt)

	fmt.Println("Какое содержание углеводов в вашем молоке?")
	fmt.Scan(&CarbMilk)

	var CreamProt, CarbCream float32 //Вводим переменные белок сливок и углеводы сливок

	fmt.Println("Какое содержание белка в ваших сливках?")
	fmt.Scan(&CreamProt)

	fmt.Println("Какое содержание углеводов в ваших сливках?")
	fmt.Scan(&CarbCream)

	var CarbSug, CarbDex float32 //Вводим переменные углеводы сахара и декстрозы

	CarbSug = 100
	CarbDex = 91.5

	var ProtLMP, CarbLMP float32 //Добавляем БЖУ СОМ

	ProtLMP = 31.8
	CarbLMP = 50

	//СЧИТАЕМ ПРОМЕЖУТОЧНЫЕ УГЛИ
	var InCarbMilk float32 //Рассчитывает промежуточные угли молока
	InCarbMilk = Milk * CarbMilk / 100

	var InCarbCream float32 //Рассчитывает промежуточные угли сливок
	InCarbCream = Cream * CarbCream / 100

	var InCarbLMP float32 //Рассчитывает промежуточные угли СОМ'а
	InCarbLMP = LMP * CarbLMP / 100

	var InCarbSug float32 //Рассчитывает промежуточные угли Сахара
	InCarbSug = Sugar * CarbSug / 100

	var InCarbDex float32 // Рассчитывает промежуточные угли Декстрозы
	InCarbDex = Dex * CarbDex / 100

	// СЧИТАЕМ ПРОМЕЖУТОЧНЫЙ БЕЛОК

	var InProtMilk float32 //Рассчитывает промежуточный белок молока
	InProtMilk = Milk * MilkProt / 100

	var InProtCream float32 //Рассчитывает промежуточный белок сливок
	InProtCream = Cream * CreamProt / 100

	var InProtLMP float32 //Рассчитывает промежуточный белок СОМ'а
	InProtLMP = LMP * ProtLMP / 100

	//СЧИТАЕМ ОБЩИЙ БЕЛОК БАЗЫ

	var TotalProt float32
	TotalProt = (InProtMilk + InProtCream + InProtLMP) / TotalW * 100

	//СЧИТАЕМ ОБЩИЕ УГЛИ БАЗЫ

	var TotalCarb float32
	TotalCarb = (InCarbMilk + InCarbCream + InCarbLMP + InCarbSug + InCarbDex) / TotalW * 100

	fmt.Println("Жирность вашей базы =", TotalFat, "%", "Содержание белка вашей базы =", TotalProt, "Содержание углеводов вашей базы =", TotalCarb)

}
