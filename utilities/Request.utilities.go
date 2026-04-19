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
	cookie, err := r.Cookie("access_token")

    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return ""
    }

	return cookie.Value
}