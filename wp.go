package main

import (
    "fmt"
    "html/template"
    "net/http"
)

func isTor(ip string) bool{
    /*
        Function takes and IP and evaluates if it is a Tor node or not
        This is done by checking the user's IP against a list of exit nodes
        at https://check.torproject.org/
    */

    return false
}

func shipData(r *http.Request){
    /*
        shipData collects data from the request and stores it in our data store
        Currently, this is only designed to work with MongoDB
        The data we collect is:
            + URI Host and Path
            + User-Agent
            - Cookies
            - Form-Data
            - Tor (is the IP a Tor exit node?)
    */
    reqData := make(map[string]string)
    reqData["host"] = r.Host
    reqData["uri"] = r.URL.Path
    reqData["user-agent"] = r.UserAgent()
    //reqData["cookies"] = strings.Join(r.Cookies(), ",")
    reqData["ip"] = r.Header.Get("X-Forwarded-For")
    //reqData["headers"] = r.Header
    //reqData["form-data"] = strings.Join(r.PostForm,",")

    /*
    make(map[string]string)

    if err := r.ParseForm(); err != nil{
        // nothing to do here
    } else {
        for key, value := range r.PostForm {
            reqData["form-data"][key] := value
            fmt.Println("FORM DATA: ", key, value)
        }
    }
    */
    fmt.Printf("%v", reqData)
}

func handler(w http.ResponseWriter, r *http.Request) {
    /*
        handler just supports the index page. It is also used to pass
        the request to shipData in order to store data in our DB.

        First we make sure the URL is either the root url or isn't a static
        asset, then we store data about the request in shipData
    */

    u := len(r.URL.Path)
    if (u > 3) {
        if (r.URL.Path[u-3:] != "css") && (r.URL.Path[u-2:] != "js") {
            shipData(r)
        }
    } else if (r.URL.Path == "/") || (r.URL.Path == "") {
        shipData(r)
    }

    t, err := template.New("").ParseFiles("tmpl/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    // Data is a map that we use to ship data to our template
    Data := map[string]string {
        "Title": "DeadBeef",
        "Subtitle": "Cattle never been better",
        "RequestURL": r.Host,
    }
    err = t.ExecuteTemplate(w, "Data", Data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func adminHandler(w http.ResponseWriter, r *http.Request){
    /*
        adminRedirect should provide a 302 redirect that mimics wp-admin to fool scripts
    */

    t, err := template.New("").ParseFiles("tmpl/wp-login.php")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    // Data is a map that we use to ship data to our template
    Data := map[string]string {
        "Title": "DeadBeef",
        "Subtitle": "Cattle never been better",
        "RequestURL": r.Host,
    }
    err = t.ExecuteTemplate(w, "Data", Data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}

func main(){
    /*
        Main sets up our http Router and listens for requests
        Important routes are:
            - /wp-content/ - Our static file directory
            - /wp-admin/ - The administrative panel expected from Wordpress
            - /xmlrpc.php - The XML RPC service, most common attack vector for Wordpress
    */
    http.HandleFunc("/", handler)
    http.HandleFunc("/wp-login.php", adminHandler)
    http.HandleFunc("/wp-content/", func(w http.ResponseWriter, r *http.Request){
        uri := "./static/" + r.URL.Path[1:]

        fmt.Println(uri)
        http.ServeFile(w, r, uri)
    })
    http.HandleFunc("/wp-admin/", func(w http.ResponseWriter, r *http.Request){
         uri := "./static/" + r.URL.Path[1:]

        fmt.Println(uri)
        http.ServeFile(w, r, uri)
    })
    fmt.Println("Listening..")
    http.ListenAndServe(":80", nil)
}
