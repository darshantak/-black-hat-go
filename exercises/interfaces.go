package exercises

import (
	"log"
)

type dbContract interface {
	Close()
	InsertUser(userName string) error
	SelectSingleUser(userName string) (string, error)
}

type Application struct {
	db dbContract
}

func (this Application) Run() {
	err := this.db.InsertUser("")
	if err != nil {
		log.Println(err)
	}

}
func NewApplication(db dbContract) *Application {
	return &Application{db: db}
}

func InterfaceMain() {
	// db, err := mysqldb.New("", "", "", "", "")
	// app := NewApplication(db)
}
