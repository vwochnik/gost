package template

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"html/template"

	"github.com/vwochnik/gost/asset"
)

type errorTplVars struct {
	Code    int
	Message string
}

type directory struct {
	Name  string
	IsDir bool
	URL   string
}

type dirTplVars struct {
	Root        string
	Directories []directory
}

var (
	funcMap      = template.FuncMap{}
	errorTpl     = template.Must(template.New("error").Funcs(funcMap).Parse(string(asset.MustAsset("error.tpl"))))
	dirTpl       = template.Must(template.New("directory").Funcs(funcMap).Parse(string(asset.MustAsset("directory.tpl"))))
	htmlReplacer = strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		`"`, "&#34;",
		"'", "&#39;",
	)
)

func ErrorTemplate(w io.Writer, msg string, code int) {
	errorTpl.Execute(w, &errorTplVars{code, msg})
}

func DirectoryTemplate(w io.Writer, f http.File) {
	stat, err := f.Stat()
	if err != nil {
		return
	}

	vars := &dirTplVars{stat.Name(), make([]directory, 0, 0)}
	for {
		dirs, err := f.Readdir(100)
		if err != nil || len(dirs) == 0 {
			break
		}
		for _, d := range dirs {
			name := d.Name()
			url := url.URL{Path: name}
			vars.Directories = append(vars.Directories, directory{
				htmlReplacer.Replace(name), d.IsDir(), url.String(),
			})
		}
	}

	dirTpl.Execute(w, vars)
}
