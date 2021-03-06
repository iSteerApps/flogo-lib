package mapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
)

func TestLiteralMapper(t *testing.T) {

	factory := GetFactory()

	mapping1 := &data.MappingDef{Type: data.MtLiteral, Value: "1", MapTo: "Simple"}
	mapping2 := &data.MappingDef{Type: data.MtLiteral, Value: 2, MapTo: "Obj.key"}
	mapping3 := &data.MappingDef{Type: data.MtLiteral, Value: 3, MapTo: "Array[2]"}
	mapping4 := &data.MappingDef{Type: data.MtLiteral, Value: "4", MapTo: "Params.paramKey"}

	mappings := []*data.MappingDef{mapping1, mapping2, mapping3, mapping4}

	mapper := factory.NewMapper(&data.MapperDef{Mappings: mappings}, nil)

	attr1, _ := data.NewAttribute("Simple", data.INTEGER, nil)
	attr2, _ := data.NewAttribute("Obj", data.OBJECT, nil)
	attr3, _ := data.NewAttribute("Array", data.ARRAY, nil)
	attr4, _ := data.NewAttribute("Params", data.PARAMS, nil)

	md := []*data.Attribute{attr1, attr2, attr3, attr4}
	outScope := data.NewFixedScope(md)

	objVal, _ := data.CoerceToObject("{\"key1\":5}")
	outScope.SetAttrValue("Obj", objVal)

	objVal, _ = data.CoerceToObject("{\"key1\":6}")
	outScope.SetAttrValue("Obj2", objVal)

	arrVal, _ := data.CoerceToArray("[1,6,3]")
	outScope.SetAttrValue("Array", arrVal)

	arrVal, _ = data.CoerceToArray("[7,8,9]")
	outScope.SetAttrValue("Array2", arrVal)

	paramVal, _ := data.CoerceToParams("{\"param1\":\"val\"}")
	outScope.SetAttrValue("Params", paramVal)

	paramVal, _ = data.CoerceToParams("{\"param1\":\"val2\"}")
	outScope.SetAttrValue("Params2", paramVal)

	err := mapper.Apply(nil, outScope)
	assert.Nil(t, err)

	resolver := &data.BasicResolver{}

	newVal, err := resolver.Resolve("Obj.key", outScope)
	assert.Nil(t, err)
	assert.Equal(t, 2, newVal)

	newVal, err = resolver.Resolve("Array[2]", outScope)
	assert.Nil(t, err)
	assert.Equal(t, 3, newVal)

	newVal, err = resolver.Resolve("Params.paramKey", outScope)
	assert.Nil(t, err)
	assert.Equal(t, "4", newVal)
}

func TestAssignMapper(t *testing.T) {

	factory := GetFactory()

	mapping1 := &data.MappingDef{Type: data.MtAssign, Value: "SimpleI", MapTo: "SimpleO"}
	mapping2 := &data.MappingDef{Type: data.MtAssign, Value: "ObjI.key", MapTo: "ObjO.key"}
	mapping3 := &data.MappingDef{Type: data.MtAssign, Value: "ArrayI[2]", MapTo: "ArrayO[2]"}
	mapping4 := &data.MappingDef{Type: data.MtAssign, Value: "ParamsI.paramKey", MapTo: "ParamsO.paramKey"}

	mappings := []*data.MappingDef{mapping1, mapping2, mapping3, mapping4}

	mapper := factory.NewMapper(&data.MapperDef{Mappings: mappings}, nil)

	attrI1, _ := data.NewAttribute("SimpleI", data.INTEGER, nil)
	attrI2, _ := data.NewAttribute("ObjI", data.OBJECT, nil)
	attrI3, _ := data.NewAttribute("ArrayI", data.ARRAY, nil)
	attrI4, _ := data.NewAttribute("ParamsI", data.PARAMS, nil)

	mdI := []*data.Attribute{attrI1, attrI2, attrI3, attrI4}
	inScope := data.NewFixedScope(mdI)

	attrO1, _ := data.NewAttribute("SimpleO", data.INTEGER, nil)
	attrO2, _ := data.NewAttribute("ObjO", data.OBJECT, nil)
	attrO3, _ := data.NewAttribute("ArrayO", data.ARRAY, nil)
	attrO4, _ := data.NewAttribute("ParamsO", data.PARAMS, nil)

	mdO := []*data.Attribute{attrO1, attrO2, attrO3, attrO4}
	outScope := data.NewFixedScope(mdO)

	inScope.SetAttrValue("SimpleI", 1)

	objVal, _ := data.CoerceToObject("{\"key\":1}")
	inScope.SetAttrValue("ObjI", objVal)

	arrVal, _ := data.CoerceToArray("[1,2,3]")
	inScope.SetAttrValue("ArrayI", arrVal)

	paramVal, _ := data.CoerceToParams("{\"paramKey\":\"val1\"}")
	inScope.SetAttrValue("ParamsI", paramVal)

	objVal, _ = data.CoerceToObject("{\"key1\":5}")
	outScope.SetAttrValue("ObjO", objVal)

	arrVal, _ = data.CoerceToArray("[4,5,6]")
	outScope.SetAttrValue("ArrayO", arrVal)

	paramVal, _ = data.CoerceToParams("{\"param1\":\"val\"}")
	outScope.SetAttrValue("ParamsO", paramVal)

	err := mapper.Apply(inScope, outScope)
	assert.Nil(t, err)

	resolver := &data.BasicResolver{}

	newVal, err := resolver.Resolve("ObjO.key", outScope)
	assert.Nil(t, err)
	assert.Equal(t, 1.0, newVal)

	newVal, err = resolver.Resolve("ArrayO[2]", outScope)
	assert.Nil(t, err)
	assert.Equal(t, 3.0, newVal)

	newVal, err = resolver.Resolve("ParamsO.paramKey", outScope)
	assert.Nil(t, err)
	assert.Equal(t, "val1", newVal)
}
