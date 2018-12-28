package api

import (
	"github.com/go-pg/pg/orm"
)

func paginate(
	query *orm.Query,
	paginate Paginate,
) (*orm.Query, int) {
	pager := orm.Pager{
		Offset: 0,
		Limit:  100,
	}

	if paginate.First == nil && paginate.Last == nil {
		panic(ErrNoPaginationLimits())
	}

	if paginate.First != nil {
		pager.Limit = int(*paginate.First) + 1
		if paginate.After != nil {
			pager.Offset = int(*paginate.After) + 1
		}
	}

	if paginate.Last != nil {
		pager.Limit = int(*paginate.Last) + 1
		if paginate.Before != nil {
			pager.Offset = int(*paginate.Before) - int(*paginate.Last) - 1
			if pager.Offset < 0 {
				pager.Limit += pager.Offset
				pager.Offset = 0
			}
		} else {
			totalCount, err := query.Copy().Apply(pager.Paginate).Count()
			if err != nil {
				panic(err)
			}
			pager.Offset = totalCount - int(*paginate.Last) - 1
		}
	}

	query = query.Apply(pager.Paginate)

	return query, pager.Offset
}

func getPaginationData(paginationInput PaginationInput) Paginate {
	paginate := Paginate{}

	if paginationInput.First != nil {
		paginate.First = paginationInput.First
	}

	if paginationInput.Last != nil {
		paginate.Last = paginationInput.Last
	}

	if paginationInput.After != nil {
		cursor, err := fromGlobalCursor(*paginationInput.After)
		if err == nil {
			paginate.After = &cursor
		}
	}

	if paginationInput.Before != nil {
		cursor, err := fromGlobalCursor(*paginationInput.Before)
		if err == nil {
			paginate.Before = &cursor
		}
	}

	return paginate
}

type pageCursor int

type pageInfoResolver struct {
	data     string
	pageInfo PageInfo
}

func (res pageInfoResolver) EndCursor() *string {
	if res.pageInfo.endCursor != nil {
		cursor := toGlobalCursor(*res.pageInfo.endCursor)
		return &cursor
	}
	return nil
}

func (res pageInfoResolver) StartCursor() *string {
	if res.pageInfo.startCursor != nil {
		cursor := toGlobalCursor(*res.pageInfo.startCursor)
		return &cursor
	}
	return nil
}

func (res pageInfoResolver) HasNextPage() bool {
	return res.pageInfo.hasNextPage
}

func (res pageInfoResolver) HasPreviousPage() bool {
	return res.pageInfo.hasPreviousPage
}
