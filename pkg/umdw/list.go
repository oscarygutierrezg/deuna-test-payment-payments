package umdw

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	skip     = "skip"
	limit    = "limit"
	sort     = "sort"
	by       = "by"
	sortAsc  = "asc"
	sortDesc = "desc"
)

type List struct {
	Limit int    `json:"limit"`
	Skip  int    `json:"skip"`
	Sort  string `json:"sort"`
	By    string `json:"by"`
}

func (l *List) Set(skip, limit, sort, column string) error {

	limitNumber, limitErr := strconv.Atoi(limit)
	if limitErr != nil || limitNumber < 1 {
		return errors.New("limit must be a number major or equal than 1")
	}

	skipNumber, skipErr := strconv.Atoi(skip)
	if skipErr != nil || skipNumber < 0 {
		return errors.New("skip must be a number major or equal than 0")
	}

	if sort != sortAsc && sort != sortDesc {
		return errors.New("sort must be asc or desc")
	}

	l.Limit = limitNumber
	l.Skip = skipNumber
	l.Sort = sort
	l.By = column

	return nil
}

func ListContext(c *gin.Context) (*List, error) {
	sk, skFound := c.GetQuery(skip)
	l, lFound := c.GetQuery(limit)
	st, stFound := c.GetQuery(sort)
	cl, _ := c.GetQuery(by)

	if !skFound {
		sk = "0"
	}
	if !lFound {
		l = "10"
	}
	if !stFound {
		st = "asc"
	}

	var list List
	if err := list.Set(sk, l, st, cl); err != nil {
		return nil, err
	}

	return &list, nil
}
