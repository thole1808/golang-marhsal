package kontrakvacontroller

import (
	"golang-web-service-api/helper"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := []map[string]interface{}{
		{
			"id":         10,
			"no_invoice": "KS-2323232",
		},
		{
			"id":         1,
			"no_invoice": "KS-2323232",
		},
		{
			"id":         4,
			"no_invoice": "KS-2323232",
		},
	}
	helper.ResponseJSON(w, http.StatusOK, data)
}
