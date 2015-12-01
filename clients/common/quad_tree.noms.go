// This file was generated by nomdl/codegen.

package common

import (
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __commonPackageInFile_quad_tree_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.Type{
		types.MakeStructType("Node",
			[]types.Field{
				types.Field{"Geoposition", types.MakeType(ref.Parse("sha1-6d5e1c54214264058be9f61f4b4ece0368c8c678"), 0), false},
				types.Field{"Reference", types.MakeCompoundType(types.RefKind, types.MakePrimitiveType(types.ValueKind)), false},
			},
			types.Choices{},
		),
		types.MakeStructType("QuadTree",
			[]types.Field{
				types.Field{"Nodes", types.MakeCompoundType(types.ListKind, types.MakeType(ref.Ref{}, 0)), false},
				types.Field{"Tiles", types.MakeCompoundType(types.MapKind, types.MakePrimitiveType(types.StringKind), types.MakeType(ref.Ref{}, 1)), false},
				types.Field{"Depth", types.MakePrimitiveType(types.UInt8Kind), false},
				types.Field{"NumDescendents", types.MakePrimitiveType(types.UInt32Kind), false},
				types.Field{"Path", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"Georectangle", types.MakeType(ref.Parse("sha1-6d5e1c54214264058be9f61f4b4ece0368c8c678"), 1), false},
			},
			types.Choices{},
		),
		types.MakeStructType("SQuadTree",
			[]types.Field{
				types.Field{"Nodes", types.MakeCompoundType(types.ListKind, types.MakeCompoundType(types.RefKind, types.MakePrimitiveType(types.ValueKind))), false},
				types.Field{"Tiles", types.MakeCompoundType(types.MapKind, types.MakePrimitiveType(types.StringKind), types.MakeCompoundType(types.RefKind, types.MakeType(ref.Ref{}, 2))), false},
				types.Field{"Depth", types.MakePrimitiveType(types.UInt8Kind), false},
				types.Field{"NumDescendents", types.MakePrimitiveType(types.UInt32Kind), false},
				types.Field{"Path", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"Georectangle", types.MakeType(ref.Parse("sha1-6d5e1c54214264058be9f61f4b4ece0368c8c678"), 1), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{
		ref.Parse("sha1-6d5e1c54214264058be9f61f4b4ece0368c8c678"),
	})
	__commonPackageInFile_quad_tree_CachedRef = types.RegisterPackage(&p)
}

// Node

type Node struct {
	_Geoposition Geoposition
	_Reference   RefOfValue

	cs  chunks.ChunkStore
	ref *ref.Ref
}

func NewNode(cs chunks.ChunkStore) Node {
	return Node{
		_Geoposition: NewGeoposition(cs),
		_Reference:   NewRefOfValue(ref.Ref{}),

		cs:  cs,
		ref: &ref.Ref{},
	}
}

type NodeDef struct {
	Geoposition GeopositionDef
	Reference   ref.Ref
}

func (def NodeDef) New(cs chunks.ChunkStore) Node {
	return Node{
		_Geoposition: def.Geoposition.New(cs),
		_Reference:   NewRefOfValue(def.Reference),
		cs:           cs,
		ref:          &ref.Ref{},
	}
}

func (s Node) Def() (d NodeDef) {
	d.Geoposition = s._Geoposition.Def()
	d.Reference = s._Reference.TargetRef()
	return
}

var __typeForNode types.Type

func (m Node) Type() types.Type {
	return __typeForNode
}

func init() {
	__typeForNode = types.MakeType(__commonPackageInFile_quad_tree_CachedRef, 0)
	types.RegisterStruct(__typeForNode, builderForNode, readerForNode)
}

func builderForNode(cs chunks.ChunkStore, values []types.Value) types.Value {
	i := 0
	s := Node{ref: &ref.Ref{}, cs: cs}
	s._Geoposition = values[i].(Geoposition)
	i++
	s._Reference = values[i].(RefOfValue)
	i++
	return s
}

func readerForNode(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(Node)
	values = append(values, s._Geoposition)
	values = append(values, s._Reference)
	return values
}

func (s Node) Equals(other types.Value) bool {
	return other != nil && __typeForNode.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s Node) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s Node) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeForNode.Chunks()...)
	chunks = append(chunks, s._Geoposition.Chunks()...)
	chunks = append(chunks, s._Reference.Chunks()...)
	return
}

func (s Node) ChildValues() (ret []types.Value) {
	ret = append(ret, s._Geoposition)
	ret = append(ret, s._Reference)
	return
}

func (s Node) Geoposition() Geoposition {
	return s._Geoposition
}

func (s Node) SetGeoposition(val Geoposition) Node {
	s._Geoposition = val
	s.ref = &ref.Ref{}
	return s
}

func (s Node) Reference() RefOfValue {
	return s._Reference
}

func (s Node) SetReference(val RefOfValue) Node {
	s._Reference = val
	s.ref = &ref.Ref{}
	return s
}

// QuadTree

type QuadTree struct {
	_Nodes          ListOfNode
	_Tiles          MapOfStringToQuadTree
	_Depth          uint8
	_NumDescendents uint32
	_Path           string
	_Georectangle   Georectangle

	cs  chunks.ChunkStore
	ref *ref.Ref
}

func NewQuadTree(cs chunks.ChunkStore) QuadTree {
	return QuadTree{
		_Nodes:          NewListOfNode(cs),
		_Tiles:          NewMapOfStringToQuadTree(cs),
		_Depth:          uint8(0),
		_NumDescendents: uint32(0),
		_Path:           "",
		_Georectangle:   NewGeorectangle(cs),

		cs:  cs,
		ref: &ref.Ref{},
	}
}

type QuadTreeDef struct {
	Nodes          ListOfNodeDef
	Tiles          MapOfStringToQuadTreeDef
	Depth          uint8
	NumDescendents uint32
	Path           string
	Georectangle   GeorectangleDef
}

func (def QuadTreeDef) New(cs chunks.ChunkStore) QuadTree {
	return QuadTree{
		_Nodes:          def.Nodes.New(cs),
		_Tiles:          def.Tiles.New(cs),
		_Depth:          def.Depth,
		_NumDescendents: def.NumDescendents,
		_Path:           def.Path,
		_Georectangle:   def.Georectangle.New(cs),
		cs:              cs,
		ref:             &ref.Ref{},
	}
}

func (s QuadTree) Def() (d QuadTreeDef) {
	d.Nodes = s._Nodes.Def()
	d.Tiles = s._Tiles.Def()
	d.Depth = s._Depth
	d.NumDescendents = s._NumDescendents
	d.Path = s._Path
	d.Georectangle = s._Georectangle.Def()
	return
}

var __typeForQuadTree types.Type

func (m QuadTree) Type() types.Type {
	return __typeForQuadTree
}

func init() {
	__typeForQuadTree = types.MakeType(__commonPackageInFile_quad_tree_CachedRef, 1)
	types.RegisterStruct(__typeForQuadTree, builderForQuadTree, readerForQuadTree)
}

func builderForQuadTree(cs chunks.ChunkStore, values []types.Value) types.Value {
	i := 0
	s := QuadTree{ref: &ref.Ref{}, cs: cs}
	s._Nodes = values[i].(ListOfNode)
	i++
	s._Tiles = values[i].(MapOfStringToQuadTree)
	i++
	s._Depth = uint8(values[i].(types.UInt8))
	i++
	s._NumDescendents = uint32(values[i].(types.UInt32))
	i++
	s._Path = values[i].(types.String).String()
	i++
	s._Georectangle = values[i].(Georectangle)
	i++
	return s
}

func readerForQuadTree(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(QuadTree)
	values = append(values, s._Nodes)
	values = append(values, s._Tiles)
	values = append(values, types.UInt8(s._Depth))
	values = append(values, types.UInt32(s._NumDescendents))
	values = append(values, types.NewString(s._Path))
	values = append(values, s._Georectangle)
	return values
}

func (s QuadTree) Equals(other types.Value) bool {
	return other != nil && __typeForQuadTree.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s QuadTree) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s QuadTree) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeForQuadTree.Chunks()...)
	chunks = append(chunks, s._Nodes.Chunks()...)
	chunks = append(chunks, s._Tiles.Chunks()...)
	chunks = append(chunks, s._Georectangle.Chunks()...)
	return
}

func (s QuadTree) ChildValues() (ret []types.Value) {
	ret = append(ret, s._Nodes)
	ret = append(ret, s._Tiles)
	ret = append(ret, types.UInt8(s._Depth))
	ret = append(ret, types.UInt32(s._NumDescendents))
	ret = append(ret, types.NewString(s._Path))
	ret = append(ret, s._Georectangle)
	return
}

func (s QuadTree) Nodes() ListOfNode {
	return s._Nodes
}

func (s QuadTree) SetNodes(val ListOfNode) QuadTree {
	s._Nodes = val
	s.ref = &ref.Ref{}
	return s
}

func (s QuadTree) Tiles() MapOfStringToQuadTree {
	return s._Tiles
}

func (s QuadTree) SetTiles(val MapOfStringToQuadTree) QuadTree {
	s._Tiles = val
	s.ref = &ref.Ref{}
	return s
}

func (s QuadTree) Depth() uint8 {
	return s._Depth
}

func (s QuadTree) SetDepth(val uint8) QuadTree {
	s._Depth = val
	s.ref = &ref.Ref{}
	return s
}

func (s QuadTree) NumDescendents() uint32 {
	return s._NumDescendents
}

func (s QuadTree) SetNumDescendents(val uint32) QuadTree {
	s._NumDescendents = val
	s.ref = &ref.Ref{}
	return s
}

func (s QuadTree) Path() string {
	return s._Path
}

func (s QuadTree) SetPath(val string) QuadTree {
	s._Path = val
	s.ref = &ref.Ref{}
	return s
}

func (s QuadTree) Georectangle() Georectangle {
	return s._Georectangle
}

func (s QuadTree) SetGeorectangle(val Georectangle) QuadTree {
	s._Georectangle = val
	s.ref = &ref.Ref{}
	return s
}

// SQuadTree

type SQuadTree struct {
	_Nodes          ListOfRefOfValue
	_Tiles          MapOfStringToRefOfSQuadTree
	_Depth          uint8
	_NumDescendents uint32
	_Path           string
	_Georectangle   Georectangle

	cs  chunks.ChunkStore
	ref *ref.Ref
}

func NewSQuadTree(cs chunks.ChunkStore) SQuadTree {
	return SQuadTree{
		_Nodes:          NewListOfRefOfValue(cs),
		_Tiles:          NewMapOfStringToRefOfSQuadTree(cs),
		_Depth:          uint8(0),
		_NumDescendents: uint32(0),
		_Path:           "",
		_Georectangle:   NewGeorectangle(cs),

		cs:  cs,
		ref: &ref.Ref{},
	}
}

type SQuadTreeDef struct {
	Nodes          ListOfRefOfValueDef
	Tiles          MapOfStringToRefOfSQuadTreeDef
	Depth          uint8
	NumDescendents uint32
	Path           string
	Georectangle   GeorectangleDef
}

func (def SQuadTreeDef) New(cs chunks.ChunkStore) SQuadTree {
	return SQuadTree{
		_Nodes:          def.Nodes.New(cs),
		_Tiles:          def.Tiles.New(cs),
		_Depth:          def.Depth,
		_NumDescendents: def.NumDescendents,
		_Path:           def.Path,
		_Georectangle:   def.Georectangle.New(cs),
		cs:              cs,
		ref:             &ref.Ref{},
	}
}

func (s SQuadTree) Def() (d SQuadTreeDef) {
	d.Nodes = s._Nodes.Def()
	d.Tiles = s._Tiles.Def()
	d.Depth = s._Depth
	d.NumDescendents = s._NumDescendents
	d.Path = s._Path
	d.Georectangle = s._Georectangle.Def()
	return
}

var __typeForSQuadTree types.Type

func (m SQuadTree) Type() types.Type {
	return __typeForSQuadTree
}

func init() {
	__typeForSQuadTree = types.MakeType(__commonPackageInFile_quad_tree_CachedRef, 2)
	types.RegisterStruct(__typeForSQuadTree, builderForSQuadTree, readerForSQuadTree)
}

func builderForSQuadTree(cs chunks.ChunkStore, values []types.Value) types.Value {
	i := 0
	s := SQuadTree{ref: &ref.Ref{}, cs: cs}
	s._Nodes = values[i].(ListOfRefOfValue)
	i++
	s._Tiles = values[i].(MapOfStringToRefOfSQuadTree)
	i++
	s._Depth = uint8(values[i].(types.UInt8))
	i++
	s._NumDescendents = uint32(values[i].(types.UInt32))
	i++
	s._Path = values[i].(types.String).String()
	i++
	s._Georectangle = values[i].(Georectangle)
	i++
	return s
}

func readerForSQuadTree(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(SQuadTree)
	values = append(values, s._Nodes)
	values = append(values, s._Tiles)
	values = append(values, types.UInt8(s._Depth))
	values = append(values, types.UInt32(s._NumDescendents))
	values = append(values, types.NewString(s._Path))
	values = append(values, s._Georectangle)
	return values
}

func (s SQuadTree) Equals(other types.Value) bool {
	return other != nil && __typeForSQuadTree.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s SQuadTree) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s SQuadTree) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeForSQuadTree.Chunks()...)
	chunks = append(chunks, s._Nodes.Chunks()...)
	chunks = append(chunks, s._Tiles.Chunks()...)
	chunks = append(chunks, s._Georectangle.Chunks()...)
	return
}

func (s SQuadTree) ChildValues() (ret []types.Value) {
	ret = append(ret, s._Nodes)
	ret = append(ret, s._Tiles)
	ret = append(ret, types.UInt8(s._Depth))
	ret = append(ret, types.UInt32(s._NumDescendents))
	ret = append(ret, types.NewString(s._Path))
	ret = append(ret, s._Georectangle)
	return
}

func (s SQuadTree) Nodes() ListOfRefOfValue {
	return s._Nodes
}

func (s SQuadTree) SetNodes(val ListOfRefOfValue) SQuadTree {
	s._Nodes = val
	s.ref = &ref.Ref{}
	return s
}

func (s SQuadTree) Tiles() MapOfStringToRefOfSQuadTree {
	return s._Tiles
}

func (s SQuadTree) SetTiles(val MapOfStringToRefOfSQuadTree) SQuadTree {
	s._Tiles = val
	s.ref = &ref.Ref{}
	return s
}

func (s SQuadTree) Depth() uint8 {
	return s._Depth
}

func (s SQuadTree) SetDepth(val uint8) SQuadTree {
	s._Depth = val
	s.ref = &ref.Ref{}
	return s
}

func (s SQuadTree) NumDescendents() uint32 {
	return s._NumDescendents
}

func (s SQuadTree) SetNumDescendents(val uint32) SQuadTree {
	s._NumDescendents = val
	s.ref = &ref.Ref{}
	return s
}

func (s SQuadTree) Path() string {
	return s._Path
}

func (s SQuadTree) SetPath(val string) SQuadTree {
	s._Path = val
	s.ref = &ref.Ref{}
	return s
}

func (s SQuadTree) Georectangle() Georectangle {
	return s._Georectangle
}

func (s SQuadTree) SetGeorectangle(val Georectangle) SQuadTree {
	s._Georectangle = val
	s.ref = &ref.Ref{}
	return s
}

// RefOfValue

type RefOfValue struct {
	target ref.Ref
	ref    *ref.Ref
}

func NewRefOfValue(target ref.Ref) RefOfValue {
	return RefOfValue{target, &ref.Ref{}}
}

func (r RefOfValue) TargetRef() ref.Ref {
	return r.target
}

func (r RefOfValue) Ref() ref.Ref {
	return types.EnsureRef(r.ref, r)
}

func (r RefOfValue) Equals(other types.Value) bool {
	return other != nil && __typeForRefOfValue.Equals(other.Type()) && r.Ref() == other.Ref()
}

func (r RefOfValue) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, r.Type().Chunks()...)
	chunks = append(chunks, r.target)
	return
}

func (r RefOfValue) ChildValues() []types.Value {
	return nil
}

// A Noms Value that describes RefOfValue.
var __typeForRefOfValue types.Type

func (m RefOfValue) Type() types.Type {
	return __typeForRefOfValue
}

func init() {
	__typeForRefOfValue = types.MakeCompoundType(types.RefKind, types.MakePrimitiveType(types.ValueKind))
	types.RegisterRef(__typeForRefOfValue, builderForRefOfValue)
}

func builderForRefOfValue(r ref.Ref) types.Value {
	return NewRefOfValue(r)
}

func (r RefOfValue) TargetValue(cs chunks.ChunkStore) types.Value {
	return types.ReadValue(r.target, cs)
}

func (r RefOfValue) SetTargetValue(val types.Value, cs chunks.ChunkSink) RefOfValue {
	return NewRefOfValue(types.WriteValue(val, cs))
}

// ListOfNode

type ListOfNode struct {
	l   types.List
	cs  chunks.ChunkStore
	ref *ref.Ref
}

func NewListOfNode(cs chunks.ChunkStore) ListOfNode {
	return ListOfNode{types.NewTypedList(cs, __typeForListOfNode), cs, &ref.Ref{}}
}

type ListOfNodeDef []NodeDef

func (def ListOfNodeDef) New(cs chunks.ChunkStore) ListOfNode {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = d.New(cs)
	}
	return ListOfNode{types.NewTypedList(cs, __typeForListOfNode, l...), cs, &ref.Ref{}}
}

func (l ListOfNode) Def() ListOfNodeDef {
	d := make([]NodeDef, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = l.l.Get(i).(Node).Def()
	}
	return d
}

func (l ListOfNode) Equals(other types.Value) bool {
	return other != nil && __typeForListOfNode.Equals(other.Type()) && l.Ref() == other.Ref()
}

func (l ListOfNode) Ref() ref.Ref {
	return types.EnsureRef(l.ref, l)
}

func (l ListOfNode) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, l.Type().Chunks()...)
	chunks = append(chunks, l.l.Chunks()...)
	return
}

func (l ListOfNode) ChildValues() []types.Value {
	return append([]types.Value{}, l.l.ChildValues()...)
}

// A Noms Value that describes ListOfNode.
var __typeForListOfNode types.Type

func (m ListOfNode) Type() types.Type {
	return __typeForListOfNode
}

func init() {
	__typeForListOfNode = types.MakeCompoundType(types.ListKind, types.MakeType(__commonPackageInFile_quad_tree_CachedRef, 0))
	types.RegisterValue(__typeForListOfNode, builderForListOfNode, readerForListOfNode)
}

func builderForListOfNode(cs chunks.ChunkStore, v types.Value) types.Value {
	return ListOfNode{v.(types.List), cs, &ref.Ref{}}
}

func readerForListOfNode(v types.Value) types.Value {
	return v.(ListOfNode).l
}

func (l ListOfNode) Len() uint64 {
	return l.l.Len()
}

func (l ListOfNode) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfNode) Get(i uint64) Node {
	return l.l.Get(i).(Node)
}

func (l ListOfNode) Slice(idx uint64, end uint64) ListOfNode {
	return ListOfNode{l.l.Slice(idx, end), l.cs, &ref.Ref{}}
}

func (l ListOfNode) Set(i uint64, val Node) ListOfNode {
	return ListOfNode{l.l.Set(i, val), l.cs, &ref.Ref{}}
}

func (l ListOfNode) Append(v ...Node) ListOfNode {
	return ListOfNode{l.l.Append(l.fromElemSlice(v)...), l.cs, &ref.Ref{}}
}

func (l ListOfNode) Insert(idx uint64, v ...Node) ListOfNode {
	return ListOfNode{l.l.Insert(idx, l.fromElemSlice(v)...), l.cs, &ref.Ref{}}
}

func (l ListOfNode) Remove(idx uint64, end uint64) ListOfNode {
	return ListOfNode{l.l.Remove(idx, end), l.cs, &ref.Ref{}}
}

func (l ListOfNode) RemoveAt(idx uint64) ListOfNode {
	return ListOfNode{(l.l.RemoveAt(idx)), l.cs, &ref.Ref{}}
}

func (l ListOfNode) fromElemSlice(p []Node) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

type ListOfNodeIterCallback func(v Node, i uint64) (stop bool)

func (l ListOfNode) Iter(cb ListOfNodeIterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(v.(Node), i)
	})
}

type ListOfNodeIterAllCallback func(v Node, i uint64)

func (l ListOfNode) IterAll(cb ListOfNodeIterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(v.(Node), i)
	})
}

func (l ListOfNode) IterAllP(concurrency int, cb ListOfNodeIterAllCallback) {
	l.l.IterAllP(concurrency, func(v types.Value, i uint64) {
		cb(v.(Node), i)
	})
}

type ListOfNodeFilterCallback func(v Node, i uint64) (keep bool)

func (l ListOfNode) Filter(cb ListOfNodeFilterCallback) ListOfNode {
	out := l.l.Filter(func(v types.Value, i uint64) bool {
		return cb(v.(Node), i)
	})
	return ListOfNode{out, l.cs, &ref.Ref{}}
}

// MapOfStringToQuadTree

type MapOfStringToQuadTree struct {
	m   types.Map
	cs  chunks.ChunkStore
	ref *ref.Ref
}

func NewMapOfStringToQuadTree(cs chunks.ChunkStore) MapOfStringToQuadTree {
	return MapOfStringToQuadTree{types.NewTypedMap(cs, __typeForMapOfStringToQuadTree), cs, &ref.Ref{}}
}

type MapOfStringToQuadTreeDef map[string]QuadTreeDef

func (def MapOfStringToQuadTreeDef) New(cs chunks.ChunkStore) MapOfStringToQuadTree {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), v.New(cs))
	}
	return MapOfStringToQuadTree{types.NewTypedMap(cs, __typeForMapOfStringToQuadTree, kv...), cs, &ref.Ref{}}
}

func (m MapOfStringToQuadTree) Def() MapOfStringToQuadTreeDef {
	def := make(map[string]QuadTreeDef)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = v.(QuadTree).Def()
		return false
	})
	return def
}

func (m MapOfStringToQuadTree) Equals(other types.Value) bool {
	return other != nil && __typeForMapOfStringToQuadTree.Equals(other.Type()) && m.Ref() == other.Ref()
}

func (m MapOfStringToQuadTree) Ref() ref.Ref {
	return types.EnsureRef(m.ref, m)
}

func (m MapOfStringToQuadTree) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, m.Type().Chunks()...)
	chunks = append(chunks, m.m.Chunks()...)
	return
}

func (m MapOfStringToQuadTree) ChildValues() []types.Value {
	return append([]types.Value{}, m.m.ChildValues()...)
}

// A Noms Value that describes MapOfStringToQuadTree.
var __typeForMapOfStringToQuadTree types.Type

func (m MapOfStringToQuadTree) Type() types.Type {
	return __typeForMapOfStringToQuadTree
}

func init() {
	__typeForMapOfStringToQuadTree = types.MakeCompoundType(types.MapKind, types.MakePrimitiveType(types.StringKind), types.MakeType(__commonPackageInFile_quad_tree_CachedRef, 1))
	types.RegisterValue(__typeForMapOfStringToQuadTree, builderForMapOfStringToQuadTree, readerForMapOfStringToQuadTree)
}

func builderForMapOfStringToQuadTree(cs chunks.ChunkStore, v types.Value) types.Value {
	return MapOfStringToQuadTree{v.(types.Map), cs, &ref.Ref{}}
}

func readerForMapOfStringToQuadTree(v types.Value) types.Value {
	return v.(MapOfStringToQuadTree).m
}

func (m MapOfStringToQuadTree) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToQuadTree) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToQuadTree) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToQuadTree) Get(p string) QuadTree {
	return m.m.Get(types.NewString(p)).(QuadTree)
}

func (m MapOfStringToQuadTree) MaybeGet(p string) (QuadTree, bool) {
	v, ok := m.m.MaybeGet(types.NewString(p))
	if !ok {
		return NewQuadTree(m.cs), false
	}
	return v.(QuadTree), ok
}

func (m MapOfStringToQuadTree) Set(k string, v QuadTree) MapOfStringToQuadTree {
	return MapOfStringToQuadTree{m.m.Set(types.NewString(k), v), m.cs, &ref.Ref{}}
}

// TODO: Implement SetM?

func (m MapOfStringToQuadTree) Remove(p string) MapOfStringToQuadTree {
	return MapOfStringToQuadTree{m.m.Remove(types.NewString(p)), m.cs, &ref.Ref{}}
}

type MapOfStringToQuadTreeIterCallback func(k string, v QuadTree) (stop bool)

func (m MapOfStringToQuadTree) Iter(cb MapOfStringToQuadTreeIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(QuadTree))
	})
}

type MapOfStringToQuadTreeIterAllCallback func(k string, v QuadTree)

func (m MapOfStringToQuadTree) IterAll(cb MapOfStringToQuadTreeIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), v.(QuadTree))
	})
}

func (m MapOfStringToQuadTree) IterAllP(concurrency int, cb MapOfStringToQuadTreeIterAllCallback) {
	m.m.IterAllP(concurrency, func(k, v types.Value) {
		cb(k.(types.String).String(), v.(QuadTree))
	})
}

type MapOfStringToQuadTreeFilterCallback func(k string, v QuadTree) (keep bool)

func (m MapOfStringToQuadTree) Filter(cb MapOfStringToQuadTreeFilterCallback) MapOfStringToQuadTree {
	out := m.m.Filter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(QuadTree))
	})
	return MapOfStringToQuadTree{out, m.cs, &ref.Ref{}}
}

// ListOfRefOfValue

type ListOfRefOfValue struct {
	l   types.List
	cs  chunks.ChunkStore
	ref *ref.Ref
}

func NewListOfRefOfValue(cs chunks.ChunkStore) ListOfRefOfValue {
	return ListOfRefOfValue{types.NewTypedList(cs, __typeForListOfRefOfValue), cs, &ref.Ref{}}
}

type ListOfRefOfValueDef []ref.Ref

func (def ListOfRefOfValueDef) New(cs chunks.ChunkStore) ListOfRefOfValue {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = NewRefOfValue(d)
	}
	return ListOfRefOfValue{types.NewTypedList(cs, __typeForListOfRefOfValue, l...), cs, &ref.Ref{}}
}

func (l ListOfRefOfValue) Def() ListOfRefOfValueDef {
	d := make([]ref.Ref, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = l.l.Get(i).(RefOfValue).TargetRef()
	}
	return d
}

func (l ListOfRefOfValue) Equals(other types.Value) bool {
	return other != nil && __typeForListOfRefOfValue.Equals(other.Type()) && l.Ref() == other.Ref()
}

func (l ListOfRefOfValue) Ref() ref.Ref {
	return types.EnsureRef(l.ref, l)
}

func (l ListOfRefOfValue) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, l.Type().Chunks()...)
	chunks = append(chunks, l.l.Chunks()...)
	return
}

func (l ListOfRefOfValue) ChildValues() []types.Value {
	return append([]types.Value{}, l.l.ChildValues()...)
}

// A Noms Value that describes ListOfRefOfValue.
var __typeForListOfRefOfValue types.Type

func (m ListOfRefOfValue) Type() types.Type {
	return __typeForListOfRefOfValue
}

func init() {
	__typeForListOfRefOfValue = types.MakeCompoundType(types.ListKind, types.MakeCompoundType(types.RefKind, types.MakePrimitiveType(types.ValueKind)))
	types.RegisterValue(__typeForListOfRefOfValue, builderForListOfRefOfValue, readerForListOfRefOfValue)
}

func builderForListOfRefOfValue(cs chunks.ChunkStore, v types.Value) types.Value {
	return ListOfRefOfValue{v.(types.List), cs, &ref.Ref{}}
}

func readerForListOfRefOfValue(v types.Value) types.Value {
	return v.(ListOfRefOfValue).l
}

func (l ListOfRefOfValue) Len() uint64 {
	return l.l.Len()
}

func (l ListOfRefOfValue) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfRefOfValue) Get(i uint64) RefOfValue {
	return l.l.Get(i).(RefOfValue)
}

func (l ListOfRefOfValue) Slice(idx uint64, end uint64) ListOfRefOfValue {
	return ListOfRefOfValue{l.l.Slice(idx, end), l.cs, &ref.Ref{}}
}

func (l ListOfRefOfValue) Set(i uint64, val RefOfValue) ListOfRefOfValue {
	return ListOfRefOfValue{l.l.Set(i, val), l.cs, &ref.Ref{}}
}

func (l ListOfRefOfValue) Append(v ...RefOfValue) ListOfRefOfValue {
	return ListOfRefOfValue{l.l.Append(l.fromElemSlice(v)...), l.cs, &ref.Ref{}}
}

func (l ListOfRefOfValue) Insert(idx uint64, v ...RefOfValue) ListOfRefOfValue {
	return ListOfRefOfValue{l.l.Insert(idx, l.fromElemSlice(v)...), l.cs, &ref.Ref{}}
}

func (l ListOfRefOfValue) Remove(idx uint64, end uint64) ListOfRefOfValue {
	return ListOfRefOfValue{l.l.Remove(idx, end), l.cs, &ref.Ref{}}
}

func (l ListOfRefOfValue) RemoveAt(idx uint64) ListOfRefOfValue {
	return ListOfRefOfValue{(l.l.RemoveAt(idx)), l.cs, &ref.Ref{}}
}

func (l ListOfRefOfValue) fromElemSlice(p []RefOfValue) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

type ListOfRefOfValueIterCallback func(v RefOfValue, i uint64) (stop bool)

func (l ListOfRefOfValue) Iter(cb ListOfRefOfValueIterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(v.(RefOfValue), i)
	})
}

type ListOfRefOfValueIterAllCallback func(v RefOfValue, i uint64)

func (l ListOfRefOfValue) IterAll(cb ListOfRefOfValueIterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(v.(RefOfValue), i)
	})
}

func (l ListOfRefOfValue) IterAllP(concurrency int, cb ListOfRefOfValueIterAllCallback) {
	l.l.IterAllP(concurrency, func(v types.Value, i uint64) {
		cb(v.(RefOfValue), i)
	})
}

type ListOfRefOfValueFilterCallback func(v RefOfValue, i uint64) (keep bool)

func (l ListOfRefOfValue) Filter(cb ListOfRefOfValueFilterCallback) ListOfRefOfValue {
	out := l.l.Filter(func(v types.Value, i uint64) bool {
		return cb(v.(RefOfValue), i)
	})
	return ListOfRefOfValue{out, l.cs, &ref.Ref{}}
}

// MapOfStringToRefOfSQuadTree

type MapOfStringToRefOfSQuadTree struct {
	m   types.Map
	cs  chunks.ChunkStore
	ref *ref.Ref
}

func NewMapOfStringToRefOfSQuadTree(cs chunks.ChunkStore) MapOfStringToRefOfSQuadTree {
	return MapOfStringToRefOfSQuadTree{types.NewTypedMap(cs, __typeForMapOfStringToRefOfSQuadTree), cs, &ref.Ref{}}
}

type MapOfStringToRefOfSQuadTreeDef map[string]ref.Ref

func (def MapOfStringToRefOfSQuadTreeDef) New(cs chunks.ChunkStore) MapOfStringToRefOfSQuadTree {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), NewRefOfSQuadTree(v))
	}
	return MapOfStringToRefOfSQuadTree{types.NewTypedMap(cs, __typeForMapOfStringToRefOfSQuadTree, kv...), cs, &ref.Ref{}}
}

func (m MapOfStringToRefOfSQuadTree) Def() MapOfStringToRefOfSQuadTreeDef {
	def := make(map[string]ref.Ref)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = v.(RefOfSQuadTree).TargetRef()
		return false
	})
	return def
}

func (m MapOfStringToRefOfSQuadTree) Equals(other types.Value) bool {
	return other != nil && __typeForMapOfStringToRefOfSQuadTree.Equals(other.Type()) && m.Ref() == other.Ref()
}

func (m MapOfStringToRefOfSQuadTree) Ref() ref.Ref {
	return types.EnsureRef(m.ref, m)
}

func (m MapOfStringToRefOfSQuadTree) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, m.Type().Chunks()...)
	chunks = append(chunks, m.m.Chunks()...)
	return
}

func (m MapOfStringToRefOfSQuadTree) ChildValues() []types.Value {
	return append([]types.Value{}, m.m.ChildValues()...)
}

// A Noms Value that describes MapOfStringToRefOfSQuadTree.
var __typeForMapOfStringToRefOfSQuadTree types.Type

func (m MapOfStringToRefOfSQuadTree) Type() types.Type {
	return __typeForMapOfStringToRefOfSQuadTree
}

func init() {
	__typeForMapOfStringToRefOfSQuadTree = types.MakeCompoundType(types.MapKind, types.MakePrimitiveType(types.StringKind), types.MakeCompoundType(types.RefKind, types.MakeType(__commonPackageInFile_quad_tree_CachedRef, 2)))
	types.RegisterValue(__typeForMapOfStringToRefOfSQuadTree, builderForMapOfStringToRefOfSQuadTree, readerForMapOfStringToRefOfSQuadTree)
}

func builderForMapOfStringToRefOfSQuadTree(cs chunks.ChunkStore, v types.Value) types.Value {
	return MapOfStringToRefOfSQuadTree{v.(types.Map), cs, &ref.Ref{}}
}

func readerForMapOfStringToRefOfSQuadTree(v types.Value) types.Value {
	return v.(MapOfStringToRefOfSQuadTree).m
}

func (m MapOfStringToRefOfSQuadTree) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToRefOfSQuadTree) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToRefOfSQuadTree) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToRefOfSQuadTree) Get(p string) RefOfSQuadTree {
	return m.m.Get(types.NewString(p)).(RefOfSQuadTree)
}

func (m MapOfStringToRefOfSQuadTree) MaybeGet(p string) (RefOfSQuadTree, bool) {
	v, ok := m.m.MaybeGet(types.NewString(p))
	if !ok {
		return NewRefOfSQuadTree(ref.Ref{}), false
	}
	return v.(RefOfSQuadTree), ok
}

func (m MapOfStringToRefOfSQuadTree) Set(k string, v RefOfSQuadTree) MapOfStringToRefOfSQuadTree {
	return MapOfStringToRefOfSQuadTree{m.m.Set(types.NewString(k), v), m.cs, &ref.Ref{}}
}

// TODO: Implement SetM?

func (m MapOfStringToRefOfSQuadTree) Remove(p string) MapOfStringToRefOfSQuadTree {
	return MapOfStringToRefOfSQuadTree{m.m.Remove(types.NewString(p)), m.cs, &ref.Ref{}}
}

type MapOfStringToRefOfSQuadTreeIterCallback func(k string, v RefOfSQuadTree) (stop bool)

func (m MapOfStringToRefOfSQuadTree) Iter(cb MapOfStringToRefOfSQuadTreeIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(RefOfSQuadTree))
	})
}

type MapOfStringToRefOfSQuadTreeIterAllCallback func(k string, v RefOfSQuadTree)

func (m MapOfStringToRefOfSQuadTree) IterAll(cb MapOfStringToRefOfSQuadTreeIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), v.(RefOfSQuadTree))
	})
}

func (m MapOfStringToRefOfSQuadTree) IterAllP(concurrency int, cb MapOfStringToRefOfSQuadTreeIterAllCallback) {
	m.m.IterAllP(concurrency, func(k, v types.Value) {
		cb(k.(types.String).String(), v.(RefOfSQuadTree))
	})
}

type MapOfStringToRefOfSQuadTreeFilterCallback func(k string, v RefOfSQuadTree) (keep bool)

func (m MapOfStringToRefOfSQuadTree) Filter(cb MapOfStringToRefOfSQuadTreeFilterCallback) MapOfStringToRefOfSQuadTree {
	out := m.m.Filter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(RefOfSQuadTree))
	})
	return MapOfStringToRefOfSQuadTree{out, m.cs, &ref.Ref{}}
}

// RefOfSQuadTree

type RefOfSQuadTree struct {
	target ref.Ref
	ref    *ref.Ref
}

func NewRefOfSQuadTree(target ref.Ref) RefOfSQuadTree {
	return RefOfSQuadTree{target, &ref.Ref{}}
}

func (r RefOfSQuadTree) TargetRef() ref.Ref {
	return r.target
}

func (r RefOfSQuadTree) Ref() ref.Ref {
	return types.EnsureRef(r.ref, r)
}

func (r RefOfSQuadTree) Equals(other types.Value) bool {
	return other != nil && __typeForRefOfSQuadTree.Equals(other.Type()) && r.Ref() == other.Ref()
}

func (r RefOfSQuadTree) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, r.Type().Chunks()...)
	chunks = append(chunks, r.target)
	return
}

func (r RefOfSQuadTree) ChildValues() []types.Value {
	return nil
}

// A Noms Value that describes RefOfSQuadTree.
var __typeForRefOfSQuadTree types.Type

func (m RefOfSQuadTree) Type() types.Type {
	return __typeForRefOfSQuadTree
}

func init() {
	__typeForRefOfSQuadTree = types.MakeCompoundType(types.RefKind, types.MakeType(__commonPackageInFile_quad_tree_CachedRef, 2))
	types.RegisterRef(__typeForRefOfSQuadTree, builderForRefOfSQuadTree)
}

func builderForRefOfSQuadTree(r ref.Ref) types.Value {
	return NewRefOfSQuadTree(r)
}

func (r RefOfSQuadTree) TargetValue(cs chunks.ChunkStore) SQuadTree {
	return types.ReadValue(r.target, cs).(SQuadTree)
}

func (r RefOfSQuadTree) SetTargetValue(val SQuadTree, cs chunks.ChunkSink) RefOfSQuadTree {
	return NewRefOfSQuadTree(types.WriteValue(val, cs))
}
