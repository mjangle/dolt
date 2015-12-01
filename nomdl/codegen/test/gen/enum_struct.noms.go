// This file was generated by nomdl/codegen.

package gen

import (
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __genPackageInFile_enum_struct_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.Type{
		types.MakeEnumType("Handedness", "right", "left", "switch"),
		types.MakeStructType("EnumStruct",
			[]types.Field{
				types.Field{"hand", types.MakeType(ref.Ref{}, 0), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{})
	__genPackageInFile_enum_struct_CachedRef = types.RegisterPackage(&p)
}

// Handedness

type Handedness uint32

const (
	Right Handedness = iota
	Left
	Switch
)

func NewHandedness() Handedness {
	return Handedness(0)
}

var __typeForHandedness types.Type

func (e Handedness) Type() types.Type {
	return __typeForHandedness
}

func init() {
	__typeForHandedness = types.MakeType(__genPackageInFile_enum_struct_CachedRef, 0)
	types.RegisterEnum(__typeForHandedness, builderForHandedness, readerForHandedness)
}

func builderForHandedness(v uint32) types.Value {
	return Handedness(v)
}

func readerForHandedness(v types.Value) uint32 {
	return uint32(v.(Handedness))
}

func (e Handedness) Equals(other types.Value) bool {
	return e == other
}

func (e Handedness) Ref() ref.Ref {
	throwaway := ref.Ref{}
	return types.EnsureRef(&throwaway, e)
}

func (e Handedness) Chunks() []ref.Ref {
	return nil
}

func (e Handedness) ChildValues() []types.Value {
	return nil
}

// EnumStruct

type EnumStruct struct {
	_hand Handedness

	cs  chunks.ChunkStore
	ref *ref.Ref
}

func NewEnumStruct(cs chunks.ChunkStore) EnumStruct {
	return EnumStruct{
		_hand: NewHandedness(),

		cs:  cs,
		ref: &ref.Ref{},
	}
}

type EnumStructDef struct {
	Hand Handedness
}

func (def EnumStructDef) New(cs chunks.ChunkStore) EnumStruct {
	return EnumStruct{
		_hand: def.Hand,
		cs:    cs,
		ref:   &ref.Ref{},
	}
}

func (s EnumStruct) Def() (d EnumStructDef) {
	d.Hand = s._hand
	return
}

var __typeForEnumStruct types.Type

func (m EnumStruct) Type() types.Type {
	return __typeForEnumStruct
}

func init() {
	__typeForEnumStruct = types.MakeType(__genPackageInFile_enum_struct_CachedRef, 1)
	types.RegisterStruct(__typeForEnumStruct, builderForEnumStruct, readerForEnumStruct)
}

func builderForEnumStruct(cs chunks.ChunkStore, values []types.Value) types.Value {
	i := 0
	s := EnumStruct{ref: &ref.Ref{}, cs: cs}
	s._hand = values[i].(Handedness)
	i++
	return s
}

func readerForEnumStruct(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(EnumStruct)
	values = append(values, s._hand)
	return values
}

func (s EnumStruct) Equals(other types.Value) bool {
	return other != nil && __typeForEnumStruct.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s EnumStruct) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s EnumStruct) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeForEnumStruct.Chunks()...)
	return
}

func (s EnumStruct) ChildValues() (ret []types.Value) {
	ret = append(ret, s._hand)
	return
}

func (s EnumStruct) Hand() Handedness {
	return s._hand
}

func (s EnumStruct) SetHand(val Handedness) EnumStruct {
	s._hand = val
	s.ref = &ref.Ref{}
	return s
}
