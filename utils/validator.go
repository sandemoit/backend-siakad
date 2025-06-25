package utils

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate = validator.New()

func ValidateStruct(data interface{}) map[string]string {
	err := validate.Struct(data)
	if err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = fmt.Sprintf("%s is %s", err.Field(), err.Tag())
		}
		return errors
	}
	return nil
}

func NormalizeKey(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, "/", "_")
	name = strings.ReplaceAll(name, "-", "_")
	return name
}

func IsValidUUID(uuid string) bool {
	regex := regexp.MustCompile(`(?i)^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`)
	return regex.MatchString(uuid)
}

func GenerateUID() string {
	return uuid.New().String()
}

func NowUnix() int64 {
	return time.Now().Unix()
}

func Slugify(s string) string {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	// Replace non-alphanumeric characters with hyphens
	re := regexp.MustCompile(`[^a-z0-9]+`)
	s = re.ReplaceAllString(s, "")
	// Remove leading/trailing hyphens
	s = strings.Trim(s, "-")
	return s
}
func GenerateInvoiceNumber() string {
	now := time.Now()
	return fmt.Sprintf("INV-%s-%d", now.Format("20060102"), now.Unix())
}
