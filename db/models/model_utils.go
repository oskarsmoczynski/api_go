package models

import (
	"fmt"
	"reflect"

	"errors"
)

// GetModelByName returns a new instance of the model by its name
func GetModelByName(modelName string, slice bool) (any, error) {
	var model any

	switch modelName {
	case "user":
		model = &User{}
	case "order":
		model = &Order{}
	case "product":
		model = &Product{}
	case "category":
		model = &Category{}
	case "orderitem":
		model = &OrderItem{}
	case "review":
		model = &Review{}
	default:
		err := fmt.Sprintf("Model %v not found", modelName)
		return nil, errors.New(err)
	}

	if slice {
		model = createModelSlice(model)
	}
	return model, nil
}

func createModelSlice(model any) any {
	modelType := reflect.TypeOf(model).Elem()
	sliceType := reflect.SliceOf(modelType)
	slice := reflect.MakeSlice(sliceType, 0, 0).Interface()
	return slice
}

func ModelFromStruct(input any, model any) error {
	inputVal := reflect.ValueOf(input)
	modelVal := reflect.ValueOf(model)

	if modelVal.Kind() != reflect.Ptr || modelVal.Elem().Kind() != reflect.Struct {
		return errors.New("model must be a pointer to a struct")
	}

	modelVal = reflect.Indirect(modelVal)

	if inputVal.Kind() == reflect.Map {
		iter := inputVal.MapRange()
		for iter.Next() {
			key := iter.Key()
			value := iter.Value()

			// Convert key to string and find corresponding field
			fieldName := key.String()
			dstField := modelVal.FieldByName(fieldName)

			if dstField.IsValid() && dstField.CanSet() {
				// Handle type conversion for common types
				if value.Kind() == reflect.Interface {
					value = reflect.ValueOf(value.Interface())
				}

				// Try to set the value, handle type conversion
				if value.Type().AssignableTo(dstField.Type()) {
					dstField.Set(value)
				} else {
					// Try to convert the value
					switch dstField.Type().Kind() {
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						if value.Kind() == reflect.Float64 {
							dstField.SetInt(int64(value.Float()))
						}
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						if value.Kind() == reflect.Float64 {
							dstField.SetUint(uint64(value.Float()))
						}
					case reflect.Float32, reflect.Float64:
						if value.Kind() == reflect.Int || value.Kind() == reflect.Int64 {
							dstField.SetFloat(float64(value.Int()))
						}
					case reflect.String:
						if value.Kind() != reflect.String {
							dstField.SetString(fmt.Sprintf("%v", value.Interface()))
						} else {
							dstField.SetString(value.String())
						}
					}
				}
			}
		}
		return nil
	}
	return errors.New("wrong type of input data")
}
