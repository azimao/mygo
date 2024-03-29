package src

import (
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"regexp"
)

//func TopicValidate(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
//	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
//	//topicsData, topicsOk := topStruct.Interface().(*Topics)
//	//if topicsOk && topicsData.TopicListSize == len(topicsData.TopicList){
//	//	return true
//	//}
//	topics, ok := topStruct.Interface().(*Topics)
//
//	if ok && topics.TopicListSize == len(topics.TopicList) {
//		return true
//	}
//
//	return false
//}
//
//func TopicUrl(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
//	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
//
//	_,ok1:= topStruct.Interface().(*Topics)
//	_,ok2:= topStruct.Interface().(*Topic)
//	if ok1 || ok2{
//		if matched, _ := regexp.MatchString(`^\w{4,10}$`, field.String()); matched {
//			return true
//		}
//	}
//	return false
//}

func TopicsValidate(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	topics, ok := topStruct.Interface().(*Topics)

	if ok && topics.TopicListSize == len(topics.TopicList) {
		return true
	}
	return false
}
func TopicUrl(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	_, ok1 := topStruct.Interface().(*Topics)
	_, ok2 := topStruct.Interface().(*Topic)
	if ok1 || ok2 {
		if matched, _ := regexp.MatchString(`^\w{4,10}$`, field.String()); matched {
			return true
		}

	}
	return false
}
