package docman

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"reflect"
)

var (
	templatesDir = "templates"
	templates    = make(map[string]*template.Template)
)

//go:embed templates/index.gohtml
var files embed.FS

func NewWeb() *http.ServeMux {
	r := http.NewServeMux()
	err := LoadTemplates()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/health", healthCheck)
	r.HandleFunc("/job", jobHandler)
	r.HandleFunc("/job/delete", deleteJob)
	r.HandleFunc("/jobs", jobsHandler)

	return r
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getHome(w, r)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

func getHome(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/jobs", http.StatusSeeOther)
}

func respondError(w http.ResponseWriter, errs ...string) {
	w.WriteHeader(http.StatusBadRequest)
	err := "An error occured"
	if len(errs) > 0 {
		err = ""
		for _, e := range errs {
			err += ";" + e
		}
	}
	resp := &JobResp{ServerPort: Cfg.PORT, Error: err}
	_err := tplParser(w, "index.gohtml", resp)
	if _err != nil {
		fmt.Println(_err)
		w.Write([]byte(`{"error":"internal server error"}`))
		return
	}
}

func parseRequest(r *http.Request, v interface{}) error {
	if reflect.TypeOf(v).Kind() != reflect.Ptr {
		return fmt.Errorf("v must be a pointer")
	}

	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return err
	}
	return nil
}

func tplParser(w http.ResponseWriter, tplName string, data interface{}) error {
	tpl, ok := templates[tplName]
	if !ok {
		return fmt.Errorf("template %s not found", tplName)
	}
	return tpl.ExecuteTemplate(w, tplName, data)
}

func LoadTemplates() error {
	tmplFiles, err := fs.ReadDir(files, templatesDir)
	if err != nil {
		return err
	}
	for _, f := range tmplFiles {
		if f.IsDir() {
			continue
		}
		tpl, err := template.ParseFS(files, templatesDir+"/"+f.Name())
		if err != nil {
			return err
		}
		templates[f.Name()] = tpl
	}
	return nil
}
