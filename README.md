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
p.OnTick("hello",tick.Every(1*tick.Day).At("11:58"),func() {
	request.SendMessage("hello ha", groupid)
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
启动时钩子
```go
//启动时开启定时任务，卸载时自动结束任务
p.OnBoot(func() {
	p.StartTask("定时任务")
})
//ps : 初起时在bot收到第一条消息后每个plugin的onboot才会被调用
// 所以除了在init中初始化数据，在onboot中写也可以
// 因为 是伪卸载，所以plugin中定义的全局变量并不会被清除，再次加载插件仍可用，所以不用再ontrigger中注册同样的启动函数
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
	```go
	p.OnTick("leetcode", tick.Every(1*tick.Day).At("07:00"), SendLeetcodeInfo)
	p.OnBoot(func() {
		p.StartTask("leetcode")
	})
	p.OnMessage("订阅每日一题","all",subscribe)
	p.OnMessage("取消订阅力扣","all",unsubscribe)
	```
* ##### 不可以sese（其实可以
* ##### shinobu语音包！ （从副音轨里裁的
* ##### bilibili 直播推送
	```go
    关注永雏塔菲喵
	p.OnMessage(`^关注主播喵:(?P<mid>.*)`, "regex", subscribe)
	p.OnMessage(`^取关了喵:(?P<mid>.*)`, "regex", unsubscribe)
	```
### TO NOT DO  :heart_eyes:
* ##### 就不写
