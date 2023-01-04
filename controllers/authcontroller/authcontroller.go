package authcontroller

import (
	"encoding/json"
	"golang-web-service-api/config"
	"golang-web-service-api/helper"
	"golang-web-service-api/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// fungsi untuk login
func Login(w http.ResponseWriter, r *http.Request) {
	// mengambil inputan dari json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userInput); err != nil {
		// log.Fatal("Gagal Mendecode JSON")
		// fmt.Println("Gagal Mendecode JSON")
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// close
	defer r.Body.Close()

	// ambil data user berdasarkan username
	var user models.User
	if err := models.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			// Case jika data tidak ditemukan
			response := map[string]string{"message": "Username atau Password anda salah"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			// Case Internal Server Error
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	// cek apakah password valid atau tidak
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "Username atau Password anda salah"}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// proses pembuatan token jwt
	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	//mendeklarasikan algoritma yang akan di gunakan untuk signing
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// set token yang ke cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	// SUCCESS
	response := map[string]string{"message": "Login Berhasil", "data": token}
	helper.ResponseJSON(w, http.StatusOK, response)
	return

	// costumize response JSON
	// response := map[string]interface{}{
	// 	"metadata": map[string]interface{}{
	// 		"message":       "Sukses",
	// 		"response_code": http.StatusOK,
	// 	},
	// 	"data": map[string]interface{}{
	// 		"token":    token,
	// 		"username": user.NamaLengkap,
	// 		"password": user.Password,
	// 	},
	// }
	// helper.ResponseJSON(w, http.StatusOK, response)
	// return

	// SUCCESS
	// response := map[string]string{"message": "Login Berhasil", "data": token}
	// helper.ResponseJSON(w, http.StatusOK, response)
	// return

	// response := map[string]interface{}{
	// 	"metadata": map[string]interface{}{
	// 		"message":       "Sukses",
	// 		"response_code": http.StatusOK,
	// 	},
	// 	"data": map[string]interface{}{
	// 		"token":    token,
	// 		"username": user.NamaLengkap,
	// 		"password": user.Password,
	// 	},
	// }

	// response := map[string]interface{}{
	// 	"Name": "Wednesday",
	// 	"Age":  6,
	// 	"Parents": []interface{}{
	// 		"Gomez",
	// 		"Morticia",
	// 	},
	// }
	// helper.ResponseJSON(w, http.StatusOK, response)
	// return

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
		return
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
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)

}

// fungsi untuk logiut
func Logout(w http.ResponseWriter, r *http.Request) {
	// set token yang ke cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Path:  "/",
		Value: "",
		// HttpOnly: true,
		MaxAge: -1,
	})

	// SUCCESS
	response := map[string]string{"message": "Logout Berhasil"}
	helper.ResponseJSON(w, http.StatusOK, response)
	// return
}
