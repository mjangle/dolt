// Copyright 2019 Liquidata, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// This file incorporates work covered by the following copyright and
// permission notice:
//
// Copyright 2016 Attic Labs, Inc. All rights reserved.
// Licensed under the Apache License, version 2.0:
// http://www.apache.org/licenses/LICENSE-2.0

package types

import (
	"context"
	"math"

	"github.com/liquidata-inc/ld/dolt/go/store/d"
)

type leafSequence struct {
	sequenceImpl
}

func newLeafSequence(vrw ValueReadWriter, buff []byte, offsets []uint32, len uint64) leafSequence {
	return leafSequence{newSequenceImpl(vrw, buff, offsets, len)}
}

func newLeafSequenceFromValues(kind NomsKind, vrw ValueReadWriter, vs ...Value) leafSequence {
	d.PanicIfTrue(vrw == nil)
	w := newBinaryNomsWriter()
	offsets := make([]uint32, len(vs)+sequencePartValues+1)
	offsets[sequencePartKind] = w.offset
	kind.writeTo(&w, vrw.Format())
	offsets[sequencePartLevel] = w.offset
	w.writeCount(0) // level
	offsets[sequencePartCount] = w.offset
	count := uint64(len(vs))
	w.writeCount(count)
	offsets[sequencePartValues] = w.offset
	for i, v := range vs {
		v.writeTo(&w, vrw.Format())
		offsets[i+sequencePartValues+1] = w.offset
	}
	return newLeafSequence(vrw, w.data(), offsets, count)
}

func (seq leafSequence) values() []Value {
	return seq.valuesSlice(0, math.MaxUint64)
}

func (seq leafSequence) valuesSlice(from, to uint64) []Value {
	if len := seq.Len(); to > len {
		to = len
	}

	dec := seq.decoderSkipToIndex(int(from))
	vs := make([]Value, (to-from)*uint64(getValuesPerIdx(seq.Kind())))
	for i := range vs {
		vs[i] = dec.readValue(seq.format())
	}
	return vs
}

func (seq leafSequence) getCompareFnHelper(other leafSequence) compareFn {
	dec := seq.decoder()
	otherDec := other.decoder()

	return func(idx, otherIdx int) bool {
		dec.offset = uint32(seq.getItemOffset(idx))
		otherDec.offset = uint32(other.getItemOffset(otherIdx))
		return dec.readValue(seq.format()).Equals(otherDec.readValue(seq.format()))
	}
}

func (seq leafSequence) getCompareFn(other sequence) compareFn {
	panic("unreachable")
}

func (seq leafSequence) typeOf() *Type {
	dec := seq.decoder()
	kind := dec.readKind()
	dec.skipCount() // level
	count := dec.readCount()
	ts := make(typeSlice, 0, count)
	var lastType *Type
	for i := uint64(0); i < count; i++ {
		if lastType != nil {
			offset := dec.offset
			if dec.isValueSameTypeForSure(seq.format(), lastType) {
				continue
			}
			dec.offset = offset
		}

		lastType = dec.readTypeOfValue(seq.format())
		ts = append(ts, lastType)
	}

	return makeCompoundType(kind, makeUnionType(ts...))
}

func (seq leafSequence) numLeaves() uint64 {
	return seq.len
}

func (seq leafSequence) getChildSequence(ctx context.Context, idx int) sequence {
	return nil
}

func (seq leafSequence) treeLevel() uint64 {
	return 0
}

func (seq leafSequence) isLeaf() bool {
	return true
}

func (seq leafSequence) cumulativeNumberOfLeaves(idx int) uint64 {
	return uint64(idx) + 1
}

func (seq leafSequence) getCompositeChildSequence(ctx context.Context, start uint64, length uint64) sequence {
	panic("getCompositeChildSequence called on a leaf sequence")
}

func (seq leafSequence) getItem(idx int) sequenceItem {
	dec := seq.decoderSkipToIndex(idx)
	return dec.readValue(seq.format())
}

func getValuesPerIdx(kind NomsKind) int {
	if kind == MapKind {
		return 2
	}
	return 1
}