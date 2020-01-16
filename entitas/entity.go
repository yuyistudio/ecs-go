package entitas

import (
	"errors"
	"fmt"
	"sort"
)

var (
	ErrComponentExists       = errors.New("component exists")
	ErrComponentDoesNotExist = errors.New("component does not exist")
)

type EntityID uint

type Entity interface {
	AddComponent(cs ...Component) error
	ReplaceComponent(cs ...Component)
	WillRemoveComponent(ts ...ComponentType) error
	RemoveComponent(ts ...ComponentType) error
	RemoveAllComponents()
	RemoveAllCallbacks()
	AddCallback(ev ComponentEvent, cb ComponentCallback)
	HasCallbacks() bool

	ID() EntityID
	HasComponent(ts ...ComponentType) bool
	HasAnyComponent(ts ...ComponentType) bool
	Component(t ComponentType) (Component, error)
	GetComponent(t ComponentType) Component
	Components() []Component
	ComponentIndices() []ComponentType
}

type ComponentEvent uint

const (
	ComponentAdded ComponentEvent = iota
	ComponentReplaced
	ComponentWillBeRemoved
	ComponentRemoved
)

type ComponentCallback func(Entity, Component)

type entity struct {
	id         EntityID
	components map[ComponentType]Component
	callbacks  map[ComponentEvent][]ComponentCallback
}

func NewEntity(id int) Entity {
	return &entity{
		id:         EntityID(id),
		components: make(map[ComponentType]Component),
		callbacks:  make(map[ComponentEvent][]ComponentCallback),
	}
}

func (e *entity) AddComponent(cs ...Component) error {
	for _, c := range cs {
		if e.HasComponent(c.Type()) {
			return ErrComponentExists
		}
		e.components[c.Type()] = c
		e.callback(ComponentAdded, c)
	}
	return nil
}

func (e *entity) ReplaceComponent(cs ...Component) {
	for _, c := range cs {
		has := e.HasComponent(c.Type())
		e.components[c.Type()] = c
		if has {
			e.callback(ComponentReplaced, c)
		} else {
			e.callback(ComponentAdded, c)
		}
	}
}

func (e *entity) WillRemoveComponent(ts ...ComponentType) error {
	for _, t := range ts {
		c, err := e.Component(t)
		if err != nil {
			return err
		}
		e.callback(ComponentWillBeRemoved, c)
	}
	return nil
}

func (e *entity) RemoveComponent(ts ...ComponentType) error {
	for _, t := range ts {
		c, err := e.Component(t)
		if err != nil {
			return err
		}
		e.callback(ComponentWillBeRemoved, c)
		delete(e.components, t)
		e.callback(ComponentRemoved, c)
	}
	return nil
}

func (e *entity) RemoveAllComponents() {
	components := e.components

	for _, c := range components {
		e.callback(ComponentWillBeRemoved, c)
	}

	e.components = make(map[ComponentType]Component)

	for _, c := range components {
		e.callback(ComponentRemoved, c)
	}
}

func (e *entity) ID() EntityID {
	return e.id
}

func (e *entity) AddCallback(ev ComponentEvent, cb ComponentCallback) {
	cbs, ok := e.callbacks[ev]
	if !ok {
		cbs = make([]ComponentCallback, 0)
	}
	e.callbacks[ev] = append(cbs, cb)
}

func (e *entity) HasCallbacks() bool {
	return len(e.callbacks) > 0
}

func (e *entity) RemoveAllCallbacks() {
	e.callbacks = make(map[ComponentEvent][]ComponentCallback)
}

func (e *entity) HasComponent(ts ...ComponentType) bool {
	for _, t := range ts {
		if _, ok := e.components[t]; !ok {
			return false
		}
	}
	return true
}

func (e *entity) HasAnyComponent(ts ...ComponentType) bool {
	for _, t := range ts {
		if _, ok := e.components[t]; ok {
			return true
		}
	}
	return false
}

func (e *entity) Component(t ComponentType) (Component, error) {
	c, ok := e.components[t]
	if !ok {
		return nil, ErrComponentDoesNotExist
	}
	return c, nil
}

func (e *entity) GetComponent(t ComponentType) Component {
	c, ok := e.components[t]
	if !ok {
		panic(fmt.Sprintf("component type %v doesn't exist", t))
	}
	return c
}

func (e *entity) Components() []Component {
	components := make([]Component, len(e.components))
	i := 0
	for _, c := range e.components {
		components[i] = c
		i++
	}
	sort.Sort(ComponentsByType(components))
	return components
}

func (e *entity) ComponentIndices() []ComponentType {
	types := make([]ComponentType, len(e.components))
	i := 0
	for t := range e.components {
		types[i] = t
		i++
	}
	return types
}

func (e *entity) String() string {
	return fmt.Sprintf("Entity_%d(%v)", e.id, e.Components())
}

func (e *entity) callback(ev ComponentEvent, c Component) {
	if cbs, ok := e.callbacks[ev]; ok {
		for _, cb := range cbs {
			cb(e, c)
		}
	}
}
