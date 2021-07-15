package main

import (
	"poker/define"
	"poker/src"
)

/*
	运行功能函数并记录运行时间
*/

func main() {
	////记录开始时间
	//beginTime := time.Now()

	//提取数据
	for _, path := range define.MatchSamplesPaths {
		src.GetMatchesFromMatchFile(path).PrintCompareResult()
	}

	////记录结束时间
	//finishTime := time.Now()
	////打印时间
	//fmt.Printf("共耗时：%.2f 毫秒\n", finishTime.Sub(beginTime).Seconds()*1000)
}