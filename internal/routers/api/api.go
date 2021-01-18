package api

type QueryParams struct {
	Criteria []CriteriaParam
	Order    map[string]string    `query:"order"`
	Limit    int                  `query:"limit"`
	Offset   int                  `query:"offset"`
}

type CriteriaParam struct {
	Field string     `query:"field"`
	Value string     `value:"field"`
	Condition string `condition:"field"`
}

//func ParseQueryParams(AllowedFields []string, qp *QueryParams) (map[string][2]string, map[string]string, int, int, bool) {
//	isComparisonOperator := true
//	criteria := make(map[string][2]string)
//	order := make(map[string]string)
//	for _, value := range AllowedFields {
//		queryVal, okVal := qp.Criteria[value]
//		orderVal, okOrder := qp.Order[value]
//		if !okOrder {
//			logging.Error("user doesn't have this field for order")
//			break
//		}
//		order[value] = orderVal
//		if !okVal {
//			continue
//		}
//		isComparisonOperator = repository.IsComparisonOperator(queryVal[0])
//		if !isComparisonOperator {
//			logging.Error("last element must be comparison operator")
//			break
//		}
//		criteria[value] = queryVal
//	}
//
//	return criteria, qp.Order, qp.Limit, qp.Offset, isComparisonOperator
//}
