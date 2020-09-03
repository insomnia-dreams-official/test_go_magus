package magus

import "fmt"

// создаем html разметку с не нумерованными списками <ul>
// @tree вложенное дерево
func createHtmlList(tree *Node) string {
	// аккамулируем пары html тэги тут
	s := ""
	// база рекурсии
	if len(tree.Children) == 0 {
		s += fmt.Sprintf(
			"<ul><li>%d</li></ul>",
			tree.Id)
	} else {
		// шаг рекурсии
		s += fmt.Sprintf(
			"<ul><li>%d",
			tree.Id)
		for _, childTree := range tree.Children {
			s += createHtmlList(childTree)
		}
		s += "</li></ul>"
	}
	return s
}

