package collector

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
)

type WxedgeDashboard struct {
	Code int `json:"code"`
	Data struct {
		Resource    Status `json:"resource"`
		RunTasks    []Task `json:"run_tasks"`
		OptionTasks []struct {
			Name            string `json:"name"`
			ID              string `json:"id"`
			Disk            int    `json:"disk"`
			Mem             int    `json:"mem"`
			IsRecommend     int    `json:"is_recommend"`
			NeedRunningDays int    `json:"need_running_days"`
		} `json:"option_tasks"`
		Config struct {
			AddBeginHour int `json:"add_begin_hour"`
			AddEndHour   int `json:"add_end_hour"`
			DelBeginHour int `json:"del_begin_hour"`
			DelEndHour   int `json:"del_end_hour"`
		} `json:"config"`
	} `json:"data"`
}
type Status struct {
	CPUNum        int     `json:"cpu_num"`
	CPUUsage      float64 `json:"cpu_usage"`
	CPUUsageAlarm float64 `json:"cpu_usage_alarm"`
	TotalMem      int     `json:"total_mem"`
	UsedMem       int     `json:"used_mem"`
	MemAlarm      float64 `json:"mem_alarm"`
	DiskDev       string  `json:"disk_dev"`
	DiskUsed      float64 `json:"disk_used"`
	DiskTotal     int     `json:"disk_total"`
	FsType        string  `json:"fs_type"`
	FsEnable      int     `json:"fs_enable"`
	Ioutil        float64 `json:"ioutil"`
	IoutilAlarm   float64 `json:"ioutil_alarm"`
	RAwait        float64 `json:"r_await"`
	RAwaitAlarm   float64 `json:"r_await_alarm"`
	WAwait        float64 `json:"w_await"`
	WAwaitAlarm   float64 `json:"w_await_alarm"`
	Load5         float64 `json:"load5"`
	Load5Alarm    float64 `json:"load5_alarm"`
}

var WxedgeRepo WxedgeDashboard

func GetWxedgeStatus(host string) {
	client := http.Client{}
	res, err := client.Get(fmt.Sprintf("%s/docker/dashboard", host))
	if err != nil {
		logrus.Error("请求错误，原因:", err)
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err = json.Unmarshal(body, &WxedgeRepo); err != nil {
		logrus.WithError(err).Error("格式化失败，请检查获取内容: ", string(body))
	}
}
