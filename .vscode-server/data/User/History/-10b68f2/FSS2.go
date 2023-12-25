package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

func test(m map[string]uint8, block_path string) {
    block_data, err := ioutil.ReadFile(block_path)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
    dataStr := string(block_data)
    fmt.Println(dataStr) // 将字节切片转换为字符串并打印
    fmt.Println("\n")               // 打印一个换行符

    rows := strings.Split(dataStr, "\n")
    for _, rows := range rows{
        columns := strings.Split(row,",")
        if len(columns)>0 {
            firstColumnData := columns[0] // 第一列数据
			fmt.Println(firstColumnData)
        }
    }
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
