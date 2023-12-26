package main

import (
    "bufio"
    "bytes"
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
type DataSet struct{
    
}
type Context interface{
    Read_Block() []string
    Write_Block()[]string
    Task_Info()uint32
    Datasets_By_Tag() []DataSet
}
// mock MyContext： 假设这里有一个map管理block_path
type MyContext struct{

    // myMap map[string]string
}


type Block struct{
    id string
    path string
}
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
    rows uint32
    missing uint32
}

type PSIReport struct {
    reportMap map[string]uint32
    size uint32
    dataset []DataSetReport
}


func (r PSIReport) String() string {
    var builder strings.Builder
    builder.WriteString(fmt.Sprintf("Dataset intersection size: %d\n", r.size))
    for _, dsr := range r.dataset {
        builder.WriteString(fmt.Sprintf("dataset_id: %s tag: %d rows: %d columns: %s missing: %d\n",
            dsr.id, dsr.tag, dsr.rows, strings.Join(dsr.columns, ","), dsr.missing))
    }
    return builder.String()
}
// mock
// func Get_BlockPath(dataset_id string, block_id string) ([]byte,error){
    
//     path :="t.txt"
//     block_data ,err := ioutil.ReadFile(path) //demo of read_block 问题 我从read_block拿到的数据是什么？ []byte 这里是mock的read_block()方法 假设拿到的数据叫做cursor
//     if err != nil {
//         fmt.Println("Error reading file:", err)
        
//         return nil,err
//     }

//     return block_data,nil
// }
func Read_Block(ctx *MyContext,dataset_id string, block_id string)([]byte,error){

    if (dataset_id == "0" && block_id == "1"){
        block_1_1_path :="testdata/d1_block1.txt"
        block_data1_1 ,err := ioutil.ReadFile(block_1_1_path) //demo of read_block 问题 我从read_block拿到的数据是什么？ []byte 这里是mock的read_block()方法 假设拿到的数据叫做cursor
        if err != nil {
            fmt.Println("Error reading file:", err)
            return nil,err
        }
        return block_data1_1,nil
    }
//  block_2_1
    if (dataset_id == "1" && block_id == "1"){
        block_2_1_path :="testdata/d2_block1.txt"
        block_data2_1 ,err := ioutil.ReadFile(block_2_1_path) //demo of read_block 问题 我从read_block拿到的数据是什么？ []byte 这里是mock的read_block()方法 假设拿到的数据叫做cursor
        if err != nil {
            fmt.Println("Error reading file:", err)     
            return nil,err
        }
        return block_data2_1,nil

    }    
    

}

func Task_Info(ctx *MyContext) uint32{
    return 2
}

func From_TEE_Config(ctx *MyContext)([]Arguments){//error
    args1 := new(Arguments)  
    col1 := Column{name: "name"}
    col2 := Column{name: "gender"}
    args1.columns = append(args1.columns, col1)
    args1.columns = append(args1.columns, col2)

    args2 := new(Arguments)  
    args2.columns = append(args2.columns, col1)
    args2.columns = append(args2.columns, col2)


    arguments := []Arguments{}
    arguments = append(arguments,*args1)
    arguments = append(arguments,*args2)



    return arguments
}



// mock
func psi_validate_args(args []Arguments)(bool,error) { //Result...
 
    // how to add elements in Arguments slice?
    return true,nil
}
func psi_add_dataset(ctx *MyContext, args *Arguments, report *PSIReport) (*PSIReport,error) {//ctx *Context, args *Arguments, report *PSIReport
    // get the dataset from outside: using the arguments: id, datasetid
    // psi
    // 
    // psi_read_data()
    data_label := Datasets_By_Tag()
    return report,nil
}
func psi_read_data(ctx *MyContext, args * Arguments, report * PSIReport) (*PSIReport,error){
    
    singleDatasetReport := new(DataSetReport)
    singleDatasetReport.id = "test"
    singleDatasetReport.tag = 0

    // test() get the raw data block_id,DataSet_id
    // we assume the block_data comes from outside
    block_path := "testdata/d1_block1.txt"
    //iter control 
    pos := 0 
    var header_idx []int 
    // blocks := []Block{}// 从外部拿到这个block数组 也是mock的
    //有很多个block
    // var block Block //how to anounce a struct? 
    // block := Block{id: "some_id", path: block_path}

    
    block := Block{id: "some_id", path: block_path}
    blocks = append(blocks,block)//mock
    // demo: add the block_path 
    for _,b := range blocks{
        // read_block()// get the sequence data
        counts := make(map[string]int)
        block_data,err :=Read_Block(1,1)
        // block_data, err := ioutil.ReadFile(b.path) //demo of read_block 问题 我从read_block拿到的数据是什么？ []byte 这里是mock的read_block()方法 假设拿到的数据叫做cursor
        // block_data里面储存了所有的[]bytes数据，可能数据量很大，有1e8的数据，这时候我需要使用bufio.Scanner来扫描它
        // uint8类型的字节流吗？那我在这里应该使用什么函数来处理它？
        if err != nil {
            fmt.Println("Error reading file:", err)
            
            return nil,err
        }

        // dataStr := string(block_data) //之前的这里是处理的dataStr 但是我这里应该要直接处理byte数据吧？
        // fmt.Println(dataStr) // 将字节切片转换为字符串并打印
        // fmt.Println("\n")    // 打印一个换行符
        // for the blocks loop
        // using ctx.read_block() -> get the data from the block

 
        var flag  bool = false
        // read first row in the block_1
        // assume there's an iterator..
        scanner := bufio.NewScanner(bytes.NewReader(block_data))
            // 只读取第一行
        if pos == 0 {
            if scanner.Scan(){
                headerLine := scanner.Text() //这个得到的是什么元数据？
                header_Text := strings.Split(headerLine,",") //split the data by "," []strings? 
                for _,col := range args.columns{
                    flag =false
                    for idx,Text := range header_Text{
                        singleDatasetReport.columns = append(singleDatasetReport.columns,Text)
                        if col.name == Text{
                            header_idx = append(header_idx,idx)
                            flag = true
                        }   
                    } //
                }// for loop: b
            }//scanner
        }// scan the header
        // check whether the row is properly founded or not
        if !flag {
            fmt.Println("check flag err",err)
            return nil,err
        }
        // then read the rest data
        // rows := bytes.Split(block_data, []byte("\n"))

        for scanner.Scan(){//Scanner_Iter
            line_tmp := scanner.Text()
            line := strings.Split(line_tmp,",")
            var keyPart []string
            for _,i := range header_idx{//检查元素，并且读入这一行的数据
                keyPart = append(keyPart,line[i])
                // 如果数据有缺失的话，missing++，跳出这个循环并且读下一行了
                if line[i] == "" {
                    singleDatasetReport.missing++
                    continue
                }
                key := strings.Join(keyPart,"")
                singleDatasetReport.rows++
                if counts[key] == 0{//key 查询一下 有就不加 set map扩容
                    report.reportMap[key]++
                    counts[key]++
                }                
            }
        }



        // fmt.Println("Temp map:", counts)
    }//block loop
    
    report.dataset = append(report.dataset,*singleDatasetReport)
    return report,nil
}




func psi_summary(ctx *MyContext, report *PSIReport) (*PSIReport,error){
    // mock一下taskInfo 哦哦
    thres := Task_Info(ctx)
    for _,v := range report.reportMap{
        if v > thres {
            report.size++
        }
    }
    return report,nil
}

func test_PSI(){
    // args := Arguments 

    // for i in range 0,3 { 
        // how to control this for loop?  
        // how to represent the length of args slice? args.len()? 
    //     // psi_add_dataset();
    // }   
}


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
    // fmt.Println("Temp map:", counts)

}


func main() {
    // block1_1_path := "testdata/d1_block1.txt"
    // block2_1_path := "testdata/d2_block1.txt"
    // myMap := make(map[string]uint8)
    // fmt.Println(block1_1_path,block2_1_path,myMap)
    // test(myMap, block1_1_path)
    // fmt.Println("After processing 1:", myMap)

    // test(myMap, block2_1_path)
    // fmt.Println("After processing 2:", myMap)
    // 前面省去validate操作，从外部调用tee_config拿arguments
    report := new(PSIReport)
    report.reportMap = make(map[string]uint32) // 初始化map
// mock tee_config的参数
// 
    myContext := new(MyContext)
    myContext.myMap
    arguments :=From_TEE_Config(myContext)

    
    // 省去一步 取dataset，直接readdata了
    for _,args := range arguments{
        psi_read_data(&args,report)
    }
    psi_summary(myContext,report)

    fmt.Println(*report)
}
