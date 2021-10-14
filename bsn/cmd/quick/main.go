package main

import (
	"fmt"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
)

func init() {
	stdsql.Init("mysql", "clown:1844@/bsn")
}

func main() {
	crud := stdsql.NewCRUD()
	r := flux.NewRecord(flux.GetRelation("x_choice_list", crud), crud)
	_ = r.Query()

	for {
		hasNext, _ := r.Next()

		if !hasNext {
			break
		}

		//fmt.Println( "row ", i ) // dbug
		//list = append(list, r.GetMap())
		//v, _ := r.Get("x_name")
		//fmt.Println( "wth ", v )
		//break

		r.Set("x_value", "ace rox")

		r.Update()

		break
	}
}

func main1() {
	a := flux.GetRelation( "x_user", stdsql.NewCRUD())
	if a == nil {
		fmt.Println( "no relations found" )
		return
	}

	fmt.Println( a.Fields()[0].Name )
	fmt.Println( a.Fields()[0].Type )
}


func main2() {
	/*
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
	 */
}

