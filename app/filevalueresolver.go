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

const ENV_APP_PROPERTY_FILE_CONFIG_KEY = "FLOGO_APP_PROPS_FILE_CONFIG"

func init() {

	filePaths := getExternalFiles()
	if filePaths != "" {
		// Register value resolver
		RegisterPropertyValueResolver("file", &FileValueResolver{})

		//load props from external files
		valueResolvers := config.GetAppPropertiesValueResolver()
		if valueResolvers == "" {
			//Make file default resolver since FLOGO_APP_PROPS_FILE_CONFIG is set
			os.Setenv(config.ENV_APP_PROPERTY_RESOLVER_KEY, "file")
		}

		files := strings.Split(filePaths, ",")
		if len(files) > 0 {
			for _, filePath := range files {
				props := make(map[string]interface{})
				if strings.HasSuffix(filePath, ".json") {
					// Override through file
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
				} else {
					log.Error("File must be a valid JOSN e.g. props.json")
					panic("")
				}

				for k, v := range props {
					preload[k] = v
					delete(props, k)
				}
			}
		}
	}
}

func getExternalFiles() string {
	key := os.Getenv(ENV_APP_PROPERTY_FILE_CONFIG_KEY)
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


