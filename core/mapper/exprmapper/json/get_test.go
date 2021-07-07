package json

import (
	"encoding/json"
	"strconv"
	"sync"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/json/field"

	"github.com/stretchr/testify/assert"
)

var jsonData = `{
    "City": [
        {
            "Array": [
                {
                    "id": "11111"
                },
                {
                    "id": "2222"
                }
            ],
            "InUS": true,
            "Name": "Sugar Land",
            "Park": {
                "Emails": null,
                "Location": "location",
                "Maps": {
                    "bb": "bb",
                    "cc": "cc",
                    "dd": "dd"
                },
                "Name": "Name"
            }
        }
    ],
    "Emails": [
        "123@123.com",
        "456@456.com"
    ],
    "Id": 1234,
    "In": "string222",
    "Maps": {
        "bb": "bb",
        "cc": "cc",
        "dd": "dd"
    },
    "State": "Taxes",
    "Streat": "311 wind st",
    "ZipCode": "77477",
    "hello world":"CHINA",
    "tag  world":"CHINA"
}`

func TestRootChildArray(t *testing.T) {
	mappingField := field.NewMappingField([]string{"City[0]"})
	value, err := GetFieldValue(jsonData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
}

func TestRoot(t *testing.T) {
	mappingField := field.NewMappingField([]string{"City"})
	value, err := GetFieldValue(jsonData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
}

func TestGetFieldWithSpaces(t *testing.T) {
	mappingField := field.NewMappingField([]string{"hello world"})
	value, err := GetFieldValue(jsonData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, "CHINA", value)
}

func TestGetFieldWithTag(t *testing.T) {
	mappingField := field.NewMappingField([]string{"tag  world"})
	value, err := GetFieldValue(jsonData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, "CHINA", value)
}

func TestGetArray(t *testing.T) {
	mappingField := field.NewMappingField([]string{"Emails[0]"})
	value, err := GetFieldValue(jsonData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, "123@123.com", value)

}

func TestMultipleLevelArray(t *testing.T) {
	mappingField := field.NewMappingField([]string{"City[0]", "Array[1]", "id"})
	value, err := GetFieldValue(jsonData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, "2222", value)
}

func TestMultipleLevelObject(t *testing.T) {
	mappingField := field.NewMappingField([]string{"City[0]", "Park", "Maps", "bb"})
	value, err := GetFieldValue(jsonData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, "bb", value)
}

func TestID(t *testing.T) {
	mappingField := field.NewMappingField([]string{"Id"})
	value, err := GetFieldValue(jsonData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, float64(1234), value)
}

func TestGetFieldValue(t *testing.T) {
	account := `{
    "Account": {
        "records": [
            {
                "AccountNumber": "12356",
                "AccountSource": "Test Source",
                "Active__c": "Yes",
                "AnnualRevenue": "324556",
                "BillingCity": "Palo Alto",
                "BillingCountry": "USA",
                "BillingGeocodeAccuracy": null,
                "BillingLatitude": null,
                "BillingLongitude": null,
                "BillingPostalCode": "94207",
                "BillingState": "California",
                "BillingStreet": "3330 hillview ave",
                "CleanStatus": "Pending",
                "CustomerPriority__c": "High",
                "Description": "Sample Description for the account",
                "DunsNumber": "32653",
                "Fax": "345272",
                "Industry": "Engineering",
                "Jigsaw": "Test",
                "NaicsCode": "34583",
                "NaicsDesc": "Test Description",
                "Name": "may24_a",
                "Ownership": "Private",
                "ParentId": null,
                "Phone": "1234567890",
                "Rating": "Warm",
                "SLAExpirationDate__c": "2017-08-27",
                "SLA__c": "23453",
                "ShippingCity": "San Francisco",
                "ShippingCountry": "USA",
                "ShippingGeocodeAccuracy": null,
                "ShippingLatitude": null,
                "ShippingLongitude": null,
                "ShippingPostalCode": 45692,
                "ShippingState": "California",
                "ShippingStreet": "1234 Hillview Ave",
                "Sic": "Gold",
                "SicDesc": null,
                "Site": "www.example2.com",
                "TickerSymbol": null,
                "Tradestyle": "Regular",
                "Type": "Custumer-Direct",
                "UpsellOpportunity__c": "Yes",
                "Website": "www.example.com",
                "YearStarted": "2015"
            }
        ]
    }
}
`
	mappingField := field.NewMappingField([]string{"Account", "records[0]", "Name"})
	value, err := GetFieldValue(account, mappingField)
	log.Infof("Value:%s", value)

	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, "may24_a", value)

}

func TestConcurrentGetk(t *testing.T) {
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
			mappingField := field.NewMappingField([]string{"City[0]", "Park", "Maps", "bb"})
			value, err := GetFieldValue(jsonData, mappingField)
			assert.Nil(t, err)
			assert.NotNil(t, value)
		}(r)

	}
	w.Wait()
	assert.Nil(t, recovered)
}

func TestRootArray(t *testing.T) {
	jsonArray := `[
    {
        "Body": "test from WI",
        "MD5OfBody": "ec7d5c27e25bcd3d6a65362b71bd0525",
        "MD5OfMessageAttributes": "50df80e5fea57210bb8167abfd053899",
        "MessageAttributes": {
            "MA1": "test"
        },
        "MessageId": "1c0483d9-8166-4df0-be9f-cd03177a38c6",
        "ReceiptHandle": "AQEBE6elNqdJKrTz5A2X/gQJETxPdtJgAktTAuT4pvBTjQgnJpSEPhfMI08fHCMrEX6ILD0fTY2FVPCCJ8LfMvAxp+LO2/Bsi1uZhUyesFoj11Y/4jvdYSCQhqWEuAI1q1pxpSj2d2QbL5SUlX979ZG+Abv/IYeDvPO8nyuZ0IWgVhZWaGcoOwADvj3mNJZ9XJh8mS3vL8EQlUO6dhIRn9PxVet2fGRmm3iY1YI4N7bZsw9nxIqIYgl5kfuBNegSRcrrTOb6u9vTnHK2uiiCwJi+Io6WNGuJGF4fyFi3skk/AvCS7fjl+4MFqoHKsm1nR06Rel7017m0/Dg5KaOJCRAJ92gV4iuUMynG1WfmELMMg/sS19hrNvcgdKW5Vd3Snn/oNcoP2Ebb7CQA08XzVcoO0maVt2KqUWgvqf0DDxVArEE="
    }
]`

	mappingField := field.NewMappingField([]string{"[0]", "MessageId"})
	value, err := GetFieldValue(jsonArray, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, "1c0483d9-8166-4df0-be9f-cd03177a38c6", value)

}

func TestRootArrayInvalid(t *testing.T) {
	jsonArray := `[
    {
        "Body": "test from WI",
        "MD5OfBody": "ec7d5c27e25bcd3d6a65362b71bd0525",
        "MD5OfMessageAttributes": "50df80e5fea57210bb8167abfd053899",
        "MessageAttributes": {
            "MA1": "test"
        },
        "MessageId": "1c0483d9-8166-4df0-be9f-cd03177a38c6",
        "ReceiptHandle": "AQEBE6elNqdJKrTz5A2X/gQJETxPdtJgAktTAuT4pvBTjQgnJpSEPhfMI08fHCMrEX6ILD0fTY2FVPCCJ8LfMvAxp+LO2/Bsi1uZhUyesFoj11Y/4jvdYSCQhqWEuAI1q1pxpSj2d2QbL5SUlX979ZG+Abv/IYeDvPO8nyuZ0IWgVhZWaGcoOwADvj3mNJZ9XJh8mS3vL8EQlUO6dhIRn9PxVet2fGRmm3iY1YI4N7bZsw9nxIqIYgl5kfuBNegSRcrrTOb6u9vTnHK2uiiCwJi+Io6WNGuJGF4fyFi3skk/AvCS7fjl+4MFqoHKsm1nR06Rel7017m0/Dg5KaOJCRAJ92gV4iuUMynG1WfmELMMg/sS19hrNvcgdKW5Vd3Snn/oNcoP2Ebb7CQA08XzVcoO0maVt2KqUWgvqf0DDxVArEE="
    },
	    {
        "Body": "test from WI2",
        "MD5OfBody": "ec7d5c27e25bcd33d6a65362b71bd0525",
        "MD5OfMessageAttributes": "50df80e5fea57210bb8167abfd053899",
        "MessageAttributes": {
            "MA1": "test"
        },
        "MessageId": "==1c04833d9-8166-4df0-be9f-cd03177a38c6",
        "ReceiptHandle": "AQE3BE6elNqdJKrTz5A2X/gQJETxPdtJgAktTAuT4pvBTjQgnJpSEPhfMI08fHCMrEX6ILD0fTY2FVPCCJ8LfMvAxp+LO2/Bsi1uZhUyesFoj11Y/4jvdYSCQhqWEuAI1q1pxpSj2d2QbL5SUlX979ZG+Abv/IYeDvPO8nyuZ0IWgVhZWaGcoOwADvj3mNJZ9XJh8mS3vL8EQlUO6dhIRn9PxVet2fGRmm3iY1YI4N7bZsw9nxIqIYgl5kfuBNegSRcrrTOb6u9vTnHK2uiiCwJi+Io6WNGuJGF4fyFi3skk/AvCS7fjl+4MFqoHKsm1nR06Rel7017m0/Dg5KaOJCRAJ92gV4iuUMynG1WfmELMMg/sS19hrNvcgdKW5Vd3Snn/oNcoP2Ebb7CQA08XzVcoO0maVt2KqUWgvqf0DDxVArEE="
    }
]`

	mappingField := field.NewMappingField([]string{"[0]", "MessageId[0]"})
	value, err := GetFieldValue(jsonArray, mappingField)
	assert.NotNil(t, err)
	assert.Nil(t, nil, value)

}

func TestGetStructValue(t *testing.T) {
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

	v, err := GetFieldValue(value, mappingField)
	assert.Nil(t, err)
	assert.Equal(t, "12222", v)

	testMap := make(map[string]string)
	testMap["id"] = "id"
	testMap["id2"] = "id2"

	testMap2 := make(map[string]interface{})
	testMap2["id"] = value
	testMap2["id2"] = int(2)

	mappingField2 := field.NewMappingField([]string{"id"})
	v, err = GetFieldValue(testMap, mappingField2)
	assert.Nil(t, err)
	assert.Equal(t, "id", v)

	mappingField3 := field.NewMappingField([]string{"id2"})
	v, err = GetFieldValue(testMap2, mappingField3)
	assert.Nil(t, err)
	assert.Equal(t, int(2), v)

	mappingField4 := field.NewMappingField([]string{"id", "id"})
	v, err = GetFieldValue(testMap2, mappingField4)
	assert.Nil(t, err)
	assert.Equal(t, "12222", v)

	////Int64
	mappingFieldInt64 := field.NewMappingField([]string{"id", "int_64"})
	v, err = GetFieldValue(testMap2, mappingFieldInt64)
	assert.Nil(t, err)
	assert.Equal(t, int64(123), v)
	//Int
	mappingFieldint := field.NewMappingField([]string{"id", "int_v"})
	v, err = GetFieldValue(testMap2, mappingFieldint)
	assert.Nil(t, err)
	assert.Equal(t, int(12), v)
}

var SpecialData = `{
    "City": [
        {
            "Array": [
                {
                    "id": "11111"
                },
                {
                    "id": "2222"
                }
            ],
            "InUS": true,
            "Name": "Sugar Land",
            "Park": {
                "Emails": null,
                "Location": "location",
                "Maps": {
                    "bb": "bb",
                    "cc": "cc",
                    "dd": "dd"
                },
                "Name": "Name"
            }
        }
    ],
    "Emails": [
        "123@123.com",
        "456@456.com"
    ],
    "Id": 1234,
    "In": "string222",
    "Maps": {
        "bb.bb": {
            "id": "10001"
        },
        "cc#cc": "cc",
        "dd**cc": "dd"
    },
    "Maps2": {
        "bb.bb": [
            {
                "id": "10001"
            }
        ],
        "good":[{"id":"12", "x.y":"234"}],
        "cc#cc": "cc",
        "dd**cc": "dd"
    },
    "Maps3": {
        "bb.bb": [
            {
                "x.y": "10001"
            }
        ],
        "cc#cc": [
            {
                "x.y": {
                    "id": "1"
                }
            }
        ],
        "dd*cc": {
            "x.y": {
                "g%f": "hello"
            },
            "y.x":[
                {"d.d":"123"}
            ]
        }
    }
}`

func TestSpecialFieldNames(t *testing.T) {
	mappingField := field.NewMappingField([]string{"Maps", "bb.bb"})
	value, err := GetFieldValue(SpecialData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	log.Info(value)
	v, _ := json.Marshal(value)
	assert.Equal(t, `{"id":"10001"}`, string(v))
}
func TestGetSpecialFieldRoot(t *testing.T) {
	mappingField := field.NewMappingField([]string{"Maps", "bb.bb"})
	value, err := GetFieldValue(SpecialData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, map[string]interface{}{"id": "10001"}, value)
}

func TestGetSpecial2LevelObjectField(t *testing.T) {
	mappingField := field.NewMappingField([]string{"Maps", "bb.bb", "id"})
	value, err := GetFieldValue(SpecialData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, "10001", value)
}

func TestGetSpecial2LevelArrayField(t *testing.T) {
	mappingField := field.NewMappingField([]string{"Maps2", "bb.bb[0]", "id"})
	value, err := GetFieldValue(SpecialData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, "10001", value)
}

func TestGetSpecial2LevelArrayField2(t *testing.T) {
	mappingField := field.NewMappingField([]string{"Maps3", "cc#cc[0]", "x.y", "id"})
	value, err := GetFieldValue(SpecialData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, "1", value)
}

func TestGetSpecialSpecial(t *testing.T) {
	mappingField := field.NewMappingField([]string{"Maps3", "dd*cc", "x.y", "g%f"})
	value, err := GetFieldValue(SpecialData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, "hello", value)
}

func TestGetSpecialSpecial2(t *testing.T) {
	mappingField := field.NewMappingField([]string{"Maps3", "dd*cc", "y.x[0]", "d.d"})
	value, err := GetFieldValue(SpecialData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, "123", value)
}

func TestGetSpecialSpecial3(t *testing.T) {
	mappingField := field.NewMappingField([]string{"Maps2", "good[0]", "x.y"})
	value, err := GetFieldValue(SpecialData, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, "234", value)
}

func TestHasArrayFieldInArray(t *testing.T) {
	assert.False(t, hasArrayFieldInArray([]string{"[square]"}))
	assert.True(t, hasArrayFieldInArray([]string{"bb.bb[0]"}))
}

func TestSpecial(t *testing.T) {

	data := `{
        "[square]":"123",
  "array1": [
    {
      "id.1": 21907387
    },
    {
      "email": -54931037,
      "array2": [
        {
          "id.2": 3458316
        },
        {
          "id.2": 57420133
        },
        {
          "id.2": -95395610
        },
        {
          "id.2": 68243245
        }
      ]
    }
  ]
}`

	mappingField := field.NewMappingField([]string{"array1[1]", "array2[0]", "id.2"})
	value, err := GetFieldValue(data, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)

	assert.Equal(t, "3458316", strconv.FormatFloat(value.(float64), 'f', -1, 64))

	mappingField = field.NewMappingField([]string{"[square]"})
	value, err = GetFieldValue(data, mappingField)
	assert.Nil(t, err)
	assert.NotNil(t, value)

	assert.Equal(t, "123", value)

}
