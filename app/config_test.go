package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {

	s := `name="select id, name , ssss",val=ddd`
	m := getOverrideAppProperty(s)

	assert.Equal(t, "select id, name , ssss", m["name"])
	assert.Equal(t, "ddd", m["val"])

	s = `name="select id, name , ssss",val="ddd"`
	m = getOverrideAppProperty(s)

	assert.Equal(t, "select id, name , ssss", m["name"])
	assert.Equal(t, "ddd", m["val"])

	s = `name="dddd",val="ddd"`
	m = getOverrideAppProperty(s)

	assert.Equal(t, "dddd", m["name"])
	assert.Equal(t, "ddd", m["val"])

	s = `name="name=v",val="select=v"`
	m = getOverrideAppProperty(s)

	assert.Equal(t, "name=v", m["name"])
	assert.Equal(t, "select=v", m["val"])

}
