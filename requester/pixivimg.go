package requester

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func GetImg(){

	uri := "https://i.pximg.net/img-master/img/2022/04/24/06/10/00/97837406_p0_master1200.jpg"
	method := "GET"
	
	pu ,_ := url.Parse("http://172.28.0.1:4710")
	
  
	t := &http.Transport{
		MaxIdleConns:   10,
		MaxConnsPerHost: 10,
		IdleConnTimeout: time.Duration(100)*time.Second,
		Proxy: http.ProxyURL(pu),
	}
  
	client := &http.Client {
		Transport: t,
		Timeout: time.Duration(10)*time.Second,
	}
	req, err := http.NewRequest(method, uri, nil)
  
	if err != nil {
	  fmt.Println(err)
	  return
	}
	req.Header.Add("Referer", "https://www.pixiv.net/")
  
	res, err := client.Do(req)
	if err != nil {
	  fmt.Println(err)
	  return
	}
	defer res.Body.Close()
  
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
	  fmt.Println(err)
	  return
	}
  
	er := ioutil.WriteFile("/home/shinoshina/gocode/src/gocqserver/sese/bukeyisese!2.jpg",body,0666)
	if er != nil {
		fmt.Println(er)
	}
}