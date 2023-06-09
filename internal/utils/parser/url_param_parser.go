package parser

import (
	"net/http"

	"github.com/gorilla/schema"

	"assessment/internal/utils/validators"
)

type UrlParamsParserInt interface {
	ParseUrlParams(r *http.Request, data interface{}) error
}

type UrlParamsParser struct {
	structValidator validators.StructValidatorInt
	schemaDecoder   *schema.Decoder
}

func NewUrlParamsParser(structValidator validators.StructValidatorInt, schemaDecoder *schema.Decoder) *UrlParamsParser {
	return &UrlParamsParser{
		structValidator: structValidator,
		schemaDecoder:   schemaDecoder,
	}
}

func (upp *UrlParamsParser) ParseUrlParams(r *http.Request, data interface{}) error {
	if err := upp.schemaDecoder.Decode(data, r.URL.Query()); err != nil {
		return &UrlParamParserError{err.Error()}
	}

	if err := upp.structValidator.Validate(data); err != nil {
		return err
	}

	return nil
}

type UrlParamParserError struct {
	ErrorMessage string
}

func (upe *UrlParamParserError) Error() string {
	return "Error while parsing url params : " + upe.ErrorMessage
}
