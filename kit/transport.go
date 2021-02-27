package kit

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// http 服务
func DecodeRequest(c context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println("EncodeResponse: ", err)
	}
	fmt.Println(response)
	return err
}

