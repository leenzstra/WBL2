package pattern

/*
	Реализовать паттерн «фасад».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
	Паттерн Facade - структурный паттерн проектирования.
	Предоставляет интерфейс для взаимодействия с несколькими системами сразу через один интерфейс.

	+ упрощает работу со сложной системой, скрывает сложность
	+ можно расширять

	- уменьшение гибкости
	- может стать слишком большим

	Примеры:
	* ORM библиотеки
	* Logging
*/

// Описание части дома
type HousePart struct {
}

// Сервис строитель крыши
type RoofBuilder struct {
	roof *HousePart
}

func (rb *RoofBuilder) Build() {
	rb.roof = &HousePart{}
}

// Сервис строитель стен
type WallBuilder struct {
	walls []HousePart
}

func (wb *WallBuilder) Build(count int) {
	wb.walls = make([]HousePart, count)
}

// Сервис строитель фунтадмента
type FoundationBuilder struct {
	foundation *HousePart
}

func (fb *FoundationBuilder) Build() {
	fb.foundation = &HousePart{}
}

// Фасад - строитель дома
type HouseBuilder struct {
	rb *RoofBuilder
	wb *WallBuilder
	fb *FoundationBuilder
}

func NewHouseBuilder(rb *RoofBuilder, wb *WallBuilder, fb *FoundationBuilder) *HouseBuilder {
	return &HouseBuilder{
		rb, wb, fb,
	}
}

// Метод постройки дома
func (hb *HouseBuilder) BuildHouse(walls int) {
	hb.fb.Build()
	hb.wb.Build(walls)
	hb.rb.Build()
}

// Пример использования
func FacadeUsage() {
	fb := &FoundationBuilder{}
	wb := &WallBuilder{}
	rb := &RoofBuilder{}

	hb := NewHouseBuilder(rb, wb, fb)

	hb.BuildHouse(4)
}
