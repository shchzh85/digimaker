//Author xc, Created on 2019-03-28 21:00
//{COPYRIGHTS}

package contenttype

import (
	"github.com/xc/digimaker/core/util"
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	util.SetPackageName("github.com/xc/digimaker/admin")
	err := LoadDefinition()
	if err != nil {
		t.Fail()
	}
	report := contentTypeDefinition["report"]
	// report.Init()
	fmt.Println(report.FieldMap)
	t.Log(fmt.Printf(contentTypeDefinition["report"].TableName + "\n"))

	t.Log(fmt.Printf(contentTypeDefinition["folder"].TableName + "\n"))
	t.Log(fmt.Printf(contentTypeDefinition["article"].TableName + "\n"))

}
