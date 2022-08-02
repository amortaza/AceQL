package main

import (
	"fmt"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"strconv"
)

func init() {
	stdsql.Init("mysql", "clown:1844@/bsn")
}

func main() {
	//auth := authn.NewAuthN("mysecret")
	//token, _ := auth.CreateToken("ace", time.Second*5)
	//_, ok := auth.VerifyToken(token)
}

// 5
func main2() {
	//encodedQuery = "where age=\"45\" and name=\"afshin \\\"the clown\\\"\" and thats it"
	//bootstrap.Run()
	//parseMappings("first >> u_first, last >> u_last")
	//parseMappings("first>>u_first,last>>u_last")

	testGetUser()
}

func testGetUser() {
	name := "x_user"

	orderByAscending := true
	orderBy := "x_name"

	paginationIndex := "0"
	paginationSize := "3"

	crud := stdsql.NewCRUD()

	tableschema, err := flux.GetTableSchema(name, crud)
	if err != nil {
		fmt.Println("error!")
	}

	r := flux.NewRecord(tableschema, crud)
	defer r.Close()

	index, err := strconv.Atoi(paginationIndex)
	if err != nil {
		fmt.Println("error!")
	}

	size, err := strconv.Atoi(paginationSize)
	if err != nil {
		fmt.Println("error!")
	}

	r.Pagination(index, size)

	if orderBy != "" {
		if orderByAscending {
			if err := r.SetOrderBy(orderBy); err != nil {
				fmt.Println("error!")
			}
		} else {
			if err := r.SetOrderByDesc(orderBy); err != nil {
				fmt.Println("error!")
			}
		}
	}

	_, err = r.Query()
	if err != nil {
		fmt.Println("error!")
	}

	list := make([]*flux.RecordMap, 0)

	for {
		hasNext, err := r.Next()

		if err != nil {
			fmt.Println("error!")
		}

		if !hasNext {
			break
		}

		list = append(list, r.GetMap())

		v := r.GetMap()
		d := v.Data
		for a, b := range d {
			fmt.Println(a, b.GetValue())
		}
	}

	fmt.Println("done")
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
