# Tree-sitter Query Builder

**It is still an experimental library, you shouldn't expect it to produce valid queries for now**

TQB is a golang library that sould be used in conjonction with the go-tree-sitter library.

It aims to ease the construction of Pattern Macthing queries, and provide you a way create queries dynamically, like such:

```go
func Query() string {
    query := querier.Init()
    i := 0
    query.
    NewSExpression("call_expression").
        If(
            i == 0,
            func(e *SExpression) {
                e.For(func() bool {return i < 10},
                func(e2 *SExpression){

                    e2.NewSExpression("string_literal").
                        If(i < 9,
                        func(e *SExpression) {e.Prop("field")},
                        ).
                        NewSExpression("string_value")
                },
                func(){i += 1})
            }).
        Var("myexpression", query).
    End()
    return query.Marshal()
}
```

This is the equivalent of that string :

```go
    const QUERY = `(
(call_expression
    field: (string_literal
        (string_value))
    field: (string_literal
        (string_value))
    field: (string_literal
        (string_value))
    field: (string_literal
        (string_value))
    field: (string_literal
        (string_value))
    field: (string_literal
        (string_value))
    field: (string_literal
        (string_value))
    field: (string_literal
        (string_value))
    field: (string_literal
        (string_value))
    (string_literal
        (string_value))) @myexpression
)`
```

That way, you souldn't have to worry about syntax errors in your string queries anymore.

If your really care a lot about performances and do not have to generate the string base on input, you can still use this library to
construct the query, print the result in your terminal, and copy paste it in your code.

Benchmarks between different ways of constructing a query will soon be added.
