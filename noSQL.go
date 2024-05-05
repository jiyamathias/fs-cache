package fscache

import (
	"fmt"
	"reflect"
)

type (
	// Collection object
	Collection struct {
		CollectionName string
		NoSQL          []map[string]interface{}
	}
)

// Collection defines the collection(table) name to perform an operations on
func (ns *NoSQL) Collection(col interface{}) *Collection {
	t := reflect.TypeOf(col)
	// validate t.Name()
	if len(t.Name()) == 0 {
		ns.logger.Error().Msg("Collection cannot be empty...")
	}

	if t.Kind() != reflect.Struct || t.Kind() != reflect.String {
		ns.logger.Error().Msg("Collection must either be a [string] or an [object]")
	}

	var colName string
	// check if the ending string ends with the letter s
	// if not, append (s) to it e.g user = users
	if len(t.Name()) > 0 && string(t.Name()[len(t.Name())-1]) != "s" {
		colName = fmt.Sprintf("%ss", t.Name())
	} else {
		colName = t.Name()
	}

	return &Collection{
		CollectionName: colName,
		NoSQL:          ns.Storage,
	}
}
