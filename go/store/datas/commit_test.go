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

package datas

import (
	"context"
	"testing"

	"github.com/liquidata-inc/ld/dolt/go/store/chunks"
	"github.com/liquidata-inc/ld/dolt/go/store/nomdl"
	"github.com/liquidata-inc/ld/dolt/go/store/types"
	"github.com/stretchr/testify/assert"
)

func mustHead(ds Dataset) types.Struct {
	s, ok := ds.MaybeHead()
	if !ok {
		panic("no head")
	}

	return s
}

func mustHeadRef(ds Dataset) types.Ref {
	hr, ok := ds.MaybeHeadRef()
	if !ok {
		panic("no head")
	}

	return hr
}

func mustHeadValue(ds Dataset) types.Value {
	val, ok := ds.MaybeHeadValue()
	if !ok {
		panic("no head")
	}

	return val
}

func TestNewCommit(t *testing.T) {
	assert := assert.New(t)

	assertTypeEquals := func(e, a *types.Type) {
		assert.True(a.Equals(e), "Actual: %s\nExpected %s", a.Describe(context.Background()), e.Describe(context.Background()))
	}

	storage := &chunks.TestStorage{}
	db := NewDatabase(storage.NewView())
	defer db.Close()

	commit := NewCommit(types.Float(1), types.NewSet(context.Background(), db), types.EmptyStruct(types.Format_7_18))
	at := types.TypeOf(commit)
	et := makeCommitStructType(
		types.EmptyStructType,
		types.MakeSetType(types.MakeUnionType()),
		types.FloaTType,
	)
	assertTypeEquals(et, at)

	// Committing another Float
	commit2 := NewCommit(types.Float(2), types.NewSet(context.Background(), db, types.NewRef(commit, types.Format_7_18)), types.EmptyStruct(types.Format_7_18))
	at2 := types.TypeOf(commit2)
	et2 := nomdl.MustParseType(`Struct Commit {
                meta: Struct {},
                parents: Set<Ref<Cycle<Commit>>>,
                value: Float,
        }`)
	assertTypeEquals(et2, at2)

	// Now commit a String
	commit3 := NewCommit(types.String("Hi"), types.NewSet(context.Background(), db, types.NewRef(commit2, types.Format_7_18)), types.EmptyStruct(types.Format_7_18))
	at3 := types.TypeOf(commit3)
	et3 := nomdl.MustParseType(`Struct Commit {
                meta: Struct {},
                parents: Set<Ref<Cycle<Commit>>>,
                value: Float | String,
        }`)
	assertTypeEquals(et3, at3)

	// Now commit a String with MetaInfo
	meta := types.NewStruct(types.Format_7_18, "Meta", types.StructData{"date": types.String("some date"), "number": types.Float(9)})
	metaType := nomdl.MustParseType(`Struct Meta {
                date: String,
                number: Float,
	}`)
	assertTypeEquals(metaType, types.TypeOf(meta))
	commit4 := NewCommit(types.String("Hi"), types.NewSet(context.Background(), db, types.NewRef(commit2, types.Format_7_18)), meta)
	at4 := types.TypeOf(commit4)
	et4 := nomdl.MustParseType(`Struct Commit {
                meta: Struct {} | Struct Meta {
                        date: String,
                        number: Float,
        	},
                parents: Set<Ref<Cycle<Commit>>>,
                value: Float | String,
        }`)
	assertTypeEquals(et4, at4)

	// Merge-commit with different parent types
	commit5 := NewCommit(
		types.String("Hi"),
		types.NewSet(context.Background(), db,
			types.NewRef(commit2, types.Format_7_18),
			types.NewRef(commit3, types.Format_7_18)),
		types.EmptyStruct(types.Format_7_18))
	at5 := types.TypeOf(commit5)
	et5 := nomdl.MustParseType(`Struct Commit {
                meta: Struct {},
                parents: Set<Ref<Cycle<Commit>>>,
                value: Float | String,
        }`)
	assertTypeEquals(et5, at5)
}

func TestCommitWithoutMetaField(t *testing.T) {
	assert := assert.New(t)

	storage := &chunks.TestStorage{}
	db := NewDatabase(storage.NewView())
	defer db.Close()

	metaCommit := types.NewStruct(types.Format_7_18, "Commit", types.StructData{
		"value":   types.Float(9),
		"parents": types.NewSet(context.Background(), db),
		"meta":    types.EmptyStruct(types.Format_7_18),
	})
	assert.True(IsCommit(metaCommit))

	noMetaCommit := types.NewStruct(types.Format_7_18, "Commit", types.StructData{
		"value":   types.Float(9),
		"parents": types.NewSet(context.Background(), db),
	})
	assert.False(IsCommit(noMetaCommit))
}

// Convert list of Struct's to Set<Ref>
func toRefSet(vrw types.ValueReadWriter, commits ...types.Struct) types.Set {
	set := types.NewSet(context.Background(), vrw).Edit()
	for _, p := range commits {
		set.Insert(types.NewRef(p, types.Format_7_18))
	}
	return set.Set(context.Background())
}

func TestFindCommonAncestor(t *testing.T) {
	assert := assert.New(t)
	storage := &chunks.TestStorage{}
	db := NewDatabase(storage.NewView())
	defer db.Close()

	// Add a commit and return it
	addCommit := func(datasetID string, val string, parents ...types.Struct) types.Struct {
		ds, err := db.GetDataset(context.Background(), datasetID)
		assert.NoError(err)
		ds, err = db.Commit(context.Background(), ds, types.String(val), CommitOptions{Parents: toRefSet(db, parents...)})
		assert.NoError(err)
		return mustHead(ds)
	}

	// Assert that c is the common ancestor of a and b
	assertCommonAncestor := func(expected, a, b types.Struct) {
		if found, ok := FindCommonAncestor(context.Background(), types.NewRef(a, types.Format_7_18), types.NewRef(b, types.Format_7_18), db); assert.True(ok) {
			ancestor := found.TargetValue(context.Background(), db).(types.Struct)
			assert.True(
				expected.Equals(ancestor),
				"%s should be common ancestor of %s, %s. Got %s",
				expected.Get(ValueField),
				a.Get(ValueField),
				b.Get(ValueField),
				ancestor.Get(ValueField),
			)
		}
	}

	// Build commit DAG
	//
	// ds-a: a1<-a2<-a3<-a4<-a5<-a6
	//       ^    ^   ^          |
	//       |     \   \----\  /-/
	//       |      \        \V
	// ds-b:  \      b3<-b4<-b5
	//         \
	//          \
	// ds-c:     c2<-c3
	//              /
	//             /
	//            V
	// ds-d: d1<-d2
	//
	a, b, c, d := "ds-a", "ds-b", "ds-c", "ds-d"
	a1 := addCommit(a, "a1")
	d1 := addCommit(d, "d1")
	a2 := addCommit(a, "a2", a1)
	c2 := addCommit(c, "c2", a1)
	d2 := addCommit(d, "d2", d1)
	a3 := addCommit(a, "a3", a2)
	b3 := addCommit(b, "b3", a2)
	c3 := addCommit(c, "c3", c2, d2)
	a4 := addCommit(a, "a4", a3)
	b4 := addCommit(b, "b4", b3)
	a5 := addCommit(a, "a5", a4)
	b5 := addCommit(b, "b5", b4, a3)
	a6 := addCommit(a, "a6", a5, b5)

	assertCommonAncestor(a1, a1, a1) // All self
	assertCommonAncestor(a1, a1, a2) // One side self
	assertCommonAncestor(a2, a3, b3) // Common parent
	assertCommonAncestor(a2, a4, b4) // Common grandparent
	assertCommonAncestor(a1, a6, c3) // Traversing multiple parents on both sides

	// No common ancestor
	if found, ok := FindCommonAncestor(context.Background(), types.NewRef(d2, types.Format_7_18), types.NewRef(a6, types.Format_7_18), db); !assert.False(ok) {
		assert.Fail(
			"Unexpected common ancestor!",
			"Should be no common ancestor of %s, %s. Got %s",
			d2.Get(ValueField),
			a6.Get(ValueField),
			found.TargetValue(context.Background(), db).(types.Struct).Get(ValueField),
		)
	}
}

func TestNewCommitRegressionTest(t *testing.T) {
	storage := &chunks.TestStorage{}
	db := NewDatabase(storage.NewView())
	defer db.Close()

	c1 := NewCommit(types.String("one"), types.NewSet(context.Background(), db), types.EmptyStruct(types.Format_7_18))
	cx := NewCommit(types.Bool(true), types.NewSet(context.Background(), db), types.EmptyStruct(types.Format_7_18))
	value := types.String("two")
	parents := types.NewSet(context.Background(), db, types.NewRef(c1, types.Format_7_18))
	meta := types.NewStruct(types.Format_7_18, "", types.StructData{
		"basis": cx,
	})

	// Used to fail
	NewCommit(value, parents, meta)
}