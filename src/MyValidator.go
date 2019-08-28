package src

import (
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"regexp"
)

func TopicUrl(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if _, ok := topStruct.Interface().(*Topic); ok {
		if macthed, _ := regexp.MatchString(`^\w{4,10}$`, field.String()); macthed {
			return true
		}
	}
	return false
}
