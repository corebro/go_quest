package main

import (
	"fmt"
	"absurd-quest/internal/quest"
	"absurd-quest/internal/storage"
)

func main() {
	fmt.Println("=== Генератор абсурдных квестов ===")

	q := quest.NewGenerator()
	store := storage.NewJSONStorage("quests.json")

	quests := q.GenerateMultiple(5)

	for i, quest := range quests {
		fmt.Printf("%d) %s\n", i+1, quest)
	}

	store.Save(quests)
	fmt.Println("\nКвесты сохранены в quests.json")
}
