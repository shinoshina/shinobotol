## Hello Shinobu
自用bot到框架转型中

### 基于gocqserver的服务端框架（算是吧。。。

初始化bot，加载插件并启动 （当然是静态导入啦。。。
```go
func main(){
	a := sbot.NewBot()
	a.LoadPlugin(dplugin.Export())
	a.LoadPlugin(test.Export())
	a.LoadPlugin(shino.Export())
	a.LoadPlugin(leetcode.Export())
	a.LoadPlugin(bilibili.Export())
	a.Run()
}
```
插件定义并导出
```go
func Export() (p *route.Plugin) {
	p = route.NewPlugin("test","shut")
	return
}
```
支持多种处理
```go
p = route.NewPlugin("test","shut")
//严格匹配
p.OnMessage("all", "all", func(d route.DataMap) {
	request.SendMessage("parttest", d.GroupId())
})
//部分匹配
p.OnMessage("part", "part", func(d route.DataMap) {
	request.SendMessage("parttest", d.GroupId())
})
//支持分组的正则匹配
p.OnMessage(`^name:(?P<name>.*?) age:(?P<age>.*)`, "regex", func(d route.DataMap) {
	vmap := d["group_value"].(map[string](string))
	for k, v := range vmap {
		request.SendMessage("key: "+k+"  "+"value: "+"  "+v, d.GroupId())
	}
})
```
注册定时任务
```go
p.OnTick("定时任务",func() {
	timer := time.NewTimer(5*time.Second)
	<-timer.C
	request.SendMessage("定时任务",1012330112)
})
```
（伪）动态加载卸载插件
```go
p.OnTrigger("testload","testshut",func(d route.DataMap, pluginState string) {
	if pluginState == "loaded" {
		request.SendMessage("testonloaded",d.GroupID())
	}else if pluginState == "shut"{
		request.SendMessage("testonshut",d.GroupID())
	}
})
```
初启时/重启时钩子
```go
//插件初启/重启时启动定时任务，卸载时自动结束任务
p.OnBoot(func() {
	p.StartTask("定时任务")
})
```
### 半成品
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
### 内置功能 (都还没做完
* ##### 学人精加复读机 cosplay 帕拉斯
* ##### leetcode每日一题提醒
* ##### 不可以sese（其实可以
* ##### shinobu语音包！ （从副音轨里裁的
### TO NOT DO :sweet_smile:
* ##### 就不写
