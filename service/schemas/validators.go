package schemas

import (
	"api/db/models"
	"errors"
	"fmt"
	"strings"
)

var errorInvalidModel = errors.New("invalid model name")

func (body *ReadTableSchema) Serialize() error {
	body.Table = strings.ToLower(strings.TrimSpace(body.Table))
	if model := models.GetModelByName(body.Table, false); model == nil {
		return errorInvalidModel
	}
	var orderBy []string
	for _, v := range body.OrderBy {
		parts := strings.Split(v, ":")
		if len(parts) > 2 {
			return errors.New("invalid format of OrderBy field element " + v)
		}

		col := strings.TrimSpace(parts[0])
		dir := ""
		if len(parts) == 2 {
			dir = strings.TrimSpace(strings.ToLower(parts[1]))
			if dir != "desc" {
				return errors.New("invalid OrderBy direction for element " + v)
			}
		}

		if dir == "" {
			orderBy = append(orderBy, col)
		} else {
			orderBy = append(orderBy, fmt.Sprintf("%s %s", col, dir))
		}
	}
	body.OrderBy = orderBy
	if err := parseFilters(body.Filters); err != nil {
		return err
	}

	return nil
}

func (body *CreateEntrySchema) Serialize() error {
	body.Table = strings.ToLower(strings.TrimSpace(body.Table))
	if model := models.GetModelByName(body.Table, false); model == nil {
		return errorInvalidModel
	}
	return nil
}


func (body *UpdateEntrySchema) Serialize() error {
	body.Table = strings.ToLower(strings.TrimSpace(body.Table))
	if model := models.GetModelByName(body.Table, false); model == nil {
		return errorInvalidModel
	}

	if err := parseFilters(body.Filters); err != nil {
		return err
	}

	for _, v := range body.Values {
		parts := strings.Split(v, "=")
		if len(parts) != 2 {
			return errors.New("couldn't determine key-value assignment in values: " + v)
		}
	}

	return nil
}


func (body *DeleteEntrySchema) Serialize() error {
	body.Table = strings.ToLower(strings.TrimSpace(body.Table))
	if model := models.GetModelByName(body.Table, false); model == nil {
		return errorInvalidModel
	}
	if err := parseFilters(body.Filters); err != nil {
		return err
	}
	
	return nil
}

func GetOperator(filter string) string {
	for _, op := range filterOperators {
		if strings.Contains(filter, op) {
			return op
		}
	}
	return ""
}

func parseFilters(filters []string) error {
	for _, v := range filters {
		operator := GetOperator(v)
		if operator == "" {
			return errors.New("invalid filters operator")
		}
		parts := strings.Split(v, operator)
		if len(parts) != 2 {
			return errors.New("multiple operator usage in filters for " + v)
		}
		if operator == "between" {
			vals := strings.Split(parts[1], ",")
			if len(vals) != 2 {
				return errors.New("invalid values separator for 'between' operator")
			}
		}
	}
	return nil
}