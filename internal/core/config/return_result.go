package config

import (
	"encoding/json"
	"net/http"

	"github.com/fsnotify/fsnotify"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// RR -> for use to return result model
var (
	RR = &ReturnResult{}
)

// Language language
type Language string

const (
	// LanguageTH th
	LanguageTH Language = "th"
	// LanguageEN en
	LanguageEN Language = "en"
)

// IsValid is valid
func (l Language) IsValid() bool {
	return l == LanguageTH || l == LanguageEN
}

// String string
func (l Language) String() string {
	return string(l)
}

// Result result
type Result struct {
	Code        int               `json:"code" mapstructure:"code"`
	Description LocaleDescription `json:"message" mapstructure:"localization"`
}

// SwaggerInfoResult swagger info result
type SwaggerInfoResult struct {
	Code        int    `json:"code"`
	Description string `json:"message"`
}

// WithLocale with locale
func (rs Result) WithLocale(c *fiber.Ctx) Result {
	lacale, ok := c.Locals("lang").(Language)
	if !ok {
		rs.Description.Locale = LanguageTH
	}
	rs.Description.Locale = lacale
	return rs
}

// Error error description
func (rs Result) Error() string {
	if rs.Description.Locale == LanguageTH {
		return rs.Description.TH
	}
	return rs.Description.EN
}

// ErrorCode error code
func (rs Result) ErrorCode() int {
	return rs.Code
}

// HTTPStatusCode http status code
func (rs Result) HTTPStatusCode() int {
	switch rs.Code {
	case 0, 200: // success
		return http.StatusOK
	case 400: // bad request
		return http.StatusBadRequest
	case 404: // connection_error
		return http.StatusNotFound
	case 401: // unauthorized
		return http.StatusUnauthorized
	}

	return http.StatusInternalServerError
}

// ReturnResult return result model
type ReturnResult struct {
	JSONDuplicateOrInvalidFormat Result `mapstructure:"json_duplicate_or_invalid_format"`
	InvalidUsername              Result `mapstructure:"invalid_username"`
	InvalidPassword              Result `mapstructure:"invalid_password"`
	InvalidPermissionRole        Result `mapstructure:"invalid_permission_role"`
	InvalidCurrentPassword       Result `mapstructure:"invalid_current_password"`
	UserNotFound                 Result `mapstructure:"user_not_found"`
	InvalidPrefixUpload          Result `mapstructure:"invalid_prefix_path_upload"`
	InvalidMaximumSize           Result `mapstructure:"invalid_maximum_size"`
	UploadFileFail               Result `mapstructure:"upload_file_fail"`
	LDAP                         struct {
		GroupNotExist    Result `mapstructure:"group_not_exist"`
		UserNotExist     Result `mapstructure:"user_not_exist"`
		TooManyEntries   Result `mapstructure:"too_many_entries"`
		UserAlreadyExist Result `mapstructure:"user_already_exist"`
		DataNotExist     Result `mapstructure:"data_not_exist"`
	} `mapstructure:"ldap"`
	Internal struct {
		Success          Result `mapstructure:"success"`
		General          Result `mapstructure:"general"`
		BadRequest       Result `mapstructure:"bad_request"`
		ConnectionError  Result `mapstructure:"connection_error"`
		DatabaseNotFound Result `mapstructure:"database_not_found"`
		Unauthorized     Result `mapstructure:"unauthorized"`
	} `mapstructure:"internal"`
}

// LocaleDescription locale description
type LocaleDescription struct {
	EN     string   `mapstructure:"en"`
	TH     string   `mapstructure:"th"`
	Locale Language `mapstructure:"success"`
}

// MarshalJSON marshall json
func (ld LocaleDescription) MarshalJSON() ([]byte, error) {
	if ld.Locale == LanguageTH {
		return json.Marshal(ld.TH)
	}
	return json.Marshal(ld.EN)
}

// UnmarshalJSON unmarshal json
func (ld *LocaleDescription) UnmarshalJSON(data []byte) error {
	var res string
	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}
	ld.EN = res
	ld.Locale = LanguageEN
	return nil
}

// InitReturnResult init return result
func InitReturnResult(configPath string) error {
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName("return_result")

	if err := v.ReadInConfig(); err != nil {
		logrus.Error("read config file error:", err)
		return err
	}

	if err := bindingReturnResult(v, RR); err != nil {
		logrus.Error("binding config error:", err)
		return err
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		logrus.Info("config file changed:", e.Name)
		if err := bindingReturnResult(v, RR); err != nil {
			logrus.Error("binding error:", err)
		}
		logrus.Infof("Initial 'Return Result'. %+v", RR)
	})
	return nil
}

// bindingReturnResult binding return result
func bindingReturnResult(vp *viper.Viper, rr *ReturnResult) error {
	if err := vp.Unmarshal(&rr); err != nil {
		logrus.Error("unmarshal config error:", err)
		return err
	}
	return nil
}

// CustomMessage custom message
func (rr *ReturnResult) CustomMessage(messageEN, messageTH string, code ...int) Result {
	result := Result{
		Code: 999,
		Description: LocaleDescription{
			EN: messageEN,
			TH: messageTH,
		},
	}
	if code != nil {
		result.Code = code[0]
	}

	return result
}
