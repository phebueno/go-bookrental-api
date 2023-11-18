package helper

import (
	"net/http"
	"strconv"
)

func ConvertIDParam(idStr string) (int, *string, int) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		message := "invalid id parameter"
		return 0, &message, http.StatusBadRequest
	}
	return id, nil, 0
}

func ConvertQueryParam(idStr string, ok bool) (int, *string, int) {
	if !ok {
		message := "invalid id query parameter"
		return 0, &message, http.StatusBadRequest
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		message := "invalid id query parameter"
		return 0, &message, http.StatusBadRequest
	}
	return id, nil, 0
}
