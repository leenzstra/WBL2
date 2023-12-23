package pattern

/*
	Реализовать паттерн «строитель».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
	Паттерн Builder - порождающий паттерн.
	Применяется, когда необходимо создать сложный объект поэтапно.

	+ гибкость в создании объектов
	+ изолирование процесса создания от конечного объекта

	- растет сложность кода

	Примеры:
	* SQL библиотеки построения запросов
	* StringBuilder
*/

// Структура Игрок, которую билдим
type Player struct {
	weapon string
	armor  string
	boots  string
	gloves string
	helmet string
}

// Строитель Player
type PlayerBuilder struct {
	player Player
}

func NewPlayerBuilder() *PlayerBuilder {
	return &PlayerBuilder{
		player: Player{},
	}
}

// Методы строителя, позволяющие снова использовать его
// С каждым вызовом метода меняется одно из полей Player
func (p *PlayerBuilder) SetWeapon(v string) *PlayerBuilder {
	p.player.weapon = v
	return p
}

func (p *PlayerBuilder) SetArmor(v string) *PlayerBuilder {
	p.player.armor = v
	return p
}

func (p *PlayerBuilder) SetBoots(v string) *PlayerBuilder {
	p.player.boots = v
	return p
}

func (p *PlayerBuilder) SetGloves(v string) *PlayerBuilder {
	p.player.gloves = v
	return p
}

func (p *PlayerBuilder) SetHelmet(v string) *PlayerBuilder {
	p.player.helmet = v
	return p
}

// Возвращаем конечный объект
func (p *PlayerBuilder) Build() Player {
	return p.player
}

// Пример использования
func BuilderUsage() {
	player := NewPlayerBuilder().
		SetWeapon("Weapon").
		SetArmor("Armor").
		SetBoots("Boots").
		SetGloves("Gloves").
		SetHelmet("Helmet").
		Build()

	_ = player
}
