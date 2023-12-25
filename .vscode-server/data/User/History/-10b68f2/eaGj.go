package main
import (
    "io/ioutil"
    "fmt")
func test(){
    
}
func main(){
	block1_1_path := "testdata/d1_block1.txt"
    block2_1_path := "testdata/d2_block1.txt"

    block1_1_data, err := ioutil.ReadFile(block1_1_path)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return 
    }
    block2_1_data, err := ioutil.ReadFile(block2_1_path)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return 
    }
    // fmt.Println(block1_1_data) // 打印字节切片
    // fmt.Println("\n") // 打印字节切片

    // fmt.Println(block2_1_data) // 打印字节切片



    fmt.Println(string(block1_1_data)) // 将字节切片转换为字符串并打印
    fmt.Println("\n")                  // 打印一个换行符
    fmt.Println(string(block2_1_data)) // 将字节切片转换为字符串并打印
    

}

