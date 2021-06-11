package generater

import (
	"bufio"
	"encoding/json"
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
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

type QuestionGenerater interface {
	WriteDesc(path string) error
	WriteCode(path string) error
}

type LeetCodeDesc struct {
	name string
	desc string
	code string
}

func (l LeetCodeDesc) WriteDesc(path string) error {
	path = fmt.Sprintf("%s/%s", path, l.name)
	if err := os.MkdirAll(path, 0766); err != nil {
		panic(err)
	}
	return WriteStringToFile(l.desc, path+fmt.Sprintf("/%s.md", l.name))
}

func (l LeetCodeDesc) WriteCode(path string) error {
	path = fmt.Sprintf("%s/%s", path, l.name)
	if err := os.MkdirAll(path, 0766); err != nil {
		panic(err)
	}
	return WriteStringToFile(l.code, path+fmt.Sprintf("/%s.go", l.name))
}

func NewLeetCode(name string) (QuestionGenerater, error) {
	res, err := requestLeetcode(name)
	if err != nil {
		return nil, err
	}
	markdown, code, err := formatResponse(res)
	if err != nil {
		return nil, err
	}

	code = fmt.Sprintf(leetcodeTemp, strings.ReplaceAll(name, "-", "_"), code)
	return LeetCodeDesc{name, markdown, code}, nil
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

func formatResponse(res *http.Response) (markdown string, code string, err error) {
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return markdown, code, err
	}

	questionInfo := NewQuestionResponse()
	err = json.Unmarshal(body, questionInfo)
	if err != nil {
		return markdown, code, err
	}

	converter := md.NewConverter("", true, nil)
	markdown, err = converter.ConvertString(questionInfo.GetQuestion())
	if err != nil {
		log.Fatal(err)
		return markdown, code, err
	}

	code = questionInfo.GetCode(Language)

	return markdown, code, nil
}

func WriteStringToFile(content, path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	write.WriteString(content)
	write.Flush()

	return nil
}
