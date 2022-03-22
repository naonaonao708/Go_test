package main

import (
	"fmt"
	"math"
)

// 定数の定義
const weight = 60
const height = 165

// var 変数名　変数型　＝　値
func main() {
	var hm = height / 100.0
	var bmi = weight / math.Pow(hm, 2)
	var bestW = math.Pow(hm, 2) * 22.0
	var per = weight / bestW * 100

	fmt.Printf("BMI=%f, 肥満度=%.0f\n", bmi, per)
}

// 型を明示的に指定する場合　Goは静的言語のため本来なら型宣言が必要だが
//var hm float64 = 165 / 100.0

// 型推論という機能によって型宣言を記述する必要がない
//var height = 165 / 100.0

// また変数宣言と代入を一度に記述することもできる
//hm := 165 / 100.0

//a := 10 / 3    // aは、int型 3
//b := 10 / 3.0    // bは、float64型 3.333333
//var c float64 = 10 / 3    // cは、float64型 3.0
