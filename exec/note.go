package exec

import (
	"fmt"
	"github.com/leechichuang/lctool/question"
	"github.com/leechichuang/lctool/utils"
	"os"
	"text/template"
	"time"
)

var (
	noteTpl = `# {{.questionName}}
## 解题思路:



*{{.addDate}}*
`
	repeatTpl = `{{if .IsNew}}# Todo List
{{- else}}
- [ ] {{.Date}} | **{{.Name}}**
{{- end}}
`

	readmeTpl = `{{if .IsNew}}# leetcode题目分类
|题目|难度|标签|次数|
|--|--|--|--|
{{- else}}
|{{.Name}}|{{.Difficulty}}|{{range $index, $element := .Tag}} {{$element}} {{end}}|{{.Times}}|
{{- end}}`
	repeatInterval = [5]int{0, 1, 4, 7, 30}

	todoFile   = "../todo.md"
	readmeFile = "README.md"
)

const (
	Easy   = ":smile:"
	Medium = ":smile::smile:"
	Hard   = ":smile::smile::smile:"
)

type RepeatStruct struct {
	IsNew bool
	Date  string
	Name  string
}

type Readme struct {
	IsNew      bool
	Name       string
	Link       string
	Difficulty string
	Tag        []string
	Times      string
}

func GenerateNote(q question.QGenerater) error {
	if utils.FileExists(getNoteName(q.GetName())) {
		return nil
	}

	err := utils.MkdirIfNotExist(q.GetName())
	if err != nil {
		return err
	}

	fp, err := utils.CreateIfNotExist(getNoteName(q.GetName()))
	if err != nil {
		return err
	}

	var t = template.Must(template.New("questionNote").Parse(noteTpl))

	err = t.Execute(fp, map[string]string{
		"questionName": q.GetName(),
		"addDate":      time.Now().Format("2006-01-02 15:04:05"),
	})

	if err != nil {
		return err
	}
	return nil
}

func GenerateRepeat(q question.QGenerater) error {
	var tpl = template.Must(template.New("questionTodo").Parse(repeatTpl))
	var fp *os.File
	var err error
	if !utils.FileExists(todoFile) {
		fp, err = os.Create(todoFile)
		if err != nil {
			return err
		}
		isNew := RepeatStruct{IsNew: true}
		err = tpl.Execute(fp, isNew)
		if err != nil {
			return err
		}
	} else {
		fp, err = os.OpenFile(todoFile, os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
	}

	reps := [question.RepeatTimes]RepeatStruct{}
	t := time.Now()
	for i := 0; i < question.RepeatTimes; i++ {
		reps[i] = RepeatStruct{
			IsNew: false,
			Date:  t.Add(time.Hour * 24 * time.Duration(repeatInterval[i])).Format("2006-01-02"),
			Name:  fmt.Sprintf("%s_%d", q.GetName(), i),
		}

	}

	for _, v := range reps {
		err = tpl.Execute(fp, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func GenerateReadme(q *question.LeetCodeDesc) error {
	var tpl = template.Must(template.New("questionReadme").Parse(readmeTpl))
	if !utils.FileExists(readmeFile) {
		fd, err := os.Create(readmeFile)
		if err != nil {
			return err
		}
		r := new(Readme)
		r.IsNew = true
		err = tpl.Execute(fd, r)
		if err != nil {
			return err
		}
	}
	fd, err := os.OpenFile(readmeFile, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	row := Readme{
		IsNew:      false,
		Name:       q.GetMdName(),
		Difficulty: q.GetDifficulty(),
		Tag:        q.GetTags(),
		Times:      ":+1:​",
	}
	err = tpl.Execute(fd, row)
	if err != nil {
		return err
	}

	return nil
}

func getNoteName(name string) string {
	return fmt.Sprintf("%s/%s_note.md", name, name)
}
