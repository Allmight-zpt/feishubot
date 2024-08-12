package main

import (
	"fmt"

	"github.com/6tail/lunar-go/calendar"
	"github.com/CatchZeng/feishu/pkg/feishu"
)

// 定义阳历节日及其祝福
var solarHolidays = map[string]string{
	"1-1":   "元旦快乐！祝新的一年万事如意！",
	"2-14":  "情人节快乐！愿你与心爱的人永远幸福！",
	"3-8":   "妇女节快乐！愿所有女性都能享受幸福与尊重！",
	"5-1":   "劳动节快乐！祝您工作顺利，生活美满！",
	"6-1":   "儿童节快乐！愿每一个孩子都能健康快乐成长！",
	"10-1":  "国庆节快乐！祝愿祖国繁荣昌盛！",
	"12-25": "圣诞节快乐！愿您的生活充满爱与欢笑！",
}

// 定义农历节日及其祝福
var lunarHolidays = map[string]string{
	"1-1":  "春节",
	"1-15": "元宵节",
	"5-5":  "端午节",
	"7-7":  "七夕节",
	"8-15": "中秋节",
	"9-9":  "重阳节",
}

// 定义二十四节气及其祝福
var jieqi = map[string]string{}

//每日一句鸡汤

func sendHolidayMessage() {
	// solar := calendar.NewSolarFromDate(time.Now())
	solar := calendar.NewSolarFromYmd(2024, 8, 11)

	solar_month := solar.GetMonth()
	solar_day := solar.GetDay()
	solar_hour := solar.GetHour()
	solar_month_day_key := fmt.Sprintf("%d-%d", solar_month, solar_day)

	lunar := solar.GetLunar()
	lunar_month := lunar.GetMonth()
	lunar_day := lunar.GetDay()
	lunar_jieqi := lunar.GetJieQi()
	lunar_month_day_key := fmt.Sprintf("%d-%d", lunar_month, lunar_day)

	// 检查是否是晚上10点
	if solar_hour == 22 {
		fmt.Println("It’s 10 PM, here’s something special!")
	} else {
		// 检查是否是阳历节日
		if message, ok := solarHolidays[solar_month_day_key]; ok {
			fmt.Println(message)
		} else {
			// 检查是否是农历节日
			if message, ok := lunarHolidays[lunar_month_day_key]; ok {
				fmt.Println(message)
			} else if len(lunar_jieqi) > 0 {
				fmt.Println(lunar_jieqi)
			} else {
				fmt.Println("nothing")
			}
		}
	}
}
func SendMessage() string {
	token := "318e87b6-c6ee-4d66-8169-b60a881cac25"
	key := "wNeRboPX4HSHrrdG5l6oU"
	client := feishu.NewClient(token, key)
	msg := feishu.NewTextMessage()
	msg.Content.Text = "hello world"
	_, respone, err := client.Send(msg)
	if err != nil {
		panic(err)
	}
	return respone.Msg
}

func main() {
	SendMessage()
	/**
	1. 工作日
	2. 周末
	3. 节假日 & 纪念日
	4. 晚上10点提醒回家
	*/
}
