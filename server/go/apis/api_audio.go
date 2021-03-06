package apis

import (
	"../DbWorker"
	"../Roles"
	Model "../models"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func AudioStudentIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "audio/mpeg")
	if r.Header.Get("role") == Roles.RolesText(Roles.Student) {
		w.WriteHeader(http.StatusForbidden)
		_, err := w.Write([]byte("У вас нет полномочий для этого действия."))
		if err != nil {
			log.Println(err.Error())
		}
		log.Print(http.StatusText(http.StatusForbidden))
		return
	}
	studId, err := strconv.ParseInt(mux.Vars(r)["studentId"], 10, 64)
	var path string

	err = DbWorker.Db.Model((*Model.Audio)(nil)).
		Column("path").
		Where("student_id = ?", studId).
		Select(&path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
			log.Println(err.Error())
		return
	}
	audioFile, err := ioutil.ReadFile(path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err.Error())
		}
		return
	}
	_, err = w.Write(audioFile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
	}
	w.WriteHeader(http.StatusOK)
}

func AudioStudentIdPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Header.Get("role") == Roles.RolesText(Roles.Student) {
		w.WriteHeader(http.StatusForbidden)
		_, err := w.Write([]byte("У вас нет полномочий для этого действия."))
		if err != nil {
			log.Println(err.Error())
		}
		log.Print(http.StatusText(http.StatusForbidden))
		return
	}

	studId, err := strconv.ParseInt(mux.Vars(r)["studentId"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	err = r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	file, _, err := r.FormFile("file") //retrieve the file from form data
	defer file.Close()                 //close the file when we finish

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//this is path which  we want to store the file'
	err = os.MkdirAll("./audios/", os.ModePerm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	filename:=strings.Join([]string{strconv.FormatInt(studId, 10),time.Now().Format("02_01_2006_15_04_05")}, "_")
	fp, err := filepath.Abs("./audios/" +filename  + ".mp3")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	f, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println(fp)
	defer f.Close()
	io.Copy(f, file)
	var audio = Model.Audio{StudentId: studId,
		Path: fp,
	}

	_, err = DbWorker.Db.Model(&audio).OnConflict("(student_id) DO UPDATE").Set("path = ?",audio.Path).Insert()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
