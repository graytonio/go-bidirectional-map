package gobidirectionalmap_test

import (
   "fmt"

    bidmap "github.com/graytonio/go-bidirectional-map"
)

func ExampleNewMap() {
    myMap := bidmap.NewMap(map[string]int{
        "foo": 3,
        "bar": 4,
    })

    fmt.Println(myMap.GetP("foo"))
    // Output: 3 true
}

func ExampleMap_GetP() {
    myMap := bidmap.NewMap(map[string]int{
        "foo": 3,
        "bar": 4,
    })

    fmt.Println(myMap.GetP("foo"))
    // Output: 3 true
}

func ExampleMap_GetS() {
    myMap := bidmap.NewMap(map[string]int{
        "foo": 3,
        "bar": 4,
    })

    fmt.Println(myMap.GetS(3))
    // Output: foo true
}

