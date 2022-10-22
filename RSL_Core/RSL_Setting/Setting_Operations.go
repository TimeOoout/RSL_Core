package RSL_Setting

import (
	"encoding/json"
	"io/ioutil"
)

func GetSettings() {
	fileContent, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		logWarningGetSettings(err.Error())
		return
	}
	if err := json.Unmarshal(fileContent, &CurrentConfig); err != nil {
		logWarningGetSettings(err.Error())
	} else {
		logInfoGetSetting()
	}
}

func SetSettings() {
	fileContent, err := json.Marshal(DefaultConfig)
	if err = ioutil.WriteFile(ConfigPath, fileContent, 0666); err != nil {
		logWarningSetSettings(err.Error())
	} else {
		logInfoGetSetting()
	}
}
