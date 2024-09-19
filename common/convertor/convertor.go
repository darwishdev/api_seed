package convertor

import (
	"fmt"
	"reflect"

	"github.com/jackc/pgx/v5/pgtype"
)

func SetField(obj interface{}, name string, value interface{}) error {
	// Get the struct value from the interface
	structValue := reflect.ValueOf(obj).Elem()

	// Get the field value by name
	fieldValue := structValue.FieldByName(name)

	// Check if the field exists
	if !fieldValue.IsValid() {
		return fmt.Errorf("no such field: %s in obj", name)
	}

	// Get the type of the field
	fieldType := fieldValue.Type()

	// Get the value to be set as a reflect.Value
	newValue := reflect.ValueOf(value)

	// Check if the type of the value to be set matches the type of the field
	if !newValue.Type().AssignableTo(fieldType) {
		return fmt.Errorf("value type %v is not assignable to field type %v", newValue.Type(), fieldType)
	}

	// Set the field value to the new value
	fieldValue.Set(newValue)

	return nil
}

func ToPgTypeID(value int32) pgtype.Int4 {
	return pgtype.Int4{Int32: value, Valid: value > 0}
}
