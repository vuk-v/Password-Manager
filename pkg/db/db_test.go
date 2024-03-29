package db

import (
	"os"
	"testing"
)

func TestGetDataBase(t *testing.T) {
	_, err := GetDataBase()
	if err != nil {
		t.Error("Failed to get the database: ", err)
	}

	os.Remove("../../passwords.db")
}

func TestCreateTable(t *testing.T) {
	err := CreateTable()
	if err != nil {
		t.Error("Failed to create table: ", err)
	}

	os.Remove("passwords.db")
}
