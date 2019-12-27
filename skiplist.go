package skiplist

import (
	"fmt"
	"math/rand"
)

const (
	PROBABILITY = 0.25
)

//可比较接口
type Comparable interface {
	CompareTo(interface{}) int //大于返回1 等于返回0 小于返回-1
	String() string            //字符化
}

type SkipList struct {
	Level  int
	Head   *Node
	Tail   *Node
	Length int
}

type Node struct {
	Forward []*Node
	Key     Comparable
	Value   interface{}
}

func newNode(key Comparable, value interface{}, level int) *Node {
	return &Node{
		Forward: make([]*Node, level),
		Key:     key,
		Value:   value,
	}
}

func New(maxLevel int) *SkipList {
	skl := &SkipList{
		Level:  maxLevel,
		Length: 0,
		Head:   &Node{Forward: make([]*Node, maxLevel)},
		Tail:   &Node{},
	}

	for i, _ := range skl.Head.Forward {
		skl.Head.Forward[i] = skl.Tail
	}

	return skl

}

func (skl *SkipList) Search(key Comparable) (interface{}, bool) {
	searchNode := skl.Head
	for level := skl.Level - 1; level >= 0; level-- {
		for {
			tempNode := searchNode.Forward[level]
			if tempNode == skl.Tail {
				break
			} else if tempNode.Key.CompareTo(key) == 1 {
				break
			} else if tempNode.Key.CompareTo(key) == 0 {
				return tempNode.Value, true
			} else {
				searchNode = tempNode
			}
		}

	}
	return nil, false
}

func (skl *SkipList) randomLevel() int {
	level := 1
	for rand.Float32() < PROBABILITY && level < skl.Level {
		level++
	}

	return level
}

func (skl *SkipList) Insert(key Comparable, value interface{}) {
	level := skl.randomLevel()
	targetNodes := make([]*Node, level)
	fmt.Println(key, "level", level)
	//搜索出插入节点
	currentNode := skl.Head
	for l := level - 1; l >= 0; l-- {
		for {
			nextNode := currentNode.Forward[l]
			if nextNode == skl.Tail {
				//空
				targetNodes[l] = currentNode
				break
			} else if nextNode.Key.CompareTo(key) == 1 {
				//第一个大于的节点
				targetNodes[l] = currentNode
				break

			} else if nextNode.Key.CompareTo(key) == 0 {
				nextNode.Value = value
				return
			} else {
				currentNode = nextNode
			}
		}
	}

	//插入
	newnode := newNode(key, value, level)
	for i := 0; i < len(targetNodes); i++ {
		//在节点后加入
		newnode.Forward[i] = targetNodes[i].Forward[i]
		targetNodes[i].Forward[i] = newnode
	}

	skl.Length++

}

func (skl *SkipList) Delete(key Comparable) {
	var affectedNodes []*Node
	var targetNode *Node
	//搜索出插入节点
	currentNode := skl.Head
	for l := skl.Level - 1; l >= 0; l-- {
		for {
			nextNode := currentNode.Forward[l]
			if nextNode == skl.Tail {
				//空
				break
			} else if nextNode.Key.CompareTo(key) == 1 {
				//第一个大于的节点
				break

			} else if nextNode.Key.CompareTo(key) == 0 {
				//前一个节点
				affectedNodes = append(affectedNodes, currentNode)
				targetNode = nextNode
				break
			} else {
				currentNode = nextNode
			}
		}
	}
	if targetNode == nil {
		return
	}
	//删除
	lens := len(affectedNodes)
	for i := 0; i < lens; i++ {
		level := lens - i - 1
		affectedNodes[i].Forward[level] = targetNode.Forward[level]
	}
}

//PrintStructure 打印可视化结构
func (skl *SkipList) PrintStructure() {

	for l := skl.Level - 1; l >= 0; l-- {
		fmt.Printf("L=%d| ", l+1)
		currentNode := skl.Head.Forward[l]
		for {
			if l+1 > len(currentNode.Forward) {
				break
			}
			fmt.Printf("%d ", currentNode.Key)
			nextNode := currentNode.Forward[l]
			if nextNode == skl.Tail {

				fmt.Printf("tail")
				break
			} else {
				currentNode = nextNode
			}
		}

		fmt.Print("\n")
	}
}
