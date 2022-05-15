package models

import (
	"fmt"
	"regexp"

	"github.com/asaskevich/govalidator"
)

func validRegex() {
	govalidator.CustomTypeTagMap.Set("nameRegex", func(v interface{}, context interface{}) bool {
		r := regexp.MustCompile("^[a-zA-Zก-๙-_/]+$")
		return r.MatchString(fmt.Sprintf("%v", v))
	})
	govalidator.CustomTypeTagMap.Set("nameWithDigitRegex", func(v interface{}, context interface{}) bool {
		r := regexp.MustCompile("^[0-9a-zA-Zก-๙-_/]+$")
		return r.MatchString(fmt.Sprintf("%v", v))
	})
	govalidator.CustomTypeTagMap.Set("resourceNameRegex", func(v interface{}, context interface{}) bool {
		r := regexp.MustCompile("^[0-9a-zA-Z-]+$")
		return r.MatchString(fmt.Sprintf("%v", v))
	})
	govalidator.CustomTypeTagMap.Set("sha512Regex", func(v interface{}, context interface{}) bool {
		r := regexp.MustCompile("^[a-f0-9]{128}$")
		return r.MatchString(fmt.Sprintf("%v", v))
	})
	govalidator.CustomTypeTagMap.Set("timeRegex", func(v interface{}, context interface{}) bool {
		r := regexp.MustCompile("^([0-9]|0[0-9]|1[0-9]|2[0-3]):([0-9]|[0-5][0-9])$")
		return r.MatchString(fmt.Sprintf("%v", v))
	})
	govalidator.CustomTypeTagMap.Set("dateRegex", func(v interface{}, context interface{}) bool {
		r := regexp.MustCompile("^[0-9]{4}-([1-9]|0[1-9]|1[0-2])-([0-9]|0[0-9]|1[0-9]|2[0-9]|3[0-1])$")
		return r.MatchString(fmt.Sprintf("%v", v))
	})
	govalidator.CustomTypeTagMap.Set("telNumberRegex", func(v interface{}, context interface{}) bool {
		r := regexp.MustCompile("^0[0-9]{9}$")
		return r.MatchString(fmt.Sprintf("%v", v))
	})
	govalidator.ParamTagMap["digitRegex"] = govalidator.ParamValidator(func(str string, params ...string) bool {
		r := regexp.MustCompile(fmt.Sprintf("^[0-9]{%s}$", params[0]))
		return r.MatchString(str)
	})
}
