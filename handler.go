package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	CreateTime time.Time `db:"create_time"`
	ID         int64     `db:"user_id"`
	Email      string    `db:"user_email"`
	Fullname   string    `db:"full_name"`
	Messenger  string    `db:"messenger"`
	MSISDN     string    `db:"msisdn"`
	Sex        int       `db:"sex"`
	Status     int       `db:"status"`
	BirthDate  time.Time `db:"birth_date"`
	UpdateTime time.Time `db:"update_time"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//define template
	template, err := template.ParseFiles("template/home.html")
	if err != nil {
		fmt.Println(err)
	}
	uCount := getCount()
	user := getUser()
	data := struct {
		User   []User
		UCount string
	}{user, uCount}

	template.ExecuteTemplate(w, "home.html", data)

}

func filterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	nama := r.FormValue("nama")
	data := make(map[string][]User)
	data["data"] = getUserByName(nama)
	res, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(res)
}

func paginate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	page := r.FormValue("page")

	//convert page to int
	pageI, err := strconv.Atoi(page)
	if err != nil {
		fmt.Println(err)
	}
	nPage := 0
	if pageI != 1 {
		nPage = (pageI - 1) * 10
	}
	pageS := strconv.Itoa(nPage)
	data := make(map[string][]User)
	data["data"] = getUserByPage(pageS)
	res, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(res)
}
