package controllers

import (
	"encoding/json"
	"net/http"

	"time"
	"todolist/models"
	"todolist/utils"

	"github.com/thedevsaddam/renderer"
	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("r5u8x/A?D(G+KbPeShVmYp3s6v9y$B&E") //secret key to decode the session
	store = sessions.NewCookieStore(key)
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		rnd.JSON(w, http.StatusProcessing, err)
		return
	}

	// simple validation
	if u.Email == "" && u.Password == "" {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message": "The Email and Password field are requried",
		})
		return
	}

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		panic(err)
	}

	um := models.UserModel{
		ID:        bson.NewObjectId(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}

	if err := db.C("users").Insert(&um); err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
			"message": "Failed to save user",
			"error":   err,
		})
		return
	}

	rnd.JSON(w, http.StatusCreated, renderer.M{
		"message": "User created successfully",
		"status":  true,
	})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	var u models.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		rnd.JSON(w, http.StatusProcessing, err)
		return
	}

	// simple validation
	if u.Email == "" && u.Password == "" {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message": "The Email and Password fields are requried",
		})
		return
	}
	var um models.UserModel
	err := db.C("users").
		Find(bson.M{"email": u.Email}).One(&um)

	if err != nil {
		if err.Error() == "not found" {
			rnd.JSON(w, http.StatusUnauthorized, renderer.M{
				"status":  false,
				"message": "Invalid Email or Password",
			})
			return
		} else {
			rnd.JSON(w, http.StatusProcessing, renderer.M{
				"message": "Failed to fetch user",
				"error":   err,
			})
			return
		}

	}

	if !utils.CheckPasswordHash(u.Password, um.Password) {
		rnd.JSON(w, http.StatusUnauthorized, renderer.M{
			"status":  false,
			"message": "Invalid Email or Password",
		})
		return
	}

	session.Values["authenticated"] = true
	session.Save(r, w)
	rnd.JSON(w, http.StatusOK, renderer.M{
		"status":  true,
		"message": "Login successfully",
	})
}
func IsLoggedIn(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {

		rnd.JSON(w, http.StatusForbidden, renderer.M{
			"message": "Not logged in",
			"status":  false,
		})
		return
	}

	rnd.JSON(w, http.StatusOK, renderer.M{
		"message": "Already logged in",
		"status":  true,
	})

}
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}
