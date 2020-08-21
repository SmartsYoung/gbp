# syncer 集群

### 带日志输出 sc 
 `docker run -p 30100:30100 servicecomb/service-center`
### 映射sc端口到30101
 `docker run -p 30101:30100 servicecomb/service-center`
### 节点1
`./syncer daemon --sc-addr http://192.168.140.180:30100 --bind-addr 192.168.140.180:30190 --rpc-addr 192.168.140.180:30191 --cluster-port 30192`
### 节点2
`./syncer daemon --sc-addr http://192.168.140.180:30101 --bind-addr 192.168.140.180:40190 --rpc-addr 192.168.140.180:40191 --cluster-port 40192 --join-addr 192.168.140.180:3019`

## 拼写错误
如 `val.Set("appId", svc.AppID)` 写成 `val.Set("appID", svc.AppID) `错误极难发现