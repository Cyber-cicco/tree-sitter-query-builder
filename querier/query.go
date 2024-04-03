package querier

import (
	sitter "github.com/smacker/go-tree-sitter"
)

type Query struct {
    Query []byte
    Content []byte
    Lang *sitter.Language
    Tree *sitter.Tree
}

type CaptureFunc func(c *sitter.QueryCapture) error

//Executes the Query.Query slice of bytes as a tree sitter query.
//
//Return every matches found for the query
//Filters predicates
func (q *Query) ExecuteQuery(captureFunc CaptureFunc) error {

    query, err := sitter.NewQuery(q.Query, q.Lang)

    if err != nil {
        return err
    }

    qc := sitter.NewQueryCursor()
    qc.Exec(query, q.Tree.RootNode())

    for {

        m, ok := qc.NextMatch()

        if !ok {
            break
        }

        m = qc.FilterPredicates(m, q.Content)

        for _, c := range m.Captures {
            err := captureFunc(&c)

            if err != nil {
                return err
            }
        }
    }

    return nil 
}

//Executes the Query.Query slice of bytes as a tree sitter query.
//
//Return every matches found for the query
//Doesn't filter predicates
func (q *Query) ExecuteSimpleQuery(captureFunc CaptureFunc) error {

    query, err := sitter.NewQuery(q.Query, q.Lang)

    if err != nil {
        return err
    }

    qc := sitter.NewQueryCursor()
    qc.Exec(query, q.Tree.RootNode())

    for {

        m, ok := qc.NextMatch()

        if !ok {
            break
        }

        for _, c := range m.Captures {
           err := captureFunc(&c)

            if err != nil {
                return err
            }
        }
    }

    return nil 
}
