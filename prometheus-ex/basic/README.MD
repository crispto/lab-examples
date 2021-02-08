# 基本的 prometheus 操作

## dependences

### 运行 prometheus 程序
- [下载 prometheus 程序](https://prometheus.io/docs/introduction/first_steps/)
- 配置 prometheus 配置， cfg.yml 如下

```bash
global:
  scrape_interval: 15s  # 每次抓取的时间间隔
  evaluation_interval: 15s 

scrape_configs:
- job_name: my-count
  metrics_path: /metrics  # 服务metrics handler 的url 地址
  scheme: http
  static_configs:
  - targets: ['localhost:8083'] # 服务运行的地址
```
- 运行 ./prometheus --config.file=cfg.yml，程序启动，GUI 界面在9090 端口