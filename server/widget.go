package server

import (
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

func RenderWidget() string {
	tmpl, err := template.ParseFS(JSWidgetContent, "widget.js") //template.ParseFiles(filename)
	if err != nil {
		// return "", err
		log.Println("ERR" + err.Error())
		return "Render ERR" + err.Error()
	}

	var buf []byte
	w := &bufferWriter{&buf}

	var templateData map[string]any = map[string]any{}
	templateData["scriptIngestURL"] = "http://localhost:8089/script/ingest"

	if err := tmpl.Execute(w, templateData); err != nil {
		log.Println("ERR" + err.Error())
		return "Render ERR" + err.Error()
	}

	return string(buf)
}
