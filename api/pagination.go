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

type pageInfoResolver struct {
	data       string
	pageInfo   PageInfo
	sortColumn string
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
