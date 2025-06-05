package health

import (
	"fmt"
	"net/http"
)

// Handler responds with basic service health.
func Handler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintln(w, "ok"); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
