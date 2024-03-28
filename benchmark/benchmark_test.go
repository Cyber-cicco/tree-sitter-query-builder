package benchmark

import (
	"context"
	"fmt"
	"testing"
	"time"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/java"
	"github.com/Cyber-cicco/tree-sitter-query-builder/querier"
)

var lang *sitter.Language
var parser *sitter.Parser

func init() {
    parser = sitter.NewParser()
    lang = java.GetLanguage()
    parser.SetLanguage(lang)
}

func TestBuilder(t *testing.T) {
    numLoops := 5000
    parser := sitter.NewParser()
    parser.SetLanguage(java.GetLanguage())

    queryString := `(
(field_declaration
    declarator: (variable_declarator
        name: (identifier)
        value: (string_literal
            (string_fragment) @fragment)))
(#match? @fragment "~param")
)`


    content := []byte(`
package fr.edpurolo.freelearning.page;

/**
 * THIS CLASS CAN BE OVERWRITTEN BY N0DZCRYPT
 * IF YOU INTEND TO USE THE N0DZCRYPT CLI IN YOUT APP, TRY NOT TO CHANGE IT.
 *
 * Contains constants pointing to thymeleaf fragments
 */
public class Routes {

    public static final String ADR_BASE_LAYOUT = "layout/base";
    public static final String ADR_HOME = "page/home/home";
    public static final String ADR_ABOUT = "page/about/about";
    public static final String ADR_LOGIN = "page/about/login";
    public static final String ADR_FORM_ERROR = "components/form-error";
}
`)
    startQB := time.Now().UnixMilli()

    for i := 0; i < numLoops; i++ {

        qb := querier.NewQB()
        qb.
        NewSExpression("field_declaration").
            NewSExpression("modifiers").End().
            NewSExpression("variable_declarator").
                Prop("declarator").
                NewSExpression("identifier").
                    Prop("name").
                    End().
                NewSExpression("string_literal").
                    Prop("value").
                    NewSExpression("string_fragment").
                    Var("fragment", qb)

        qb.NewMatcher("match").Identifier("fragment").ValueAsString("page.+")
        query := qb.Marshal()
        query.Content = content
        query.Lang = lang

        tree, err := parser.ParseCtx(context.Background(), nil, content)

        if err != nil {
            t.Fatalf("Shouldn't have had error, got %s", err)
        }

        query.Tree = tree
        captures := []*sitter.QueryCapture{}
        err = query.ExecuteQuery(func(c *sitter.QueryCapture){captures = append(captures, c)}) 

        if err != nil {
            t.Fatalf("Shouldn't have had error, got %s", err)
        }

        if len(captures) != 3 {
            t.Fatalf("Wrong number of captures. Expected 3, got %d", len(captures))
        }
    }

    endQB := time.Now().UnixMilli()

    fmt.Printf("It took %d milliseconds to do %d loops with query builder and parsing\n", endQB - startQB, numLoops)

    startPQ := time.Now().UnixMilli()
    for i := 0; i < numLoops; i++ {

        pq := querier.NewPQ(queryString)
        pq.AddValue("param", "page.+")
        query, err := pq.GetQuery()

        if err != nil {
            t.Fatalf("Shouldn't have had error, got %s", err)
        }

        query.Content = content
        query.Lang = lang

        tree, err := parser.ParseCtx(context.Background(), nil, content)

        if err != nil {
            t.Fatalf("Shouldn't have had error, got %s", err)
        }

        query.Tree = tree
        captures := []*sitter.QueryCapture{}
        err = query.ExecuteQuery(func(c *sitter.QueryCapture){captures = append(captures, c)}) 

        if err != nil {
            t.Fatalf("Shouldn't have had error, got %s", err)
        }

        if len(captures) != 3 {
            t.Fatalf("Wrong number of captures. Expected 3, got %d", len(captures))
        }
    }

    endPQ := time.Now().UnixMilli()

    fmt.Printf("It took %d milliseconds to do %d loops with parameterized query and parsing\n", endPQ - startPQ, numLoops)

    startQB = time.Now().UnixMilli()

    for i := 0; i < numLoops * 1000; i++ {

        qb := querier.NewQB()
        qb.
        NewSExpression("field_declaration").
            NewSExpression("modifiers").End().
            NewSExpression("variable_declarator").
                Prop("declarator").
                NewSExpression("identifier").
                    Prop("name").
                    End().
                NewSExpression("string_literal").
                    Prop("value").
                    NewSExpression("string_fragment").
                    Var("fragment", qb)

        qb.NewMatcher("match").Identifier("fragment").ValueAsString("page.+")
        qb.Marshal()
    }

    endQB = time.Now().UnixMilli()

    fmt.Printf("It took %d milliseconds to do %d loops with query builder to get the query object\n", endQB - startQB, numLoops*1000)

    startPQ = time.Now().UnixMilli()

    for i := 0; i < numLoops * 1000; i++ {

        pq := querier.NewPQ(queryString)
        pq.AddValue("param", "page.+")
        pq.GetQuery()

    }

    endPQ = time.Now().UnixMilli()

    fmt.Printf("It took %d milliseconds to do %d loops to get query object with parameterized query\n", endPQ - startPQ, numLoops*1000)

}
