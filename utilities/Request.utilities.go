package utilities

import (
	"net/http"
	"slices"
	"strings"
)

func GetVariableFromPath(r *http.Request, functionName string) string {
	url := r.URL.Path
	
	tokens := strings.Split(url, "/")
	slices.Reverse(tokens)
	lastToken := tokens[0]

	if lastToken == functionName || lastToken == functionName+"/" {
		return ""
	}

	return lastToken
}

func GetTokenFromCookies(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie("token")

    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return ""
    }

	return cookie.Value
}

func ApplyDefaultHeaders(w http.ResponseWriter, methodType string) {
	w.Header().Set("Access-Control-Allow-Methods", methodType+", OPTIONS")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With, x-requested-with")
}