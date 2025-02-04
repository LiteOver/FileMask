package main

import (
	"FileMask/Mask"
	"fmt"
)

func main() {
	fmt.Println("Введите адрес файла, данные которого нужно считать")
	adr := ""
	fmt.Scanln(&adr)
	mProducer := Mask.NewProducer(adr)

	fmt.Println("Введите адрес файла для записи результата")
	fmt.Scanln(&adr)
	mPresenter := Mask.NewPresenter(adr)

	workNewService := Mask.NewService(mProducer, mPresenter)
	workNewService.Run()
}
