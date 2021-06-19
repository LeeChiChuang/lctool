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
	repeatTpl = `{{if .IsNew}}
# Todo List
{{- else}}
- [ ] {{.Date}} | **{{.Name}}**
{{- end}}
`
	repeatInterval = [5]int{0, 1, 4, 7, 30}

	todoFile = "../todo.md"
)

type RepeatStruct struct {
	IsNew bool
	Date  string
	Name  string
}

func GenerateNote(q question.QGenerater) error {
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
		fp, err = os.OpenFile(todoFile, os.O_RDWR, 0666)
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

func getNoteName(name string) string {
	return fmt.Sprintf("%s/%s_note.md", name, name)
}
