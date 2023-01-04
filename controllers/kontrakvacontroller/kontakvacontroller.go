package kontrakvacontroller

import (
	"golang-web-service-api/helper"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	// birdJson := `[{"species":"pigeon","decription":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
	// var birds []models.User
	// json.Unmarshal([]byte(birdJson), &birds)
	// helper.ResponseJSON(w, http.StatusOK, birds)
	// fmt.Printf("Birds : %+v", birds)

	// birdData := map[string]any{
	// 	"birdSounds": map[string]string{
	// 		"pigeon": "coo",
	// 		"eagle":  "squak",
	// 	},
	// 	"total birds": 2,
	// }

	// data := []map[string]interface{}{
	// 	{
	// 		"id":         10,
	// 		"no_invoice": "KS-2323232",
	// 	},
	// 	{
	// 		"id":         1,
	// 		"no_invoice": "KS-2323232",
	// 	},
	// 	{
	// 		"id":         4,
	// 		"no_invoice": "KS-2323232",
	// 	},
	// }

	data := map[string]any{
		"metadata": map[string]string{
			"code":    "200",
			"message": "sukses",
		},
		"total birds": 2,
	}
	helper.ResponseJSON(w, http.StatusOK, data)

	// birdData := map[string]any{
	// 	"birdSounds": map[string]string{
	// 		"pigeon": "coo",
	// 		"eagle":  "squak",
	// 	},
	// 	"total birds": 2,
	// }

	// JSON encoding is done the same way as before
	// datas, _ := json.Marshal(birdData)
	// helper.ResponseJSON(w, http.StatusOK, datas)
	// fmt.Println(string(data))
}
