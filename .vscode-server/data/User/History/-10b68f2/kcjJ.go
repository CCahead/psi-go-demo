package main
import "fmt"

func swap(x, y, product *int) {
    if *x > *y {
        *x, *y = *y, *x
    }
    *product = *x * *y
}
type Exchanger interface{
    Exchange()
}
type StringPair struct{
    first,
    second string
}
func (pair *StringPair) Exchange(){
    pair.first,pair.second= pair.second,pair.first
}

type Point [2]int 
func (point *Point) Exchange(){
    point[0],point[1]=point[1],point[0]
}
func (pair StringPair) String() string{ // args, FunName, ReturnType
    return  fmt.Sprintf("%q,%q", pair.first, pair.second) //"%q" Sprintf
}

func main(){
	var a int = 10
	var b int = 1
    pa, pb := &a, &b

    var res int
    swap(pa, pb, &res)
    fmt.Println(*pa, *pb, res)

    jekyll :=StringPair{"Herry","Jekyll"} 
    hyde :=StringPair{"Edward", "Hyde"}

    point := Point{3,2}

    fmt.Println("Before:" ,jekyll,hyde,point)
}

