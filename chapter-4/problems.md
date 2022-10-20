1. There are two ways to create a new variable:
    1. By defining first, assigning later: var x string; string = "Hello World"
    2. By defining and assigning at a time: var x string = "Hello World"
2. The value of x after running: x := 5; x += 1 is 6.
3. Scope is a range of places we allowed a variable to be available. We determine the scope of a variable based on the block where it's defined.
4. We use the var keyword to create a variable whose value can be changed multiple times, to create a variable whose only can be assigned to a value once, we use the const keyword.
5. See ./fahrenheit-to-celcius-converter.go
6. See ./feet-to-meter-converter.go
