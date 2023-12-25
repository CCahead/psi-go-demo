package main

import (
    // "bytes"
    "fmt"
    "io/ioutil"
    "strings"
    // "worker"
    // "psi/execute"
    // "github.com/rocketlaunchr/dataframe-go"
	// "github.com/rocketlaunchr/dataframe-go/imports"
)
// type Context interface{
    
// }
// how to implement Context interface?
type Column struct{
    name string
    idx int32
}
type Arguments struct{
    tag int32
    columns []Column
}
type Result struct{
    res bool

}
type DataSetReport struct {
    id string
    tag uint32
    columns []string
    row uint32
    missing uint32
}

type PSIReport struct {
    report_map map[string]uint32
    dataset []DataSetReport
}

func psi_validate_args(args *Arguments) bool { //Result...
    
    // how to add elements in Arguments slice?
    return true
}
func psi_add_dataset( args *Arguments, report *PSIReport) bool {//ctx *Context, args *Arguments, report *PSIReport
    // get the dataset from outside: using the arguments: id, datasetid
    psi
    return true
}
func psi_read_data(args * Arguments, report * PSIReport)

func psi_summary() bool{
    return true
}

func test_PSI(){
    // args := Arguments 

    for i in range 0,3 {
        // psi_add_dataset();
    }
}

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
    for _, row := range rows[1:] {
        row = strings.TrimSpace(row) // 去除行两端的空格
        columns := strings.Split(row, ",")
        for i, col := range columns[:] { //（这里做了一个简化操作，去掉了头）          
            
            columns[i] = strings.TrimSpace(col) // 去除列两端的空格
        }
        if len(columns) >= 2 { // 确保至少有两列
            firstColumnData := columns[0]  // 第一列数据
            secondColumnData := columns[1] // 第二列数据
            fmt.Println(firstColumnData, secondColumnData)
            if counts[firstColumnData] == 0{//key 查询一下 有就不加 set map扩容
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
    fmt.Println(block1_1_path,block2_1_path,myMap)
    // test(myMap, block1_1_path)
    // fmt.Println("After processing 1:", myMap)

    // test(myMap, block2_1_path)
    // fmt.Println("After processing 2:", myMap)

    // test_dataframe()
}
