package magus

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

// миграция из комментариев
// *поскольку в игре первичный ключ, запрос выполнится только один раз
// !!!дальше будет выпадать ошибка в логах
/*
CREATE DATABASE database_name;
CREATE TABLE IF NOT EXISTS tree
(
    id        INT PRIMARY KEY,
    parent_id INT
);
*/
type Ids struct {
	id       int
	parentId int
}

// рекурсивно собираем структуру данных для записи в БД (массив из пар id)
// @tree вложенное дерево
func collectIds(tree *Node) []*Ids {
	// аккамулируем пары id тут
	ids := make([]*Ids, 0)
	// база рекурсии
	if len(tree.Children) == 0 {
		ids = append(ids, &Ids{id: tree.Id, parentId: tree.ParentId})
	} else {
		// шаг рекурсии
		ids = append(ids, &Ids{id: tree.Id, parentId: tree.ParentId})
		for _, childTree := range tree.Children {
			ids = append(ids, collectIds(childTree)...)
		}
	}
	return ids
}

// собираем запрос
// @tree вложенное дерево
func createQuery(tree *Node) string {
	q := "INSERT INTO tree (id, parent_id) VALUES "
	ids := collectIds(tree)

	for _, pair := range ids {
		q += fmt.Sprintf("(%d, %d),", pair.id, pair.parentId)
	}

	return strings.TrimSuffix(q, ",")
}

// сохраняем в базу,
// *поскольку функция асинхронная (запуская ее в отдельной горутине), ничего возвращать из нее не стал
// @db объект БД
// @tree вложенное дерево
func saveToDB(db *sql.DB, tree *Node) {
	q := createQuery(tree)

	_, err := db.Exec(q)

	if err != nil {
		log.Println("Ошибка, данные не были сохранены в БД")
	}
}
