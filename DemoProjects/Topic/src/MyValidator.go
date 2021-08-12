package src

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
)

//func TopicUrl(v *validator.Validate,topStruct reflect.Value,currentStructOrField reflect.Value,
//	field reflect.Value,fieldType reflect.Type,fieldKind reflect.Kind,param string) bool {
//
//	fmt.Println(topStruct)
//	fmt.Println(topStruct.Interface())
//	return false
//}

func TopicUrl(fl validator.FieldLevel) bool {
	if str := fl.Field().String(); len(str) != 0 {
		//fmt.Println(str)
		if matched, _ := regexp.MatchString(`^\w{4,10}$`, str); matched {
			return true
		}
	}
	return false
}

//TODO:TopicList.Len()==TopicListSize
func TopicsValidate(fl validator.FieldLevel) bool {

	fmt.Println(fl.Field().Len())
	if fl.Field().Len() !=0{
		return true
	}
	return false
}
