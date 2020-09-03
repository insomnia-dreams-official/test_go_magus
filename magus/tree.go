package magus

import "fmt"

// структура для элемента дерева
type Node struct {
	Id       int
	ParentId int
	Children []*Node
}

// функция создающая полное дерево по параметрам:
// @max_lvl глубина дерева
// @n арность дерева
func createTree(maxLvl int, n int) *Node {
	// хардкодим корневой элемент (кейсы когда его нет отловлены в проверке query params)
	// *для удобства сделал карту, крутая штука для подзадачи с БД
	// *могли бы пройти ней (карте) одним циклом и собрать данные для insert'а, но я решил не портить интерфейсы *Node
	// *и пить чашу с рекурсивными функциями до дна :)
	nodeMap := map[int]*Node{
		1: &Node{Id: 1, ParentId: 0, Children: []*Node{}},
	}

	// для создания близнецов текущего уровня, я сохранил их родителей с предэдущего
	previousLevelParentIds := []int{1}
	// эту переменную мы свопнем с предэдущей после сборки текущего уровня,
	// *как в задаче Кантора про факториалы
	currentLevelParentIds := []int{}
	// будем сохранять максимальый id элемента дерева
	currentMaxId := 1

	// поскольку первый уровень мы уже захардкодили, залетаем в игру сразу со второго
	for lvl := 2; lvl <= maxLvl; lvl++ {
		// начинаем собирать близнецов на каждого родителя с предэдущего уровня
		for _, parentId := range previousLevelParentIds {
			for i := 1; i <= n; i++ {
				// обновляем текущий id
				currentMaxId += 1
				currentLevelParentIds = append(currentLevelParentIds, currentMaxId)
				// *тут могли бы не проверять на наличие предка, так как сами их генерим
				// *но если задача была бы хитрее, такая проверка бы не помешала
				parent, ok := nodeMap[parentId]
				if !ok {
					fmt.Printf("Ошибка, parentId=%v: не найден\n", parentId)
					return nil
				}
				// создаем текущий элемент
				node := &Node{Id: currentMaxId, ParentId: parentId, Children: []*Node{}}
				// сохраняем его в карту
				nodeMap[currentMaxId] = node
				// добавляем элемент в родителя
				parent.Children = append(parent.Children, node)
			}
		}
		// готовим переменные к следующему уровню (reset)
		previousLevelParentIds = currentLevelParentIds
		currentLevelParentIds = []int{}
	}

	// отдаем корень, выше я описал почему взял карту
	return nodeMap[1]
}