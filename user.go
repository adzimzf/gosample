package main

import (
	"fmt"
	"log"
)

func getUser() []User {
	//start query
	var user []User
	query := fmt.Sprintf(`
		SELECT
			 create_time,
			 COALESCE(update_time, create_time) as update_time,
			 user_id,
			 user_email,
			 full_name,
			 msisdn
			 
		FROM ws_user 
		LIMIT 10
		`)
	db := getDB()
	rows, _ := db.Queryx(query)
	for rows.Next() {
		var u User
		err := rows.StructScan(&u)
		if err != nil {
			log.Fatalln(err)
		}
		user = append(user, u)
	}
	return user
}

func getUserByName(name string) []User {
	var user []User
	query := fmt.Sprintf(`
		SELECT
			 create_time,
			 COALESCE(update_time, create_time) as update_time,
			 user_id,
			 user_email,
			 full_name,
			 msisdn
			 
		FROM ws_user 
		WHERE full_name LIKE '%s%s'
		LIMIT 10
		`, name, "%")
	db := getDB()
	fmt.Println(query)
	rows, _ := db.Queryx(query)
	for rows.Next() {
		var u User
		err := rows.StructScan(&u)
		if err != nil {
			log.Fatalln(err)
		}
		user = append(user, u)
	}
	return user
}

func getUserByPage(page string) []User {
	var user []User
	query := fmt.Sprintf(`
		SELECT
			 create_time,
			 COALESCE(update_time, create_time) as update_time,
			 user_id,
			 user_email,
			 full_name,
			 msisdn
			 
		FROM ws_user 
		LIMIT 10
		OFFSET %s
		`, page)
	db := getDB()
	rows, _ := db.Queryx(query)
	for rows.Next() {
		var u User
		err := rows.StructScan(&u)
		if err != nil {
			log.Fatalln(err)
		}
		user = append(user, u)
	}
	return user
}
