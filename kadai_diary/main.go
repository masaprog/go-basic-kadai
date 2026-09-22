package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

//日記データ1件分を表す構造体
type Entry struct{
	Timestamp time.Time `json:"timestamp"`
	Title string `json:"title"`
	Content string `json:"content"`
}

//日記をファイルに書き込む関数
func saveDiary(e []Entry) error {
	j1, _ := json.Marshal(e)
	err := os.WriteFile("diary.json", j1, 0644)
	if err != nil {
		fmt.Println("ファイルの書き込みに失敗しました：", err)
		return err
	}
	return nil
}

//既にある日記を読み込む関数
func loadDiary() ([]Entry, error) {
	data, err := os.ReadFile("diary.json")
	if err != nil{
		if os.IsNotExist(err) {
			// 初回はファイルがないため、空のスライスを返す
			return []Entry{}, nil
		}
		fmt.Println("読み込みに失敗しました：", err)
		return nil, err
	}

	//JSONから変換して表示する用のスライスの用意
	var entries []Entry
	err2 := json.Unmarshal(data, &entries)
	if err2 != nil{
		fmt.Println("データの変換に失敗しました。処理を終了します。")
		return nil, err2
	}
	return entries, nil
}

func main() {
	//実行時の引数がプログラム名のみの場合は、使い方を表示して終了する
	if len(os.Args) < 2{
		fmt.Println("使い方：go run main.go [add|list]...")
		return
	}
	//add or listを変数に保存する
	command := os.Args[1]

	switch command{
	case "add":
		addTitle := os.Args[2]
		addContent := os.Args[3]
		s1, err := loadDiary()
		if err != nil {
			fmt.Println("読み込みに失敗しました：", err)
			return
		}
		newEntry := Entry{
			Timestamp: time.Now(),
			Title:     addTitle,
			Content:   addContent,
		}

		//読み込んだ日記スライスに新規で受け付けたデータを追加
		s1 = append(s1, newEntry)
		err2 := saveDiary(s1)
		if err2 != nil{
			fmt.Println("保存に失敗しました。：", err2)
			return
		}

	case "list":
		s1, err := loadDiary()
		if err != nil{
			fmt.Println("読み込みに失敗しました：", err)
			return
		}
		fmt.Println("--- 日記データ ---")
		for _, entry := range s1 {
			fmt.Println("日時：", entry.Timestamp.Format("2006-01-02 15:04:05"))
			fmt.Println("タイトル：", entry.Title)
			fmt.Println("本文：", entry.Content)
		}
	default:
		fmt.Println("不正なコマンドです：", command)
	}
}