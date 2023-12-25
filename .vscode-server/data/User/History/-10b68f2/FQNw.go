package main
import (
    "io/ioutil"
    "fmt")
func test(m map [string]u8, block1_1_path string) {
    // block1_1_path := "testdata/d1_block1.txt"
    // block2_1_path := "testdata/d2_block1.txt"

    block1_1_data, err := ioutil.ReadFile(block1_1_path)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return 
    }

    for

    // block2_1_data, err := ioutil.ReadFile(block2_1_path)
    // if err != nil {
    //     fmt.Println("Error reading file:", err)
    //     return 
    // }

    // fmt.Println(block1_1_data) // 打印字节切片
    // fmt.Println("\n") // 打印字节切片

    // fmt.Println(block2_1_data) // 打印字节切片



    fmt.Println(string(block1_1_data)) // 将字节切片转换为字符串并打印
    fmt.Println("\n")                  // 打印一个换行符
    fmt.Println(string(block2_1_data)) // 将字节切片转换为字符串并打印
    

}
func main(){
    block1_1_path := "testdata/d1_block1.txt"
    block2_1_path := "testdata/d2_block1.txt"
	mp  := make (map[string]u8)
    test(mp,block1_1_path)

    fmt.Println("After processing1:", myMap)

    test(mp,block2_1_path)

    fmt.Println("After processing2:", myMap)

}

