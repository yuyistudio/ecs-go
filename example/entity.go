package main

import (
	"sort"
	"github.com/yuyistudio/ecs-go/entitas"
)

type Entity struct {
	id         int
	sortedComponents []entitas.Component
}

func NewEntity(id int) *Entity {
	return &Entity{
		id:         id,
	}
}

func (e *Entity) AddComponent(cs ...entitas.Component) error {
	e.sortedComponents = append(e.sortedComponents, cs...)
	sort.Slice(e.sortedComponents, func(i, j int) bool {
		return e.sortedComponents[i].Type() < e.sortedComponents[j].Type()
	})
	return nil
}

func (e *Entity) ID() int {
	return e.id
}

func (e *Entity) GetComponent(targetType entitas.ComponentType) entitas.Component {
	coms := e.sortedComponents
	startIdx := 0
	endIdx := len(coms) - 1
	for startIdx <= endIdx {
		midIdx := (startIdx + endIdx) / 2
		v := coms[midIdx].Type()
		// fmt.Printf("checking[(%v,%v)%v]=%v for %v\n", startIdx, endIdx, midIdx, v, targetType)
		if v < targetType {
			startIdx = midIdx + 1
		} else if v > targetType {
			endIdx = midIdx - 1
		} else {
			return coms[midIdx]
		}
	}
	return nil
}
