package session

import (
	"fmt"
	"testing"

	_ "github.com/glebarez/go-sqlite"
)

func TestDriver_Open(t *testing.T) {
	drv := Session{}
	db, err := drv.Open("./database.db")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(db)
}

func TestDriver_Close(t *testing.T) {
	drv := Session{}
	drv.Open("./database.db")
	defer drv.Close()

}
