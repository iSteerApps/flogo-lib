package app

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/config"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var preload = make(map[string]interface{})

var log = logger.GetLogger("app-props-resolver-file")

// Comma separated list of json files overriding default application property values
// e.g. FLOGO_APP_PROPS_FILE_CONFIG=app1.json,common.json
const EnvAppPropertyFileConfigKey = "FLOGO_APP_PROPS_FILE_OVERRIDE"

func init() {

	filePaths := getExternalFiles()
	if filePaths != "" {
		// Register value resolver
		RegisterPropertyValueResolver("file", &FileValueResolver{})

		if config.GetAppPropertiesValueResolver() == "" {
			//Make file resolver default since FLOGO_APP_PROPS_FILE_CONFIG is set
			os.Setenv(config.ENV_APP_PROPERTY_RESOLVER_KEY, "file")
		}

		// preload props from files
		files := strings.Split(filePaths, ",")
		if len(files) > 0 {
			for _, filePath := range files {
				props := make(map[string]interface{})

				file, e := ioutil.ReadFile(filePath)
				if e != nil {
					log.Errorf("Can not read - %s due to error - %v", filePath, e)
					panic("")
				}
				e = json.Unmarshal(file, &props)
				if e != nil {
					log.Errorf("Can not read - %s due to error - %v", filePath, e)
					panic("")
				}
			}
		}
	}
}

func getExternalFiles() string {
	key := os.Getenv(EnvAppPropertyFileConfigKey)
	if len(key) > 0 {
		return key
	}
	return ""
}

// Resolve property value from external files
type FileValueResolver struct {
}

func (resolver *FileValueResolver) LookupValue(toResolve string) (interface{}, bool) {
	val, found := preload[toResolve]
	return val, found
}
