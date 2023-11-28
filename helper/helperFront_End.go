package helper

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

func RenderTemplate(w http.ResponseWriter, tmplName string, tmplDir string, data interface{}) {
	templateCache, err := createTemplateCache(tmplDir)

	if err != nil {
		panic(err)
	}
	// templateCache["home.page.tmpl"]
	tmpl, ok := templateCache[tmplName+".page.tmpl"]

	if !ok {
		http.Error(w, "le template n'existe pas", http.StatusInternalServerError)
		return
	}

	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, data)
	buffer.WriteTo(w)
}

func createTemplateCache(tmplDir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./template/pages/" + tmplDir + "/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.ParseFiles(page))

		layouts, err := filepath.Glob("./template/layouts/*.layout.tmpl")
		if err != nil {
			return cache, err
		}
		if len(layouts) > 0 {
			tmpl.ParseGlob("./template/layouts/*.layout.tmpl")
		}
		cache[name] = tmpl
	}
	return cache, nil
}

func RenderError(w http.ResponseWriter, tmplName string, tmplDir string) {
	log.Println("RenderErrorStart")
	templateCache, err := createTemplateCache(tmplDir)

	if err != nil {

		return
	}
	// templateCache["home.page.tmpl"]

	tmpl, ok := templateCache[tmplName+".page.tmpl"]

	if !ok {
		RenderTemplate(w, "404", "error", 404)
		//http.Error(w, "le template n'existe pas", http.StatusInternalServerError)
		return
	}

	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, nil)
	buffer.WriteTo(w)
	log.Println("RenderError end")
}

func ErrorPage(w http.ResponseWriter, i int) {
	DataError := struct {
		Code    string
		Message string
	}{
		Code:    strconv.Itoa(i),
		Message: http.StatusText(i),
	}
	w.WriteHeader(i)
	RenderTemplate(w, "error", "error", DataError)

}
