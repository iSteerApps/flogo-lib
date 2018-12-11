package propertyresolver

import (
	"fmt"
	"os"
"strings"

"github.com/TIBCOSoftware/flogo-lib/app"
"github.com/iancoleman/strcase"
)

// Enable to resolve app props value from env variable.
// e.g. FLOGO_APP_PROPS_CONFIG_ENV=true
const EnvAppPropertyEnvConfigKey = "FLOGO_APP_PROPS_CONFIG_ENV"

func init() {
	if getValue() == "true" {
		app.RegisterPropertyValueResolver("env", &EnvVariableValueResolver{})
	}
}

func getValue() string {
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
	value, exists = os.LookupEnv(getCanonicalEnv(key)) // if not found try with the canonical form
	return value, exists
}

func getCanonicalEnv(key string) string {
	result := strcase.ToScreamingSnake(key)
	result = strings.Replace(result, ".", "_", -1)
	result = strings.Replace(result, "__", "_", -1)
	fmt.Println(result)
	return result
}
