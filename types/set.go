package types

import (
	"runtime"
	"sort"
	"sync"

	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/ref"
)

type Set struct {
	data    setData // sorted by Ref()
	indexOf indexOfSetFn
	t       Type
	ref     *ref.Ref
	cs      chunks.ChunkStore
}

type setData []Value

type indexOfSetFn func(m setData, v Value) int

var setType = MakeCompoundType(SetKind, MakePrimitiveType(ValueKind))

func NewSet(cs chunks.ChunkStore, v ...Value) Set {
	return NewTypedSet(cs, setType, v...)
}

func NewTypedSet(cs chunks.ChunkStore, t Type, v ...Value) Set {
	return newSetFromData(cs, buildSetData(setData{}, v, t), t)
}

func (s Set) Empty() bool {
	return s.Len() == uint64(0)
}

func (s Set) Len() uint64 {
	return uint64(len(s.data))
}

func (s Set) Has(v Value) bool {
	idx := s.indexOf(s.data, v)
	return idx < len(s.data) && s.data[idx].Equals(v)
}

func (s Set) Insert(values ...Value) Set {
	assertType(s.elemType(), values...)
	return newSetFromData(s.cs, buildSetData(s.data, values, s.t), s.t)
}

func (s Set) Remove(values ...Value) Set {
	data := copySetData(s.data)
	for _, v := range values {
		if v != nil {
			idx := s.indexOf(s.data, v)
			if idx < len(s.data) && s.data[idx].Equals(v) {
				data = append(data[:idx], data[idx+1:]...)
			}
		}
	}
	return newSetFromData(s.cs, data, s.t)
}

func (s Set) Union(others ...Set) Set {
	assertSetsSameType(s, others...)
	result := s
	for _, other := range others {
		other.Iter(func(v Value) (stop bool) {
			result = result.Insert(v)
			return
		})
	}
	return result
}

func (s Set) Subtract(others ...Set) Set {
	result := s
	for _, other := range others {
		other.Iter(func(v Value) (stop bool) {
			result = result.Remove(v)
			return
		})
	}
	return result
}

type setIterCallback func(v Value) bool

func (s Set) Iter(cb setIterCallback) {
	for _, v := range s.data {
		if cb(v) {
			break
		}
	}
}

type setIterAllCallback func(v Value)

func (s Set) IterAll(cb setIterAllCallback) {
	for _, v := range s.data {
		cb(v)
	}
}

func (s Set) IterAllP(concurrency int, f setIterAllCallback) {
	if concurrency == 0 {
		concurrency = runtime.NumCPU()
	}
	sem := make(chan int, concurrency)

	wg := sync.WaitGroup{}

	for idx := range s.data {
		wg.Add(1)

		sem <- 1
		go func(idx int) {
			defer wg.Done()
			f(s.data[idx])
			<-sem
		}(idx)
	}

	wg.Wait()
}

type setFilterCallback func(v Value) (keep bool)

func (s Set) Filter(cb setFilterCallback) Set {
	data := setData{}
	for _, v := range s.data {
		if cb(v) {
			data = append(data, v)
		}
	}
	// Already sorted.
	return newSetFromData(s.cs, data, s.t)
}

func (s Set) Any() Value {
	for _, v := range s.data {
		return v
	}
	return nil
}

func (s Set) Ref() ref.Ref {
	return EnsureRef(s.ref, s)
}

func (s Set) Equals(other Value) bool {
	return other != nil && s.t.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s Set) Chunks() (chunks []ref.Ref) {
	for _, v := range s.data {
		chunks = append(chunks, v.Chunks()...)
	}
	return
}

func (s Set) ChildValues() []Value {
	return append([]Value{}, s.data...)
}

func (s Set) Type() Type {
	return s.t
}

func (s Set) elemType() Type {
	return s.t.Desc.(CompoundDesc).ElemTypes[0]
}

func newSetFromData(cs chunks.ChunkStore, m setData, t Type) Set {
	return Set{m, getIndexFnForSetType(t), t, &ref.Ref{}, cs}
}

func copySetData(m setData) setData {
	r := make(setData, len(m))
	copy(r, m)
	return r
}

func buildSetData(old setData, values []Value, t Type) setData {
	idxFn := getIndexFnForSetType(t)
	elemType := t.Desc.(CompoundDesc).ElemTypes[0]

	data := make(setData, len(old), len(old)+len(values))
	copy(data, old)
	for _, v := range values {
		assertType(elemType, v)
		idx := idxFn(data, v)
		if idx < len(data) && data[idx].Equals(v) {
			// We already have this fellow.
			continue
		}
		// TODO: These repeated copies suck. We're not allocating more memory (because we made the slice with the correct capacity to begin with above - yay!), but still, this is more work than necessary. Perhaps we should use an actual BST for the in-memory state, rather than a flat list.
		data = append(data, nil)
		copy(data[idx+1:], data[idx:])
		data[idx] = v
	}
	return data
}

func getIndexFnForSetType(t Type) indexOfSetFn {
	orderByValue := t.Desc.(CompoundDesc).ElemTypes[0].IsOrdered()
	if orderByValue {
		return indexOfOrderedSetValue
	}

	return indexOfSetValue
}

func indexOfSetValue(m setData, v Value) int {
	return sort.Search(len(m), func(i int) bool {
		return !m[i].Ref().Less(v.Ref())
	})
}

func indexOfOrderedSetValue(m setData, v Value) int {
	ov := v.(OrderedValue)

	return sort.Search(len(m), func(i int) bool {
		return !m[i].(OrderedValue).Less(ov)
	})
}
