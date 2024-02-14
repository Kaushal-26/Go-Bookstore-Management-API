// Usage to make functions use

package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func PrintBook(w http.ResponseWriter, print any) {
	res, _ := json.Marshal(print)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func BookIdFind(r *http.Request) int64 {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing book by Id")
	}
	return ID
}

// Error Handling for the code
func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}
