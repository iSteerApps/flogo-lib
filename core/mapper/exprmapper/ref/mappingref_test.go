package ref

import (
	"encoding/json"
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/json/field"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMappingRef(t *testing.T) {

	mappingref := MappingRef{ref: "$activity[name].input.query.address.city"}

	resu, _ := data.GetResolutionDetails(mappingref.ref)

	v, _ := json.Marshal(resu)

	fmt.Println(string(v))

	mapField, err := field.ParseMappingField(resu.Path)
	//fields, err := mappingref.GetFields(mapField)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, []string{"query", "address", "city"}, mapField.Getfields())

}

//func TestMappingRefFuntion(t *testing.T) {
//
//	mappingref := MappingRef{ref: `concat($activity[activityname].input.query.address.city, "name"`}
//
//	_, err := mappingref.GetActivityId()
//
//	if err != nil {
//		t.Log(err)
//	}
//
//	_, err2 := mappingref.GetActivtyRootField()
//	if err2 != nil {
//		t.Log(err2)
//	}
//
//}
//
//func TestMappingRefExpression(t *testing.T) {
//
//	mappingref := MappingRef{ref: `$activity[activityname].input.query.address.zipcode > 77477`}
//
//	_, err := mappingref.GetActivityId()
//
//	if err != nil {
//		t.Log(err)
//	}
//
//	_, err2 := mappingref.GetActivtyRootField()
//	if err2 != nil {
//		t.Log(err2)
//	}
//	//
//	//assert.Equal(t, true, mappingref.IsExpression())
//	//
//	//assert.Equal(t, false, mappingref.IsFunction())
//
//}

//func TestMappingRef_GetPath(t *testing.T) {
//
//	mappingref := MappingRef{ref: `$activity[activityname].input.query.address.zipcode`}
//
//	name, _ := mappingref.GetFields()
//	assert.Equal(t, "query.address.zipcode", strings.Join(name.Fields, "."))
//	fmt.Println(name)
//
//	mappingref = MappingRef{ref: `input.query.address.zipcode`}
//
//	name, _ = mappingref.GetFields()
//	assert.Equal(t, "query.address.zipcode", strings.Join(name.Fields, "."))
//	fmt.Println(name)
//
//}

func TestGetActivtyRootField(t *testing.T) {

	mappingref := MappingRef{ref: `input[0].query.address.zipcode`}
	mapField, err := field.ParseMappingField(mappingref.ref)

	name, err := GetMapToAttrName(mapField)

	assert.Nil(t, err)
	assert.Equal(t, "input", name)

	mappingref = MappingRef{ref: `input.query.address.zipcode`}

	name, err = GetMapToAttrName(mapField)
	assert.Nil(t, err)
	assert.Equal(t, "input", name)

}

func TestGetFieldsMapFrom(t *testing.T) {
	ref := &MappingRef{ref: "Message[0].MessageId"}
	mapField, _ := field.ParseMappingField(ref.ref)
	mappingFields, err := GetMapToPathFields(mapField)
	assert.Nil(t, err)
	assert.Equal(t, []string{"[0]", "MessageId"}, mappingFields.Getfields())
	fmt.Println(fmt.Printf("%+v", mappingFields))

	//Special one
	ref = &MappingRef{ref: `name[0]["name&name"]`}
	mapField, _ = field.ParseMappingField(ref.ref)

	mappingFields, err = GetMapToPathFields(mapField)
	assert.Nil(t, err)
	assert.Equal(t, []string{"[0]", "name&name"}, mappingFields.Getfields())
	fmt.Println(fmt.Printf("%+v", mappingFields))

}

func TestGetFieldsMapFrom2(t *testing.T) {
	//Special one
	ref := &MappingRef{ref: `ReceiveSQSMessage.["x.y"][0]["name&name"]`}
	mapField, _ := field.ParseMappingField(ref.ref)

	s, err := GetMapToAttrName(mapField)
	assert.Equal(t, "ReceiveSQSMessage", s)

	mappingFields, err := GetMapToPathFields(mapField)
	assert.Nil(t, err)
	assert.Equal(t, []string{"x.y[0]", "name&name"}, mappingFields.Getfields())
	fmt.Println(fmt.Printf("%+v", mappingFields))
}

func TestGetFieldsMapTo(t *testing.T) {
	//Special one
	ref := &MappingRef{ref: `["x.y"][0]["name&name"]`}
	mapField, _ := field.ParseMappingField(ref.ref)

	s, err := GetMapToAttrName(mapField)
	assert.Nil(t, err)
	assert.Equal(t, "x.y", s)
}
