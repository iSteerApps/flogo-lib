package app

import (
	"testing"
	"os"
	"github.com/TIBCOSoftware/flogo-lib/config"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/stretchr/testify/assert"
)

func TestEnvValueResolver(t *testing.T) {
	os.Setenv(config.ENV_APP_PROPERTY_RESOLVER_KEY, "env")
	os.Setenv("TEST_PROP", "testprop")
	defer func() {
		os.Unsetenv(config.ENV_APP_PROPERTY_RESOLVER_KEY)
		os.Unsetenv("TEST_PROP")
	}()

	resolver := GetPropertyValueResolver("env")
	assert.NotNil(t, resolver)
	resolvedVal, found := resolver.LookupValue("TEST_PROP")
	assert.True(t, true, found)
	assert.Equal(t, "testprop", resolvedVal)

	_, found = resolver.LookupValue("TEST_PROP1")
	assert.False(t, false, found)
}

func TestExternalPropResolution(t *testing.T) {
	os.Setenv(config.ENV_APP_PROPERTY_RESOLVER_KEY, "env")
	os.Setenv("MyProp", "env_myprop_value")
	defer func() {
		os.Unsetenv(config.ENV_APP_PROPERTY_RESOLVER_KEY)
		os.Unsetenv("MyProp")
	}()

	var attrs []*data.Attribute
	attr, _ := data.NewAttribute("MyProp", data.TypeString, "")
	attrs = append(attrs, attr)
	resolvedProps, err := loadExternalProperties(attrs)
	assert.Nil(t, err)
	assert.Equal(t, "env_myprop_value", resolvedProps["MyProp"])
}