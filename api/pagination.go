package api

import (
	"github.com/go-pg/pg/orm"
)

func paginate(query *orm.Query, paginate *Paginate) *orm.Query {
	if paginate != nil {
		if paginate.First != nil {
			query = query.Limit(int(*paginate.First) + 1)
		}

		if paginate.Last != nil {
			query = query.Limit(int(*paginate.Last))
		}

		if paginate.After != nil {
			cursor, err := fromGlobalID(gqlCursor, *paginate.After)
			if err == nil {
				query = query.Offset(cursor + 1)
			}
		}

		if paginate.Before != nil && paginate.Last != nil {
			cursor, err := fromGlobalID(gqlCursor, *paginate.Before)
			if err == nil {
				offset := cursor - int(*paginate.Last) - 1
				if offset < 0 {
					offset = 0
				}
				query = query.Offset(offset)
			}
		}
	}

	return query
}
