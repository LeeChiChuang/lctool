package question

import "fmt"

var LinkPrev = "https://leetcode-cn.com/problems/"

type QuestionResponse struct {
	Data Data `json:"data"`
}

type TopicTags struct {
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	TranslatedName string `json:"translatedName"`
	Typename       string `json:"__typename"`
}

type CodeSnippets struct {
	Lang     string `json:"lang"`
	LangSlug string `json:"langSlug"`
	Code     string `json:"code"`
	Typename string `json:"__typename"`
}

type Question struct {
	QuestionID            string         `json:"questionId"`
	QuestionFrontendID    string         `json:"questionFrontendId"`
	Title                 string         `json:"title"`
	TitleSlug             string         `json:"titleSlug"`
	TranslatedTitle       string         `json:"translatedTitle"`
	TranslatedContent     string         `json:"translatedContent"`
	IsPaidOnly            bool           `json:"isPaidOnly"`
	Difficulty            string         `json:"difficulty"`
	Contributors          []interface{}  `json:"contributors"`
	LangToValidPlayground string         `json:"langToValidPlayground"`
	TopicTags             []TopicTags    `json:"topicTags"`
	CodeSnippets          []CodeSnippets `json:"codeSnippets"`
	SampleTestCase        string         `json:"sampleTestCase"`
}

type Data struct {
	Question Question `json:"question"`
}

func NewQuestionResponse() *QuestionResponse {
	return &QuestionResponse{}
}

func (q *QuestionResponse) GetQuestion() string {
	return q.Data.Question.TranslatedContent
}

func (q *QuestionResponse) GetCode(lang string) string {
	for _, c := range q.Data.Question.CodeSnippets {
		if c.Lang == lang {
			return c.Code
		}
	}

	return ""
}

func (q *QuestionResponse) GetDifficulty() string {
	return q.Data.Question.Difficulty
}

func (q *QuestionResponse) GetTags() []string {
	tags := q.Data.Question.TopicTags
	res := make([]string, 0)
	for _, v := range tags {
		res = append(res, v.TranslatedName)
	}

	return res
}

func (q *QuestionResponse) GetMdName() string {
	return fmt.Sprintf("[%s *%s*](%s)", q.getQuestion().TranslatedTitle, q.getQuestion().TitleSlug, q.GetLink())
}

func (q *QuestionResponse) GetLink() string {
	return LinkPrev+q.getQuestion().TitleSlug
}

func (q *QuestionResponse) getQuestion() Question {
	return q.Data.Question
}