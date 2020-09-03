package magus

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/stdlib"
	"log"
	"net/http"
	"strconv"
)

// Запускаем http server (будем использовать роутер gorilla/mux)
// Создаем объект подключения к БД (postgres)
func RunServer() {
	// Ваш покорный слуга выбрал postgres, в качестве базы
	connStr := fmt.Sprintf("host=* port=5432 user=* password=* dbname=database_name sslmode=disable")
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/gen_tree", func(w http.ResponseWriter, r *http.Request) {
		// пробросываем базу в обработчик
		genTreeHandler(w, r, db)
	})
	http.Handle("/", r)
	http.ListenAndServe(":7777", nil)
}

// http обработчик для сборки дерева по заданным параметрам
// @max_lvl глубина дерева
// @n арность дерева
func genTreeHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// получаем и проверяем параметры (могли бы делать это в отдельном мидлваре)
	var maxLvl, n int
	var err error
	maxLvlQry := r.URL.Query().Get("max_lvl")
	if maxLvl, err = strconv.Atoi(maxLvlQry); err != nil || maxLvl == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Вы ввели параметр max_lvl='%s'. Пожалуйста, введите число больше нуля.", maxLvlQry)
		return
	}
	nQry := r.URL.Query().Get("n")
	if n, err = strconv.Atoi(nQry); err != nil || n == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Вы ввели параметр n='%s'. Пожалуйста, введите число больше нуля.", nQry)
		return
	}

	// создаем дерево
	tree := createTree(maxLvl, n)

	// создаем html разметку
	htmlList := createHtmlList(tree)

	// сохраняем дерево в базу
	go saveToDB(db, tree)

	// отправляем ответ клиенту
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<div>%s</div>", htmlList)
}

