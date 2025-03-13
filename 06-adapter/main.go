package main

import (
	"fmt"
	"golang-practice/adapter"
)

func main() {
	// 建立 LegacyLogger 實例（舊系統的 Logger）
	legacyLogger := &adapter.LegacyLogger{}

	// 建立 LoggerAdapter，將 legacyLogger 包裝起來，使其符合 Logger 介面
	logger := &adapter.LoggerAdapter{Legacy: legacyLogger}

	// 客戶端使用 Logger 介面進行記錄
	result := logger.Log("Hello, adapter pattern!")
	fmt.Println(result)
}
