Argstogram
----------

Prints a histogram of the cardinality of function arguments in a Go project. It hunts recursively from the current directory of execution.

This is currently a very hacky project and is mostly me practising with the AST libs.

usage:

`argstogram [-skip-tests]`

example output of argstogram on Go itself:

```
(00) [06279] ===================================================================================
(01) [13857] ========================================================================================================================================================================================
(02) [06156] =================================================================================
(03) [02465] ================================
(04) [01043] =============
(05) [00459] ======
(06) [00174] ==
(07) [00061]
(08) [00038]
(09) [00010]
(10) [00015]
(11) [00006]
(12) [00001]
(13) [00000]
(14) [00002]
(15) [00000]
(16) [00000]
(17) [00002]
```

In the spirit of [Clean Code](http://www.goodreads.com/book/show/3735293-clean-code) we want our functions to take as few arguments as possible. This helps to visualise how well a project is doing in the pursuit of reducing the number of arguments their functions take in a project. The aim is `up and to the left`.
