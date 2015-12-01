// This file was generated by nomdl/codegen.

package common

import (
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __commonPackageInFile_incident_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.Type{
		types.MakeStructType("Incident",
			[]types.Field{
				types.Field{"ID", types.MakePrimitiveType(types.Int64Kind), false},
				types.Field{"Category", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"Description", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"DayOfWeek", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"Date", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"Time", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"PdDistrict", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"Resolution", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"Address", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"Geoposition", types.MakeType(ref.Parse("sha1-6d5e1c54214264058be9f61f4b4ece0368c8c678"), 0), false},
				types.Field{"PdID", types.MakePrimitiveType(types.StringKind), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{
		ref.Parse("sha1-6d5e1c54214264058be9f61f4b4ece0368c8c678"),
	})
	__commonPackageInFile_incident_CachedRef = types.RegisterPackage(&p)
}

// Incident

type Incident struct {
	_ID          int64
	_Category    string
	_Description string
	_DayOfWeek   string
	_Date        string
	_Time        string
	_PdDistrict  string
	_Resolution  string
	_Address     string
	_Geoposition Geoposition
	_PdID        string

	cs  chunks.ChunkStore
	ref *ref.Ref
}

func NewIncident(cs chunks.ChunkStore) Incident {
	return Incident{
		_ID:          int64(0),
		_Category:    "",
		_Description: "",
		_DayOfWeek:   "",
		_Date:        "",
		_Time:        "",
		_PdDistrict:  "",
		_Resolution:  "",
		_Address:     "",
		_Geoposition: NewGeoposition(cs),
		_PdID:        "",

		cs:  cs,
		ref: &ref.Ref{},
	}
}

type IncidentDef struct {
	ID          int64
	Category    string
	Description string
	DayOfWeek   string
	Date        string
	Time        string
	PdDistrict  string
	Resolution  string
	Address     string
	Geoposition GeopositionDef
	PdID        string
}

func (def IncidentDef) New(cs chunks.ChunkStore) Incident {
	return Incident{
		_ID:          def.ID,
		_Category:    def.Category,
		_Description: def.Description,
		_DayOfWeek:   def.DayOfWeek,
		_Date:        def.Date,
		_Time:        def.Time,
		_PdDistrict:  def.PdDistrict,
		_Resolution:  def.Resolution,
		_Address:     def.Address,
		_Geoposition: def.Geoposition.New(cs),
		_PdID:        def.PdID,
		cs:           cs,
		ref:          &ref.Ref{},
	}
}

func (s Incident) Def() (d IncidentDef) {
	d.ID = s._ID
	d.Category = s._Category
	d.Description = s._Description
	d.DayOfWeek = s._DayOfWeek
	d.Date = s._Date
	d.Time = s._Time
	d.PdDistrict = s._PdDistrict
	d.Resolution = s._Resolution
	d.Address = s._Address
	d.Geoposition = s._Geoposition.Def()
	d.PdID = s._PdID
	return
}

var __typeForIncident types.Type

func (m Incident) Type() types.Type {
	return __typeForIncident
}

func init() {
	__typeForIncident = types.MakeType(__commonPackageInFile_incident_CachedRef, 0)
	types.RegisterStruct(__typeForIncident, builderForIncident, readerForIncident)
}

func builderForIncident(cs chunks.ChunkStore, values []types.Value) types.Value {
	i := 0
	s := Incident{ref: &ref.Ref{}, cs: cs}
	s._ID = int64(values[i].(types.Int64))
	i++
	s._Category = values[i].(types.String).String()
	i++
	s._Description = values[i].(types.String).String()
	i++
	s._DayOfWeek = values[i].(types.String).String()
	i++
	s._Date = values[i].(types.String).String()
	i++
	s._Time = values[i].(types.String).String()
	i++
	s._PdDistrict = values[i].(types.String).String()
	i++
	s._Resolution = values[i].(types.String).String()
	i++
	s._Address = values[i].(types.String).String()
	i++
	s._Geoposition = values[i].(Geoposition)
	i++
	s._PdID = values[i].(types.String).String()
	i++
	return s
}

func readerForIncident(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(Incident)
	values = append(values, types.Int64(s._ID))
	values = append(values, types.NewString(s._Category))
	values = append(values, types.NewString(s._Description))
	values = append(values, types.NewString(s._DayOfWeek))
	values = append(values, types.NewString(s._Date))
	values = append(values, types.NewString(s._Time))
	values = append(values, types.NewString(s._PdDistrict))
	values = append(values, types.NewString(s._Resolution))
	values = append(values, types.NewString(s._Address))
	values = append(values, s._Geoposition)
	values = append(values, types.NewString(s._PdID))
	return values
}

func (s Incident) Equals(other types.Value) bool {
	return other != nil && __typeForIncident.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s Incident) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s Incident) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeForIncident.Chunks()...)
	chunks = append(chunks, s._Geoposition.Chunks()...)
	return
}

func (s Incident) ChildValues() (ret []types.Value) {
	ret = append(ret, types.Int64(s._ID))
	ret = append(ret, types.NewString(s._Category))
	ret = append(ret, types.NewString(s._Description))
	ret = append(ret, types.NewString(s._DayOfWeek))
	ret = append(ret, types.NewString(s._Date))
	ret = append(ret, types.NewString(s._Time))
	ret = append(ret, types.NewString(s._PdDistrict))
	ret = append(ret, types.NewString(s._Resolution))
	ret = append(ret, types.NewString(s._Address))
	ret = append(ret, s._Geoposition)
	ret = append(ret, types.NewString(s._PdID))
	return
}

func (s Incident) ID() int64 {
	return s._ID
}

func (s Incident) SetID(val int64) Incident {
	s._ID = val
	s.ref = &ref.Ref{}
	return s
}

func (s Incident) Category() string {
	return s._Category
}

func (s Incident) SetCategory(val string) Incident {
	s._Category = val
	s.ref = &ref.Ref{}
	return s
}

func (s Incident) Description() string {
	return s._Description
}

func (s Incident) SetDescription(val string) Incident {
	s._Description = val
	s.ref = &ref.Ref{}
	return s
}

func (s Incident) DayOfWeek() string {
	return s._DayOfWeek
}

func (s Incident) SetDayOfWeek(val string) Incident {
	s._DayOfWeek = val
	s.ref = &ref.Ref{}
	return s
}

func (s Incident) Date() string {
	return s._Date
}

func (s Incident) SetDate(val string) Incident {
	s._Date = val
	s.ref = &ref.Ref{}
	return s
}

func (s Incident) Time() string {
	return s._Time
}

func (s Incident) SetTime(val string) Incident {
	s._Time = val
	s.ref = &ref.Ref{}
	return s
}

func (s Incident) PdDistrict() string {
	return s._PdDistrict
}

func (s Incident) SetPdDistrict(val string) Incident {
	s._PdDistrict = val
	s.ref = &ref.Ref{}
	return s
}

func (s Incident) Resolution() string {
	return s._Resolution
}

func (s Incident) SetResolution(val string) Incident {
	s._Resolution = val
	s.ref = &ref.Ref{}
	return s
}

func (s Incident) Address() string {
	return s._Address
}

func (s Incident) SetAddress(val string) Incident {
	s._Address = val
	s.ref = &ref.Ref{}
	return s
}

func (s Incident) Geoposition() Geoposition {
	return s._Geoposition
}

func (s Incident) SetGeoposition(val Geoposition) Incident {
	s._Geoposition = val
	s.ref = &ref.Ref{}
	return s
}

func (s Incident) PdID() string {
	return s._PdID
}

func (s Incident) SetPdID(val string) Incident {
	s._PdID = val
	s.ref = &ref.Ref{}
	return s
}

// ListOfRefOfIncident

type ListOfRefOfIncident struct {
	l   types.List
	cs  chunks.ChunkStore
	ref *ref.Ref
}

func NewListOfRefOfIncident(cs chunks.ChunkStore) ListOfRefOfIncident {
	return ListOfRefOfIncident{types.NewTypedList(cs, __typeForListOfRefOfIncident), cs, &ref.Ref{}}
}

type ListOfRefOfIncidentDef []ref.Ref

func (def ListOfRefOfIncidentDef) New(cs chunks.ChunkStore) ListOfRefOfIncident {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = NewRefOfIncident(d)
	}
	return ListOfRefOfIncident{types.NewTypedList(cs, __typeForListOfRefOfIncident, l...), cs, &ref.Ref{}}
}

func (l ListOfRefOfIncident) Def() ListOfRefOfIncidentDef {
	d := make([]ref.Ref, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = l.l.Get(i).(RefOfIncident).TargetRef()
	}
	return d
}

func (l ListOfRefOfIncident) Equals(other types.Value) bool {
	return other != nil && __typeForListOfRefOfIncident.Equals(other.Type()) && l.Ref() == other.Ref()
}

func (l ListOfRefOfIncident) Ref() ref.Ref {
	return types.EnsureRef(l.ref, l)
}

func (l ListOfRefOfIncident) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, l.Type().Chunks()...)
	chunks = append(chunks, l.l.Chunks()...)
	return
}

func (l ListOfRefOfIncident) ChildValues() []types.Value {
	return append([]types.Value{}, l.l.ChildValues()...)
}

// A Noms Value that describes ListOfRefOfIncident.
var __typeForListOfRefOfIncident types.Type

func (m ListOfRefOfIncident) Type() types.Type {
	return __typeForListOfRefOfIncident
}

func init() {
	__typeForListOfRefOfIncident = types.MakeCompoundType(types.ListKind, types.MakeCompoundType(types.RefKind, types.MakeType(__commonPackageInFile_incident_CachedRef, 0)))
	types.RegisterValue(__typeForListOfRefOfIncident, builderForListOfRefOfIncident, readerForListOfRefOfIncident)
}

func builderForListOfRefOfIncident(cs chunks.ChunkStore, v types.Value) types.Value {
	return ListOfRefOfIncident{v.(types.List), cs, &ref.Ref{}}
}

func readerForListOfRefOfIncident(v types.Value) types.Value {
	return v.(ListOfRefOfIncident).l
}

func (l ListOfRefOfIncident) Len() uint64 {
	return l.l.Len()
}

func (l ListOfRefOfIncident) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfRefOfIncident) Get(i uint64) RefOfIncident {
	return l.l.Get(i).(RefOfIncident)
}

func (l ListOfRefOfIncident) Slice(idx uint64, end uint64) ListOfRefOfIncident {
	return ListOfRefOfIncident{l.l.Slice(idx, end), l.cs, &ref.Ref{}}
}

func (l ListOfRefOfIncident) Set(i uint64, val RefOfIncident) ListOfRefOfIncident {
	return ListOfRefOfIncident{l.l.Set(i, val), l.cs, &ref.Ref{}}
}

func (l ListOfRefOfIncident) Append(v ...RefOfIncident) ListOfRefOfIncident {
	return ListOfRefOfIncident{l.l.Append(l.fromElemSlice(v)...), l.cs, &ref.Ref{}}
}

func (l ListOfRefOfIncident) Insert(idx uint64, v ...RefOfIncident) ListOfRefOfIncident {
	return ListOfRefOfIncident{l.l.Insert(idx, l.fromElemSlice(v)...), l.cs, &ref.Ref{}}
}

func (l ListOfRefOfIncident) Remove(idx uint64, end uint64) ListOfRefOfIncident {
	return ListOfRefOfIncident{l.l.Remove(idx, end), l.cs, &ref.Ref{}}
}

func (l ListOfRefOfIncident) RemoveAt(idx uint64) ListOfRefOfIncident {
	return ListOfRefOfIncident{(l.l.RemoveAt(idx)), l.cs, &ref.Ref{}}
}

func (l ListOfRefOfIncident) fromElemSlice(p []RefOfIncident) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

type ListOfRefOfIncidentIterCallback func(v RefOfIncident, i uint64) (stop bool)

func (l ListOfRefOfIncident) Iter(cb ListOfRefOfIncidentIterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(v.(RefOfIncident), i)
	})
}

type ListOfRefOfIncidentIterAllCallback func(v RefOfIncident, i uint64)

func (l ListOfRefOfIncident) IterAll(cb ListOfRefOfIncidentIterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(v.(RefOfIncident), i)
	})
}

func (l ListOfRefOfIncident) IterAllP(concurrency int, cb ListOfRefOfIncidentIterAllCallback) {
	l.l.IterAllP(concurrency, func(v types.Value, i uint64) {
		cb(v.(RefOfIncident), i)
	})
}

type ListOfRefOfIncidentFilterCallback func(v RefOfIncident, i uint64) (keep bool)

func (l ListOfRefOfIncident) Filter(cb ListOfRefOfIncidentFilterCallback) ListOfRefOfIncident {
	out := l.l.Filter(func(v types.Value, i uint64) bool {
		return cb(v.(RefOfIncident), i)
	})
	return ListOfRefOfIncident{out, l.cs, &ref.Ref{}}
}

// RefOfIncident

type RefOfIncident struct {
	target ref.Ref
	ref    *ref.Ref
}

func NewRefOfIncident(target ref.Ref) RefOfIncident {
	return RefOfIncident{target, &ref.Ref{}}
}

func (r RefOfIncident) TargetRef() ref.Ref {
	return r.target
}

func (r RefOfIncident) Ref() ref.Ref {
	return types.EnsureRef(r.ref, r)
}

func (r RefOfIncident) Equals(other types.Value) bool {
	return other != nil && __typeForRefOfIncident.Equals(other.Type()) && r.Ref() == other.Ref()
}

func (r RefOfIncident) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, r.Type().Chunks()...)
	chunks = append(chunks, r.target)
	return
}

func (r RefOfIncident) ChildValues() []types.Value {
	return nil
}

// A Noms Value that describes RefOfIncident.
var __typeForRefOfIncident types.Type

func (m RefOfIncident) Type() types.Type {
	return __typeForRefOfIncident
}

func init() {
	__typeForRefOfIncident = types.MakeCompoundType(types.RefKind, types.MakeType(__commonPackageInFile_incident_CachedRef, 0))
	types.RegisterRef(__typeForRefOfIncident, builderForRefOfIncident)
}

func builderForRefOfIncident(r ref.Ref) types.Value {
	return NewRefOfIncident(r)
}

func (r RefOfIncident) TargetValue(cs chunks.ChunkStore) Incident {
	return types.ReadValue(r.target, cs).(Incident)
}

func (r RefOfIncident) SetTargetValue(val Incident, cs chunks.ChunkSink) RefOfIncident {
	return NewRefOfIncident(types.WriteValue(val, cs))
}
