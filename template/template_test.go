package template

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"html/template"
	"testing"
)

func TestIfElseBlock(t *testing.T) {
	// 用一点小技巧来实现 for i 循环
	tpl := template.New("hello-world")
	tpl, err := tpl.Parse(`
{{- if and (gt .Age 0) (le .Age 6) }}
儿童 0<age<6
{{ else if and (gt .Age 6) (le .Age 18) }}
少年 6<age<=18
{{ else }}
成人 > 18
{{ end -}}
`)
	assert.Nil(t, err)
	bs := &bytes.Buffer{}
	err = tpl.Execute(bs, map[string]any{"Age": 5})
	assert.Nil(t, err)
	assert.Equal(t, `
儿童 0<age<6
`, bs.String())
}
