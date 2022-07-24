package leetcode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"shinobot/sbot/repo/datas"
	"shinobot/sbot/request"
	"shinobot/sbot/route"
	"strconv"
)

var (
	url    = "https://leetcode.cn/graphql/"
	method = "POST"
	purl   = "https://leetcode.cn/problems/"
)
var (
	subscribeList []float64
	db         *datas.Db
)

func init() {

	db = datas.CreateDb("assets/leetcode")
	subscribeList = make([]float64, 0)
	db.IterateAll(func(key, value string) {
		groupid, err := strconv.ParseFloat(key, 64)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("KEY", key, "VALUE", value)
			subscribeList = append(subscribeList, groupid)
		}
	})
	fmt.Println("submit list: ",subscribeList)
}
func getDailyName() (name string, id string) {
	queryMap := map[string]interface{}{
		"operationName": "questionOfToday",
		"variables":     "{}",
		"query":         "query questionOfToday { todayRecord {   question {     questionFrontendId     questionTitleSlug     __typename   }   lastSubmission {     id     __typename   }   date   userStatus   __typename }}",
	}
	queryJson, _ := json.Marshal(&queryMap)
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, bytes.NewReader(queryJson))

	req.Header.Add("authority", "leetcode.cn")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "zh-CN")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("origin", "https://leetcode.cn")
	req.Header.Add("referer", "https://leetcode.cn/problems/asteroid-collision/")

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

	var qn questionName
	json.Unmarshal(body, &qn)
	fmt.Println(qn)
	name = qn.Data.TodayRecord[0].Question.QuestionTitleSlug
	id = qn.Data.TodayRecord[0].Question.QuestionFrontendID
	return
}
func dailyQuestionInfo()string {

	name, _ := getDailyName()
	qu := "\"" + name + "\""
	queryMap := map[string]interface{}{
		"operationName": "questionData",
		"variables":     `{"titleSlug":` + qu + `}`,
		"query":         "query questionData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    categoryTitle\n    boundTopicId\n    title\n    titleSlug\n    content\n    translatedTitle\n    translatedContent\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    isLiked\n    similarQuestions\n    contributors {\n      username\n      profileUrl\n      avatarUrl\n      __typename\n    }\n    langToValidPlayground\n    topicTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n    companyTagStats\n    codeSnippets {\n      lang\n      langSlug\n      code\n      __typename\n    }\n    stats\n    hints\n    solution {\n      id\n      canSeeDetail\n      __typename\n    }\n    status\n    sampleTestCase\n    metaData\n    judgerAvailable\n    judgeType\n    mysqlSchemas\n    enableRunCode\n    envInfo\n    book {\n      id\n      bookName\n      pressName\n      source\n      shortDescription\n      fullDescription\n      bookImgUrl\n      pressImgUrl\n      productUrl\n      __typename\n    }\n    isSubscribed\n    isDailyQuestion\n    dailyRecordStatus\n    editorType\n    ugcQuestionId\n    style\n    exampleTestcases\n    jsonExampleTestcases\n    __typename\n  }\n}\n",
	}
	queryJson, _ := json.Marshal(&queryMap)
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, bytes.NewReader(queryJson))

	req.Header.Add("authority", "leetcode.cn")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "zh-CN")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("origin", "https://leetcode.cn")
	req.Header.Add("referer", "https://leetcode.cn/problems/asteroid-collision/")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var qinfo questionInfo
	json.Unmarshal(body, &qinfo)

	tags := ""
	for _, v := range qinfo.Data.Question.TopicTags {
		tags += v.TranslatedName
		tags += " "
	}
	message := "哦哈哟！" + "\n" + 
	    "今日题目: " + qinfo.Data.Question.TranslatedTitle + "\n" +
		"难度: " + qinfo.Data.Question.Difficulty + "\n" +
		"tags: " + tags + "\n" +
		"详情这里哦: " + purl + name + "/"
	return message
}
func subscribe(d route.DataMap) {
	idstr := Float64toString(d.GroupID())
	if db.Has(idstr) {
		request.SendMessage("不可以！！！！！", d.GroupID())
	} else {
		subscribeList = append(subscribeList, d.GroupID())
		db.Put(idstr, "useless")
		request.SendMessage("订阅成功", d.GroupID())
	}
}
func unsubscribe(d route.DataMap){
	idstr := Float64toString(d.GroupID())
	if db.Has(idstr){
		request.SendMessage("干嘛取消！！?",d.GroupID())
		db.Delete(idstr)
	}else{
		request.SendMessage("你订阅了吗[CQ:face,id=11][CQ:face,id=11]",d.GroupID())
	}

}

func SendLeetcodeInfo(){
	msg := dailyQuestionInfo()
	for _,v := range subscribeList{
		request.SendMessage(msg,v)
	}
}
func Float64toString(a float64) string {
	s := strconv.FormatFloat(a, 'f', -1, 64)
	fmt.Println(s)
	return s
}
func InttoString(a int) string {
	fmt.Println("int", a)
	return strconv.Itoa(a)
}
