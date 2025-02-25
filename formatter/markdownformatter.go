package formatter

import (
	"bytes"
	"text/template"
)

type MarkdownFormatter struct{}

func (f *MarkdownFormatter) Format(data interface{}, templateText string) (string, error) {
	tmpl := template.New("markdown")

	if _, err := tmpl.Parse(templateText); err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (f *MarkdownFormatter) FormatFile(data interface{}, templateFile string) (string, error) {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
