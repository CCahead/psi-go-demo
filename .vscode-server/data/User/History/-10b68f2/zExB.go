package main

import (
    // "bytes"
    "fmt"
    "io/ioutil"
    "strings"
    // "github.com/rocketlaunchr/dataframe-go"
	// "github.com/rocketlaunchr/dataframe-go/imports"
)

// func test_dataframe(){
//     csvStr := `Country,Population,Year
//     Brazil,211000000,2019
//     India,1366000000,2019
//     USA,329500000,2019
//     Indonesia,270600000,2019
//     Pakistan,216600000,2019`
    
//         // 将CSV字符串转换为Reader
//         r := bytes.NewBufferString(csvStr)
    
//         // 读取CSV数据
//         df, err := imports.LoadFromCSV(r)
//         if err != nil {
//             panic(err)
//         }
    
//         // 打印数据帧
//         fmt.Println(df.Table())
// }

func test(m map[string]uint8, block_path string) {
    counts := make(map[string]int)
    block_data, err := ioutil.ReadFile(block_path)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
    dataStr := string(block_data)
    fmt.Println(dataStr) // 将字节切片转换为字符串并打印
    fmt.Println("\n")    // 打印一个换行符

    rows := strings.Split(dataStr, "\n")
    for _, row := range rows {
        row = strings.TrimSpace(row) // 去除行两端的空格
        columns := strings.Split(row, ",")
        for 1, col := range columns { //（这里做了一个简化操作，去掉了头）
            columns[i] = strings.TrimSpace(col) // 去除列两端的空格
        }
        if len(columns) >= 2 { // 确保至少有两列
            firstColumnData := columns[0]  // 第一列数据
            secondColumnData := columns[1] // 第二列数据
            fmt.Println(firstColumnData, secondColumnData)
            if counts[firstColumnData] == 0{
                m[firstColumnData]++
                counts[firstColumnData]++
            }
        }
    }
    fmt.Println("Temp map:", counts)



}

func main() {
    block1_1_path := "testdata/d1_block1.txt"
    block2_1_path := "testdata/d2_block1.txt"
    myMap := make(map[string]uint8)

    test(myMap, block1_1_path)
    fmt.Println("After processing 1:", myMap)

    test(myMap, block2_1_path)
    fmt.Println("After processing 2:", myMap)

    // test_dataframe()
}
