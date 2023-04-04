package apperror

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
	"testing"
)

var (
	validate = validator.New()
)

func init() {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func TestCustomValidationError(t *testing.T) {
	type RequestParams struct {
		ID string `json:"id" form:"id" validate:"required,uuid"`
	}

	type TestCase struct {
		data    RequestParams
		errType string
		want    []map[string]string
	}
	cases := map[string]TestCase{
		"validation error occurred": {
			data: RequestParams{
				ID: "askjbaskdb1129",
			},
			errType: "validation",
			want: []map[string]string{
				{
					"id": "has to be a uuid",
				},
			},
		},
		"validation error missing param occurred": {
			data:    RequestParams{},
			errType: "validation",
			want: []map[string]string{
				{
					"id": "is required",
				},
			},
		},
	}

	for k, v := range cases {
		t.Run(k, func(t *testing.T) {
			var err error
			var got []map[string]string
			switch v.errType {
			case "validation":
				err = validate.Struct(v.data)
				got = CustomValidationError(&v.data, err)
			case "unmarshal":
				jsonData := []byte(`{"Limit": true}`)
				params := &RequestParams{}
				err = json.Unmarshal(jsonData, params)
				got = CustomValidationError(&jsonData, err)
			case "other":
				data := struct {
					email string
				}{
					email: "1",
				}

				err = validate.Struct(data)
				got = CustomValidationError(&data, err)
			case "nocustom":
				type X struct {
					ID string `validate:"required,uuid"`
				}
				data := X{
					ID: "notUUID",
				}
				err = validate.Struct(data)
				got = CustomValidationError(&data, err)
			}

			if !reflect.DeepEqual(v.want, got) {
				t.Errorf("expectations mismatched: \n want: %v \n got: %v", v.want, got)
			}
		})
	}
}
