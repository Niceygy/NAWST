package utils

import (
	"embed"
	"log"
	"text/template"
)

type bufferWriter struct {
	buf *[]byte
}

func (w *bufferWriter) Write(p []byte) (int, error) {
	*w.buf = append(*w.buf, p...)
	return len(p), nil
}

func RenderTemplate(vfs embed.FS, filename string, renderData map[string]any) string {
	tmpl, err := template.ParseFS(vfs, filename)
	if err != nil {
		// return "", err
		log.Println("ERR" + err.Error())
		return "Render ERR" + err.Error()
	}

	var buf []byte
	w := &bufferWriter{&buf}

	if err := tmpl.Execute(w, renderData); err != nil {
		log.Println("ERR" + err.Error())
		return "Render ERR" + err.Error()
	}

	return string(buf)
}
