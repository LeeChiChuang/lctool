package question

import (
	"encoding/json"
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/leechichuang/lctool/utils"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	leetcodePayload string = `{
    "operationName": "questionData",
    "variables": {
        "titleSlug": "%s"
    },
	"query": "query questionData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n       title\n    titleSlug\n    translatedTitle\n    translatedContent\n    isPaidOnly\n    difficulty\n   contributors {\n      username\n      profileUrl\n      avatarUrl\n      __typename\n    }\n    langToValidPlayground\n    topicTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n     codeSnippets {\n      lang\n      langSlug\n      code\n      __typename\n    }\n         sampleTestCase\n     }\n}\n"
	}`
	Api =  "https://leetcode-cn.com/graphql/"
	Language = "Go"
)

var leetcodeTemp string = `package %s
	
%s
`

const RepeatTimes  = 5

type QGenerater interface {
	WriteDesc(path string) error
	WriteCode(path string, repeat bool) error
	GetName() string
}

type LeetCodeDesc struct {
	*QuestionResponse
	name string
	desc string
	code string
}

func (l *LeetCodeDesc) WriteDesc(path string) error {
	path = fmt.Sprintf("%s/%s", path, l.name)

	return utils.WriteStringToFile(l.desc, path+fmt.Sprintf("/%s.md", l.name))
}

func (l *LeetCodeDesc) WriteCode(path string, repeat bool) error {
	path = fmt.Sprintf("%s/%s", path, l.name)
	if err := os.MkdirAll(path, 0766); err != nil {
		panic(err)
	}
	var n int
	if repeat {
		n = RepeatTimes
	} else {
		n = 1
	}

	for i:=0; i<n; i++ {
		err := utils.WriteStringToFile(l.code, path+fmt.Sprintf("/%s_%d.go", l.name, i))
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *LeetCodeDesc) GetName() string {
	return l.name
}

func NewLeetCode(name string) (*LeetCodeDesc, error) {
	res, err := requestLeetcode(name)
	if err != nil {
		return nil, err
	}

	questionInfo, err := res2QuestionInfo(res)
	if err != nil {
		return nil, err
	}

	markdown, code, err := formatResponse(questionInfo)
	if err != nil {
		return nil, err
	}

	code = fmt.Sprintf(leetcodeTemp, strings.ReplaceAll(name, "-", "_"), code)
	return &LeetCodeDesc{questionInfo, name, markdown, code }, nil
}

func requestLeetcode(q string) (*http.Response, error) {
	method := "POST"
	s := fmt.Sprintf(leetcodePayload, q)
	payload := strings.NewReader(s)

	client := &http.Client{}
	req, err := http.NewRequest(method, Api, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	return client.Do(req)
}

func res2QuestionInfo(res *http.Response) (*QuestionResponse, error) {
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	questionInfo := NewQuestionResponse()
	err = json.Unmarshal(body, questionInfo)
	if err != nil {
		return nil, err
	}

	return  questionInfo, nil
}

func formatResponse(q *QuestionResponse) (markdown string, code string, err error) {
	converter := md.NewConverter("", true, nil)
	markdown, err = converter.ConvertString(q.GetQuestion())
	if err != nil {
		log.Fatal(err)
		return markdown, code, err
	}
	code = q.GetCode(Language)

	return markdown, code, nil
}
