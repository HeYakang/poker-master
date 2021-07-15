package src

import (
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
	"time"
)

type Matches struct {
	Matches []*Match `json:"matches"`
}

type Match struct {
	PlayerA string `json:"alice"`
	PlayerB string `json:"bob"`
	Result int `json:"result"`
}



//获取牌组
func GetMatchesFromMatchFile(path string) *Matches {
	//beginTime := time.Now()
	var file []byte
	var err error

	if file,err = ioutil.ReadFile(path); err !=nil{
		panic("panic: "+ err.Error() )
	}

	matches := Matches{}

	//获取json数据绑定结构体
	if err := jsoniter.Unmarshal(file , &matches);err != nil {
		panic("panic: "+ err.Error() )
	}
	////记录结束时间
	//finishTime := time.Now()
	////打印时间
	//fmt.Printf("读取数据共耗时：%.2f 毫秒\n", finishTime.Sub(beginTime).Seconds()*1000)
	return &matches
}

//输出结果
// 打印牌组比较结果
func (matches *Matches) PrintCompareResult() {
	beginTime := time.Now()

	for _, v := range matches.Matches {
		res :=getWinner(v.PlayerA, v.PlayerB)

		if res != v.Result {
			fmt.Printf("%s, %s , %d, %d\n", v.PlayerA, v.PlayerB, res, v.Result)
		}
	}
	fmt.Printf("total：%d line\n", len(matches.Matches))
	//记录结束时间
	finishTime := time.Now()
	//打印时间
	fmt.Printf("比较共耗时：%.2f 毫秒\n", finishTime.Sub(beginTime).Seconds()*1000)
}








