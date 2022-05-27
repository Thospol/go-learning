package context

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/Thospol/go-learning/internal/core/config"
	"github.com/Thospol/go-learning/internal/core/sql"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	pathKey            = "path"
	compositeFormDepth = 3
	// UserKey user key
	UserKey = "user"
	// LangKey lang key
	LangKey = "lang"
	// ParametersKey parameters key
	ParametersKey = "parameters"
	// DatabaseKey database  key
	DatabaseKey = "sql_database"
)

// Context context
type Context struct {
	*fiber.Ctx
}

// New new custom fiber context
func New(c *fiber.Ctx) *Context {
	return &Context{c}
}

// BindValue bind value
func (c *Context) BindValue(i interface{}, validate bool) error {
	switch c.Method() {
	case http.MethodGet:
		_ = c.QueryParser(i)

	default:
		_ = c.BodyParser(i)
	}

	c.PathParser(i, 1)
	c.Locals(ParametersKey, i)
	c.TrimSpace(i, 1)
	return nil
}

// PathParser parse path param
func (c *Context) PathParser(i interface{}, depth int) {
	formValue := reflect.ValueOf(i)
	if formValue.Kind() == reflect.Ptr {
		formValue = formValue.Elem()
	}
	t := reflect.TypeOf(formValue.Interface())
	for i := 0; i < t.NumField(); i++ {
		fieldName := t.Field(i).Name
		paramValue := formValue.FieldByName(fieldName)
		if paramValue.IsValid() {
			if depth < compositeFormDepth && paramValue.Kind() == reflect.Struct {
				depth++
				c.PathParser(paramValue.Addr().Interface(), depth)
			}
			tag := t.Field(i).Tag.Get(pathKey)
			if tag != "" {
				setValue(paramValue, c.Params(tag))
			}
		}
	}
}

func setValue(paramValue reflect.Value, value string) {
	if paramValue.IsValid() && value != "" {
		switch paramValue.Kind() {
		case reflect.Uint:
			number, _ := strconv.ParseUint(value, 10, 32)
			paramValue.SetUint(number)

		case reflect.String:
			paramValue.SetString(value)

		default:
			number, err := strconv.Atoi(value)
			if err != nil {
				paramValue.SetString(value)
			} else {
				paramValue.SetInt(int64(number))
			}
		}
	}
}

// TrimSpace trim space
func (c *Context) TrimSpace(i interface{}, depth int) {
	e := reflect.ValueOf(i).Elem()
	for i := 0; i < e.NumField(); i++ {
		if depth <= compositeFormDepth && e.Type().Field(i).Type.Kind() == reflect.Struct {
			depth++
			c.TrimSpace(e.Field(i).Addr().Interface(), depth)
		}

		if e.Type().Field(i).Type.Kind() != reflect.String {
			continue
		}

		value := e.Field(i).String()
		e.Field(i).SetString(strings.TrimSpace(value))
	}
}

// GetDatabase get connection database
func (c *Context) GetDatabase() *gorm.DB {
	val := c.Locals(DatabaseKey)
	if val == nil {
		return sql.Database
	}

	return val.(*gorm.DB)
}

// Localizer localizer
type Localizer interface {
	WithLocale(c *Context)
	GetLanguage() config.Language
}

// Localization localization
func (c *Context) Localization(i interface{}, depth int) {
	const (
		key = "Localization"
	)
	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice:
		s := reflect.ValueOf(i)
		for i := 0; i < s.Len(); i++ {
			c.Localization(s.Index(i).Interface(), depth)
		}

	default:
		formValue := reflect.ValueOf(i)
		if !equal(formValue.Kind(), reflect.Map) {
			localizer, ok := formValue.Interface().(Localizer)
			if ok {
				localizer.WithLocale(c)
			}
			if equal(formValue.Kind(), reflect.Ptr) {
				formValue = formValue.Elem()
			}
			if equal(formValue.Kind(), reflect.Struct) {
				t := reflect.TypeOf(formValue.Interface())
				for i := 0; i < t.NumField(); i++ {
					fieldName := t.Field(i).Name
					value := formValue.FieldByName(fieldName)
					if value.IsValid() {
						if fieldName == key && equal(value.Kind(), reflect.Struct) {
							t := reflect.TypeOf(value.Interface())
							for i := 0; i < t.NumField(); i++ {
								fieldName := t.Field(i).Name
								value := formValue.FieldByName(fieldName)
								if value.IsValid() {
									switch localizer.GetLanguage() {
									case "th":
										v := formValue.FieldByName(fmt.Sprintf("%sTH", fieldName))
										if v.IsValid() {
											value.Set(v)
										}
									default:
										v := formValue.FieldByName(fmt.Sprintf("%sEN", fieldName))
										if v.IsValid() {
											value.Set(v)
										}
									}
								}
							}
						}
						if depth <= compositeFormDepth && equal(value.Kind(), reflect.Slice) {
							depth++
							c.Localization(value.Interface(), depth)
						}
					}
				}
			}
		}
	}
}

func equal(a, b reflect.Kind) bool {
	return a == b
}
