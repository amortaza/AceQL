package main

import (
	"github.com/amortaza/aceql/flux-drivers/stdsql"
)

func init() {
	stdsql.Init("mysql", "clown:1844@/bsn")
}

// 5
func main() {
	//encodedQuery = "where age=\"45\" and name=\"afshin \\\"the clown\\\"\" and thats it"
	//bootstrap.Run()
	//parseMappings("first >> u_first, last >> u_last")
	//parseMappings("first>>u_first,last>>u_last")
}

func main4() {
	/*
		r := stdsql.New Record("x_choice_list")
		_ = r.Add("x_type", query.Equals, "field")
		_ = r.Add("x_table", query.Equals, table)
		_ = r.Query()

		_, _ = r.Next()
	*/
}
func main3() {
	//crud := stdsql.NewCRUD()
	//r := flux.New Record(flux.GetTableSchema("x_choice_list", crud), crud)
	//_ = r.Query()
	//
	//for {
	//	hasNext, _ := r.Next()
	//
	//	if !hasNext {
	//		break
	//	}

	//fmt.Println( "row ", i ) // dbug
	//list = append(list, r.GetMap())
	//v, _ := r.Get("x_name")
	//fmt.Println( "wth ", v )
	//break

	//r.Set("x_value", "ace rox")
	//
	//r.Update()
	//
	//break
	//}
}

/*
func main2() {
	a := flux.GetTableSchema("x_user", stdsql.NewCRUD())
	if a == nil {
		fmt.Println("no table found")
		return
	}

	//f mt.Println( a.Fields()[0].Name )
	//f mt.Println( a.Fields()[0].Type )
}
*/
func main1() {
	/*
		stdsql.Init( "mysql", "clown:1844@/bsn")

		rec := flux.New Record("x_choice_list", stdsql.NewCRUD())
		_ = rec.AddPK("0c8e07932620473ab290b781911dbe9f")
		_ = rec.Query()

		b, _ := rec.Next()

		if b {
			rec.Set("x_order", "12.0")
			//rec.Set("x_name", "new name")
			rec.Set("x_table", "new table")
			rec.Set("x_field", "new field")
			rec.Set("x_value", "new value")
			rec.Set("x_active", "0")

			rec.Update()
		}
	*/
}
