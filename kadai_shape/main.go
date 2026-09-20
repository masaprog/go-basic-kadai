package main

import (
	"errors"
	"fmt"
	"math"
)

//エラー情報の構造体
type ValidationError struct{
	Message string
	Field string
	Value float64
}

//エラーメソッドの定義
func (v *ValidationError) Error() string{
	return fmt.Sprintf("[バリデーションエラー] %s: %s (Value: %.2f)", v.Message, v.Field, v.Value )
	}

//図形のインターフェース
type Shape interface{
	Area() float64
	Perimeter() float64
}

//長方形の構造体
type Rectangle struct{
	Width float64
	Height float64
}

//円の構造体
type Circle struct{
	Radius float64
}

//長方形の面積を求めるメソッド
func (r Rectangle) Area() float64{
	return r.Width * r.Height
}

//長方形の周長を求めるメソッド
func (r Rectangle) Perimeter() float64{
	return 2 * (r.Width + r.Height)
}

//円の面積を求めるメソッド
func (c Circle) Area() float64{
	return math.Pi * c.Radius * c.Radius
}

//円の周長を求めるメソッド
func (c Circle) Perimeter() float64{
	return 2 * c.Radius * math.Pi
}

//長方形、円それぞれでエラー処理を返す。問題なければnilを返す。
func validateShape(s Shape) error{
	switch val := s.(type){
	case Rectangle: 
		if val.Width <= 0 {
			return &ValidationError{
				Message: "幅は正の数を入力してください。",
				Field: "Rectangle.Width",
				Value: val.Width,
			}
		}
		if val.Height <= 0{
			return &ValidationError{
				Message: "高さは正の数を入力してください。",
				Field: "Rectangle.Height",
				Value: val.Height,
		}
	}
	case Circle:
		if val.Radius <= 0{
			return &ValidationError{
				Message: "半径は正の数を入力してください。",
				Field: "Circle .Radius",
				Value: val.Radius,
			}
		}
	}
	return nil
}

func main(){
		shapes := []Shape{
			Rectangle{Width: 10.0, Height: 5.0},
			Circle{Radius: 3.0},
			Rectangle{Width: -2.0, Height: 5.0},
			Circle{Radius: -1.0},
			Rectangle{Width: 1.0, Height: 0.0},
		}

		//バリデーションチェックがOKの図形のみを格納するスライスを宣言
		var calcShapes []Shape

		fmt.Println("--- 図形のバリデーション ---")
		for i, s := range shapes{
			err := validateShape(s)
			if err == nil{
				fmt.Printf("#%d: %T => バリデーションOK\n", i+1, s)
				calcShapes = append(calcShapes, s)
				continue
			}

			var  ve *ValidationError
			if errors.As(err, &ve){
				//エラーの詳細
				fmt.Printf("#%d: %T => NG: %s\n", i+1, s, err)
				fmt.Printf("[詳細] Field=%s, Value=%.2f\n", ve.Field, ve.Value)
			}
			}
			fmt.Println()

			//クリアした図形のみに対する計算
			fmt.Println("---検証済み図形の一括計算---")
			for _, s := range calcShapes{
				fmt.Printf("図形: %T | 面積: %.2f | 周長: %.2f\n", s, s.Area(), s.Perimeter())
			}
		}