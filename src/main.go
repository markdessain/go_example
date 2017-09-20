package main

import "fmt"

import "github.com/kniren/gota/dataframe"
import "github.com/kniren/gota/series"

func main() {
    fmt.Println("hello world")

    df := dataframe.New(
    	series.New([]string{"b", "a"}, series.String, "COL.1"),
    	series.New([]int{1, 2}, series.Int, "COL.2"),
    	series.New([]float64{3.0, 4.0}, series.Float, "COL.3"),
    )

    fmt.Println(df)
}
