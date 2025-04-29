package tools

import (
	"fmt"
	"strings"
	"time"
)

func ProcessBar() {
	totalTime := 30
	processBarWidth := 50

	for i := 1; i <= totalTime; i++ {
		time.Sleep(1 * time.Second)

		//计算进度条长度
		percent := float64(i) / float64(totalTime) * 100
		barLengeth := int(percent * float64(processBarWidth) / 100)
		bar := strings.Repeat("█", barLengeth) + strings.Repeat(" ", processBarWidth-barLengeth)

		// 打印进度条和倒计时
		fmt.Printf("\r[%s] %.0f%% (%d/%d 秒)", bar, percent, i, totalTime)
	}
	fmt.Printf("程序运行完成,花费时间: %ds", totalTime)
}
