package core

import (
	"dm/core/contenttype"
	"dm/core/fieldtype"
	"dm/core/permission"
	"dm/core/util"
	"dm/core/util/log"
)

func Bootstrap(packageName string) bool {
	log.Info("Starting from " + packageName)
	util.SetPackageName(packageName)
	err := contenttype.LoadDefinition()
	if err != nil {
		return false
	}
	err = fieldtype.LoadDefinition()
	if err != nil {
		return false
	}
	err = permission.LoadPolicies()
	if err != nil {
		return false
	}
	return true
}

//Initialize db
func InitDB() bool {
	return true
}

func Reload() {

}
