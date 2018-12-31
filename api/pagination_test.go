package api

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
)

func TestPagination(t *testing.T) {
	query := `
		query Pages($paginate: PaginationInput!) {
			pages(sort: { field: NAME, order: ASC }, paginate:$paginate) {
				edges{
					cursor
					node {
						id
						name
					}
				}
				pageInfo{
					endCursor
					hasNextPage
					hasPreviousPage
					startCursor
				}
			}
		}
	`
	type QueryVariables struct {
		Paginate PaginationInput `json:"paginate"`
	}

	limit1 := int32(2)
	limit2 := int32(3)
	limit3 := int32(10)
	cursor1 := toGlobalCursor(Cursor(2))
	cursor2 := toGlobalCursor(Cursor(3))

	paginationInputs := []PaginationInput{
		PaginationInput{
			First: &limit1,
		},
		PaginationInput{
			Last: &limit1,
		},
		PaginationInput{
			First: &limit1,
			After: &cursor1,
		},
		PaginationInput{
			Last:   &limit1,
			Before: &cursor1,
		},
		PaginationInput{
			Last:   &limit2,
			Before: &cursor1,
		},
		PaginationInput{
			Last:   &limit1,
			Before: &cursor2,
		},
		PaginationInput{
			Last: &limit3,
		},
	}

	for index, paginationInput := range paginationInputs {
		t.Run(
			fmt.Sprintf("Test pagination %d", index), func(t *testing.T) {
				structVariables := QueryVariables{
					Paginate: paginationInput,
				}
				variables, err := json.Marshal(&structVariables)
				if err != nil {
					t.Fatal(err)
				}
				data, err := execQuery(query, string(variables), nil)
				if err != nil {
					t.Fatal(err)
				}
				cupaloy.SnapshotT(t, data)
			},
		)
	}
}
