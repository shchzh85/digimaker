//Author xc, Created on 2019-05-10 22:48
//{COPYRIGHTS}

//Package handlers implements build-in action callbacks.
package handlers

import (
	"database/sql"
	"github.com/xc/digimaker/core/contenttype"
	"github.com/xc/digimaker/core/handler"
)

type ImageHandler struct {
}

func (ih ImageHandler) New(content contenttype.ContentTyper, tx *sql.Tx, parentID ...int) error {
	content.SetValue("parent_id", parentID[0]) //todo: validate more
	return nil
}

func (ih ImageHandler) Validate(inputs map[string]interface{}, result *handler.ValidationResult) {

}

func init() {
	handler.RegisterContentTypeHandler("image", ImageHandler{})
}
