package helper

import "reflect"

func GenerateFilter(input any) map[string]any {
	filters := map[string]any{}

	value := reflect.ValueOf(input)
	typ := reflect.TypeOf(input)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := typ.Field(i)

		// get tag form
		formTag := fieldType.Tag.Get("form")

		// skip if value empty
		if !field.IsZero() {
			filters[formTag] = field.Interface()
		}
	}

	return filters
}
