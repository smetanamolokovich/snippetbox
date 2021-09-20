package main

import (
	"path/filepath"
	"text/template"
	"time"

	"github.com/smetanamolokovich/snippetbox/pkg/forms"
	"github.com/smetanamolokovich/snippetbox/pkg/models"
)

type templateData struct {
	CSRFToken       string
	CurrentYear     int
	Form            *forms.Form
	Flash           string
	IsAuthenticated bool
	Snippet         *models.Snippet
	User            *models.User
	Snippets        []*models.Snippet
}

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.UTC().Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	chache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		chache[name] = ts
	}
	return chache, nil
}
