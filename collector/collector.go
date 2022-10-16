package collector

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"wxedge_exporter/config"
)

type Metrics struct {
	metrics map[string]*prometheus.Desc
	mutex   sync.Mutex
}

func newGlobalMetric(namespace string, metricName string, docString string, labels []string) *prometheus.Desc {
	return prometheus.NewDesc(namespace+"_"+metricName, docString, labels, nil)
}

func NewMetrics(namespace string) *Metrics {
	return &Metrics{
		metrics: map[string]*prometheus.Desc{
			"cpu_num":         newGlobalMetric(namespace, "cpu_num", "", []string{"host"}),
			"cpu_usage":       newGlobalMetric(namespace, "cpu_usage", "", []string{"host"}),
			"cpu_usage_alarm": newGlobalMetric(namespace, "cpu_usage_alarm", "", []string{"host"}),
			"disk_total":      newGlobalMetric(namespace, "disk_total", "", []string{"host"}),
			"disk_used":       newGlobalMetric(namespace, "disk_used", "", []string{"host"}),
			"ioutil":          newGlobalMetric(namespace, "ioutil", "", []string{"host"}),
			"ioutil_alarm":    newGlobalMetric(namespace, "ioutil_alarm", "", []string{"host"}),
			"load5":           newGlobalMetric(namespace, "load5", "", []string{"host"}),
			"load5_alarm":     newGlobalMetric(namespace, "load5_alarm", "", []string{"host"}),
			"mem_alarm":       newGlobalMetric(namespace, "mem_alarm", "", []string{"host"}),
			"r_await":         newGlobalMetric(namespace, "r_await", "", []string{"host"}),
			"r_await_alarm":   newGlobalMetric(namespace, "r_await_alarm", "", []string{"host"}),
			"total_mem":       newGlobalMetric(namespace, "total_mem", "", []string{"host"}),
			"used_mem":        newGlobalMetric(namespace, "used_mem", "", []string{"host"}),
			"w_await":         newGlobalMetric(namespace, "w_await", "", []string{"host"}),
			"w_await_alarm":   newGlobalMetric(namespace, "w_await_alarm", "", []string{"host"}),
			"total_tasks":     newGlobalMetric(namespace, "total_tasks", "", []string{"host"}),

			"task_cpu_usage":  newGlobalMetric(namespace, "task_cpu_usage", "", []string{"id", "mac", "name"}),
			"task_disk":       newGlobalMetric(namespace, "task_disk", "", []string{"id", "mac", "name"}),
			"task_mem":        newGlobalMetric(namespace, "task_mem", "", []string{"id", "mac", "name"}),
			"task_speed":      newGlobalMetric(namespace, "task_speed", "", []string{"id", "mac", "name"}),
			"task_state_code": newGlobalMetric(namespace, "task_state_code", "", []string{"id", "mac", "name"}),
		},
	}
}

func (c *Metrics) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range c.metrics {
		ch <- m
	}
}

func (c *Metrics) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	hosts := config.GetHost()
	for _, host := range hosts {
		GetWxedgeStatus(host)

		ch <- prometheus.MustNewConstMetric(c.metrics["cpu_num"], prometheus.GaugeValue, float64(WxedgeRepo.Data.Resource.CPUNum), host)
		ch <- prometheus.MustNewConstMetric(c.metrics["cpu_usage"], prometheus.GaugeValue, WxedgeRepo.Data.Resource.CPUUsage, host)
		ch <- prometheus.MustNewConstMetric(c.metrics["cpu_usage_alarm"], prometheus.GaugeValue, WxedgeRepo.Data.Resource.CPUUsageAlarm, host)
		ch <- prometheus.MustNewConstMetric(c.metrics["disk_total"], prometheus.GaugeValue, float64(WxedgeRepo.Data.Resource.DiskTotal), host)
		ch <- prometheus.MustNewConstMetric(c.metrics["disk_used"], prometheus.GaugeValue, WxedgeRepo.Data.Resource.DiskUsed, host)
		ch <- prometheus.MustNewConstMetric(c.metrics["ioutil"], prometheus.GaugeValue, WxedgeRepo.Data.Resource.Ioutil, host)
		ch <- prometheus.MustNewConstMetric(c.metrics["ioutil_alarm"], prometheus.GaugeValue, WxedgeRepo.Data.Resource.IoutilAlarm, host)
		ch <- prometheus.MustNewConstMetric(c.metrics["load5"], prometheus.GaugeValue, WxedgeRepo.Data.Resource.Load5, host)
		ch <- prometheus.MustNewConstMetric(c.metrics["load5_alarm"], prometheus.GaugeValue, WxedgeRepo.Data.Resource.Load5Alarm, host)
		ch <- prometheus.MustNewConstMetric(c.metrics["mem_alarm"], prometheus.GaugeValue, WxedgeRepo.Data.Resource.MemAlarm, host)
		ch <- prometheus.MustNewConstMetric(c.metrics["r_await"], prometheus.GaugeValue, WxedgeRepo.Data.Resource.RAwait, host)
		ch <- prometheus.MustNewConstMetric(c.metrics["r_await_alarm"], prometheus.GaugeValue, WxedgeRepo.Data.Resource.RAwaitAlarm, host)
		ch <- prometheus.MustNewConstMetric(c.metrics["total_mem"], prometheus.GaugeValue, float64(WxedgeRepo.Data.Resource.TotalMem), host)
		ch <- prometheus.MustNewConstMetric(c.metrics["used_mem"], prometheus.GaugeValue, float64(WxedgeRepo.Data.Resource.UsedMem), host)
		ch <- prometheus.MustNewConstMetric(c.metrics["w_await"], prometheus.GaugeValue, WxedgeRepo.Data.Resource.WAwait, host)
		ch <- prometheus.MustNewConstMetric(c.metrics["w_await_alarm"], prometheus.GaugeValue, WxedgeRepo.Data.Resource.WAwaitAlarm, host)
		ch <- prometheus.MustNewConstMetric(c.metrics["total_tasks"], prometheus.GaugeValue, float64(len(WxedgeRepo.Data.RunTasks)), host)

		for _, task := range WxedgeRepo.Data.RunTasks {
			ch <- prometheus.MustNewConstMetric(c.metrics["task_cpu_usage"], prometheus.GaugeValue, task.CPUUsage, task.ID, task.Mac, task.Name)
			ch <- prometheus.MustNewConstMetric(c.metrics["task_disk"], prometheus.GaugeValue, float64(task.Disk), task.ID, task.Mac, task.Name)
			ch <- prometheus.MustNewConstMetric(c.metrics["task_mem"], prometheus.GaugeValue, task.Mem, task.ID, task.Mac, task.Name)
			ch <- prometheus.MustNewConstMetric(c.metrics["task_speed"], prometheus.GaugeValue, task.Speed, task.ID, task.Mac, task.Name)
			ch <- prometheus.MustNewConstMetric(c.metrics["task_state_code"], prometheus.GaugeValue, float64(task.StateCode), task.ID, task.Mac, task.Name)
		}
	}
}
