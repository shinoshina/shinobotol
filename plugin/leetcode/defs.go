package leetcode

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"shinobot/sbot/request"
	"shinobot/sbot/route"
)

var (
	url    = "https://leetcode.cn/graphql/"
	method = "POST"
)

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
func dailyQuestionInfo(d route.DataMap) {

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
		return
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var qinfo questionInfo
	json.Unmarshal(body, &qinfo)
	fmt.Println(qinfo.Data.Question.TranslatedTitle)
	fmt.Println(qinfo.Data.Question.TranslatedContent)

	request.SendMessage(qinfo.Data.Question.TranslatedTitle, d.GroupID())
	// request.SendMessage(qinfo.Data.Question.TranslatedContent,d.GroupID())
	filePath := "assets/leetcode/question.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("fail to open", err)
	}

	defer file.Close()
	write := bufio.NewWriter(file)

	write.WriteString(qinfo.Data.Question.TranslatedContent)

	write.Flush()

}
