package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
	Паттерн Factory method - порождающий паттерн.
	Предоставляется интерфейс для создания объектов, позволяя его наследникам определять конкретный тип создаваемого объекта.

	+ универсальный код создания объектов
	+ нет привязки к конкретным типам

	- создание фабрики для каждого нового объекта

	Примеры:
	* loggers

*/


type FPlayer struct {
	Health int
}

// Интерфейс создаваемых объектов
type Mob interface {
	Attack(*FPlayer)
}

// Реализация Mob
type Skeleton struct {
	attackPower int
}

func (s *Skeleton) Attack(player *FPlayer) {
	player.Health -= s.attackPower
	fmt.Printf("Skeleton атаковал игрока: здоровье = %d\n", player.Health)
}

// Реализация Mob
type Zombie struct {
	attackPower int
}

func (z *Zombie) Attack(player *FPlayer) {
	player.Health -= z.attackPower
	fmt.Printf("Zombie атаковал игрока: здоровье = %d\n", player.Health)
}

// Интерфейс фабрики
type MobFactory interface {
	CreateMob() Mob
}

// фабрика скелетов
type SkeletonFactory struct {
	baseAttackPower int
}

// Создание конкретного типа Skeleton
func (sf *SkeletonFactory) CreateMob() Mob {
	return &Skeleton{sf.baseAttackPower}
}

// фабрика зомби
type ZombieFactory struct {
	baseAttackPower int
}

// Создание конкретного типа Zombie
func (zf *ZombieFactory) CreateMob() Mob {
	return &Zombie{zf.baseAttackPower}
}

// Пример использования
func FactoryUsage() {
	zombieFactory := ZombieFactory{10}
	skeletonFactory := SkeletonFactory{5}

	zombie := zombieFactory.CreateMob()     // Mob
	skeleton := skeletonFactory.CreateMob() // Mob

	player := &FPlayer{100}

	zombie.Attack(player)
	skeleton.Attack(player)
}
