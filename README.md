Argstogram
----------

Prints a histogram of the cardinality of function arguments in a Go project. It hunts recursively from the current directory of execution.

This is currently a very hacky project and is mostly me practising with the AST libs.

usage:

`argstogram`

example output of argstogram on argstogram:

```
(00) [1] ============================================================================================
(01) [0]
(02) [2] ========================================================================================================================================================================================
(03) [1] ============================================================================================
```

In the spirit of [Clean Code](http://www.goodreads.com/book/show/3735293-clean-code) we want our functions to take as few arguments as possible. This helps to visualise how well a project is doing in the persuit of reducing the number of arguments their functions take in a project. The aim is `up and to the left`.
