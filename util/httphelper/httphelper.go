package httphelper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/schema"
	"github.com/mochammadshenna/saza-aluminium/model/api"
	"github.com/mochammadshenna/saza-aluminium/state"
	"github.com/mochammadshenna/saza-aluminium/util/exceptioncode"
	"github.com/mochammadshenna/saza-aluminium/util/helper"
	"github.com/mochammadshenna/saza-aluminium/util/logger"
)

var Decoder = schema.NewDecoder()

func init() {
	Decoder.RegisterConverter([]string{}, convertStringCommaSeparated)
}

func convertStringCommaSeparated(value string) reflect.Value {
	// fix bug if empty value should return []string{} instead of []string{""}
	if value == "" {
		return reflect.ValueOf([]string{})
	}
	return reflect.ValueOf(strings.Split(value, ","))
}

func Close(ctx context.Context, body io.Closer) {
	if body != nil {
		err := body.Close()
		if err != nil {
			logger.Errorf(ctx, "got an error while closing the body; err=%+v", err)
			return
		}
	}
}

func Read(request *http.Request, result interface{}) error {
	defer Close(request.Context(), request.Body)
	if request.Method == http.MethodPost || request.Method == http.MethodPut || request.Method == http.MethodPatch {
		var b bytes.Buffer
		_, err := io.Copy(&b, request.Body)
		if err != nil {
			logger.Error(request.Context(), strings.Replace(fmt.Sprintf("request body: %+v", b.String()), "\u0026", "", 1), err)
			return api.ErrorResponse{
				Code:    exceptioncode.CodeInvalidRequest,
				Message: err.Error(),
			}
		}

		jsonDecoder := json.NewDecoder(&b)
		err = jsonDecoder.Decode(&result)
		if err != nil && err != io.EOF {
			logger.Error(request.Context(), strings.Replace(fmt.Sprintf("request body: %+v", b.String()), "\u0026", "", 1), err)
			return api.ErrorResponse{
				Code:    exceptioncode.CodeInvalidRequest,
				Message: err.Error(),
			}
		}

		return nil
	}

	err := Decoder.Decode(result, request.URL.Query())
	if err != nil {
		logger.Error(request.Context(), strings.Replace(fmt.Sprintf("request params: %s", request.URL.RawQuery), "\u0026", ", ", -1), err)
		return parseError(err)
	}

	return nil
}

func Write(ctx context.Context, writer http.ResponseWriter, data interface{}) {
	response := api.ApiResponse{
		Header: getHeader(writer),
		Data:   data,
	}
	write(ctx, writer, response)
}

func WriteError(ctx context.Context, writer http.ResponseWriter, errorResponse error) {
	httpError, ok := errorResponse.(api.HttpError)
	if ok {
		WriteErrorWithStatusCode(ctx, writer, errorResponse, httpError.StatusCode())
		return
	}
	writer.WriteHeader(http.StatusBadRequest)
	response := api.ApiResponse{
		Header: getHeader(writer),
		Error:  errorResponse,
	}
	write(ctx, writer, response)
}

func WriteErrorWithStatusCode(ctx context.Context, writer http.ResponseWriter, errorResponse error, statusCode int) {
	defaultStatusCode := http.StatusBadRequest
	if statusCode >= 400 && statusCode <= 599 {
		defaultStatusCode = statusCode
	}
	writer.WriteHeader(defaultStatusCode)
	response := api.ApiResponse{
		Header: getHeader(writer),
		Error:  errorResponse,
	}
	write(ctx, writer, response)
}

func write(ctx context.Context, writer http.ResponseWriter, response api.ApiResponse) {
	data, err := response.ToJSON()
	helper.PanicError(err)
	writer.Write(data)
}

func getHeader(writer http.ResponseWriter) api.HeaderResponse {
	headerResponse := api.HeaderResponse{
		ServerTimeMs: time.Now().Unix(),
		RequestId:    writer.Header().Get(state.HttpHeaders().RequestId.String()),
	}

	startTimeHeader := writer.Header().Get(state.HttpHeaders().StartTime.String())
	if len(startTimeHeader) > 0 {
		startTime, _ := strconv.ParseInt(startTimeHeader, 10, 64)
		headerResponse.ProcessTimeMs = time.Since(time.Unix(0, startTime)).Milliseconds()
	}

	return headerResponse
}

func parseError(err error) error {
	errors := []api.ErrorValidate{}
	new := err.(schema.MultiError)
	for i, a := range new {
		errors = append(errors, api.ErrorValidate{
			Key:     i,
			Code:    "VALIDATION",
			Message: a.Error(),
		})
	}
	return api.ErrorResponse{
		Code:    exceptioncode.CodeInvalidValidation,
		Message: "validation error",
		Errors:  errors,
	}
}
