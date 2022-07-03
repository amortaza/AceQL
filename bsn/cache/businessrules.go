package cache

import (
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/query"
)

//todo: if gliderecord query is referencing a column that does not exist, do a nice comment and exit
//todo: if gliderecord Get is referencing a column that does not exist, do a nice comment and exit

var gOnAfterUpdateScriptNames = make(map[string][]string)

func ClearAll() {
	gOnAfterUpdateScriptNames = make(map[string][]string)
}

func GetOnAfterUpdate_ScriptNames(tablename string) []string {
	if _, ok := gOnAfterUpdateScriptNames[tablename]; !ok {
		names := make([]string, 0)

		gr := stdsql.NewRecord("x_business_rule")
		if gr == nil {
			return nil
		}

		gr.Add("x_table_name", query.Equals, tablename)
		gr.Query()

		for {
			hasNext, _ := gr.Next()
			if !hasNext {
				break
			}

			name, _ := gr.Get("x_script_name")
			names = append(names, name)
		}

		gOnAfterUpdateScriptNames[tablename] = names
	}

	if val, ok := gOnAfterUpdateScriptNames[tablename]; ok {
		return val
	}

	panic("This can't happen in cache")
}
