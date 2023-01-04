package authcontroller

import (
	"encoding/json"
	"golang-web-service-api/helper"
	"golang-web-service-api/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// fungsi untuk login
func Login(w http.ResponseWriter, r *http.Request) {

}

// fungsi untuk register
func Register(w http.ResponseWriter, r *http.Request) {
	// mengambil inputan dari json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userInput); err != nil {
		// log.Fatal("Gagal Mendecode JSON")
		// fmt.Println("Gagal Mendecode JSON")
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
	}

	// close
	defer r.Body.Close()

	// hasing paassword menggunakan bycypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)
	// log.Fatal("data berhasil ke post", &userInput)

	// insert ke database
	if err := models.DB.Create(&userInput).Error; err != nil {
		// log.Fatal("Gagal Menyimpan data")
		// fmt.Println("Gagal Menyimpan data")
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)

}

// fungsi untuk logiut
func Logout(w http.ResponseWriter, r *http.Request) {

}
