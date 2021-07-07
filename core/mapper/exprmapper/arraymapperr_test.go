package exprmapper

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//7. array mapping with leaf field
func TestArrayMapping(t *testing.T) {
	mappingValue := `{
    "fields": [
        {
            "from": "$.street",
            "to": "$.street",
            "type": "primitive"
        },
        {
            "from": "$.zipcode",
            "to": "$.zipcode",
            "type": "primitive"
        },
        {
            "from": "$.state",
            "to": "$.state",
            "type": "primitive"
        }
    ],
    "from": "$activity[a1].field.addresses",
    "to": "field.addresses",
    "type": "foreach"
}`

	arrayData := `{
    "person": "name",
    "addresses": [
        {
            "street": "street",
            "zipcode": 77479,
            "state": "tx"
        }
    ]
}`

	array, err := ParseArrayMapping(mappingValue)
	assert.Nil(t, array.Validate())

	inputScope := GetObjectFieldScope("_A.a1.field", arrayData)
	outputScope := GetObjectFieldScope("field", "")
	err = array.DoArrayMapping(inputScope, outputScope, GetTestResolver())
	assert.Nil(t, err)

	arr, ok := outputScope.GetAttr("field")
	assert.True(t, ok)
	assert.Equal(t, "street", arr.Value().(map[string]interface{})["addresses"].([]interface{})[0].(map[string]interface{})["street"])
	assert.Equal(t, float64(77479), arr.Value().(map[string]interface{})["addresses"].([]interface{})[0].(map[string]interface{})["zipcode"])
	assert.Equal(t, "tx", arr.Value().(map[string]interface{})["addresses"].([]interface{})[0].(map[string]interface{})["state"])

}

//8. array mappping with static function and leaf field
func TestArrayMappingWithFunction(t *testing.T) {
	mappingValue := `{
    "fields": [
        {
            "from": "string.concat(\"this stree name: \", $.street)",
            "to": "$.street",
            "type": "primitive"
        },
        {
            "from": "string.concat(\"The zipcode is: \",$.zipcode)",
            "to": "$.zipcode",
            "type": "primitive"
        },
        {
            "from": "$.state",
            "to": "$.state",
            "type": "primitive"
        }
    ],
    "from": "$activity[a1].field.addresses",
    "to": "field.addresses",
    "type": "foreach"
}`

	arrayData := `{
    "person": "name",
    "addresses": [
        {
            "street": "street",
            "zipcode": 77479,
            "state": "tx"
        }
    ]
}`

	array, err := ParseArrayMapping(mappingValue)
	assert.Nil(t, err)
	assert.Nil(t, array.Validate())

	inputScope := GetObjectFieldScope("_A.a1.field", arrayData)
	outputScope := GetObjectFieldScope("field", "")
	err = array.DoArrayMapping(inputScope, outputScope, GetTestResolver())
	assert.Nil(t, err)

	arr, ok := outputScope.GetAttr("field")
	assert.True(t, ok)
	assert.Equal(t, "this stree name: street", arr.Value().(map[string]interface{})["addresses"].([]interface{})[0].(map[string]interface{})["street"])
	assert.Equal(t, "The zipcode is: 77479", arr.Value().(map[string]interface{})["addresses"].([]interface{})[0].(map[string]interface{})["zipcode"])
	assert.Equal(t, "tx", arr.Value().(map[string]interface{})["addresses"].([]interface{})[0].(map[string]interface{})["state"])

}

//9. array mapping with other activity output
func TestArrayMappingWithUpstreamingOutput(t *testing.T) {
	mappingValue := `{
    "fields": [
        {
            "from": "string.concat(\"this stree name: \", $activity[a1].field.person)",
            "to": "$.street",
            "type": "primitive"
        },
        {
            "from": "string.concat(\"The zipcode is: \",$.zipcode)",
            "to": "$.zipcode",
            "type": "primitive"
        },
        {
            "from": "$.state",
            "to": "$.state",
            "type": "primitive"
        }
    ],
    "from": "$activity[a1].field.addresses",
    "to": "field.addresses",
    "type": "foreach"
}`

	arrayData := `{
    "person": "name",
    "addresses": [
        {
            "street": "street",
            "zipcode": 77479,
            "state": "tx"
        }
    ]
}`

	array, err := ParseArrayMapping(mappingValue)
	assert.Nil(t, err)
	assert.Nil(t, array.Validate())

	inputScope := GetObjectFieldScope("_A.a1.field", arrayData)
	outputScope := GetObjectFieldScope("field", "")
	err = array.DoArrayMapping(inputScope, outputScope, GetTestResolver())
	assert.Nil(t, err)

	arr, ok := outputScope.GetAttr("field")
	assert.True(t, ok)
	assert.Equal(t, "this stree name: name", arr.Value().(map[string]interface{})["addresses"].([]interface{})[0].(map[string]interface{})["street"])
	assert.Equal(t, "The zipcode is: 77479", arr.Value().(map[string]interface{})["addresses"].([]interface{})[0].(map[string]interface{})["zipcode"])
	assert.Equal(t, "tx", arr.Value().(map[string]interface{})["addresses"].([]interface{})[0].(map[string]interface{})["state"])

}

//9. array mapping with other activity output
func TestArrayMappingWithNest(t *testing.T) {
	mappingValue := `{
    "fields": [
        {
            "from": "string.concat(\"this stree name: \", $activity[a1].field.person)",
            "to": "$.street",
            "type": "primitive"
        },
        {
            "from": "string.concat(\"The zipcode is: \",$.zipcode)",
            "to": "$.zipcode",
            "type": "primitive"
        },
        {
            "from": "$.state",
            "to": "$.state",
            "type": "primitive"
        },
		{
    		"from": "$.array",
    		"to": "$.array",
            "type": "foreach",
			"fields":[
				{
           			 "from": "$.field1",
           			 "to": "$.tofield1",
           			 "type": "assign"
        		},
				{
            		"from": "$.field2",
					"to": "$.tofield2",
            		"type": "assign"
        		},
				{
            		"from": "wangzai",
					"to": "$.tofield3",
            		"type": "assign"
        		}
			]

		}
    ],
    "from": "$activity[a1].field.addresses",
    "to": "field.addresses",
    "type": "foreach"
}`

	arrayData := `{
    "person": "name",
    "addresses": [
        {
            "street": "street",
            "zipcode": 77479,
            "state": "tx",
			"array":[
				{
					"field1":"field1value",
					"field2":"field2value",
					"field3":"field3value"
				}
			]
        }
    ]
}`

	array, err := ParseArrayMapping(mappingValue)
	assert.Nil(t, err)
	assert.Nil(t, array.Validate())

	inputScope := GetObjectFieldScope("_A.a1.field", arrayData)
	outputScope := GetObjectFieldScope("field", "")
	err = array.DoArrayMapping(inputScope, outputScope, GetTestResolver())
	assert.Nil(t, err)

	arr, ok := outputScope.GetAttr("field")
	assert.True(t, ok)
	v, _ := json.Marshal(arr.Value())
	fmt.Println(string(v))
	assert.Equal(t, "this stree name: name", arr.Value().(map[string]interface{})["addresses"].([]interface{})[0].(map[string]interface{})["street"])
	assert.Equal(t, "The zipcode is: 77479", arr.Value().(map[string]interface{})["addresses"].([]interface{})[0].(map[string]interface{})["zipcode"])
	assert.Equal(t, "tx", arr.Value().(map[string]interface{})["addresses"].([]interface{})[0].(map[string]interface{})["state"])

	assert.Equal(t, "field1value", arr.Value().(map[string]interface{})["addresses"].([]interface{})[0].(map[string]interface{})["array"].([]interface{})[0].(map[string]interface{})["tofield1"])
	assert.Equal(t, "field2value", arr.Value().(map[string]interface{})["addresses"].([]interface{})[0].(map[string]interface{})["array"].([]interface{})[0].(map[string]interface{})["tofield2"])
	assert.Equal(t, "wangzai", arr.Value().(map[string]interface{})["addresses"].([]interface{})[0].(map[string]interface{})["array"].([]interface{})[0].(map[string]interface{})["tofield3"])

}
