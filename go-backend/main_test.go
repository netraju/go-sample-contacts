// main_test.go

package main

import (
	"log"
	"os"
	"testing"

	"bytes"
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"strconv"

	"github.com/joho/godotenv"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	godotenv.Load("../.env")
	a.Initialize(
		os.Getenv("TEST_DB_SERVER"),
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"))

	//ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM contact_phones ;DELETE FROM contacts;")
	a.DB.Exec("ALTER SEQUENCE contacts RESTART WITH 1ALTER SEQUENCE contact_phones RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE contacts
(
    id SERIAL PRIMARY KEY,
    first_name varchar(30) NOT NULL,
    last_name varchar(30),
    email varchar(30),
    UNIQUE(email) 
);
CREATE TABLE contact_phones
(
    id SERIAL,
    contact_id int NOT NULL REFERENCES contacts,
    phone_number varchar(20) NOT NULL
    
);`

// tom: next functions added later, these require more modules: net/http net/http/httptest
func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/contacts", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}
func TestGetNonExistentContact(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/contacts/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Contact not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Contact not found'. Got '%s'", m["error"])
	}
}

// tom: rewritten function
func TestCreateContact(t *testing.T) {

	clearTable()

	var jsonStr = []byte(`{"first_name":"john", "last_name":"doe", "email":"john.doe@example.com","phone_numbers":["+61XXXXXXXXX","+61XXXXXXYYY"]}`)
	req, _ := http.NewRequest("POST", "/contact", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["first_name"] != "john" {
		t.Errorf("Expected Contact name to be 'john'. Got '%v'", m["name"])
	}

	if m["email"] != "john.doe@example.com" {
		t.Errorf("Expected Contact email to be 'john.doe@example.com'. Got '%v'", m["price"])
	} 
}

func TestGetContacts(t *testing.T) {
	clearTable()
	var c contact 
	c.First_name = "john"
	c.Last_name = "doe"
	c.Email = "john.doe@example.com"
	c.Phones = ["+61xxxxxxxxx","+61yyyyyyyyy"]
	addContacts(c)

	req, _ := http.NewRequest("GET", "/contacts/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func addContacts(c contact) {
	if count < 1 {
		count = 1
	} 
	//TODO Complete the save contact and phones
	a.DB.Exec("INSERT INTO contacts")
	 
} 

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
