package entitas

import "fmt"

type Context interface {
	CreateEntity(cs ...Component) Entity // 创建entity
	Entities() []Entity                  // 获取pool创建的所有还在的entity
	Count() int                          // entity数量
	HasEntity(e Entity) bool             // 是否包含某个entity
	DestroyEntity(e Entity)              // 删除entity
	DestroyAllEntities()                 // -
	Group(m Matcher) Group               // 获取包含满足条件的所有entities的group. group其实就是一个增强版的entities list.
}

type pool struct {
	entityMinID      int
	// componentsLength ComponentType  // 没啥用
	entities         map[EntityID]Entity
	cache            []Entity
	matcher2group    map[MatcherHash]Group
	com2groups       map[ComponentType][]Group
	unused           []Entity
}

func NewContext(startIndex int) Context {
	return &pool{
		entityMinID: startIndex,
		// componentsLength: componentsLength,
		entities:         make(map[EntityID]Entity),
		matcher2group:    make(map[MatcherHash]Group),
		com2groups:       make(map[ComponentType][]Group),
		unused:           make([]Entity, 0),
	}
}

func (p *pool) CreateEntity(cs ...Component) Entity {
	e := p.getEntity()
	e.AddComponent(cs...)
	p.entities[e.ID()] = e
	p.cache = append(p.cache, e)
	for _, g := range p.matcher2group {
		g.HandleEntity(e)
	}
	return e
}

func (p *pool) Entities() []Entity {
	if p.cache == nil {
		entities := make([]Entity, len(p.entities))
		i := 0
		for _, e := range p.entities {
			entities[i] = e
		}
		p.cache = entities
	}
	return p.cache
}

func (p *pool) Count() int {
	return len(p.entities)
}

func (p *pool) HasEntity(e Entity) bool {
	if entity, ok := p.entities[e.ID()]; ok && entity == e {
		return true
	}
	return false
}

func (p *pool) DestroyEntity(e Entity) {
	if entity, ok := p.entities[e.ID()]; ok && entity == e {
		e.RemoveAllComponents()
		e.RemoveAllCallbacks()
		delete(p.entities, e.ID())
		p.cache = nil
		for _, g := range p.matcher2group {
			g.HandleEntity(e)
		}
		p.unused = append(p.unused, e)
		return
	}
	panic("unknown entity")
}

func (p *pool) DestroyAllEntities() {
	for _, e := range p.entities {
		e.RemoveAllComponents()
		e.RemoveAllCallbacks()
	}
	p.entities = make(map[EntityID]Entity)
	p.cache = nil
}

func (p *pool) Group(m Matcher) Group {
	if g, ok := p.matcher2group[m.Hash()]; ok {
		return g
	}

	g := NewGroup(m)
	for _, e := range p.entities {
		g.HandleEntity(e)
	}
	p.matcher2group[m.Hash()] = g

	for _, component := range m.ComponentTypes() {
		p.com2groups[component] = append(p.com2groups[component], g)
	}

	return g
}

func (p *pool) String() string {
	return fmt.Sprintf("Context(%v)", p.Entities())
}

func (p *pool) componentAddedCallback(e Entity, c Component) {
	p.forMatchingGroups(e, c, func(g Group) {
		g.HandleEntity(e)
	})
}

func (p *pool) componentReplacedCallback(e Entity, c Component) {
	p.forMatchingGroups(e, c, func(g Group) {
		g.UpdateEntity(e)
	})
}

func (p *pool) componentWillBeRemovedCallback(e Entity, c Component) {
	p.forMatchingGroups(e, c, func(g Group) {
		delete(e.(*entity).components, c.Type())
		if !g.Matches(e) {
			e.(*entity).components[c.Type()] = c
			g.WillRemoveEntity(e)
		}
	})
}

func (p *pool) componentRemovedCallback(e Entity, c Component) {
	p.forMatchingGroups(e, c, func(g Group) {
		g.HandleEntity(e)
	})
}

func (p *pool) getEntity() Entity {
	var e Entity
	if len(p.unused) > 0 {
		e = p.unused[0]
		p.unused = p.unused[1:]
	} else {
		e = NewEntity(p.entityMinID)
		p.entityMinID++
	}
	e.AddCallback(ComponentAdded, p.componentAddedCallback)
	e.AddCallback(ComponentReplaced, p.componentReplacedCallback)
	e.AddCallback(ComponentWillBeRemoved, p.componentWillBeRemovedCallback)
	e.AddCallback(ComponentRemoved, p.componentRemovedCallback)
	return e
}

func (p *pool) forMatchingGroups(e Entity, c Component, f func(g Group)) {
	if p.HasEntity(e) {
		for _, g := range p.com2groups[c.Type()] {
			f(g)
		}
	}
}
