package libs

import (
	"fmt"
	"reflect"
)

func CopyStruct(src interface{}, dest interface{}) error {
	vsrc := reflect.ValueOf(src)
	vdest := reflect.ValueOf(dest)
	if vsrc.Kind() != reflect.Ptr ||
		vdest.Kind() != reflect.Ptr {
		return fmt.Errorf("Invalid input")
	}
	esrc := vsrc.Elem()
	edest := vdest.Elem()
	if esrc.Kind() != reflect.Struct ||
		esrc.Kind() != reflect.Struct {
		return fmt.Errorf("Invalid input")
	}
	tsrc := esrc.Type()
	for i := 0; i < esrc.NumField(); i++ {
		sourceField := esrc.Field(i)
		sourceFieldName := tsrc.Field(i).Name
		destFiled := edest.FieldByName(sourceFieldName)
		if !destFiled.IsValid() {
			continue
		}
		if destFiled.Kind() != sourceField.Kind() {
			continue
		}
		switch sourceField.Kind() {
		//    int...
		case reflect.Int:
			fallthrough
		case reflect.Int8:
			fallthrough
		case reflect.Int16:
			fallthrough
		case reflect.Int32:
			fallthrough
		case reflect.Int64:
			{
				destFiled.SetInt(sourceField.Int())
			}
		// uint...
		case reflect.Uint:
			fallthrough
		case reflect.Uint8:
			fallthrough
		case reflect.Uint16:
			fallthrough
		case reflect.Uint32:
			fallthrough
		case reflect.Uint64:
			{
				destFiled.SetUint(sourceField.Uint())
			}
		// string
		case reflect.String:
			{
				destFiled.SetString(sourceField.String())
			}
		}
	}
	return nil
}
