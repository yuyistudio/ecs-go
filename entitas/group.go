package entitas

type Group interface {
	Entities() []Entity                        // 获取所有的组件。
	HandleEntity(e Entity)                     // 将组件添加到或者移除出当前group（判断标准是group.matcher）
	UpdateEntity(e Entity)                     // 触发事件，触发一下Remove再触发一下Add事件
	WillRemoveEntity(e Entity)                 // 触发WillRemove事件
	Matches(e Entity) bool                     // 判断是不是应该包含参数entity（判断标准是group.matcher）
	ContainsEntity(e Entity) bool              // 判断是不是已经包含entity
	AddCallback(e GroupEvent, c GroupCallback) // -
}

type GroupEvent uint

const (
	EntityAdded GroupEvent = iota
	EntityWillBeRemoved
	EntityRemoved
)

type GroupCallback func(Group, Entity)

type group struct {
	entities         map[EntityID]Entity
	cache            []Entity
	cacheInvalidated bool
	matcher          Matcher
	callbacks        map[GroupEvent][]GroupCallback
}

func NewGroup(matcher Matcher) Group {
	return &group{
		entities:         make(map[EntityID]Entity),
		cache:            make([]Entity, 0),
		cacheInvalidated: false,
		matcher:          matcher,
		callbacks:        make(map[GroupEvent][]GroupCallback),
	}
}

func (g *group) Entities() []Entity {
	if g.cacheInvalidated {
		cache := make([]Entity, len(g.entities))
		i := 0
		for _, e := range g.entities {
			cache[i] = e
			i++
		}
		g.cache = cache
		g.cacheInvalidated = false
	}
	return g.cache
}

func (g *group) HandleEntity(e Entity) {
	if g.matcher.Matches(e) {
		g.addEntity(e)
	} else {
		g.removeEntity(e)
	}
}

func (g *group) UpdateEntity(e Entity) {
	if _, ok := g.entities[e.ID()]; ok {
		g.callback(EntityRemoved, e)
		g.callback(EntityAdded, e)
	}
}

func (g *group) WillRemoveEntity(e Entity) {
	if _, ok := g.entities[e.ID()]; ok {
		g.callback(EntityWillBeRemoved, e)
	}
}

func (g *group) Matches(e Entity) bool {
	return g.matcher.Matches(e)
}

func (g *group) ContainsEntity(e Entity) bool {
	if _, ok := g.entities[e.ID()]; ok {
		return true
	}
	return false
}

func (g *group) AddCallback(ev GroupEvent, c GroupCallback) {
	cs, ok := g.callbacks[ev]
	if !ok {
		cs = make([]GroupCallback, 0)
	}
	g.callbacks[ev] = append(cs, c)
}

func (g *group) addEntity(e Entity) {
	if _, ok := g.entities[e.ID()]; !ok {
		g.entities[e.ID()] = e
		if g.cacheInvalidated == false {
			g.cache = append(g.cache, e)
		}
		g.callback(EntityAdded, e)
	}
}

func (g *group) removeEntity(e Entity) {
	if _, ok := g.entities[e.ID()]; ok {
		delete(g.entities, e.ID())
		if i := findIndex(g.cache, e); i != -1 {
			g.cache = nil
			g.cacheInvalidated = true
		}
		g.callback(EntityRemoved, e)
	}
}

func (g *group) callback(ev GroupEvent, e Entity) {
	if cs, ok := g.callbacks[ev]; ok {
		for _, c := range cs {
			c(g, e)
		}
	}
}

func findIndex(entities []Entity, e Entity) int {
	for i, entity := range entities {
		if entity == e {
			return i
		}
	}
	return -1
}

func removeIndexed(entities []Entity, i int) []Entity {
	copy(entities[i:], entities[i+1:])
	entities[len(entities)-1] = nil
	new := entities[:len(entities)-1]
	return new
}
