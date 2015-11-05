// This file was generated by nomdl/codegen.

package main

import (
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

// MapOfStringToSetOfRefOfRemotePhoto

type MapOfStringToSetOfRefOfRemotePhoto struct {
	m   types.Map
	ref *ref.Ref
}

func NewMapOfStringToSetOfRefOfRemotePhoto() MapOfStringToSetOfRefOfRemotePhoto {
	return MapOfStringToSetOfRefOfRemotePhoto{types.NewMap(), &ref.Ref{}}
}

type MapOfStringToSetOfRefOfRemotePhotoDef map[string]SetOfRefOfRemotePhotoDef

func (def MapOfStringToSetOfRefOfRemotePhotoDef) New() MapOfStringToSetOfRefOfRemotePhoto {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), v.New())
	}
	return MapOfStringToSetOfRefOfRemotePhoto{types.NewMap(kv...), &ref.Ref{}}
}

func (m MapOfStringToSetOfRefOfRemotePhoto) Def() MapOfStringToSetOfRefOfRemotePhotoDef {
	def := make(map[string]SetOfRefOfRemotePhotoDef)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = v.(SetOfRefOfRemotePhoto).Def()
		return false
	})
	return def
}

func (m MapOfStringToSetOfRefOfRemotePhoto) Equals(other types.Value) bool {
	return other != nil && __typeRefForMapOfStringToSetOfRefOfRemotePhoto.Equals(other.TypeRef()) && m.Ref() == other.Ref()
}

func (m MapOfStringToSetOfRefOfRemotePhoto) Ref() ref.Ref {
	return types.EnsureRef(m.ref, m)
}

func (m MapOfStringToSetOfRefOfRemotePhoto) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, m.TypeRef().Chunks()...)
	chunks = append(chunks, m.m.Chunks()...)
	return
}

// A Noms Value that describes MapOfStringToSetOfRefOfRemotePhoto.
var __typeRefForMapOfStringToSetOfRefOfRemotePhoto types.TypeRef

func (m MapOfStringToSetOfRefOfRemotePhoto) TypeRef() types.TypeRef {
	return __typeRefForMapOfStringToSetOfRefOfRemotePhoto
}

func init() {
	__typeRefForMapOfStringToSetOfRefOfRemotePhoto = types.MakeCompoundTypeRef(types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakeCompoundTypeRef(types.SetKind, types.MakeCompoundTypeRef(types.RefKind, types.MakeTypeRef(ref.Parse("sha1-00419ebbb418539af67238164b20341913efeb4d"), 0))))
	types.RegisterValue(__typeRefForMapOfStringToSetOfRefOfRemotePhoto, builderForMapOfStringToSetOfRefOfRemotePhoto, readerForMapOfStringToSetOfRefOfRemotePhoto)
}

func builderForMapOfStringToSetOfRefOfRemotePhoto(v types.Value) types.Value {
	return MapOfStringToSetOfRefOfRemotePhoto{v.(types.Map), &ref.Ref{}}
}

func readerForMapOfStringToSetOfRefOfRemotePhoto(v types.Value) types.Value {
	return v.(MapOfStringToSetOfRefOfRemotePhoto).m
}

func (m MapOfStringToSetOfRefOfRemotePhoto) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToSetOfRefOfRemotePhoto) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToSetOfRefOfRemotePhoto) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToSetOfRefOfRemotePhoto) Get(p string) SetOfRefOfRemotePhoto {
	return m.m.Get(types.NewString(p)).(SetOfRefOfRemotePhoto)
}

func (m MapOfStringToSetOfRefOfRemotePhoto) MaybeGet(p string) (SetOfRefOfRemotePhoto, bool) {
	v, ok := m.m.MaybeGet(types.NewString(p))
	if !ok {
		return NewSetOfRefOfRemotePhoto(), false
	}
	return v.(SetOfRefOfRemotePhoto), ok
}

func (m MapOfStringToSetOfRefOfRemotePhoto) Set(k string, v SetOfRefOfRemotePhoto) MapOfStringToSetOfRefOfRemotePhoto {
	return MapOfStringToSetOfRefOfRemotePhoto{m.m.Set(types.NewString(k), v), &ref.Ref{}}
}

// TODO: Implement SetM?

func (m MapOfStringToSetOfRefOfRemotePhoto) Remove(p string) MapOfStringToSetOfRefOfRemotePhoto {
	return MapOfStringToSetOfRefOfRemotePhoto{m.m.Remove(types.NewString(p)), &ref.Ref{}}
}

type MapOfStringToSetOfRefOfRemotePhotoIterCallback func(k string, v SetOfRefOfRemotePhoto) (stop bool)

func (m MapOfStringToSetOfRefOfRemotePhoto) Iter(cb MapOfStringToSetOfRefOfRemotePhotoIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(SetOfRefOfRemotePhoto))
	})
}

type MapOfStringToSetOfRefOfRemotePhotoIterAllCallback func(k string, v SetOfRefOfRemotePhoto)

func (m MapOfStringToSetOfRefOfRemotePhoto) IterAll(cb MapOfStringToSetOfRefOfRemotePhotoIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), v.(SetOfRefOfRemotePhoto))
	})
}

type MapOfStringToSetOfRefOfRemotePhotoFilterCallback func(k string, v SetOfRefOfRemotePhoto) (keep bool)

func (m MapOfStringToSetOfRefOfRemotePhoto) Filter(cb MapOfStringToSetOfRefOfRemotePhotoFilterCallback) MapOfStringToSetOfRefOfRemotePhoto {
	nm := NewMapOfStringToSetOfRefOfRemotePhoto()
	m.IterAll(func(k string, v SetOfRefOfRemotePhoto) {
		if cb(k, v) {
			nm = nm.Set(k, v)
		}
	})
	return nm
}

// SetOfRefOfRemotePhoto

type SetOfRefOfRemotePhoto struct {
	s   types.Set
	ref *ref.Ref
}

func NewSetOfRefOfRemotePhoto() SetOfRefOfRemotePhoto {
	return SetOfRefOfRemotePhoto{types.NewSet(), &ref.Ref{}}
}

type SetOfRefOfRemotePhotoDef map[ref.Ref]bool

func (def SetOfRefOfRemotePhotoDef) New() SetOfRefOfRemotePhoto {
	l := make([]types.Value, len(def))
	i := 0
	for d, _ := range def {
		l[i] = NewRefOfRemotePhoto(d)
		i++
	}
	return SetOfRefOfRemotePhoto{types.NewSet(l...), &ref.Ref{}}
}

func (s SetOfRefOfRemotePhoto) Def() SetOfRefOfRemotePhotoDef {
	def := make(map[ref.Ref]bool, s.Len())
	s.s.Iter(func(v types.Value) bool {
		def[v.(RefOfRemotePhoto).TargetRef()] = true
		return false
	})
	return def
}

func (s SetOfRefOfRemotePhoto) Equals(other types.Value) bool {
	return other != nil && __typeRefForSetOfRefOfRemotePhoto.Equals(other.TypeRef()) && s.Ref() == other.Ref()
}

func (s SetOfRefOfRemotePhoto) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s SetOfRefOfRemotePhoto) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, s.TypeRef().Chunks()...)
	chunks = append(chunks, s.s.Chunks()...)
	return
}

// A Noms Value that describes SetOfRefOfRemotePhoto.
var __typeRefForSetOfRefOfRemotePhoto types.TypeRef

func (m SetOfRefOfRemotePhoto) TypeRef() types.TypeRef {
	return __typeRefForSetOfRefOfRemotePhoto
}

func init() {
	__typeRefForSetOfRefOfRemotePhoto = types.MakeCompoundTypeRef(types.SetKind, types.MakeCompoundTypeRef(types.RefKind, types.MakeTypeRef(ref.Parse("sha1-00419ebbb418539af67238164b20341913efeb4d"), 0)))
	types.RegisterValue(__typeRefForSetOfRefOfRemotePhoto, builderForSetOfRefOfRemotePhoto, readerForSetOfRefOfRemotePhoto)
}

func builderForSetOfRefOfRemotePhoto(v types.Value) types.Value {
	return SetOfRefOfRemotePhoto{v.(types.Set), &ref.Ref{}}
}

func readerForSetOfRefOfRemotePhoto(v types.Value) types.Value {
	return v.(SetOfRefOfRemotePhoto).s
}

func (s SetOfRefOfRemotePhoto) Empty() bool {
	return s.s.Empty()
}

func (s SetOfRefOfRemotePhoto) Len() uint64 {
	return s.s.Len()
}

func (s SetOfRefOfRemotePhoto) Has(p RefOfRemotePhoto) bool {
	return s.s.Has(p)
}

type SetOfRefOfRemotePhotoIterCallback func(p RefOfRemotePhoto) (stop bool)

func (s SetOfRefOfRemotePhoto) Iter(cb SetOfRefOfRemotePhotoIterCallback) {
	s.s.Iter(func(v types.Value) bool {
		return cb(v.(RefOfRemotePhoto))
	})
}

type SetOfRefOfRemotePhotoIterAllCallback func(p RefOfRemotePhoto)

func (s SetOfRefOfRemotePhoto) IterAll(cb SetOfRefOfRemotePhotoIterAllCallback) {
	s.s.IterAll(func(v types.Value) {
		cb(v.(RefOfRemotePhoto))
	})
}

func (s SetOfRefOfRemotePhoto) IterAllP(concurrency int, cb SetOfRefOfRemotePhotoIterAllCallback) {
	s.s.IterAllP(concurrency, func(v types.Value) {
		cb(v.(RefOfRemotePhoto))
	})
}

type SetOfRefOfRemotePhotoFilterCallback func(p RefOfRemotePhoto) (keep bool)

func (s SetOfRefOfRemotePhoto) Filter(cb SetOfRefOfRemotePhotoFilterCallback) SetOfRefOfRemotePhoto {
	ns := NewSetOfRefOfRemotePhoto()
	s.IterAll(func(v RefOfRemotePhoto) {
		if cb(v) {
			ns = ns.Insert(v)
		}
	})
	return ns
}

func (s SetOfRefOfRemotePhoto) Insert(p ...RefOfRemotePhoto) SetOfRefOfRemotePhoto {
	return SetOfRefOfRemotePhoto{s.s.Insert(s.fromElemSlice(p)...), &ref.Ref{}}
}

func (s SetOfRefOfRemotePhoto) Remove(p ...RefOfRemotePhoto) SetOfRefOfRemotePhoto {
	return SetOfRefOfRemotePhoto{s.s.Remove(s.fromElemSlice(p)...), &ref.Ref{}}
}

func (s SetOfRefOfRemotePhoto) Union(others ...SetOfRefOfRemotePhoto) SetOfRefOfRemotePhoto {
	return SetOfRefOfRemotePhoto{s.s.Union(s.fromStructSlice(others)...), &ref.Ref{}}
}

func (s SetOfRefOfRemotePhoto) Subtract(others ...SetOfRefOfRemotePhoto) SetOfRefOfRemotePhoto {
	return SetOfRefOfRemotePhoto{s.s.Subtract(s.fromStructSlice(others)...), &ref.Ref{}}
}

func (s SetOfRefOfRemotePhoto) Any() RefOfRemotePhoto {
	return s.s.Any().(RefOfRemotePhoto)
}

func (s SetOfRefOfRemotePhoto) fromStructSlice(p []SetOfRefOfRemotePhoto) []types.Set {
	r := make([]types.Set, len(p))
	for i, v := range p {
		r[i] = v.s
	}
	return r
}

func (s SetOfRefOfRemotePhoto) fromElemSlice(p []RefOfRemotePhoto) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

// RefOfRemotePhoto

type RefOfRemotePhoto struct {
	target ref.Ref
	ref    *ref.Ref
}

func NewRefOfRemotePhoto(target ref.Ref) RefOfRemotePhoto {
	return RefOfRemotePhoto{target, &ref.Ref{}}
}

func (r RefOfRemotePhoto) TargetRef() ref.Ref {
	return r.target
}

func (r RefOfRemotePhoto) Ref() ref.Ref {
	return types.EnsureRef(r.ref, r)
}

func (r RefOfRemotePhoto) Equals(other types.Value) bool {
	return other != nil && __typeRefForRefOfRemotePhoto.Equals(other.TypeRef()) && r.Ref() == other.Ref()
}

func (r RefOfRemotePhoto) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, r.TypeRef().Chunks()...)
	chunks = append(chunks, r.target)
	return
}

// A Noms Value that describes RefOfRemotePhoto.
var __typeRefForRefOfRemotePhoto types.TypeRef

func (m RefOfRemotePhoto) TypeRef() types.TypeRef {
	return __typeRefForRefOfRemotePhoto
}

func init() {
	__typeRefForRefOfRemotePhoto = types.MakeCompoundTypeRef(types.RefKind, types.MakeTypeRef(ref.Parse("sha1-00419ebbb418539af67238164b20341913efeb4d"), 0))
	types.RegisterFromValFunction(__typeRefForRefOfRemotePhoto, func(v types.Value) types.Value {
		return NewRefOfRemotePhoto(v.(types.Ref).TargetRef())
	})
}

func (r RefOfRemotePhoto) TargetValue(cs chunks.ChunkSource) RemotePhoto {
	return types.ReadValue(r.target, cs).(RemotePhoto)
}

func (r RefOfRemotePhoto) SetTargetValue(val RemotePhoto, cs chunks.ChunkSink) RefOfRemotePhoto {
	return NewRefOfRemotePhoto(types.WriteValue(val, cs))
}
