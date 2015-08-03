package main

import (
    "encoding/json"
    "flag"
    "fmt"
    //"log"
    //"io/ioutil"
    "net/http"
    //"net/url"
    "bytes"
)

type SessionData struct {
        Username	string `json:"username"`
        Token	string `json:"private_token"`
    }
type ProjectsData []struct {
        Name         string `json:"name"`
            RepoSSH      string `json:"ssh_url_to_repo"`
    }


func getSession(username string, password string) SessionData {
	

    var sessionData = SessionData{}

    //TODO: move to a config file
	url := "https://git.permissiondata.com/api/v3/session/"
    fmt.Println("URL:>", url)

    var dataStr = fmt.Sprintf("login=%s&password=%s", username, password)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(dataStr)))
    //req.Header.Set("X-Custom-Header", "myvalue")
    //req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    //body, _ := ioutil.ReadAll(resp.Body)
    //fmt.Println("response Body:", string(body))

    dec := json.NewDecoder(resp.Body)
    dec.Decode(&sessionData)
    fmt.Printf("sessionData:%+v \n", sessionData)

    return sessionData
}

func getProjects(token string) ProjectsData {
	

    var data = ProjectsData{}

    //TODO: move to a config file
	url := fmt.Sprintf("https://git.permissiondata.com/api/v3/projects/owned/?private_token=%s", token)
    fmt.Println("URL:>", url)

    req, err := http.NewRequest("GET", url, nil)
    //req.Header.Set("X-Custom-Header", "myvalue")
    //req.Header.Set("Content-Type", "application/json")
    //req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    //body, _ := ioutil.ReadAll(resp.Body)
    //fmt.Println("response Body:", string(body))

    dec := json.NewDecoder(resp.Body)
    dec.Decode(&data)
    fmt.Printf("data:%+v \n", data)

    return data
}


func main() {


	repoSiteCode := flag.String("registryType","gitlab", "possible values: gitlab (github in the future)")
	destinationPath := flag.String("d", ".", "clone to this location")
	username := flag.String("u", "", "")
	password := flag.String("p", "", "")

	flag.Parse()

	fmt.Println("Hello")
	fmt.Printf("repoSiteCode:%s destinationPath:%s", *repoSiteCode, *destinationPath)
	
	session := getSession(*username, *password)
	fmt.Println("Token: %+v \n", session)

	projects := getProjects(session.Token)
	fmt.Println("Projects: ", projects)



	//TODO: get the html
	//TODO: parse the repo links
	//TODO: pull the repos to destination
}
