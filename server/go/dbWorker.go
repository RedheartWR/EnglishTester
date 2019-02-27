package swagger

import (
	"encoding/json"
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
)

var db *pg.DB

const dbPASS = "aaa"          //CHANGE HERE
const addr = "localhost:5432" //"138.68.78.205:5432"

//TODO override

func InitDB() {
	db = pg.Connect(&pg.Options{
		Addr:     addr,
		User:     "postgres",
		Password: dbPASS,
	})
}

func TokenExists(token string) (bool, error) {
	var student Student
	var teacher Teacher

	exists, err := db.Model(&student).Where("access_token = ?", token).Exists()
	if err == nil {
		return exists, err
	}

	exists, err = db.Model(&teacher).Where("access_token = ?", token).Exists()
	if err == nil {
		return exists, err
	}
	return false, nil
}

func GiveStudentToken(student *Student, token string) error {
	student.AccessToken = token
	_, err := db.Model(student).WherePK().Update()
	return err
}

func GiveTeacherToken(teacher *Teacher, token string) error {
	teacher.AccessToken = token
	_, err := db.Model(teacher).WherePK().Update()
	return err
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*Student)(nil), (*Test)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateSchemaStudents() error {
	for _, model := range []interface{}{(*Student)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateSchemaTeachers() error {
	for _, model := range []interface{}{(*Teacher)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateSchemaTest() error {
	for _, model := range []interface{}{(*Test)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateSchemaAudio() error {
	for _, model := range []interface{}{(*Audio)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertTests() error {

	var tests []Test = []Test{
		Test{Id: 1,
			Questions: []Question{{1, "dsds"}, {1, "vopros"}, {2, "kekus"}},
			Answers:   []string{"otvet1", "otvet2", "otvet3"},
		},
		Test{Id: 2,
			Questions: []Question{{1, "vopros1"}, {1, "vopros2"}, {2, "vopros3"}},
			Answers:   []string{"otv1", "otv2", "otv3"},
		},
	}
	jsoned, err := json.Marshal(tests)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(string(jsoned))
	_, err = db.Model(&tests).Insert()
	if err != nil {
		log.Print(err)
	}
	return err
}
