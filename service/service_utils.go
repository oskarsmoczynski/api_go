package service

import (
	"api/service/schemas"
	"fmt"
	"strings"
)

func buildFilters(filters string) (string, string, string) {
    operator := schemas.GetOperator(filters)
    parts := strings.Split(filters, operator)
    col := parts[0]
    if operator == "between" {
        vals := strings.Split(parts[1], ",")
        expression := fmt.Sprintf("%v %v ? AND ?", col, operator)
        return expression, vals[0], vals[1]
    }
    
    expression := fmt.Sprintf("%v %v ?", col, operator)
    return expression, parts[1], ""
}