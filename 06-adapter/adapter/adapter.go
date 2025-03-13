package adapter

import "fmt"

/**
情境描述

假設你有一個舊系統提供的 Logger，它只能用 LegacyLog(message string) 方法來記錄訊息，並回傳一個字串。
但你的新應用程式期望使用一個標準化的 Logger 介面，該介面提供 Log(message string) 方法。
這時候，我們可以使用 Adapter 模式，讓舊 Logger 符合新 Logger 的介面。
*/

// LegacyLogger 是舊系統的 Logger，只有 LegacyLog 方法
type LegacyLogger struct{}

// LegacyLog 用於記錄訊息，返回格式化後的結果
func (l *LegacyLogger) LegacyLog(message string) string {
	return fmt.Sprintf("Legacy log: %s", message)
}

// Logger 是客戶端期望的新 Logger 介面
type Logger interface {
	Log(message string) string
}

// LoggerAdapter 是適配器，負責將 LegacyLogger 轉換為符合 Logger 介面的實作
type LoggerAdapter struct {
	Legacy *LegacyLogger
}

// Log 實作 Logger 介面的方法，內部轉呼 LegacyLog
func (a *LoggerAdapter) Log(message string) string {
	// 這裡你可以添加額外邏輯，例如格式轉換、錯誤處理等等
	return a.Legacy.LegacyLog(message)
}
