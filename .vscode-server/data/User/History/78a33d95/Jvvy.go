
import "fmt"

func swap(x, y, product *int) {
    if *x > *y {
        *x, *y = *y, *x
    }
    *product = *x * *y
}


func main(){
	var a int = 10
	var b int = 1
    pa, pb := &a, &b

    var res int
    swap(pa, pb, &res)
    fmt.Println(*pa, *pb, res)}
