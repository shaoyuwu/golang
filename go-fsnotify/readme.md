# 说明
    该工具用于动态监控对应文件出现变化时，发送钉钉消息，产生告警信息发送。


### 编译
```
go build
[root@ go-fsnotify]# cat start.sh
#!/bin/bash
 
cd /data/go-fsnotify
nohup ./go-fsnotify canal-test1 canal-test2 canal-test3 canal-test4 canal-test5 canal-test6 &> ./fsnotify.log &
```