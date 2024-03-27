package querier

import (
	"testing"
)

func TestBinaryExpression(t *testing.T) {
    expected := `(
(binary_expression
    (number_literal)
    (number_literal))
)`
    query := Init()
    query.
    NewSExpression("binary_expression").
        NewSExpression("number_literal").End().
        NewSExpression("number_literal").End().
    End()
    result := query.Marshal()
    if result != expected {
        t.Fatalf("Expected %s, got %s", expected, result)
    }
}

func TestParamNegation(t *testing.T) {
    expected := `(
(class_declaration
    name: (identifier) @class_name
    !type_parameters)
)`
    query := Init()
    query.
    NewSExpression("class_declaration").
        NewSExpression("identifier").
            Prop("name").
            Var("class_name", query).
            End().
        NewSExpression("").
            Not().
            Prop("type_parameters").
            End().
    End()
    result := query.Marshal()
    if result != expected {
        t.Fatalf("Expected %s, got %s", expected, result)
    }
}

func TestIfFor(t *testing.T) {
    expected := `(
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
    query := Init()
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
    result := query.Marshal()
    if result != expected {
        t.Fatalf("Expected %s, got %s", expected, result)
    }
}

func TestMultipleChildsForGroup(t *testing.T) {
    expected := `(
(call_expression
    args: (
        (string_literal)
        (integer_literal)))
)`
    query := Init()
    query.
    NewSExpression("call_expression").
        NewSExpression("").
        Group().
        Prop("args").
            NewSExpression("string_literal").End().
            NewSExpression("integer_literal").End().
        End().
    End()
    result := query.Marshal()
    if expected != result {
        t.Fatalf("Expected %s, got %s", expected, result)
    }
}

