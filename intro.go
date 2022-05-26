package main

import (
	"archive/zip"
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func sumThreeNums() {
	var num int
	fmt.Scan(&num)
	fmt.Println(num/100 + num%100/10 + num%10)
}

func reverseThreeNum() {
	var num, result int
	fmt.Scan(&num)
	result = num%10*100 + num%100/10*10 + num/100*1
	fmt.Println(result)

	// another version
	var a, b, c int
	fmt.Scanf("%1d%1d%1d", &a, &b, &c)
	fmt.Printf("%d%d%d", c, b, a)
}

func hoursMinutesGone() {
	var seconds, hours, minutes int
	fmt.Scan(&seconds)
	hours = seconds / 3600
	minutes = seconds % 3600 / 60
	fmt.Printf("It is %d hours %d minutes.\n", hours, minutes)

	// another version
	var k int
	_, _ = fmt.Scan(&k)
	d := time.Date(0, 0, 0, 0, 0, k, 0, time.UTC)

	fmt.Printf("It is %d hours %d minutes.\n", d.Hour(), d.Minute())

}

func rightTriangle() {
	var a, b, c int
	fmt.Scanf("%d%d%d", &a, &b, &c)
	if a*a+b*b == c*c {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}

	// another version
	val, a, b, c := map[bool]string{
		true:  "Прямоугольный",
		false: "Непрямоугольный",
	}, 0, 0, 0
	fmt.Scan(&a, &b, &c)
	fmt.Println(val[a*a+b*b == c*c])
}

func existTriangle() {
	exist_map, a, b, c := map[bool]string{
		true:  "существует",
		false: "не существует",
	}, 0, 0, 0
	fmt.Scan(&a, &b, &c)
	fmt.Println(exist_map[a+b > c && a+c > b && b+c > a])
}

func averageNum() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(((float32(a) + float32(b)) / 2))
}

func zeroOfAll() {
	var n, d, c int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&d)
		if d == 0 {
			c += 1
		}
	}
	fmt.Println(c)
}

func minCounter() {
	n, r, min, j := 0, 1, 0, 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&j)
		switch {
		case j == min:
			r++
		case j < min || min == 0:
			r, min = 1, j
		}
	}
	fmt.Println(r)
}

func digitRootOfNum() {
	var n int
	fmt.Scan(&n)
	for n >= 10 {
		n_ := 0
		for n > 0 {
			n_ += n % 10
			n /= 10
		}
		n = n_
	}
	fmt.Println(n)

	// another version
	var a int
	fmt.Scan(&a)
	fmt.Println((a-1)%9 + 1)
}

func multiSeven() {
	var a_, b_ int
	fmt.Scan(&a_, &b_)
	for i := b_; i != a_; i-- {
		if i%7 == 0 {
			fmt.Println(i)
			return
		}
	}
	fmt.Println("NO")

	// another version
	var a, b int
	fmt.Scan(&a, &b)
	b = b / 7 * 7
	if b >= a {
		fmt.Print(b)
	} else {
		fmt.Print("NO")
	}
}

func cows() {
	var n int
	fmt.Scan(&n)
	if n >= 11 && n <= 14 {
		fmt.Println(n, "korov")
	} else {
		temp := n % 10
		if temp == 0 || (temp >= 5 && temp <= 9) {
			fmt.Println(n, "korov")
		}
		if temp == 1 {
			fmt.Println(n, "korova")
		}
		if temp >= 2 && temp <= 4 {
			fmt.Println(n, "korovy")
		}
	}
}

func powMax() {
	var n int
	fmt.Scan(&n)
	for i := 0; math.Pow(float64(2), float64(i)) <= float64(n); i++ {
		fmt.Println(math.Pow(float64(2), float64(i)))
	}
}

func fibNum() {
	var n int
	fmt.Scan(&n)

	sliceFib := make([]int, n+5)
	sliceFib[0] = 0
	sliceFib[1] = 1
	for i := 2; i < n+5; i++ {
		sliceFib[i] = sliceFib[i-1] + sliceFib[i-2]
	}
	for i, v := range sliceFib {
		if v == n {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
}

func toBin() {
	var n int
	fmt.Scan(&n)
	fmt.Println(strconv.FormatInt(int64(n), 2))
}

func removeDigit() {
	var str_n string
	var n string
	fmt.Scan(&str_n)
	fmt.Scan(&n)

	fmt.Println(strings.ReplaceAll(str_n, n, ""))
}

func fn() {
	fn := func(i uint) (r uint) {
		s := []rune(strconv.Itoa(int(i)))
		var str_r string
		for _, v := range s {
			tmp, _ := strconv.Atoi(string(v))
			if tmp%2 == 0 && tmp != 0 {
				str_r += string(v)
			}
		}
		int_r, _ := strconv.Atoi(str_r)
		r = uint(int_r)
		if r == 100 {
			return 0
		}
		return r
	}

	fn(1234)
}

func interfaceType(i interface{}) bool {
	sign := func(s string) bool {
		if s == "+" || s == "*" || s == "-" || s == "/" {
			return true
		}
		return false
	}

	switch i.(type) {
	case float64:
		return true
	case string:
		if sign(i.(string)) {
			return true
		}
		return false
	default:
		return false
	}
}

func interfaceValues() {
	var (
		value1    interface{} = 42.2
		value2    interface{} = 21.21
		operation interface{} = "+"
	)
	v1, v2, op_ := interfaceType(value1), interfaceType(value2), interfaceType(operation)
	switch {
	case v1 == false:
		fmt.Printf("value=%v:%T", value1, value1)
		return
	case v2 == false:
		fmt.Printf("value=%v:%T", value2, value2)
		return
	case op_ == false:
		fmt.Printf("неизвестная операция")
		return
	case operation.(string) == "+":
		fmt.Printf("%.4f", value1.(float64)+value2.(float64))
		return
	case operation.(string) == "-":
		fmt.Printf("%.4f", value1.(float64)-value2.(float64))
		return
	case operation.(string) == "*":
		fmt.Printf("%.4f", value1.(float64)*value2.(float64))
		return
	case operation.(string) == "/":
		fmt.Printf("%.4f", value1.(float64)/value2.(float64))
		return
	default:
		return
	}
}

type Stringer interface {
	String() string
}

type Battery struct {
	volume []string
}

func (b Battery) String() string {
	return fmt.Sprintf("%v", b.volume)
}

func BatteryTask() {
	var str string
	fmt.Scan(&str)

	vol := []string{}
	for _, v := range str {
		if string(v) == "0" {
			vol = append(vol, "")
		}
	}
	for _, v := range str {
		if string(v) == "1" {
			vol = append(vol, "X")
		}
	}
	fmt.Println(vol)
}

func files() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		tmp, _ := strconv.Atoi(scanner.Text())
		sum += tmp
	}
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	f.Write([]byte(strconv.Itoa(sum)))

	// another version
	scanner_ := bufio.NewScanner(os.Stdin)
	amount := 0
	for scanner_.Scan() {
		a, _ := strconv.Atoi(scanner_.Text())
		amount += a
	}
	io.WriteString(os.Stdout, strconv.Itoa(amount))
}

func csvFile() {
	zip, _ := zip.OpenReader("task.zip")
	defer zip.Close()

	for _, file := range zip.File {
		if !file.FileInfo().IsDir() {
			txt, _ := file.Open()
			if data, _ := csv.NewReader(txt).ReadAll(); len(data) == 10 {
				fmt.Println(data[4][2])
			}
			txt.Close()
		}

	}
}

func iNum() {
	f, err := os.ReadFile("numbers")
	if err != nil {
		panic(err)
	}
	s := strings.Split(string(f), ";")
	for i, v := range s {
		if v == "0" {
			fmt.Println(i + 1)
		}
	}
}

func main() {
	// sumThreeNums()
	// reverseThreeNum()
	// hoursMinutesGone()
	// rightTriangle()
	// existTriangle()
	// averageNum()
	// zeroOfAll()
	// minCounter()
	// digitRootOfNum()
	// multiSeven()
	// cows()
	// powMax()
	// fibNum()
	// toBin()
	// removeDigit()
	// fn()
	// interfaceValues()
	// BatteryTask()
	// files()
	// csvFile()
	// iNum()
}
