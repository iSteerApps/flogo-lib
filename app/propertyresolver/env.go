package propertyresolver

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/app"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var logEnv = logger.GetLogger("app-props-env-resolver")

const EnvAppPropertyEnvConfigKey = "FLOGO_APP_PROPS_ENV"

type PropToEnvMapping struct {
	Mappings map[string]string `json:"mappings"`
}

var mapping PropToEnvMapping

func init() {
	app.RegisterPropertyValueResolver("env", &EnvVariableValueResolver{})
	mappings := getEnvValue()
	if mappings != "" {
		e := json.Unmarshal([]byte(mappings), &mapping)
		if e != nil {
			logEnv.Errorf("Can not parse value set to '%s' due to error - '%v'", EnvAppPropertyEnvConfigKey, e)
			panic("")
		}
	}
}

func getEnvValue() string {
	key := os.Getenv(EnvAppPropertyEnvConfigKey)
	if len(key) > 0 {
		return key
	}
	return ""
}

// Resolve property value from environment variable
type EnvVariableValueResolver struct {
}

func (resolver *EnvVariableValueResolver) LookupValue(key string) (interface{}, bool) {
	value, exists := os.LookupEnv(key) // first try with the name of the property as is
	if exists {
		return value, exists
	}

	// Lookup based on mapping defined
	keyMapping, ok := mapping.Mappings[key]
	if ok {
		return os.LookupEnv(keyMapping)
	}

	// Try to canonical form
	value, exists = os.LookupEnv(getCanonicalEnv(key)) // if not found try with the canonical form
	return value, exists
}

func getCanonicalEnv(key string) string {
	result := strings.Replace(key, ".", "_", -1)
	result = strings.ToUpper(result)
	return result
}
