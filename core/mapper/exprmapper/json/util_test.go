package json

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestStruct struct {
	Id      string `json:"id"`
	name    string `json:"name"`
	Zipcode string `json:"zipcode"`
	Test    string `json:",omitempty`
}

func TestGetFieldByName(t *testing.T) {
	test := &TestStruct{Id: "idssss", name: "dddddd", Zipcode: "55555", Test: "ddd"}
	field, _ := GetFieldByName(test, "id")
	fmt.Println(field.Interface())
}

func TestIsMapperableType(t *testing.T) {
	assert.True(t, IsMapperableType(map[string]interface{}{}))
	assert.True(t, IsMapperableType(map[string]string{}))
	assert.True(t, IsMapperableType([]int{1, 2, 3, 4}))
	assert.True(t, IsMapperableType([]interface{}{"string"}))
	assert.True(t, IsMapperableType([]interface{}{map[string]interface{}{}}))

	str := struct {
		ID   string
		Name string
	}{
		ID:   "11111",
		Name: "22222",
	}

	assert.False(t, IsMapperableType(&str))
	assert.False(t, IsMapperableType(str))

	array := []interface{}{str}
	array2 := []interface{}{&str}
	assert.False(t, IsMapperableType(array))
	assert.False(t, IsMapperableType(array2))

}
