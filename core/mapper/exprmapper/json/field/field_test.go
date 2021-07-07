package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSpecialFields(t *testing.T) {
	path := `Object.Maps3["dd*cc"]["y.x"]["d.d"]`
	fields, err := GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{"Object", "Maps3", "dd*cc", "y.x", "d.d"}, fields)
}

func TestMapTo(t *testing.T) {
	path := `Object`
	fields, err := GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{"Object"}, fields)
}

func TestMapToArrayIndex(t *testing.T) {
	path := `[0]`
	fields, err := GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{"[0]"}, fields)
}

func TestGetSpecialFields2(t *testing.T) {
	path := `Id_dd.name-sss.test["dd*cc"]["y.x"]["d.d"]`
	fields, err := GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{"Id_dd", "name-sss", "test", "dd*cc", "y.x", "d.d"}, fields)
}

func TestGetSpecialFields3(t *testing.T) {
	path := `Object.Maps3[0]["dd*cc"]["y.x"]['d.d']`
	fields, err := GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{"Object", "Maps3[0]", "dd*cc", "y.x", "d.d"}, fields)
}

func TestGetSpecialFields4(t *testing.T) {
	path := `["message.id"]`
	fields, err := GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{"message.id"}, fields)
}

func TestGetAllSpecialWithRootSpecial(t *testing.T) {
	path := `Maps2["bb.bb"][0].id.name`
	fields, err := GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{"Maps2", "bb.bb[0]", "id", "name"}, fields)
}

func TestGetAllSpecialFields(t *testing.T) {
	path := `Object.Maps3["dd.cc"][0]["y.x"]['d.d'].name`
	res, err := GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{"Object", "Maps3", "dd.cc[0]", "y.x", "d.d", "name"}, res)
}

func TestApostoph(t *testing.T) {
	path := `["'apostoph"]`
	res, err := GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{"'apostoph"}, res)

	path = `[''apostoph']`
	res, err = GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{"'apostoph"}, res)

	path = `['"apo"stoph"']`
	res, err = GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{`"apo"stoph"`}, res)

	path = `["apo'"stoph''"]`
	res, err = GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{`apo'"stoph''`}, res)

	path = "[\"apo`stoph\"]"
	res, err = GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{"apo`stoph"}, res)

}

func TestGetAllSpecialFields2(t *testing.T) {
	path := `ReceiveSQSMessage.["x.y"][0]["name&name"]`
	res, err := GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{"ReceiveSQSMessage", "x.y[0]", "name&name"}, res)

}

func TestGetAllSpecialEmpty(t *testing.T) {
	path := `["Output**&&&&$$$%%%@(){ String"]`
	res, err := GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{"Output**&&&&$$$%%%@(){ String"}, res)

	path = `["Output ** &&&&$$$%%%@(){ String"]`
	res, err = GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{"Output ** &&&&$$$%%%@(){ String"}, res)

	path = `["Output String"]`
	res, err = GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{"Output String"}, res)

	path = `["Output	String"]`
	res, err = GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{"Output	String"}, res)

}

func TestGetAllSpecialSingleQuote(t *testing.T) {
	path := "['data']['Array1']"

	res, err := GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{"data", "Array1"}, res)

}

func TestGetAllSpecial(t *testing.T) {
	path := "['data']['Array1']"

	res, err := GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{"data", "Array1"}, res)

	path = "['data']['Array1']"

	res, err = GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{"data", "Array1"}, res)

	path = "['data']['Array1']"

	res, err = GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{"data", "Array1"}, res)

	path = "['data']['Array1']"

	res, err = GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{"data", "Array1"}, res)

	path = "['data']['Array1']"

	res, err = GetAllspecialFields(path)
	assert.Nil(t, err)
	assert.Equal(t, []string{"data", "Array1"}, res)
}

func GetAllspecialFields(path string) ([]string, error) {
	field, err := ParseMappingField(path)
	return field.fields, err

}
