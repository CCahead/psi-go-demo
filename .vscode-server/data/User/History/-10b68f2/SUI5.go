package main
import (
    "io/ioutil"
    "fmt")

func main(){
	block1_1_path := "testdata/d1_block1.txt"
    block2_1_path := "testdata/d2_block1.txt"

    block1_1_data, err := ioutil.ReadFile(block1_1_path)
    block2_1_data, err := ioutil.ReadFile(block2_1_path)

    // fmt.Println("After:" ,jekyll,hyde,point)

}

