package main

import (
	"fmt"
	//"html/template"
	"bytes"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"text/template"
	"strings"
)

type answerPageData struct {
	Question string
	Answer   string
}

func main() {

	http.HandleFunc("/jsoneditor.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("dist", "jsoneditor.css"))
	})

	http.HandleFunc("/jsoneditor.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("dist", "jsoneditor.js"))
	})
	
	http.HandleFunc("/img/jsoneditor-icons.svg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("dist/img", "jsoneditor-icons.svg"))
	})
	

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles(filepath.Join("html", "index.html")))
		msg := "test"
		templ.Execute(w, msg)
	})

	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		query := r.PostFormValue("query")
		//query_esc := url.QueryEscape(query)
		query_esc := url.PathEscape(query)
		stargraph_server := os.Getenv("STARGRAPH_SERVER_IP")
		request_url := fmt.Sprintf("http://%s:8917/_kb/dbpedia-2016/query?q=%s?", stargraph_server, query_esc)

		client := &http.Client{}

		resp, err := client.Get(request_url)
		if err != nil {
			fmt.Printf("ERROR %s", err)
			return
		}
		defer resp.Body.Close()

		fmt.Printf("response status : %s\n", resp.Status)
		fmt.Printf("response code : %d\n", resp.StatusCode)

		buf := new(bytes.Buffer)
		if _, err := buf.ReadFrom(resp.Body); err != nil {
			fmt.Printf("ERROR reading body : %s", err)
			return
		}

		answer := strings.Replace(buf.String(), "\"", "'", -1)
		answer = strings.Replace(answer, "\\n", " ", -1)

		data := answerPageData{
			Question: query,
			Answer:   answer,
		}

		templ := template.Must(template.ParseFiles(filepath.Join("html", "answer.html")))
		templ.Execute(w, data)

	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("ERROR : %s\n", err)
		os.Exit(1)

	}

}
