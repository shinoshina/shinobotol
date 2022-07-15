## Hello Shinobu
自用bot到框架转型中

#### 半成品
```
docker pull shinoshina/shinobot
```
运行
```
docker run -idt shinoshina/shinobot /bin/bash /root/start.sh
```
查看容器
```
docker ps -a
```
停止运行并删除
```
docker stop containerid 
docker rm -v containerid
```
不给你玩