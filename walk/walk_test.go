package walk

import (
	"testing"

	"github.com/attic-labs/noms/Godeps/_workspace/src/github.com/stretchr/testify/suite"
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

func TestWalkTestSuite(t *testing.T) {
	suite.Run(t, &WalkAllTestSuite{})
	suite.Run(t, &WalkTestSuite{})
}

type WalkAllTestSuite struct {
	suite.Suite
	cs *chunks.TestStore
}

func (suite *WalkAllTestSuite) SetupTest() {
	suite.cs = chunks.NewTestStore()
}

func (suite *WalkAllTestSuite) walkWorker(r ref.Ref, expected int) {
	actual := 0
	AllP(types.NewRef(r), suite.cs, func(c types.Value) {
		actual++
	}, 1)
	suite.Equal(expected, actual)
}

func (suite *WalkAllTestSuite) storeAndRef(v types.Value) ref.Ref {
	return types.WriteValue(v, suite.cs)
}

func (suite *WalkAllTestSuite) TestWalkPrimitives() {
	suite.walkWorker(suite.storeAndRef(types.Float64(0.0)), 2)
	suite.walkWorker(suite.storeAndRef(types.NewString("hello")), 2)
}

func (suite *WalkAllTestSuite) TestWalkComposites() {
	cs := chunks.NewMemoryStore()
	suite.walkWorker(suite.storeAndRef(types.NewList(cs)), 2)
	suite.walkWorker(suite.storeAndRef(types.NewList(cs, types.Bool(false), types.Int32(8))), 4)
	suite.walkWorker(suite.storeAndRef(types.NewSet(cs)), 2)
	suite.walkWorker(suite.storeAndRef(types.NewSet(cs, types.Bool(false), types.Int32(8))), 4)
	suite.walkWorker(suite.storeAndRef(types.NewMap(cs)), 2)
	suite.walkWorker(suite.storeAndRef(types.NewMap(cs, types.Int32(8), types.Bool(true), types.Int32(0), types.Bool(false))), 6)
}

func (suite *WalkAllTestSuite) NewList(cs chunks.ChunkStore, vs ...types.Value) types.Ref {
	v := types.NewList(cs, vs...)
	r := types.WriteValue(v, suite.cs)
	return types.NewRef(r)
}

func (suite *WalkAllTestSuite) NewMap(cs chunks.ChunkStore, vs ...types.Value) types.Ref {
	v := types.NewMap(cs, vs...)
	r := types.WriteValue(v, suite.cs)
	return types.NewRef(r)
}

func (suite *WalkAllTestSuite) NewSet(cs chunks.ChunkStore, vs ...types.Value) types.Ref {
	v := types.NewSet(cs, vs...)
	r := types.WriteValue(v, suite.cs)
	return types.NewRef(r)
}

func (suite *WalkAllTestSuite) TestWalkNestedComposites() {
	cs := chunks.NewMemoryStore()
	suite.walkWorker(suite.storeAndRef(types.NewList(cs, suite.NewSet(cs), types.Int32(8))), 5)
	suite.walkWorker(suite.storeAndRef(types.NewSet(cs, suite.NewList(cs), suite.NewSet(cs))), 6)
	// {"string": "string",
	//  "list": [false true],
	//  "map": {"nested": "string"}
	//  "mtlist": []
	//  "set": [5 7 8]
	//  []: "wow"
	// }
	nested := types.NewMap(cs,
		types.NewString("string"), types.NewString("string"),
		types.NewString("list"), suite.NewList(cs, types.Bool(false), types.Bool(true)),
		types.NewString("map"), suite.NewMap(cs, types.NewString("nested"), types.NewString("string")),
		types.NewString("mtlist"), suite.NewList(cs),
		types.NewString("set"), suite.NewSet(cs, types.Int32(5), types.Int32(7), types.Int32(8)),
		suite.NewList(cs), types.NewString("wow"), // note that the dupe list chunk is skipped
	)
	suite.walkWorker(suite.storeAndRef(nested), 25)
}

type WalkTestSuite struct {
	WalkAllTestSuite
	shouldSeeItem types.Value
	shouldSee     types.Value
	mustSkip      types.Value
	deadValue     types.Value
}

func (suite *WalkTestSuite) SetupTest() {
	cs := chunks.NewMemoryStore()
	suite.cs = chunks.NewTestStore()
	suite.shouldSeeItem = types.NewString("zzz")
	suite.shouldSee = types.NewList(cs, suite.shouldSeeItem)
	suite.deadValue = types.UInt64(0xDEADBEEF)
	suite.mustSkip = types.NewList(cs, suite.deadValue)
}

func (suite *WalkTestSuite) TestStopWalkImmediately() {
	cs := chunks.NewMemoryStore()
	actual := 0
	SomeP(types.NewList(cs, types.NewSet(cs), types.NewList(cs)), suite.cs, func(v types.Value) bool {
		actual++
		return true
	}, 1)
	suite.Equal(1, actual)
}

func (suite *WalkTestSuite) skipWorker(composite types.Value) (reached []types.Value) {
	SomeP(composite, suite.cs, func(v types.Value) bool {
		suite.False(v.Equals(suite.deadValue), "Should never have reached %+v", suite.deadValue)
		reached = append(reached, v)
		return v.Equals(suite.mustSkip)
	}, 1)
	return
}

// Skipping a sub-tree must allow other items in the list to be processed.
func (suite *WalkTestSuite) SkipTestSkipListElement() {
	cs := chunks.NewMemoryStore()
	wholeList := types.NewList(cs, suite.mustSkip, suite.shouldSee, suite.shouldSee)
	reached := suite.skipWorker(wholeList)
	for _, v := range []types.Value{wholeList, suite.mustSkip, suite.shouldSee, suite.shouldSeeItem} {
		suite.Contains(reached, v, "Doesn't contain %+v", v)
	}
	suite.Len(reached, 6)
}

func (suite *WalkTestSuite) SkipTestSkipSetElement() {
	cs := chunks.NewMemoryStore()
	wholeSet := types.NewSet(cs, suite.mustSkip, suite.shouldSee).Insert(suite.shouldSee)
	reached := suite.skipWorker(wholeSet)
	for _, v := range []types.Value{wholeSet, suite.mustSkip, suite.shouldSee, suite.shouldSeeItem} {
		suite.Contains(reached, v, "Doesn't contain %+v", v)
	}
	suite.Len(reached, 4)
}

func (suite *WalkTestSuite) SkipTestSkipMapValue() {
	cs := chunks.NewMemoryStore()
	shouldAlsoSeeItem := types.NewString("Also good")
	shouldAlsoSee := types.NewSet(cs, shouldAlsoSeeItem)
	wholeMap := types.NewMap(cs, suite.shouldSee, suite.mustSkip, shouldAlsoSee, suite.shouldSee)
	reached := suite.skipWorker(wholeMap)
	for _, v := range []types.Value{wholeMap, suite.shouldSee, suite.shouldSeeItem, suite.mustSkip, shouldAlsoSee, shouldAlsoSeeItem} {
		suite.Contains(reached, v, "Doesn't contain %+v", v)
	}
	suite.Len(reached, 8)
}

func (suite *WalkTestSuite) SkipTestSkipMapKey() {
	cs := chunks.NewMemoryStore()
	wholeMap := types.NewMap(cs, suite.mustSkip, suite.shouldSee, suite.shouldSee, suite.shouldSee)
	reached := suite.skipWorker(wholeMap)
	for _, v := range []types.Value{wholeMap, suite.mustSkip, suite.shouldSee, suite.shouldSeeItem} {
		suite.Contains(reached, v, "Doesn't contain %+v", v)
	}
	suite.Len(reached, 8)
}
