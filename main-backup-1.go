package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

func getAllData(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	if r.Method == "GET" {
		res, err := json.Marshal(albums)

		// jika ada error dalam get maka keluar perintah ini
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(res)
		return
	}

	// jika ada request selain method Get maka akan tampil error perintah ini
	http.Error(w, "400 status bad request", http.StatusBadRequest)
}

func getByTitle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	if r.Method == "GET" {
		title := r.FormValue("title")

		for _, v := range albums {
			if v.Title == title {
				result, err := json.Marshal(v)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}
	}

	// jika ada request selain method Get maka akan tampil error perintah ini
	http.Error(w, "400 status bad request", http.StatusBadRequest)
}

func main1() {
	// Ini dengan Framework Gin
	// router := gin.Default()
	// router.GET("/albums", getAlbums)

	// router.Run("localhost:8989")

	// Route
	http.HandleFunc("/albums", getAllData)
	http.HandleFunc("/albums/title", getByTitle)
	http.ListenAndServe(":8000", nil)

}
