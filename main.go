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
import "gopkg.in/libgit2/git2go.v22"
import "os"
import "log"

type SessionData struct {
        Username    string `json:"username"`
        Token   string `json:"private_token"`
    }
type ProjectsData []struct {
        ID         int `json:"id"`
        Name         string `json:"name"`
        RepoSSH      string `json:"ssh_url_to_repo"`
        Path string `json:"path"`
        //Namespace map[string]string `json:"namespace"`
        Namespace struct {
            Path string
        }
    }
type ProjectData struct {
        ID         int `json:"id"`
        Name         string `json:"name"`
        RepoSSH      string `json:"ssh_url_to_repo"`
        Path string `json:"path"`
        //Namespace map[string]string `json:"namespace"`
        Namespace struct {
            Path string
        }
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
    url := fmt.Sprintf("https://git.permissiondata.com/api/v3/projects/?private_token=%s", token)
    //url := fmt.Sprintf("https://git.permissiondata.com/api/v3/projects/108/?private_token=%s", token)
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


func getProject(token string, id int) ProjectData {
    

    var data = ProjectData{}

    //TODO: move to a config file
    url := fmt.Sprintf("https://git.permissiondata.com/api/v3/projects/%d/?private_token=%s", id, token)
    //url := fmt.Sprintf("https://git.permissiondata.com/api/v3/projects/108/?private_token=%s", token)
    fmt.Println("URL: ", url)

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

    //fmt.Println("response Status:", resp.Status)
    //fmt.Println("response Headers:", resp.Header)
    //body, _ := ioutil.ReadAll(resp.Body)
    //fmt.Println("response Body:", string(body))

    dec := json.NewDecoder(resp.Body)
    dec.Decode(&data)
    fmt.Printf("data:%+v \n", data)

    return data
}


func credentialsCallback(url string, username string, allowedTypes git.CredType) (git.ErrorCode, *git.Cred) {
    ret, cred := git.NewCredSshKey("git", "/home/witek/.ssh/for_allan.pub", "/home/witek/.ssh/for_allan", "")
    return git.ErrorCode(ret), &cred
}

// Made this one just return 0 during troubleshooting...
func certificateCheckCallback(cert *git.Certificate, valid bool, hostname string) git.ErrorCode {
    return 0
}

func main() {


    repoSiteCode := flag.String("registryType","gitlab", "possible values: gitlab (github in the future)")
    destinationPath := flag.String("d", "~/tmp_git", "clone to this location")
    username := flag.String("u", "", "")
    password := flag.String("p", "", "")

    flag.Parse()

    fmt.Println("Hello")
    fmt.Printf("repoSiteCode:%s destinationPath:%s", *repoSiteCode, *destinationPath)
    
    session := getSession(*username, *password)
    fmt.Println("Token: %+v \n", session)

    projects := getProjects(session.Token)
    fmt.Println("Projects: ", projects)


    //fmt.Println("Project0-RepoSSH: ", projects[0].RepoSSH)
    //fmt.Printf("Project0: %+v \n", projects[0])

    cbs := &git.RemoteCallbacks{
        CredentialsCallback:      credentialsCallback,
        CertificateCheckCallback: certificateCheckCallback,
    }

    cloneOptions := &git.CloneOptions{}
    cloneOptions.RemoteCallbacks = cbs

    //create folder name
    //projectDestinationPath := fmt.Sprintf("%s/%s_%s", *destinationPath, projects[0].Namespace["Path"], projects[0].Path)
    //fmt.Println("projectDestinationPath: ", projectDestinationPath)
    //_, err := git.Clone(projects[0].RepoSSH, projectDestinationPath, cloneOptions)


    //project := getProject(session.Token, 10)
    //fmt.Println("\n\n\n\nProject10: ", project)

    //project = getProject(session.Token, 11)
    //fmt.Println("\n\n\n\nProject11: ", project)


    //TODO: remove
    //clone
            //fmt.Println("clone")
            //_, err := git.Clone(project.RepoSSH, projectDestinationPath, cloneOptions)    
            //if err != nil {
            //    panic(err)
            //}
    //return


    for idx, project := range projects {
        //TODO: remove
        fmt.Println("index: ", idx)
        //if(idx == 3){
        //    break;
        //}

        projectDestinationPath := fmt.Sprintf("%s/%s/%s", *destinationPath, project.Namespace.Path, project.Path)
        

        //TODO: remove
        //projects[0].RepoSSH = "git@git.permissiondata.com:wciemiega/test1.git"
        //projectDestinationPath = fmt.Sprintf("%s/%s/%s", *destinationPath, "wciemiega", "test1")

        //TODO: pull the repos to destination
        fmt.Println("projectDestinationPath: ", projectDestinationPath)
        if _, err := os.Stat(projectDestinationPath); os.IsNotExist(err) {
            //clone
            fmt.Println("clone")
            _, err := git.Clone(project.RepoSSH, projectDestinationPath, cloneOptions)    
            if err != nil {
                panic(err)
            }
        }else{
            //fetch all
            fmt.Println("fetch all")
            repo, err:= git.OpenRepository(projectDestinationPath)
            log.Println(repo) //print repo address

            remote, err:= repo.LookupRemote("origin")
            //fmt.Printf("remote: %+v \n", remote)
            
            //callbacks := git.RemoteCallbacks{
            //      
            //}
            //err = remote.SetCallbacks(&callbacks)
            //err = remote.Connect(git.ConnectDirectionFetch)
            //err = remote.ConnectFetch()
            err = remote.SetCallbacks(cbs)
            if err != nil {
                panic(err)
            }

            err = remote.Fetch([]string{}, nil, "")
            if err != nil {
                panic(err)
            }
            //refs := []string{"+refs/heads/*:refs/remotes/origin/*"}
            //remote.Fetch(refs, nil, "")

            //remote_master, err := repo.LookupReference("refs/remotes/origin/master")
            //mergeRemoteHead, err := repo.AnnotatedCommitFromRef(remote_master)
            //mergeHeads := make([]*git.AnnotatedCommit, 1)
            //mergeHeads[0] = mergeRemoteHead
            //err = repo.Merge(mergeHeads, nil, nil)
            //repo.StateCleanup()
        }
    }
}
