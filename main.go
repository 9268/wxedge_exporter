package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"strconv"
	"wxedge_exporter/collector"
	"wxedge_exporter/config"
)

func main() {
	logrus.Infof("欢迎使用wxedge容器监控prometheus客户端，项目名wxedge_exporter，作者：9268，欢迎提交issues、PullRequest")
	logrus.Infof("初始化程序")
	config.ReadConfig()

	logrus.Infof("初始化完成")

	logrus.Infof("初始化监控指标")
	metrics := collector.NewMetrics(viper.GetString("namespace"))
	registry := prometheus.NewRegistry()
	registry.MustRegister(metrics)
	logrus.Infof("监控指标初始化注册完成")

	logrus.Infof("启动服务器，监听端口为:" + strconv.Itoa(config.Configs.Port))
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`<html>
			<head><title>A Prometheus Exporter</title></head>
			<body>
			<h1>A Prometheus Exporter</h1>
			<p><a href='/metrics'>Metrics</a></p>
			</body>
			</html>`))
		if err != nil {
			logrus.WithError(err).Errorf("运行exporter错误")
			os.Exit(1)
		}
	})

	logrus.Infof("监控Metrics位置： http://localhost:%d%s", config.Configs.Port, "/metrics")
	logrus.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Configs.Port), nil))
}
