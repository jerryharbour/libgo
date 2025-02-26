package formatter

import (
	"bytes"
	"text/template"
)

type HtmlFormatter struct{}

func (f *HtmlFormatter) Format(data interface{}, templateText string) (string, error) {
	tmpl := template.New("html")

	if _, err := tmpl.Parse(templateText); err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (f *HtmlFormatter) FormatFile(data interface{}, templateFile string) (string, error) {
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
