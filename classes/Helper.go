package Classes

import(
	"regexp"
	// "fmt"
	"strings"
	"strconv"
	"sort"
)

type NgPosition struct{
	From float64
	To float64
	Kind string
}

type body_wrapper struct {
    Bodys []NgPosition
    by func(p,q *NgPosition) bool //内部Less()函数会用到
}
func (acw body_wrapper) Len() int  {
    return len(acw.Bodys)
}
func (acw body_wrapper) Swap(i,j int){
    acw.Bodys[i],acw.Bodys[j] = acw.Bodys[j],acw.Bodys[i]
}

//比较函数，使用外部传入的by比较函数
func (acw body_wrapper) Less(i,j int) bool {
    return acw.by(&acw.Bodys[i],&acw.Bodys[j])
}


//////////////////////////////////////////////////////////////////
type Helper struct{ 

}

func (this *Helper)justSort(bodys []NgPosition){
	sort.Sort(body_wrapper{bodys,func(p,q *NgPosition) bool{
		return p.From < q.From
	}})
}



func (this *Helper) Start(val string) []NgPosition  {
	var rets []NgPosition

	// 針點:350mm、800mm
	// 白點:275.0mm、刮傷:56.0mm、69.0mm、761.0mm

	reg := regexp.MustCompile("[\u4e00-\u9fa5]:?")
	val = reg.ReplaceAllString(val, "")
	val = strings.Replace(val, "mm", "", -1)

	// fmt.Println(val)
	
	arr := strings.Split(val, "、")

	for _, str := range arr{
		var model = NgPosition{}
		if(strings.Contains(str, "~")){
			t1 := strings.Split(str, "~")
			model.From, _ = strconv.ParseFloat(t1[0], 32) 
			model.To, _ = strconv.ParseFloat(t1[1], 32) 
			model.Kind = "duan"
		}else{
			model.Kind = "dian"
			model.From, _ = strconv.ParseFloat(str, 32) 
			model.To = model.From
		}

		rets = append(rets, model)
	}

	// fmt.Println("排序前")
	// fmt.Println(rets)
	// fmt.Println("排序后")
	this.justSort(rets)
	// fmt.Println(rets)

	return rets
}