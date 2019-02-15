package main

// Map Store
type mapStore struct {
	ms    map[uint]Puppy
	mapID uint
}

func newMapStore() *mapStore {
	return &mapStore{ms: make(map[uint]Puppy)}
}

func (m *mapStore) createPuppy(in Puppy) uint {
	m.mapID++
	in.id = m.mapID
	m.ms[in.id] = in
	return m.mapID
}

func (m *mapStore) readPuppy(id uint) Puppy {
	p, ok := m.ms[id]
	if !ok {
		return Puppy{}
	}
	return p
}

func (m *mapStore) updatePuppy(id uint, in Puppy) {
	m.ms[id] = in
}

func (m *mapStore) deletePuppy(id uint) {
	delete(m.ms, id)
}
