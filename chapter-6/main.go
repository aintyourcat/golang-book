package main

import "fmt"

func main() {
    /*
    An array is a numbered sequence of elements of SINGLE TYPE with a FIXED LENGTH.
    Array index starts from 0.
    */
    var x [5]int // An array composed of 5 ints.
    x[4] = 100
    fmt.Println("x:", x)

    // var y [5]float64
    // y[0] = 98
    // y[1] = 93
    // y[2] = 77
    // y[3] = 82
    // y[4] = 83
    // Or:
    // y := [5]float64{ 98, 93, 77, 82, 83 }
    // Or:
    y := [5]float64{
        98,
        93,
        77,
        82,
        83, // The trailing , is required by Go
    }

    var total float64 = 0
    // for i := 0; i < len(y); i++ {
    //     total += y[i]
    // }
    // Or:
    /*
    Loop over the y array,
    _ holds the current element's index, while value holds it's value.
    _ tells the Go compiler that we don't need this variable,
    since the compiler won't let us create a variable we never use.
    */
    for _, value := range y {
        total += value
    }
    // The len return type is int so we need to convert it to float64.
    // Type conversion works by using the type name like a function.
    fmt.Println("Average:", total / float64(len(y)))

    /*
    Slice is segment of an array.
    Slices are indexable and have a length, unlike array this LENGTH IS ALLOWED TO CHANGE.
    */
    // var z []float64
    // Or:
    /*
    We should use the builtin make function to create a slice.
    The following create a slice with an underlying float64 array with length 0.
    Slices are always associated with some array,
    although they can never be longer than the array, they can be smaller.
    */
    z := make([]float64, 0)
    fmt.Println("z:", z)

    // The third parameter of the make function specify the length of the underlying array the z slice points to.
    a := make([]float64, 5, 10)
    fmt.Println("a:", a)

    /*
    Slice can be created using the [low : high] expression
    low is the index where to start the slice,
    high is the index where to end it (Not including the index itself).
    We can omit low, high, or both of them.
    */
    array := [5]float64{1, 2, 3, 4, 5}
    b := array[2:4] // [3 4]
    fmt.Println("array:", array)
    fmt.Println("b:", b)
    fmt.Println("array[:2]:", array[:2]) // [1 2], same as: array[0:2]
    fmt.Println("array[2:]:", array[2:]) // [3 4 5], same as: array[2:len(array)]
    fmt.Println("array[:]:", array[:]) // [1 2 3 4 5], same as: array[0:len(array)]

    slice1 := []int{1, 2, 3}
    slice2 := append(slice1, 4, 5) // Returns a new slice
    slice3 := make([]int, 2)
    /*
    Copy slice2's content into slice3,
    since the latter's length is 2,
    only the first 2 elements are copied.
    */
    copy(slice3, slice2)
    fmt.Println("slice1:", slice1)
    fmt.Println("slice2:", slice2)
    fmt.Println("slice3:", slice3)

    /*
    A map is an ordered collection of key-value pairs.
    a.k.a an associative array, a hash table, or a dictionary.
    Map is used to look up a value by it's associated key.
    */

    /*
    c is a map of strings to ints.
    The following lines produces runtime error: assignment to entry in nil map.
    An runtime error is a type of error that happen when we run the program,
    unlike compile-time error which happen when we try to compile the program.
    */
    // var c map[string]int
    // To fix it, we have to initialize the map before using it:
    c := make(map[string]int)
    // Or:
    // var c map[string]int = map[string]int{}
    c["key"] = 10
    fmt.Println("c:", c)
    fmt.Println(`c["key"]:`, c["key"])

    /*
    d is a map of ints to ints.
    It looks like an array, however it's length can change as we add new items to it.
    Moreover, maps are not sequential. We could have d[1], but not d[0].
    */
    d := make(map[int]int)
    d[1] = 10
    fmt.Println("d[1]:", d)

    delete(c, "key") // Deleting an item from a map.
    fmt.Println("c:", c)

    // elements is a map of strings to strings.
    // elements := make(map[string]string)
    // elements["H"] = "Hydrogen"
    // elements["He"] = "Helium"
    // elements["Li"] = "Lithium"
    // elements["Be"] = "Beryllium"
    // elements["B"] = "Boron"
    // elements["C"] = "Carbon"
    // elements["N"] = "Nitrogen"
    // elements["O"] = "Oxygen"
    // elements["F"] = "Fluorine"
    // elements["Ne"] = "Neon"
    // Or:
    elements := map[string]string{
        "H": "Hydrogen",
        "He": "Helium",
        "Li": "Lithium",
        "Be": "Beryllium",
        "B": "Boron",
        "C": "Carbon",
        "N": "Nitrogen",
        "O": "Oxygen",
        "F": "Fluorine",
        "Ne": "Neon",
    }

    fmt.Println(`elements["Li"]:`, elements["Li"])
    /*
    Accessing an element that doesn't exist returns the zero value for the value type,
    in our case because the map's value type is string, an empty string returned.
    */
    fmt.Println(`elements["Un"]:`, elements["Un"])

    /*
    Accessing an element of map can return two values.
    The first value is the result of the lookup.
    The second one tells us whether or not the lookup was successful.
    */
    // name, ok := elements["Un"]
    // fmt.Println("name:", name)
    // fmt.Println("ok:", ok)

    if name, ok := elements["Un"]; ok {
        fmt.Println("name:", name)
        fmt.Println("ok:", ok)
    }

    // improvedElements is a map of strings to maps of strings to strings
    improvedElements := map[string]map[string]string{
        "H": map[string]string{
            "name": "Hydrogen",
            "state": "gas",
        },
        "He": map[string]string{
            "name": "Helium",
            "state": "gas",
        },
        "Li": map[string]string{
            "name": "Lithium",
            "state": "solid",
        },
        "Be": map[string]string{
            "name": "Beryllium",
            "state": "solid",
        },
        "B": map[string]string{
            "name": "Boron",
            "state": "solid",
        },
        "C": map[string]string{
            "name": "Carbon",
            "state": "solid",
        },
        "N": map[string]string{
            "name": "Nitrogen",
            "state": "gas",
        },
        "O": map[string]string{
            "name": "Oxygen",
            "state": "gas",
        },
        "F": map[string]string{
            "name": "Fluorine",
            "state": "gas",
        },
        "Ne": map[string]string{
            "name": "Neon",
            "state": "gas",
        },
    }

    const elementSymbol string = "O"

    if element, ok := improvedElements[elementSymbol]; ok {
        fmt.Println(`improvedElements["` + elementSymbol + `]["name"]:`, element["name"])
        fmt.Println(`improvedElements["` + elementSymbol + `"]["state"]:`, element["state"])
    }
}
