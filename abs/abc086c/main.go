package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	const max = 1024 * 1024
	var buf = make([]byte, max)
	sc.Buffer(buf, max)
}

// 偶数: (偶, 偶), (奇, 奇)
// 奇数: (偶, 奇), (奇, 偶)

func main() {
	var n, t1, t2 int
	var x1, y1, x2, y2 float64
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&t2, &x2, &y2)
		dt := t2 - t1
		dist := math.Abs(x2-x1) + math.Abs(y2-y1)
		if dt < int(dist) || int(dist)%2 != dt%2 {
			fmt.Println("No")
			return
		}
		x1, y1, t1 = x2, y2, t2
	}
	fmt.Println("Yes")
}

// Reverse 文字列を反転
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// NextLine buinfo.Scanのポインタを渡し、標準入力の次の行を読み込み
// ex. sc := buinfo.NewScanner(os.stdin)
//
//	GetNextLine(sc)
func NextLine(sc *bufio.Scanner) string {
	sc.Scan()
	s := sc.Text()
	return strings.TrimSpace(s)
}

// SplitStrList 文字列を空白区切りの文字列のリストに変換して返却
func SplitStrList(s string) []string {
	return strings.Split(s, " ")
}

// SplitIntlist 文字列を空白区切りの整数値に変換して返却
func SplitIntlist(s string) []int {
	strList := strings.Split(s, " ")
	return StrListToIntList(strList)
}

// StrListToIntList string型のスライスを渡してint型の配列に変換
func StrListToIntList(strList []string) (intList []int) {
	for _, str := range strList {
		str = strings.TrimRight(str, "\n")
		i := StrToInt(str)
		intList = append(intList, i)
	}
	return
}

// StrToInt string型をint型に変換
func StrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// Sort int型スライスの並び替え
func Sort(slice []int, order string) []int {
	sort.SliceStable(slice, func(i, j int) bool {
		if order == "desc" {
			return slice[i] > slice[j]
		} else {
			return slice[i] < slice[j]
		}
	})
	return slice
}

// FindMaxAndMin 最大値最小値を返す
func FindMaxAndMin(slice []int) (max, min int) {
	max = slice[0]
	min = slice[0]
	for _, elm := range slice {
		if elm > max {
			max = elm
		}
		if elm < min {
			min = elm
		}
	}
	return max, min
}

// Sum 合計値を返す
func Sum(slice []int) (sum int) {
	for _, i := range slice {
		sum += i
	}
	return sum
}

// Permutation Pの値を計算
func Permutation(n int, k int) *big.Int {
	v := big.NewInt(1)
	if 0 < k && k <= n {
		for i := 0; i < k; i++ {
			k := big.NewInt(int64(n - i))
			v = v.Mul(v, k)
		}
	} else if k > n {
		v = big.NewInt(0)
	}
	return v
}

// Factorial Fの値を計算
func Factorial(n int) *big.Int {
	return Permutation(n, n-1)
}

// Combination Cの計算
func Combination(n int, k int) *big.Int {
	child := Permutation(n, k)
	mother := Factorial(k)
	return child.Div(child, mother)
}

// Homogeneous Hの計算
func Homogeneous(n int, k int) *big.Int {
	return Combination(n+k-1, k)
}
