# Tree-sitter Query Builder

TQB is a golang library that sould be used in conjonction with the go-tree-sitter library.

You can get more details on tree sitter queries from [here](https://tree-sitter.github.io/tree-sitter/using-parsers#pattern-matching-with-queries)

It aims to ease the construction of Pattern Macthing queries, and provide you a way create queries dynamically, like such:

```go
func Query() string {
    query := querier.NewQB()
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

You can find example of building queries from the tree sitter documentation in [this file](./querier/query_test.go)

That way, you souldn't have to worry about syntax errors in your string queries anymore.

## Performances

Benchmarking shows that it takes approximately 5x times to get a query from the query builder than the parameterized query for a medium sized query. However, doing the query takes approximately 1000x more time than getting the query object with a query builder, so the difference is slim.
