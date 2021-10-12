package main

import (
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
)

func main() {
	stdsql.Init("mysql", "clown:1844@/bsn")

	r := flux.NewRecord("x_choice_list", stdsql.NewCRUD())
	_ = r.Query()

	//list := make([]*flux.RecordMap, 0)

	for {
		hasNext, _ := r.Next()

		if !hasNext {
			break
		}

		//list = append(list, r.GetMap())
		//v, _ := r.Get("x_name")
		//fmt.Println( "wth ", v )
		//break

		r.Update()
	}
}
func main2() {
	stdsql.Init( "mysql", "clown:1844@/bsn")

	rec := flux.NewRecord("x_choice_list", stdsql.NewCRUD())
	_ = rec.AddPrimaryKey("0c8e07932620473ab290b781911dbe9f")
	_ = rec.Query()

	b, _ := rec.Next()

	if b {
		rec.Set("x_order", "12.0")
		//rec.Set("x_name", "new name")
		rec.Set("x_table", "new table")
		rec.Set("x_field", "new field")
		rec.Set("x_value", "new value")
		rec.Set("x_enabled", "0")

		rec.Update()
	}
}

