package main
import (
    "encoding/json"
    "log"
    "net/http"
    "fmt"
    "github.com/gorilla/mux"
    "io/ioutil"
)
// Define the type Struct.

type Server struct {
    ID        string   `json:"id,omitempty"`
    Name      string   `json:"name,omitempty"`
    Cores     string   `json:"cores,omitempty"`
    Memory    string   `json:"memory,omitempty"`
    Disk      string   `json:"disk,omitempty"`
}

var servers []Server
func GetServer(w http.ResponseWriter, req *http.Request) {

    //Show the server information with the ID that we pass on header.

    params := mux.Vars(req)
    fmt.Println("The server with ID", params["id"], "has been requested")
    for _, item := range servers {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Server{})
}
func GetAllservers(w http.ResponseWriter, req *http.Request) {

    //Show the all servers's information on web.

    fmt.Println("All servers have been requested")
    json.NewEncoder(w).Encode(servers)
}
func CreateServer(w http.ResponseWriter, req *http.Request) {

    //Create the server whith the dates that we pass on the body.

    w.Header().Set("Content-Type", "application/json")

    var newserver Server
    reader, _ := ioutil.ReadAll(req.Body)
         json.Unmarshal(reader, &newserver)

    servers = append(servers, newserver)

    json.NewEncoder(w).Encode(servers)
    fmt.Println("The new server has been created")
}

func DeleteServer(w http.ResponseWriter, req *http.Request) {

    //Delete the server that we pass on header.

    params := mux.Vars(req)
    for index, item := range servers {
        if item.ID == params["id"] {
            servers = append(servers[:index], servers[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(servers)
    fmt.Println("The server with ID", params["id"], "has been deleted")
}

func Test(w http.ResponseWriter, req *http.Request) {

    //For check that the api is goingo OK run this test that answere OK on web and console.

	response := "the API is OK"
        fmt.Println("Test executed, the API is OK")
	json.NewEncoder(w).Encode(response)
}
func wellcomeMessage() {

  //When we start the api shows the missage in the console.

  fmt.Println("")
  fmt.Println("API STARTED SUCCESSFULLY!")
  fmt.Println("")
}
func main() {
    wellcomeMessage()
    router := mux.NewRouter()


    //Generate the routes for access.
    router.HandleFunc("/test", Test)
    router.HandleFunc("/servers", GetAllservers).Methods("GET")
    router.HandleFunc("/servers/{id}", GetServer).Methods("GET")
    router.HandleFunc("/servers/{id}", CreateServer).Methods("POST")
    router.HandleFunc("/servers/{id}", DeleteServer).Methods("DELETE")

    //Generate 3 example servers.
    servers = append(servers, Server{ID: "1", Name: "vxhcc-10", Cores: "5", Memory: "10 GB", Disk: "1024"})
    servers = append(servers, Server{ID: "2", Name: "vxhcc-20", Cores: "1", Memory: "2 GB", Disk: "2048"})
    servers = append(servers, Server{ID: "3", Name: "vxhcc-30", Cores: "3", Memory: "6 GB", Disk: "1024"})

    //the API is listened a server on the 8080 port.
    log.Fatal(http.ListenAndServe(":8080", router))
}
