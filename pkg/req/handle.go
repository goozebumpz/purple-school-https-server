package req

import (
	"fmt"
	"net/http"
	"purple-school/pkg/res"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)

	if err != nil {
		res.JSON(w, err.Error(), 400)
		return nil, err
	}

	err = IsValid[T](body)

	if err != nil {
		res.JSON(w, err.Error(), 400)
		fmt.Println(err.Error())
		return nil, err
	}

	return &body, nil
}
