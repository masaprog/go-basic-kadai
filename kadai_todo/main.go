package main

import "fmt"

//ToDoリストを管理する構造体
type Todo struct{
	ID int
	Title string
	Completed bool
}

//ToDoを更新するメソッド
func (d *Todo) Complete() {
	fmt.Printf("ID：%dのToDoを完了に更新します\n", d.ID)
	d.Completed = true
}

//ToDoを全て表示する関数。ステータスがtrueなら完了、falseなら未完了を表示する。
func printTools(t []Todo) {
	for i :=range t {
	var status string 
	if t[i].Completed {
		status = "完了" 
	} else {
		status = "未完了"
	}
	fmt.Printf("[%s] (ID: %d) %s\n", status, t[i].ID, t[i].Title)
	}
}

func main() {　
	todos := []Todo{
		{1, "学習計画", false},
		{2, "環境構築", false},
		{3, "基礎文法", false},
	}
	fmt.Println("--- ToDoリスト（初期状態） ---")
	printTools(todos)
	fmt.Println() //空白の行を追加
	todos[0].Complete()
	todos[1].Complete()
	fmt.Println() //空白の行を追加
	fmt.Println("--- ToDoリスト（最終状態） ---")
	printTools(todos)
}

