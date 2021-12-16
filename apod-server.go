package main
import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)
var templates = template.Must(template.ParseFiles("templates/index.html"))
type ApodData struct {
	Url string `json:"url"`
	Date string `json:"date"`
	Title string `json:"title"`
	Copyright string `json:"copyright"`
	Explanation string `json:"explanation"`
	HDUrl string `json:"hdurl"`
	ServiceVer string `json:"service_version"`
	MediaType string `json:"media_type"`
}
func fetchApodData(r *http.Request) (*ApodData, error) {
	myClient := http.Client{}
	query := r.URL.Query()
	key := "your_key_here"
	url := fmt.Sprintf("https://api.nasa.gov/planetary/apod?api_key=%s&date=%s", key, query["date"][0])
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil { return nil, err }
	res, err := myClient.Do(req)
	if err != nil { return nil, err }
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil { return nil, err }
	a := ApodData{}
	err = json.Unmarshal(body, &a)
	if err != nil { return nil, err }
	return &a, nil
}
func getApod(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if len(query) == 0 {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	a, err := fetchApodData(r)
	if err != nil {
		return
	}
	renderPage(w, "index", a)
}
func renderPage(w http.ResponseWriter, tmpl string, a *ApodData) {
	err := templates.ExecuteTemplate(w, tmpl+".html", a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}
func main() {
	public := http.StripPrefix("/public", http.FileServer(http.Dir("public")))
	http.Handle("/public/", public)
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		foo := fmt.Sprintf("%s", currentTime)
		url := fmt.Sprintf("/apodtoday/?date=%s", foo[:10])
		http.Redirect(w, r, url, http.StatusFound)
	})
	http.HandleFunc("/apodtoday/", makeHandler(getApod))
	log.Fatal(http.ListenAndServe("127.0.0.1:9999", nil))
}