package controllers

import (
	"WebApp/database"
	"WebApp/structures"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	tmpl = template.Must(template.ParseFiles("./templates/index.html"))
	db   = database.DataBase()

	id   int
	item string
	done bool
)

func Error(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func Todo(w http.ResponseWriter, r *http.Request) {
	d, e := db.Query(`SELECT id, item, done FROM todolist`)
	Error(e)

	var tasks []structures.Todo

	for d.Next() {
		d.Scan(&id, &item, &done)

		task := structures.Todo{
			Id:   id,
			Item: item,
			Done: done,
		}

		tasks = append(tasks, task)
	}

	data := structures.Data{
		Todos: tasks,
	}

	tmpl.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")

	_, e := db.Exec(`INSERT INTO todolist (item) VALUES (?)`, item)
	Error(e)

	http.Redirect(w, r, "/todo", 307)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	id := parameters["id"]

	_, e := db.Exec(`DELETE FROM todolist WHERE id = ?`, id)
	Error(e)

	http.Redirect(w, r, "/todo", 307)
}

func Done(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	id := parameters["id"]

	_, e := db.Exec("UPDATE todolist SET done = 1 WHERE id = ?", id)
	Error(e)

	http.Redirect(w, r, "/todo", 307)
}
