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

func ApplyDefaultHeaders(w http.ResponseWriter, r *http.Request, methodType string) {
	allowedOrigins := map[string]bool{
        "http://localhost:5000":    true,
        "https://app.linkwire.cc": true,
    }

    origin := r.Header.Get("Origin")

    if allowedOrigins[origin] {
        w.Header().Set("Access-Control-Allow-Origin", origin)
    }

	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", methodType+", OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, x-requested-with")
	w.Header().Set("Access-Control-Max-Age", "3600")
}