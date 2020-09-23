package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParseQueryParams(fields []string, c *gin.Context) (map[string][2]string, map[string]string, int, int, bool) {
	var (
		criteriaValue [2]string
	)
	criteria := make(map[string][2]string)
	order, orderOk := c.GetQueryMap("order")
	if !orderOk {
		order = nil
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 0
	}
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 0
	}
	for _, value := range fields {
		queryVal, okQuery := c.GetQueryMap(value)
		if !okQuery {
			continue
		}
		// only one query for one field
		for k, v := range queryVal {
			criteriaValue[0] = k
			criteriaValue[1] = v
			break
		}
		criteria[value] = criteriaValue
	}

	return criteria, order, limit, offset, true
}
