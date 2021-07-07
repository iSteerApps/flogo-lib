package json

import (
	"encoding/json"
	"sync"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/json/field"

	"fmt"
	"github.com/stretchr/testify/assert"
)

//1. without array
//2. root array, array with path field, array with
func TestSetArrayObject(t *testing.T) {
	mappingField := field.NewMappingField([]string{"City[0]", "Array[1]", "id"})
	v, err := SetStringValue("setvalue", jsonData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, v)

	d, err := getValue(v, []string{"City[0]", "Array[1]", "id"})
	assert.Nil(t, err)
	assert.Equal(t, "setvalue", d)

}

func TestSetArrayObjectEMpty(t *testing.T) {
	mappingField := field.NewMappingField([]string{"City[0]", "Array[1]", "id"})
	v, err := SetStringValue("setvalue", "{}", mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	printvalue(v)
	d, err := getValue(v, []string{"City[0]", "Array[1]", "id"})
	assert.Nil(t, err)
	assert.Equal(t, "setvalue", d)

}

func printvalue(in interface{}) {
	v, _ := json.Marshal(in)
	fmt.Println(string(v))
}

func TestSetRootChildArray(t *testing.T) {
	mappingField := field.NewMappingField([]string{"Emails[0]"})
	v, err := SetStringValue("test-cases@gmail.com", jsonData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, v)

	d, err := getValue(v, []string{"Emails[0]"})
	assert.Nil(t, err)
	assert.Equal(t, "test-cases@gmail.com", d)
}

func TestSetRootArray(t *testing.T) {
	mappingField := field.NewMappingField([]string{"[0]", "ss"})
	v, err := SetFieldValue("test-cases@gmail.com", "{}", mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, v)

	d, err := getValue(v, []string{"[0]", "ss"})
	assert.Nil(t, err)
	assert.Equal(t, "test-cases@gmail.com", d)

}

func TestSetObject(t *testing.T) {
	mappingField := field.NewMappingField([]string{"ZipCode"})
	v, err := SetStringValue("77479", jsonData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, v)

	d, err := getValue(v, []string{"ZipCode"})
	assert.Nil(t, err)
	assert.Equal(t, "77479", d)
}

func TestSetEmptyField(t *testing.T) {
	mappingField := field.NewMappingField([]string{"ZipCode"})
	jsond := "{}"
	v, err := SetStringValue("77479", jsond, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	printvalue(v)
	d, err := getValue(v, []string{"ZipCode"})
	assert.Nil(t, err)
	assert.Equal(t, "77479", d)
}

func TestSetEmptyField4(t *testing.T) {
	jsond := "{}"
	mappingField := field.NewMappingField([]string{"ZipCode[1]"})
	v, err := SetStringValue("77479", jsond, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, v)
	d, err := getValue(v, []string{"ZipCode[1]"})
	assert.Nil(t, err)
	assert.Equal(t, "77479", d)
}

func TestSetEmptyField5(t *testing.T) {
	jsond := "{}"
	mappingField := field.NewMappingField([]string{"ZipCode[1]"})
	v, err := SetStringValue("77479", jsond, mappingField)

	mappingField = field.NewMappingField([]string{"ZipCode[0]"})
	v2, err := SetFieldValue("77479", v, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, v2)
	d, err := getValue(v, []string{"ZipCode[0]"})
	assert.Nil(t, err)
	assert.Equal(t, "77479", d)

	d, err = getValue(v, []string{"ZipCode[1]"})
	assert.Nil(t, err)
	assert.Equal(t, "77479", d)
}

func TestSetEmptyNestField1(t *testing.T) {
	jsond := "{}"
	mappingField := field.NewMappingField([]string{"pet", "photoUrls[0]"})
	v, err := SetStringValue("url", jsond, mappingField)

	mappingField = field.NewMappingField([]string{"pet", "photoUrls[1]"})

	v, err = SetFieldValue("url2", v, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, v)

	d, err := getValue(v, []string{"pet", "photoUrls[0]"})
	assert.Nil(t, err)
	assert.Equal(t, "url", d)
	d, err = getValue(v, []string{"pet", "photoUrls[1]"})
	assert.Nil(t, err)
	assert.Equal(t, "url2", d)
}

func TestNameWithSpace(t *testing.T) {
	jsond := "{}"
	mappingField := field.NewMappingField([]string{"pet name", "photoUrls[0]"})

	v, err := SetFieldValue("url", jsond, mappingField)
	mappingField = field.NewMappingField([]string{"pet name", "photoUrls[1]"})

	v, err = SetFieldValue("url2", v, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, v)

	d, err := getValue(v, []string{"pet name", "photoUrls[0]"})
	assert.Nil(t, err)
	assert.Equal(t, "url", d)
	d, err = getValue(v, []string{"pet name", "photoUrls[1]"})
	assert.Nil(t, err)
	assert.Equal(t, "url2", d)

}

func TestNameNest2(t *testing.T) {
	jsond := "{}"
	mappingField := field.NewMappingField([]string{"input", "Account", "records[0]", "ID"})
	v, err := SetFieldValue("id22", jsond, mappingField)

	mappingField = field.NewMappingField([]string{"input", "Account", "records[0]", "Name"})

	v, err = SetFieldValue("namesssss", v, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, v)

	d, err := getValue(v, []string{"input", "Account", "records[0]", "ID"})
	assert.Nil(t, err)
	assert.Equal(t, "id22", d)
	d, err = getValue(v, []string{"input", "Account", "records[0]", "Name"})
	assert.Nil(t, err)
	assert.Equal(t, "namesssss", d)
}

func TestNameSameLevel(t *testing.T) {
	jsond := "{}"
	mappingField := field.NewMappingField([]string{"input", "Account", "ID"})
	v, err := SetFieldValue("id", jsond, mappingField)

	mappingField = field.NewMappingField([]string{"input", "Account", "Name"})

	v, err = SetFieldValue("namesssss", v, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, v)

	d, err := getValue(v, []string{"input", "Account", "ID"})
	assert.Nil(t, err)
	assert.Equal(t, "id", d)
	d, err = getValue(v, []string{"input", "Account", "Name"})
	assert.Nil(t, err)
	assert.Equal(t, "namesssss", d)

}

func TestNameWithTag(t *testing.T) {
	jsond := "{}"
	mappingField := field.NewMappingField([]string{"pet", "pet name", "photo	Urls[0]"})
	v, err := SetFieldValue("url", jsond, mappingField)

	mappingField = field.NewMappingField([]string{"pet", "pet name", "photo	Urls[1]"})

	v, err = SetFieldValue("url2", v, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, v)

	d, err := getValue(v, []string{"pet", "pet name", "photo	Urls[0]"})
	assert.Nil(t, err)
	assert.Equal(t, "url", d)
	d, err = getValue(v, []string{"pet", "pet name", "photo	Urls[1]"})
	assert.Nil(t, err)
	assert.Equal(t, "url2", d)
}

func TestSetEmptyNestField(t *testing.T) {
	jsond := "{}"
	mappingField := field.NewMappingField([]string{"Response", "Pet", "Tags[0]", "Name"})
	v, err := SetStringValue("tagID", jsond, mappingField)

	mappingField = field.NewMappingField([]string{"Response", "Pet", "Tags[1]", "Name"})

	v, err = SetFieldValue("tagID2", v, mappingField)

	assert.Nil(t, err)
	assert.NotNil(t, v)
	d, err := getValue(v, []string{"Response", "Pet", "Tags[0]", "Name"})
	assert.Nil(t, err)
	assert.Equal(t, "tagID", d)
	d, err = getValue(v, []string{"Response", "Pet", "Tags[1]", "Name"})
	assert.Nil(t, err)
	assert.Equal(t, "tagID2", d)
}

func TestConcurrentSet(t *testing.T) {
	w := sync.WaitGroup{}
	var recovered interface{}
	//Create factory

	for r := 0; r < 100000; r++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			defer func() {
				if r := recover(); r != nil {
					recovered = r
				}
			}()
			jsond := "{}"
			mappingField := field.NewMappingField([]string{"pet name", "photoUrls[0]"})
			v, err := SetFieldValue("url", jsond, mappingField)

			mappingField = field.NewMappingField([]string{"pet name", "photoUrls[1]"})

			v, err = SetFieldValue("url2", v, mappingField)
			assert.Nil(t, err)
			assert.NotNil(t, v)
		}(r)

	}
	w.Wait()
	assert.Nil(t, recovered)
}

func getValue(value interface{}, fields []string) (interface{}, error) {
	mapField := field.NewMappingField(fields)
	return GetFieldValue(value, mapField)
}

func TestSetSpecialObjectField(t *testing.T) {
	mappingField := field.NewMappingField([]string{"Object", "Maps3", "dd*cc", "y.x", "d.d"})

	value, err := SetStringValue("lixi", "{}", mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	log.Info(value)
	v, _ := json.Marshal(value)
	assert.Equal(t, `{"Object":{"Maps3":{"dd*cc":{"y.x":{"d.d":"lixi"}}}}}`, string(v))
}

func TestSetSpecialArrayField2(t *testing.T) {
	mappingField := field.NewMappingField([]string{"Object", "Maps3[0]", "dd*cc"})

	value, err := SetStringValue("lixi", "{}", mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	log.Info(value)
	v, _ := json.Marshal(value)
	fmt.Println(string(v))
	assert.Equal(t, `{"Object":{"Maps3":[{"dd*cc":"lixi"}]}}`, string(v))
}

func TestSetSpecialArrayFieldMultipleLEvel(t *testing.T) {
	// path := `Object.Maps3["dd.cc"][0]["y.x"][d.d].name`
	mappingField := field.NewMappingField([]string{"Object", "Maps3", "dd.cc[0]", "y.x", "d.d", "name"})
	value, err := SetStringValue("lixi", "{}", mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	log.Info(value)
	v, _ := json.Marshal(value)
	fmt.Println(string(v))
	assert.Equal(t, `{"Object":{"Maps3":{"dd.cc":[{"y.x":{"d.d":{"name":"lixi"}}}]}}}`, string(v))
}

func TestSetArrayRootOnly(t *testing.T) {
	mappingField := field.NewMappingField([]string{"[0]"})
	value, err := SetStringValue("lixi", "{}", mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	v, err := getValue(value, []string{"[0]"})
	assert.Nil(t, err)
	assert.Equal(t, "lixi", v)

	mappingField = field.NewMappingField([]string{"[0]", "tt"})
	value, err = SetStringValue("ssssss", "{}", mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	v, err = getValue(value, []string{"[0]", "tt"})
	assert.Nil(t, err)
	assert.Equal(t, "ssssss", v)
}

func TestSetStructValue(t *testing.T) {
	value := struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		IntV   int    `json:"int_v"`
		Int64V int64  `json:"int_64"`
	}{
		ID:     "12222",
		Name:   "name",
		Int64V: int64(123),
		IntV:   int(12),
	}

	mappingField := field.NewMappingField([]string{"id"})

	_, err := SetFieldValue("lixingwangid", &value, mappingField)
	assert.Nil(t, err)

	assert.Equal(t, "lixingwangid", value.ID)

}
