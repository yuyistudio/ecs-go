package main

import (
	"github.com/yuyistudio/ecs-go/entitas"
	"fmt"
	"time"
)

const (
	ComType_pos entitas.ComponentType = 1
	ComType_renderer entitas.ComponentType = 2
)

type PosCom struct {
	x float64
	y float64
}

func (p *PosCom) Type() entitas.ComponentType {
	return ComType_pos
}

type RendererCom struct {
	screen int64
}

func (p *RendererCom) Type() entitas.ComponentType {
	return ComType_renderer
}

type PosMatcher struct {
	hash     entitas.MatcherHash
}

func (b *PosMatcher) Hash() entitas.MatcherHash {
	return b.hash
}

func (b *PosMatcher) ComponentTypes() []entitas.ComponentType {
	return nil
}

func NewPosMatch() *PosMatcher {
	p := new(PosMatcher)
	p.hash = 123
	return p
}

func (m *PosMatcher) Matches(entity entitas.Entity) bool {
	return entity.HasComponent(ComType_pos)
}

func (m *PosMatcher) Equals(another entitas.Matcher) bool {
	return m == another
}

func (m *PosMatcher) String() string {
	return fmt.Sprintf("PosMatcher(%v)", m.Hash())
}

type System interface {
	OnInit()
	OnUpdate()
	OnCleanup()
}

type World struct {
	context entitas.Context
	systems []System
}

func (w *World) OnInit() {
	for _, system := range w.systems {
		system.OnInit()
	}
}

func (w *World) OnUpdate() {
	for _, system := range w.systems {
		system.OnUpdate()
	}
	for _, system := range w.systems {
		system.OnCleanup()
	}
}

func (w *World) AddSystem(system System) {
	w.systems = append(w.systems, system)
}

type MovementSystem struct {
	g entitas.Group
	context entitas.Context
}

func NewMovementSystem(context entitas.Context) System {
	m := new(MovementSystem)
	m.context = context
	return m
}
func (m *MovementSystem) OnInit() {
	posMatcher := NewPosMatch()
	m.g = m.context.Group(posMatcher)
}
func (m *MovementSystem) OnUpdate() {
	for _, entity := range m.g.Entities() {
		posCom := entity.GetComponent(ComType_pos).(*PosCom)
		posCom.y += 1
		fmt.Printf("entity[%d].pos.y = %v\n", entity.ID(), posCom.y)
	}
}

func (m *MovementSystem) OnCleanup() {
	fmt.Printf("cleaning frame data\n")
}


func main() {
	context := entitas.NewContext(888)
	context.CreateEntity(&PosCom{x:1, y:2})
	context.CreateEntity(&RendererCom{screen:888}, &PosCom{x:3, y:4})

	world := new(World)
	world.context = context

	world.AddSystem(NewMovementSystem(context))
	world.OnInit()

	for {
		// frame
		world.OnUpdate()
		time.Sleep(500 * time.Millisecond)
	}
}
