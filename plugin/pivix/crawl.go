package pv

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)


// shabi pixiv wocaonima
func GetImg() {

	uri := "https://i.pximg.net/img-master/img/2022/04/24/06/10/00/97837406_p0_master1200.jpg"
	method := "GET"

	//every time you start the wsl   ip will change  fuck!!!
	pu, _ := url.Parse("http://172.31.0.1:4710")

	t := &http.Transport{
		MaxIdleConns:    10,
		MaxConnsPerHost: 10,
		IdleConnTimeout: time.Duration(100) * time.Second,
		Proxy:           http.ProxyURL(pu),
	}

	client := &http.Client{
		Transport: t,
		Timeout:   time.Duration(10) * time.Second,
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

	er := ioutil.WriteFile("/home/shinoshina/gocode/src/gocqserver/sese/bukeyisese!2.jpg", body, 0666)
	if er != nil {
		fmt.Println(er)
	}
}

// dont just get by tag pages like  GET // uri := "https://www.pixiv.net/tags/%25E5%258E%259F%25E7%25A5%259E/artworks?s_mode=s_tag"
// or you get a dynamic pages source code different from you see on Chrome
// just GET by search throug ajax request
// https://www.pixiv.net/ajax/search/artworks/%E3%81%B5%E3%81%9F%E3%81%AA%E3%82%8A?
// word=%E3%81%B5%E3%81%9F%E3%81%AA%E3%82%8A&order=date_d&mode=r18&p=1&s_mode=s_tag_full&type=all&lang=zh
// type = illust
// mode = r18
// p means page
// word is tag
// pictures you get from this url are mostly ugly,unless you get a vip
func GetHTML() {
	uri := "https://www.pixiv.net/ajax/search/artworks/%E3%81%B5%E3%81%9F%E3%81%AA%E3%82%8A?word=%E3%81%B5%E3%81%9F%E3%81%AA%E3%82%8A&order=date_d&mode=r18&p=1&s_mode=s_tag_full&type=all&lang=zh"
	method := "GET"

	pu, _ := url.Parse("http://172.31.0.1:4710")

	t := &http.Transport{
		MaxIdleConns:    10,
		MaxConnsPerHost: 10,
		IdleConnTimeout: time.Duration(100) * time.Second,
		Proxy:           http.ProxyURL(pu),
	}

	client := &http.Client{
		Transport: t,
		Timeout:   time.Duration(10) * time.Second,
	}
	req, err := http.NewRequest(method, uri, nil)

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36")

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Referer", "https://www.pixiv.net")

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
	fmt.Println(string(body))

}

func getEverydayRecomPage(){

	t := &http.Transport{
		MaxIdleConns:    10,
		MaxConnsPerHost: 10,
		IdleConnTimeout: time.Duration(1000) * time.Second,
		Proxy:           http.ProxyURL(proxyUrl),
	}
	client := &http.Client{
		Transport: t,
		Timeout:   time.Duration(1000) * time.Second,
	}


	getRequest,err4Getreq := http.NewRequest("GET",everydayRecommendPage, nil)
    if err4Getreq!= nil{
		fmt.Println(err4Getreq.Error())
	}
	getRequest.Header.Add("User-Agent", user_agent)
	getRequest.Header.Add("Cookie", cookie)

	response,err4Res := client.Do(getRequest)
	if err4Res != nil{
		fmt.Println(err4Res.Error())
	}
	defer response.Body.Close()

	body,_ := ioutil.ReadAll(response.Body)

	imgRule := regexp.MustCompile(rawJpgRule)

	result := imgRule.FindAllStringSubmatch(string(body),-1)


	for i := range result{

		getImgRequest,err4ImgGet := http.NewRequest("GET",stringCut(result[i][1]),nil)
		if err4ImgGet != nil{
			fmt.Println(err4ImgGet.Error())
		}

		getImgRequest.Header.Add("User-Agent", user_agent)
		getImgRequest.Header.Add("Referer", referer)

		response4Img,err4Imgresp := client.Do(getImgRequest)
		if err4Imgresp != nil{
			fmt.Println(err4Imgresp.Error())
		}
		defer response4Img.Body.Close()

		imgBody,err4Img := ioutil.ReadAll(response4Img.Body)
		if err4Img != nil{
			fmt.Println(err4Img.Error())
		}

		num := strconv.Itoa(i)
		path := "/home/shinoshina/gocode/src/gocqserver/sese/bukeyisese!"+num+".jpg"
		fmt.Println(path)
		err := ioutil.WriteFile(path,imgBody,0666)
		if err != nil{
			fmt.Println(err)
		}
	}



}

func stringCut(url string) string{
	index := strings.Index(url,"/img-master")

	src1 := url[0:len("https://i.pximg.net")]
	src2 := url[index:]

	newUrl := src1+src2

	fmt.Println(newUrl)
	return newUrl
}


