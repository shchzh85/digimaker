// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package contenttype

import (
	"database/sql"
	"strings"

	"github.com/xc/digimaker/core/db"
	. "github.com/xc/digimaker/core/db"
	"github.com/xc/digimaker/core/util"
)

// Location is an object representing the database table.
// Implement dm.contenttype.ContentTyper interface
type Location struct {
	ID             int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	ParentID       int    `boil:"parent_id" json:"parent_id" toml:"parent_id" yaml:"parent_id"`
	MainID         int    `boil:"main_id" json:"main_id" toml:"main_id" yaml:"main_id"`
	IdentifierPath string `boil:"identifier_path" json:"identifier_path" toml:"identifier_path" yaml:"identifier_path"`
	Hierarchy      string `boil:"hierarchy" json:"hierarchy" toml:"hierarchy" yaml:"hierarchy"`
	Depth          int    `boil:"depth" json:"depth" toml:"depth" yaml:"depth"`
	ContentType    string `boil:"content_type" json:"content_type" toml:"content_type" yaml:"content_type"`
	ContentID      int    `boil:"content_id" json:"content_id" toml:"content_id" yaml:"content_id"`
	Language       string `boil:"language" json:"language" toml:"language" yaml:"language"`
	Name           string `boil:"name" json:"name" toml:"name" yaml:"name"`
	IsHidden       bool   `boil:"is_hidden" json:"is_hidden" toml:"is_hidden" yaml:"is_hidden"`
	IsInvisible    bool   `boil:"is_invisible" json:"is_invisible" toml:"is_invisible" yaml:"is_invisible"`
	Priority       int    `boil:"priority" json:"priority" toml:"priority" yaml:"priority"`
	UID            string `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	Section        string `boil:"section" json:"section" toml:"section" yaml:"section"`
	P              string `boil:"p" json:"p" toml:"p" yaml:"p"`
	path           []int  `boil:"-"`
}

func (c *Location) ToDBValues() map[string]interface{} {
	result := make(map[string]interface{})
	result["id"] = c.ID
	result["parent_id"] = c.ParentID
	result["identifier_path"] = c.IdentifierPath
	result["main_id"] = c.MainID
	result["hierarchy"] = c.Hierarchy
	result["depth"] = c.Depth
	result["content_type"] = c.ContentType
	result["content_id"] = c.ContentID
	result["language"] = c.Language
	result["name"] = c.Name
	result["is_hidden"] = c.IsHidden
	result["is_invisible"] = c.IsInvisible
	result["priority"] = c.Priority
	result["uid"] = c.UID
	result["section"] = c.Section
	result["p"] = c.P
	return result
}

func (c *Location) TableName() string {
	return "dm_location"
}

func (c *Location) IdentifierList() []string {
	return util.GetInternalSettings("location_columns")
}

func (c *Location) Field(name string) interface{} {
	var result interface{}
	switch name {
	case "id", "ID":
		result = c.ID
	case "parent_id", "ParentID":
		result = c.ParentID
	case "main_id", "MainID":
		result = c.MainID
	case "hierarchy", "Hierarchy":
		result = c.Hierarchy
	case "depth", "Depth":
		result = c.Depth
	case "content_type", "ContentType":
		result = c.ContentType
	case "content_id", "ContentID":
		result = c.ContentID
	case "language", "Language":
		result = c.Language
	case "name", "Name":
		result = c.Name
	case "is_hidden", "IsHidden":
		result = c.IsHidden
	case "is_invisible", "IsInvisible":
		result = c.IsInvisible
	case "priority", "Priority":
		result = c.Priority
	case "uid", "UID":
		result = c.UID
	case "section", "Section":
		result = c.Section
	case "p", "P":
		result = c.P
	default:
	}
	return result
}

//Get path array from hierarchy. eg[1, 2]
func (c *Location) Path() []int {
	if len(c.path) == 0 {
		path := strings.Split(c.Hierarchy, "/")
		c.path = util.ArrayStrToInt(path)
	}
	return c.path
}

func (c *Location) Store(transaction ...*sql.Tx) error {
	handler := db.DBHanlder()
	if c.ID == 0 {
		id, err := handler.Insert(c.TableName(), c.ToDBValues(), transaction...)
		c.ID = id
		if err != nil {
			return err
		}
	} else {
		err := handler.Update(c.TableName(), c.ToDBValues(), Cond("id", c.ID), transaction...)
		return err
	}
	return nil
}

//Count how many locations for the same content
//Note: the count is instant so not cached.
func (l *Location) CountLocations() int {
	handler := db.DBHanlder()
	count, err := handler.Count("dm_location", Cond("content_type", l.ContentType).Cond("content_id", l.ContentID))
	if err != nil {
		//todo: panic to top
	}
	return count
}

//If the location is main location
func (l *Location) IsMainLocation() bool {
	return l.MainID == l.ID
}

//Delete location only
func (l *Location) Delete(transaction ...*sql.Tx) error {
	handler := db.DBHanlder()
	contentError := handler.Delete(l.TableName(), Cond("id", l.ID), transaction...)
	return contentError
}

//Get parent id. no cache for now.
func (l *Location) GetParentLocation() (*Location, error) {
	return GetLocationByID(l.ParentID)
}

func GetLocations(contenttype string, cid int) (*[]Location, error) {
	handler := db.DBHanlder()
	locations := &[]Location{}
	err := handler.GetEntity("dm_location", Cond("content_type", contenttype).And("content_id", cid), []string{}, locations)
	if err != nil {
		return nil, err
	}
	return locations, nil
}

func GetLocationByID(locationID int) (*Location, error) {
	handler := db.DBHanlder()
	location := &Location{}
	err := handler.GetEntity("dm_location", Cond("id", locationID), []string{}, location)
	if err != nil {
		return nil, err
	}
	return location, nil
}
