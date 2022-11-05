1. There are 3 purposes of using packages:
    1. It reduces the chance of having overlapping function names. This keeps our function names short and succinct.
    2. It organizes our code so it's easier to find code we want to reuse.
    3. It speeds up the compiler by only requiring recompilation of smaller chunks of a program. Altough we use a specific package, we don't have to recompile it every time we change our program.
2. In a package, identifiers that start with a capital letter (for example: Average) are accessable by other packages or programs which using that package. That's not the case with identifiers which don't start with capital letter.
3. A package alias is an alias given to a package when that package is being imported. When two packages share a same name, we could use alias to refer one of them with a different name. For example, to refer the package `golang-book/chapter-11/math` as `m` we could do: `import m "golang-book/chapter-11/math`.
4. See ./math/math.go
5. We would use a comment before the function declaration to document a function. I'm not sure which functions the question refer by "the functions you created in #3", so i document the ones i created in #4 instead. See ./math/math.go