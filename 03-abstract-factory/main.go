package main

import "03-abstract-factory/abstract_factory"

func main() {
	// 使用 Windows 產品系列
	winFactory := &abstract_factory.WinFactory{}
	abstract_factory.Client(winFactory)

	// 使用 Mac 產品系列
	macFactory := abstract_factory.MacFactory{}
	abstract_factory.Client(&macFactory)
}
