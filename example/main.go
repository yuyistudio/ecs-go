package main

import (
	"github.com/yuyistudio/ecs-go/entitas"
	"fmt"
	"time"
	"math/rand"
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


func TestWorld(){
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

func testGetCom() {
	fmt.Printf("com count %d\n", ComTypeCount)
	context := entitas.NewContext(888)
	e1 := context.CreateEntity(&PosCom{x:1, y:2})
	/*
	*/
	e1.AddComponent(&Com0{x:0})
	e1.AddComponent(&Com1{x:1})
	e1.AddComponent(&Com2{x:2})
	e1.AddComponent(&Com3{x:3})
	e1.AddComponent(&Com4{x:4})
	e1.AddComponent(&Com5{x:5})
	e1.AddComponent(&Com6{x:6})
	e1.AddComponent(&Com7{x:7})
	e1.AddComponent(&Com8{x:8})
	e1.AddComponent(&Com9{x:9})
	e1.AddComponent(&Com10{x:10})
	e1.AddComponent(&Com11{x:11})
	e1.AddComponent(&Com12{x:12})
	e1.AddComponent(&Com13{x:13})
	e1.AddComponent(&Com14{x:14})
	e1.AddComponent(&Com15{x:15})
	e1.AddComponent(&Com16{x:16})
	e1.AddComponent(&Com17{x:17})
	e1.AddComponent(&Com18{x:18})
	e1.AddComponent(&Com19{x:19})
	e1.AddComponent(&Com20{x:20})
	e1.AddComponent(&Com21{x:21})
	e1.AddComponent(&Com22{x:22})
	e1.AddComponent(&Com23{x:23})
	e1.AddComponent(&Com24{x:24})
	e1.AddComponent(&Com25{x:25})
	e1.AddComponent(&Com26{x:26})
	e1.AddComponent(&Com27{x:27})
	e1.AddComponent(&Com28{x:28})
	e1.AddComponent(&Com29{x:29})
	e1.AddComponent(&Com30{x:30})
	e1.AddComponent(&Com31{x:31})
	e1.AddComponent(&Com32{x:32})
	e1.AddComponent(&Com33{x:33})
	e1.AddComponent(&Com34{x:34})
	e1.AddComponent(&Com35{x:35})
	e1.AddComponent(&Com36{x:36})
	e1.AddComponent(&Com37{x:37})
	e1.AddComponent(&Com38{x:38})
	e1.AddComponent(&Com39{x:39})
	e1.AddComponent(&Com40{x:40})
	e1.AddComponent(&Com41{x:41})
	e1.AddComponent(&Com42{x:42})
	e1.AddComponent(&Com43{x:43})
	e1.AddComponent(&Com44{x:44})
	e1.AddComponent(&Com45{x:45})
	e1.AddComponent(&Com46{x:46})
	e1.AddComponent(&Com47{x:47})
	e1.AddComponent(&Com48{x:48})
	e1.AddComponent(&Com49{x:49})
	e1.AddComponent(&Com50{x:50})
	e1.AddComponent(&Com51{x:51})
	e1.AddComponent(&Com52{x:52})
	e1.AddComponent(&Com53{x:53})
	e1.AddComponent(&Com54{x:54})
	e1.AddComponent(&Com55{x:55})
	e1.AddComponent(&Com56{x:56})
	e1.AddComponent(&Com57{x:57})
	e1.AddComponent(&Com58{x:58})
	e1.AddComponent(&Com59{x:59})
	e1.AddComponent(&Com60{x:60})
	e1.AddComponent(&Com61{x:61})
	e1.AddComponent(&Com62{x:62})
	e1.AddComponent(&Com63{x:63})
	e1.AddComponent(&Com64{x:64})
	e1.AddComponent(&Com65{x:65})
	e1.AddComponent(&Com66{x:66})
	e1.AddComponent(&Com67{x:67})
	e1.AddComponent(&Com68{x:68})
	e1.AddComponent(&Com69{x:69})
	e1.AddComponent(&Com70{x:70})
	e1.AddComponent(&Com71{x:71})
	e1.AddComponent(&Com72{x:72})
	e1.AddComponent(&Com73{x:73})
	e1.AddComponent(&Com74{x:74})
	e1.AddComponent(&Com75{x:75})
	e1.AddComponent(&Com76{x:76})
	e1.AddComponent(&Com77{x:77})
	e1.AddComponent(&Com78{x:78})
	e1.AddComponent(&Com79{x:79})
	e1.AddComponent(&Com80{x:80})
	e1.AddComponent(&Com81{x:81})
	e1.AddComponent(&Com82{x:82})
	e1.AddComponent(&Com83{x:83})
	e1.AddComponent(&Com84{x:84})
	e1.AddComponent(&Com85{x:85})
	e1.AddComponent(&Com86{x:86})
	e1.AddComponent(&Com87{x:87})
	e1.AddComponent(&Com88{x:88})
	e1.AddComponent(&Com89{x:89})
	e1.AddComponent(&Com90{x:90})
	e1.AddComponent(&Com91{x:91})
	e1.AddComponent(&Com92{x:92})
	e1.AddComponent(&Com93{x:93})
	e1.AddComponent(&Com94{x:94})
	e1.AddComponent(&Com95{x:95})
	e1.AddComponent(&Com96{x:96})
	e1.AddComponent(&Com97{x:97})
	e1.AddComponent(&Com98{x:98})
	e1.AddComponent(&Com99{x:99})
	e1.RebuildComponentIndex()


	e2 := context.CreateEntity(&PosCom{x:1, y:2})
	e2.AddComponent(&Com0{x:0})
	e2.AddComponent(&Com1{x:1})
	e2.AddComponent(&Com2{x:2})
	e2.AddComponent(&Com3{x:3})
	e2.AddComponent(&Com4{x:4})
	e2.AddComponent(&Com5{x:5})
	e2.AddComponent(&Com6{x:6})
	e2.AddComponent(&Com7{x:7})

	e3 := context.CreateEntity(&PosCom{x:1, y:2})
	e3.AddComponent(&Com0{x:0})

	e := e3

	if e.BinarySearchComponent(ComType_com99) != e.GetComponent(ComType_com99) {
		panic("! com-99")
	}
	if e.LinearSearchComponent(ComType_com44) != e.GetComponent(ComType_com44) {
		panic("! com-44")
	}

	const loopCount = 40000000
	const seed = 10324329
	{
		rand.Seed(seed)
		st := time.Now()
		for i := 0; i < loopCount; i++ {
			e.DictGetComponent(entitas.ComponentType(rand.Intn(int(ComTypeCount))))
		}
		et := time.Now()
		fmt.Printf("hash-index: %v\n", et.Sub(st))
	}
	{
		rand.Seed(seed)
		st := time.Now()
		for i := 0; i < loopCount; i++ {
			e.BinarySearchComponent(entitas.ComponentType(rand.Intn(int(ComTypeCount))))
		}
		et := time.Now()
		fmt.Printf("binary-search: %v\n", et.Sub(st))
	}
	{
		rand.Seed(seed)
		st := time.Now()
		for i := 0; i < loopCount; i++ {
			e.GetComponent(entitas.ComponentType(rand.Intn(int(ComTypeCount))))
		}
		et := time.Now()
		fmt.Printf("get-com: %v\n", et.Sub(st))
	}
}

func main() {
	testGetCom()
}
