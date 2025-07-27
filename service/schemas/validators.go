package schemas

import (
	"errors"
	"fmt"
	"strings"
)

func (body ReadTableSchema) Serialize() (ReadTableSchema, error) {
	body.Table = strings.TrimSpace(body.Table)

	var orderBy []string
	for _, v := range body.OrderBy {
		parts := strings.Split(v, ":")
		if len(parts) > 2 {
			return ReadTableSchema{}, errors.New("invalid format of OrderBy field element " + v)
		}

		col := strings.TrimSpace(parts[0])
		dir := ""
		if len(parts) == 2 {
			dir = strings.TrimSpace(strings.ToLower(parts[1]))
			if dir != "desc" {
				return ReadTableSchema{}, errors.New("invalid OrderBy direction for element " + v)
			}
		}

		if dir == "" {
			orderBy = append(orderBy, col)
		} else {
			orderBy = append(orderBy, fmt.Sprintf("%s %s", col, dir))
		}
	}
	body.OrderBy = orderBy

	for _, v := range body.Filters {
		operator := GetOperator(v)
		if operator == "" {
			return ReadTableSchema{}, errors.New("invalid filters operator")
		}
		parts := strings.Split(v, operator)
		if len(parts) != 2 {
			return ReadTableSchema{}, errors.New("multiple operator usage in filters for " + v)
		}
		if operator == "between" {
			vals := strings.Split(parts[1], ",")
			if len(vals) != 2 {
				return ReadTableSchema{}, errors.New("invalid values separator for 'between' operator")
			}
		}
	}

	return body, nil
}


func GetOperator(filter string) string {
	for _, op := range filterOperators {
		if strings.Contains(filter, op) {
			return op
		}
	}
	return ""
}