package template

import (
	"io"
	"net/http"

	"html/template"

	"github.com/vwochnik/gost/asset"
)

type errorTplVars struct {
	Code    int
	Message string
}

type directory struct {
	Name string
	URL  string
}

type dirTplVars struct {
	Root        string
	Directories []directory
}

var (
	funcMap  = template.FuncMap{}
	errorTpl = template.Must(template.New("error").Funcs(funcMap).Parse(string(asset.MustAsset("error.tpl"))))
)

func ErrorTemplate(w io.Writer, msg string, code int) {
	errorTpl.Execute(w, &errorTplVars{code, msg})
}

func DirectoryTemplate(w io.Writer, f http.File) {
	errorTpl.Execute(w, &errorTplVars{code, msg})
}
