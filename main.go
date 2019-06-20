package main

import (
	"zhouyuexcel/Classes"
	"strings"
	"fmt"
	// "time"
	"strconv"
	"bytes"
	"github.com/360EntSecGroup-Skylar/excelize"
	"bufio"
    "os"
)


func dealExecelData(){
	fmt.Println("正在读取目标excel文件")
	f, err := excelize.OpenFile("./整理品规.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    // 获取工作表中指定单元格的值
	// cell, err := f.GetCellValue("Sheet1", "B2")
	
	rows, err := f.GetRows("Sheet1")
	fmt.Println("文件一共有：", len(rows), "行记录")

	helper := Classes.Helper{}

	for index := 2; index <= len(rows); index++ {
		colid := fmt.Sprintf("AR%d",index) 
		instoreBatchId := fmt.Sprintf("H%d",index) 
		maxWidth := fmt.Sprintf("E%d",index) 

		beginWidth := fmt.Sprintf("N%d",index) 
		endWidth := fmt.Sprintf("O%d",index) 
		//////////////////
		updateCellId := fmt.Sprintf("F%d", index)
		// 运算结果
		var floatValue float64

		if cell, err := f.GetCellValue("Sheet1", colid); err == nil{
			batch, _ := f.GetCellValue("Sheet1", instoreBatchId)
			// 减去首尾数据
			beginwidth,_ := f.GetCellValue("Sheet1", beginWidth)
			endwidth,_ := f.GetCellValue("Sheet1", endWidth)
			maxwidth, _ := f.GetCellValue("Sheet1", maxWidth)
			tmax, _ := strconv.ParseFloat(maxwidth, 32)
			t3 , _ := strconv.ParseFloat(beginwidth, 32)
			t4 , _ := strconv.ParseFloat(endwidth, 32)

			if(len(cell) > 0){
				fmt.Println("正在解析第", index, "行")
				fmt.Println("入库批号 = ", batch, cell)

				result := helper.Start(cell)

				// 拼接字符串
				var buffer bytes.Buffer

				// var lastModel Classes.NgPosition
				for n, model := range(result){
					if(n == 0){
						buffer.WriteString( fmt.Sprintf("%v-", beginwidth) )
					}
					if(model.Kind == "duan"){
						if(n > 0){
							buffer.WriteString( fmt.Sprintf("-") )
						}
						buffer.WriteString( fmt.Sprintf("%v+%v", model.From, model.To) )
					}else{
						if(n > 0){
							buffer.WriteString( fmt.Sprintf("-") )
						}
						buffer.WriteString( fmt.Sprintf("%v+%v", model.From, model.To) )
					}

					// lastModel = model
				}
				
				buffer.WriteString( fmt.Sprintf("-%v",tmax-t4) )

				exp := buffer.String()

				arrayResult := strings.Split(exp, "+")
				expresssStr := ""

				for j, item := range(arrayResult){
					arr := strings.Split(item, "-") 
					t1 , _ := strconv.ParseFloat(arr[0], 32)
					t2 , _ := strconv.ParseFloat(arr[1], 32)
					floatValue += t2 - t1
					thiswidth := t2 - t1
					expresssStr += strconv.FormatFloat(thiswidth,'f',-1,64)
					if(j != len(arrayResult) - 1){
						expresssStr += "+"
					}
				}
			
				fmt.Println("正在更新单元格 " , updateCellId , " 值：", expresssStr ,"=",floatValue)
				f.SetCellValue("Sheet1", updateCellId, expresssStr)
			}else{
				// fmt.Println("跳过不需要处理的行：", index)
				floatValue = tmax - t3 - t4
				expresssStr := fmt.Sprintf("%s-%v-%v",maxwidth,t3, t4)
				fmt.Println("正在更新单元格 " , updateCellId , " 值：", expresssStr,"=", floatValue)
				f.SetCellValue("Sheet1", updateCellId, expresssStr)
			}
		}
	}

	if err := f.Save(); err != nil{
		fmt.Println("文件保存出错：", err)
	}
	
	fmt.Println("处理结束，请打开文件查看处理结果")
}

func main() {
    

	fmt.Println("欢迎使用本工具，请注意以下几点：") 
	fmt.Println("作者：陈日红(mail@chenrh.com)，版本 v1.2")
	fmt.Println("1.本工具仅支持 .xlsx 的文件")
	fmt.Println("2.请将待处理的文件命名为 整理品规.xlsx")
	fmt.Println("3.请把本工具拷贝至 整理品规.xlsx 目录")
	fmt.Println("4.数据必须在名为Sheet1的页签中")
	fmt.Println("")

	dealExecelData()

	var inputReader *bufio.Reader
	inputReader = bufio.NewReader(os.Stdin)
	for{
		inputReader.ReadString('\n')
	}
}

 