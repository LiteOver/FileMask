package main

import (
	"FileMask/Mask"
)

func main() {
	mProducer := Mask.NewProducer("/Users/bocman/GolandProjects/FileMask/file.txt")
	mPresenter := Mask.NewPresenter("/Users/bocman/GolandProjects/FileMask/file2.txt")
	workNewService := Mask.NewService(mProducer, mPresenter)
	workNewService.Run()
}
