package validator

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"strings"
)

func decodeRequestBody[T any](closer io.Reader) (*T, error) {

	output := new(T)

	if err := json.NewDecoder(closer).Decode(output); err != nil {
		return nil, errors.New("bad request. Invalid JSON")
	}

	return output, nil

}

func errorMsgForTag(tag string, field string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("%v is required", field)
	default:
		return fmt.Sprintf("%v is invalid", field)
	}

}

func validateInput(input any) error {

	customValidator := validator.New(validator.WithRequiredStructEnabled())

	if err := customValidator.Struct(input); err != nil {
		// TODO HAndle other errors
		validationErrors := err.(validator.ValidationErrors)

		errorArray := make([]string, len(validationErrors))

		for index, validationErr := range validationErrors {
			field := validationErr.Field()

			msg := errorMsgForTag(validationErr.Tag(), field)
			errorArray[index] = msg

		}
		return errors.New(strings.Join(errorArray, ", "))
	}

	return nil

}

func DecodeAndValidateRequestBody[T any](closer io.Reader) (*T, error) {

	input, err := decodeRequestBody[T](closer)

	if err != nil {
		return nil, err
	}

	err = validateInput(input)

	if err != nil {
		return nil, err
	}

	return input, nil

}
