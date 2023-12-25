package main

import (
    "fmt"
    "io/ioutil"
)

func test(m map[string]uint8, block_path string) {
    block_data, err := ioutil.ReadFile(block_path)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
    fmt.Println(string(block_data)) // 将字节切片转换为字符串并打印
    fmt.Println("\n")               // 打印一个换行符
}

func main() {
    block1_1_path := "testdata/d1_block1.txt"
    block2_1_path := "testdata/d2_block1.txt"
    myMap := make(map[string]uint8)

    test(myMap, block1_1_path)
    fmt.Println("After processing 1:", myMap)

    test(myMap, block2_1_path)
    fmt.Println("After processing 2:", myMap)
}
