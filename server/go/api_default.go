/*
 * Sample Password Flow OAuth2 Project
 *
 * This is an example of using OAuth2 Password Flow in a specification to describe security to your API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"bufio"
	"net/http"
)

func TestPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func TestPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func CheckCredentialsTeacherPost(w http.ResponseWriter, r *http.Request) {
	var teachers []Teacher
	scanner := bufio.NewReader(r.Body)

	res, _, _ := scanner.ReadLine()
	login := string(res)

	res, _, _ = scanner.ReadLine()
	pass := string(res)

	err := db.Model(&teachers).Where("login = ? and password = ?", login, pass).Select()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	if len(teachers) == 0 {
		w.Write([]byte("no"))
	} else {
		w.Write([]byte("yes"))
	}
	w.WriteHeader(http.StatusOK)

}

func CheckCredentialsPost(w http.ResponseWriter, r *http.Request) {
	var student []Student

	scanner := bufio.NewReader(r.Body)

	res, _, _ := scanner.ReadLine()
	login := string(res)

	res, _, _ = scanner.ReadLine()
	pass := string(res)

	err := db.Model(&student).Where("email = ? and password = ?", login, pass).Select()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	if len(student) == 0 {
		w.Write([]byte("no"))
	} else {
		w.Write([]byte("yes"))
	}
	w.WriteHeader(http.StatusOK)

}
