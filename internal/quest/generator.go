package quest

import (
	"fmt"
	"math/rand"
	"time"
)

type Generator struct{}

func NewGenerator() *Generator {
	rand.Seed(time.Now().UnixNano())
	return &Generator{}
}

var actions = []string{"Укради", "Найди", "Спрячь", "Убеди", "Разбуди"}
var objects = []string{"гуся", "фонарь", "носок", "хомяка", "утку"}
var places = []string{"в подвале", "на крыше", "в автобусе", "в шкафу"}
var rewards = []string{"золото", "ничего", "славу", "огурцы"}

func (g *Generator) GenerateOne() string {
	return fmt.Sprintf(
		"%s %d %s %s. Награда: %s.",
		random(actions),
		rand.Intn(5)+1,
		random(objects),
		random(places),
		random(rewards),
	)
}

func (g *Generator) GenerateMultiple(n int) []string {
	var result []string
	for i := 0; i < n; i++ {
		result = append(result, g.GenerateOne())
	}
	return result
}

func random(arr []string) string {
	return arr[rand.Intn(len(arr))]
}
