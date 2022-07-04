package cache

import (
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/query"
)

//todo: if gliderecord query is referencing a column that does not exist, do a nice comment and exit
//todo: if gliderecord Get is referencing a column that does not exist, do a nice comment and exit

var gCache_onAfterUpdateScriptNames = make(map[string][]string)

func ClearAll() {
	gCache_onAfterUpdateScriptNames = make(map[string][]string)
}

func GetOnAfterUpdate_ScriptNames(tablename string) ([]string, error) {
	if _, ok := gCache_onAfterUpdateScriptNames[tablename]; !ok {
		names := make([]string, 0)

		gr, err := stdsql.NewRecord("x_business_rule")
		if err != nil {
			return nil, err
		}

		if err := gr.Add("x_table_name", query.Equals, tablename); err != nil {
			return nil, err
		}

		if _, err := gr.Query(); err != nil {
			return nil, err
		}

		for {
			hasNext, err := gr.Next()

			if err != nil {
				return nil, err
			}

			if !hasNext {
				break
			}

			name, err := gr.Get("x_script_name")
			if err != nil {
				return nil, err
			}

			names = append(names, name)
		}

		gCache_onAfterUpdateScriptNames[tablename] = names
	}

	if val, ok := gCache_onAfterUpdateScriptNames[tablename]; ok {
		return val, nil
	}

	panic("This can't happen in cache")
}
