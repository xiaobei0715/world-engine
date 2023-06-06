package storage

// EntityIterator is an iterator for Ent lists in archetypes.
type EntityIterator struct {
	current      int
	archAccessor ArchetypeAccessor
	indices      []ArchetypeIndex
}

// NewEntityIterator returns an iterator for Entitys.
func NewEntityIterator(current int, archAccessor ArchetypeAccessor, indices []ArchetypeIndex) EntityIterator {
	return EntityIterator{
		current:      current,
		archAccessor: archAccessor,
		indices:      indices,
	}
}

// HasNext returns true if there are more Ent list to iterate over.
func (it *EntityIterator) HasNext() bool {
	return it.current < len(it.indices)
}

// Next returns the next Ent list.
func (it *EntityIterator) Next() []EntityID {
	archetypeIndex := it.indices[it.current]
	it.current++
	return it.archAccessor.Archetype(archetypeIndex).Entities()
}