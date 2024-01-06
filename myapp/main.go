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
type Block struct{
    block_id string
    plain_size uint64
    sha256 string
}

type DataSet struct{
    dataset_id string
    address string
    tag uint32
    blocks []Block   
}
type Context interface{
    Read_Block() []string
    Write_File()  //desc string,data []byte return ()
    Task_Info()uint32
    Datasets_By_Tag() []DataSet
}
// mock MyContext： 假设这里有一个map管理block_path
type MyContext struct{}

type Column struct{
    name string
    idx int32
}
type Arguments struct{
    tag uint32
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
func Read_Block(ctx *MyContext,dataset_id string, block_id string)(
    
    
    
    
    []byte,error){
    var block_data []byte
    var err error
    if (dataset_id == "00000" && block_id == "1"){
        block_1_1_path :="testdata/d1_block1.txt"
        fmt.Println(block_1_1_path)
        block_data ,err = ioutil.ReadFile(block_1_1_path) //demo of read_block 问题 我从read_block拿到的数据是什么？ []byte 这里是mock的read_block()方法 假设拿到的数据叫做cursor
        if err != nil {
            fmt.Println("Error reading file:", err)
            return nil,err
        }
        return block_data,nil
    }
    if (dataset_id == "00000" && block_id == "2"){
        block_1_2_path :="testdata/d1_block1_2.txt"
        fmt.Println(block_1_2_path)
        block_data ,err = ioutil.ReadFile(block_1_2_path) //demo of read_block 问题 我从read_block拿到的数据是什么？ []byte 这里是mock的read_block()方法 假设拿到的数据叫做cursor
        if err != nil {
            fmt.Println("Error reading file:", err)
            return nil,err
        }
        return block_data,nil
    }     //  block_2_1
    if (dataset_id == "11111" && block_id == "1"){
        block_2_1_path :="testdata/d2_block1.txt"
        fmt.Println(block_2_1_path)
        block_data ,err = ioutil.ReadFile(block_2_1_path) //demo of read_block 问题 我从read_block拿到的数据是什么？ []byte 这里是mock的read_block()方法 假设拿到的数据叫做cursor
        if err != nil {
            fmt.Println("Error reading file:", err)     
            return nil,err
        }
        return block_data,nil

    }
    
    return block_data,err
}

func Write_File(){

}

func Task_Info(ctx *MyContext) uint32{
    return 2
}

//一共有几个DataSet需要做

func DataSets_By_Tag(ctx * MyContext,tag uint32) ([]DataSet,error){
    // 把
    data := []DataSet{}

    if tag == 0 {
        dataset1 := new(DataSet)
        block_1_1 := Block{block_id:"1"}
        block_1_2 := Block{block_id:"2"} 

        dataset1.blocks = append(dataset1.blocks,block_1_1)
        dataset1.blocks = append(dataset1.blocks,block_1_2)
        dataset1.dataset_id="00000"
        // get dataset1： [block1,block2]blocks
        data = append(data,*dataset1)
    }
    if tag == 1{
        dataset2 := new(DataSet)
        block_2_1 := Block{block_id:"1"}
        // block_2_2 := Block{block_id:"2"}

        dataset2.blocks = append(dataset2.blocks,block_2_1)
        // dataset2.blocks = append(dataset2.blocks,block_2_2)
        dataset2.tag = 1
        dataset2.dataset_id = "11111"
        data = append(data,*dataset2)

        // get dataset2
    }
    return data,nil
}


func From_TEE_Config(ctx *MyContext)([]Arguments,error){//error
    col1 := Column{name: "name"}
    col2 := Column{name: "gender"}
// 
    args1 := new(Arguments)  
    args1.tag = 0
    args1.columns = append(args1.columns, col1)
    args1.columns = append(args1.columns, col2)

    args2 := new(Arguments)
    args2.tag = 1
    args2.columns = append(args2.columns, col1)
    args2.columns = append(args2.columns, col2)


    arguments := []Arguments{}
    arguments = append(arguments,*args1)
    arguments = append(arguments,*args2)



    return arguments,nil
}



// mock
func psi_validate_args(args []Arguments)(bool,error) { //Result...
 
    // how to add elements in Arguments slice?
    return true,nil
}
func psi_add_dataset(ctx *MyContext, args *Arguments, report *PSIReport) (*PSIReport,error) {//ctx *Context, args *Arguments, report *PSIReport
 
    dataset ,err := DataSets_By_Tag(ctx,args.tag) //[]DataSet
    if err != nil {
        fmt.Println("Error add_dataset:", err)
        return nil,err
    }

    // for _,a in range args{
    //     fmt.Println("")
    // }

    report,err =psi_read_data(ctx,&dataset[0],args.columns,report)
    if err != nil {
        fmt.Println("Error read data:", err)
        return nil,err
    }
    return report,nil
}
func psi_read_data(ctx *MyContext, dataset *DataSet, columns []Column ,report * PSIReport) (*PSIReport,error){
    
    singleDatasetReport := new(DataSetReport)
    singleDatasetReport.id = dataset.dataset_id //?
    singleDatasetReport.tag = dataset.tag
    for _,col := range columns{
        singleDatasetReport.columns = append(singleDatasetReport.columns,col.name)
    }

    // test() get the raw data block_id,DataSet_id
    // we assume the block_data comes from outside
    // block_path := "testdata/d1_block1.txt"
    //iter control 
    pos := 0 
    var header_idx []int 
    // blocks := []Block{}// 从外部拿到这个block数组 也是mock的
    //
    // var block Block //how to anounce a struct? 
    // block := Block{id: "some_id", path: block_path}

    
    blocks := dataset.blocks 
    // demo: add the block_path 
    for _,b := range blocks{
        // get the sequence data
        counts := make(map[string]int)
        // fmt.Println(dataset.dataset_id,b.block_id)
        block_data,err :=Read_Block(ctx,dataset.dataset_id,b.block_id)
        // block_data, err := ioutil.ReadFile(b.path) //demo of read_block 问题 我从read_block拿到的数据是什么？ []byte 这里是mock的read_block()方法 假设拿到的数据叫做cursor
        // block_data里面储存了所有的[]bytes数据，可能数据量很大，有1e8的数据，这时候我需要使用bufio.Scanner来扫描它
        // uint8类型的字节流吗？那我在这里应该使用什么函数来处理它？
        if err != nil {
            fmt.Println("Error reading file:", err)
            
            return nil,err
        }
        // fmt.Println(block_data)
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
                for _,col := range columns{
                    flag =false
                    for idx,Text := range header_Text{
                        if col.name == Text{
                            header_idx = append(header_idx,idx)
                            flag = true
                        }   
                    } //
                }// for loop: b
            }//scanner
            // check whether the row is properly founded or not
            if !flag {
                fmt.Println("check flag err",err)
                return nil,err
            }
        }// scan the header

        // then read the rest data
        // rows := bytes.Split(block_data, []byte("\n"))
        for scanner.Scan(){//Scanner_Iter
            line_tmp := scanner.Text()
            line := strings.Split(line_tmp,",")
            for i, value := range line {
                if value == "" { // 检查空字符串
                    line[i] = " " // 替换为单个空格
                }
            }
            var keyPart []string
            for _,i := range header_idx{//检查元素，并且读入这一行的数据
                // fmt.Println(line[i])
                // if i > len(line){
                //     singleDatasetReport.missing++
                //     continue 
                // }
                // 如果数据有缺失的话，missing++，跳出这个循环并且读下一行了
                if line[i] == " "{
                    singleDatasetReport.missing++
                    continue
                }
                keyPart = append(keyPart,line[i])       
            }
            key := strings.Join(keyPart,"")
            pos++
            if counts[key] == 0{//key 查询一下 有就不加 set map扩容
                report.reportMap[key]++
                counts[key]++
            }         

        }



        // fmt.Println("Temp map:", counts)
    }//block loop
    singleDatasetReport.rows = (uint32)(pos)
    report.dataset = append(report.dataset,*singleDatasetReport)
    return report,nil
}




func psi_summary(ctx *MyContext, report *PSIReport) (*PSIReport,error){
    // mock一下taskInfo 哦哦
    thres := Task_Info(ctx) -1
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
func execute(ctx *MyContext)(string,error){
    var err error
    report := new(PSIReport)
    report.reportMap = make(map[string]uint32) // 初始化map
// mock tee_config的参数
// 
    arguments,err :=From_TEE_Config(ctx) //step 1:get the []arguments slice
    if err != nil{
        fmt.Println("FROM TEE_CONFIG ERR",err)
        return "",err
    }
    
    // validate_args() -> get a true
// then ...
    // 省去一步 取dataset，直接readdata了
    for _,args := range arguments{
        report,err = psi_add_dataset(ctx,&args,report)
        if err != nil {
            fmt.Println("psi_add_dataset", err)
            return "",err
        }
    }
    // psi_read_data(&args,report)
    psi_summary(ctx,report)

    fmt.Println(*report)

    // serialized := 
    Write_File()
    return "",nil
}

func main()  {
    // block1_1_path := "testdata/d1_block1.txt"
    // block2_1_path := "testdata/d2_block1.txt"
    // myMap := make(map[string]uint8)
    // fmt.Println(block1_1_path,block2_1_path,myMap)
    // test(myMap, block1_1_path)
    // fmt.Println("After processing 1:", myMap)

    // test(myMap, block2_1_path)
    // fmt.Println("After processing 2:", myMap)
    // 前面省去validate操作，从外部调用tee_config拿arguments
    ctx := new(MyContext)
    execute(ctx)
}
