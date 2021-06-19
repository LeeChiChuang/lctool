package question

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
