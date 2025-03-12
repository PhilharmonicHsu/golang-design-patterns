package abstract_factory

import "fmt"

// -------------------
// 產品介面定義
// -------------------

// Button 定義了按鈕的行為
type Button interface {
	Paint() string
}

type Checkbox interface {
	Check() string
}

// -------------------
// 抽象工廠介面
// -------------------

// GUIFactory 定義了一個抽象工廠，用於創建一組相關的產品
type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// -------------------
// Windows 產品系列實作
// -------------------

type WinButton struct{}

func (b *WinButton) Paint() string {
	return "Windows Button painted"
}

type WinCheckbox struct{}

func (b *WinCheckbox) Check() string {
	return "Windows Checkbox checked"
}

// -------------------
// Mac 產品系列實作
// -------------------

type MacButton struct{}

func (b *MacButton) Paint() string {
	return "Mac Button painted"
}

type MacCheckbox struct{}

func (b *MacCheckbox) Check() string {
	return "Mac Checkbox checked"
}

// -------------------
// 具體工廠實作：Windows
// -------------------

type WinFactory struct{}

func (f *WinFactory) CreateButton() Button {
	return &WinButton{}
}

func (f *WinFactory) CreateCheckbox() Checkbox {
	return &WinCheckbox{}
}

// -------------------
// 具體工廠實作：Mac
// -------------------

type MacFactory struct{}

func (f *MacFactory) CreateButton() Button {
	return &MacButton{}
}

func (f *MacFactory) CreateCheckbox() Checkbox {
	return &MacCheckbox{}
}

// -------------------
// 客戶端函式：根據傳入的工廠創建產品家族
// -------------------

func Client(factory GUIFactory) {
	button := factory.CreateButton()
	checkbox := factory.CreateCheckbox()

	fmt.Println(button.Paint())
	fmt.Println(checkbox.Check())
}
