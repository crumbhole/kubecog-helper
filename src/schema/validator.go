package schema

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type validationError struct {
	tag             string
	actualTag       string
	namespace       string
	structNamespace string
	field           string
	structField     string
	value           interface{}
	param           string
	fieldKind       reflect.Kind
	fieldType       reflect.Type
	errorMsg        string
}

// KubecogValidator is a struct to hold things we don't want to recreate
// between validations
var (
	validate    *validator.Validate
	uni         *ut.UniversalTranslator
	trans       *ut.Translator
	initialised bool
)

func prepare() {
	if !initialised {
		validate = validator.New()
		en := en.New()
		uni = ut.New(en, en)
		ltrans, _ := uni.GetTranslator(`en`)
		trans = &ltrans
		en_translations.RegisterDefaultTranslations(validate, *trans)
		initialised = true
	}
}

func doValidation(values CogValues) []validationError {
	prepare()
	var errors []validationError
	err := validate.Struct(values)
	if err != nil {
		// // this check is only needed when your code could produce
		// // an invalid value for validation such as interface with nil
		// // value most including myself do not usually have code like this.
		// if _, ok := err.(*validator.InvalidValidationError); ok {
		// 	fmt.Println(err)
		// 	return errors
		// }
		for _, err := range err.(validator.ValidationErrors) {
			var e validationError
			e.tag = err.Tag()
			e.actualTag = err.ActualTag()
			e.namespace = err.Namespace()
			e.structNamespace = err.StructNamespace()
			e.field = err.Field()
			e.structField = err.StructField()
			e.value = err.Value()
			e.param = err.Param()
			e.fieldKind = err.Kind()
			e.fieldType = err.Type()
			e.errorMsg = err.Translate(*trans)
			errors = append(errors, e)
		}
	}
	return errors
}

func errorString(err validationError) string {
	return err.namespace + ": " + err.errorMsg
}

// ValidateToStrings returns all failures as user presentable text
// in an array. 0 length array indicates no failures
func ValidateToStrings(values CogValues) []string {
	errors := doValidation(values)
	var out []string
	for _, err := range errors {
		out = append(out, errorString(err))
	}
	return out
}

// ValidateToStrings returns all failures as user presentable single
// line breaked string. 0 length text indicates no failures
func ValidateToSingleString(values CogValues) string {
	return strings.Join(ValidateToStrings(values), "\n")
}

// ValidateToError returns all failures as a user presentable error
// or nil if no errors
func ValidateToError(values CogValues) error {
	errorString := strings.Join(ValidateToStrings(values), "\n")
	if len(errorString) > 0 {
		return errors.New(errorString)
	}
	return nil
}
