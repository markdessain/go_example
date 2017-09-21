package main

import "fmt"

import "github.com/kniren/gota/dataframe"
import "github.com/kniren/gota/series"

// import "database/sql"
// import _ "github.com/lib/pq"

import pb "proto"

// const (
//   host     = "localhost"
//   port     = 5432
//   user     = "mark.dessain"
//   password = ""
//   dbname   = "call_me_maybe_development"
// )

func main() {
    fmt.Println("hello world")

    p := pb.Person{
      Id:    1234,
      Name:  "John Doe",
      Email: "jdoe@example.com",
      Phones: []*pb.Person_PhoneNumber{
        {Number: "555-4321", Type: pb.Person_HOME},
      },
    }

    df := dataframe.New(
    	series.New([]string{"b", "a"}, series.String, "COL.1"),
    	series.New([]int{1, 2}, series.Int, "COL.2"),
    	series.New([]float64{3.0, 4.0}, series.Float, "COL.3"),
    )

    // psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    //
    // db, err := sql.Open("postgres", psqlInfo)
    // if err != nil {
  	// 	fmt.Println(err)
  	// }
    //
    // err = db.Ping()
    // if err != nil {
    //   panic(err)
    // }
    //
  	// rows, err := db.Query("SELECT full_name FROM dialler_leads")
    // if err != nil {
  	// 	fmt.Println(err)
  	// }

    fmt.Println(p)
    fmt.Println(df)
    // fmt.Println(db)
    // fmt.Println(rows)
}
