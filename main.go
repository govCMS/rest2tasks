package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type input struct {
	projectName string `yaml:"projectName"`
	branchName  string `yaml:"branchName"`
}

var DeprecationMessage = "This endpoint has been deprecated, your deployment has been actioned by the GitLab webhook."
var Port = "3000"

func bodyProcessor(w http.ResponseWriter, r *http.Request, Data *input) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}

	data := strings.Split(string(body), "&")
	for _, line := range data {
		for _, v := range strings.Split(line, "\n") {
			if strings.HasPrefix(v, "projectName") {
				if len(strings.Split(v, "=") )== 2 {
					Data.projectName = strings.Split(v, "=")[1]
				}
			}
			if strings.HasPrefix(v, "branchName") {
				if len(strings.Split(v, "=")) == 2 {
					Data.branchName = strings.Split(v, "=")[1]
				}
			}
		}
	}
}

func printOutput(w http.ResponseWriter, r *http.Request, output string) {
	fmt.Println(output)
	fmt.Fprintf(w, "{\"ok\": \"true\", \"warning\": \"welcome to rest2tasks - %v\", \"message\": %v}\n", DeprecationMessage, output)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	printOutput(w, r, "")
}

func mergeRequestDeploy(w http.ResponseWriter, r *http.Request) {
	var Data input
	bodyProcessor(w, r, &Data)
	if Data.projectName != "" && Data.branchName != "" {
		printOutput(w, r, fmt.Sprintf("rest2tasks: Attempted to deploy MR for branch '%v' on project '%v'", Data.branchName, Data.projectName))
	}
}

func deploy(w http.ResponseWriter, r *http.Request) {
	var Data input
	bodyProcessor(w, r, &Data)
	if Data.projectName != "" && Data.branchName != "" {
		printOutput(w, r, fmt.Sprintf("rest2tasks: Attempted to deploy branch '%v' on project '%v'", Data.branchName, Data.projectName))
	}
}

func promote(w http.ResponseWriter, r *http.Request) {
	var Data input
	bodyProcessor(w, r, &Data)
	if Data.projectName != "" && Data.branchName != "" {
		printOutput(w, r, fmt.Sprintf("rest2tasks: Attempted to promote branch '%v' on project '%v'", Data.branchName, Data.projectName))
	}
}

func main() {
	// Welcome message
	fmt.Println("welcome to rest2tasks - ", DeprecationMessage)
	fmt.Printf("Listening on port %v\n\n", Port)

	// Add handle funcs
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/pullrequest/deploy", mergeRequestDeploy)
	http.HandleFunc("/deploy", deploy)
	http.HandleFunc("/promote", promote)

	// Run the web server.
	log.Fatal(http.ListenAndServe(":"+Port, nil))
}
