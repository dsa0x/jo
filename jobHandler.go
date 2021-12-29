package docman

import (
	"fmt"
	"net/http"
	"strconv"
)

type JobResp struct {
	ServerPort string
	Jobs       []Job
	Error      string
}

func jobHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case http.MethodGet:
		getJob(w, r)
	case http.MethodPost:
		postJob(w, r)
	case http.MethodDelete:
		deleteJob(w, r)
	}
}
func jobsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case http.MethodGet:
		getAllJobs(w, r)
	case http.MethodPost:
		postJob(w, r)
	case http.MethodDelete:
		deleteJob(w, r)
	}
}

func getJob(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		http.Redirect(w, r, "/jobs", http.StatusSeeOther)
		return
	}

	_id, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		respondError(w, "Parse error")
		return
	}
	jobs, err := FindByID(uint(_id))
	if err != nil {
		respondError(w, err.Error())
		return
	}

	resp := &JobResp{ServerPort: Cfg.PORT, Jobs: jobs}

	err = tplParser(w, "index.gohtml", resp)
	if err != nil {
		respondError(w, err.Error())
		return
	}
}

func getAllJobs(w http.ResponseWriter, r *http.Request) {
	jobs, err := FindAll()
	if err != nil {
		respondError(w, err.Error())
		return
	}

	resp := &JobResp{ServerPort: Cfg.PORT, Jobs: jobs}

	err = tplParser(w, "index.gohtml", resp)
	if err != nil {
		respondError(w, err.Error())
		return
	}
}

func postJob(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	desc := r.FormValue("description")
	if name == "" || desc == "" {
		respondError(w, "Name or Description cannot be empty")
		return
	}

	job := &Job{Name: name, Description: desc}
	err := job.Create()
	if err != nil {
		respondError(w, err.Error())
		return
	}
	fmt.Println(job)

	http.Redirect(w, r, "/jobs", http.StatusSeeOther)
}

func deleteJob(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		respondError(w, "ID cannot be empty")
		return
	}

	_id, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		respondError(w, "Parse error")
		return
	}

	job := &Job{ID: uint(_id)}

	err = job.Delete()
	if err != nil {
		respondError(w, "DB error")
		return
	}

	http.Redirect(w, r, "/jobs", http.StatusSeeOther)
}
