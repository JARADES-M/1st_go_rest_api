package main

import (
    "fmt"
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "io/ioutil"
)

// The person Type (more like an object)
type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}
type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}

var people []Person

func handler(w http.ResponseWriter, r *http.Request) {
    stream, err := ioutil.ReadFile("home.html")
    if err != nil {
        log.Fatal(err)
    }
    htmlfile := string(stream)
    fmt.Fprintf(w, htmlfile)
}

// Display all from the people var
func GetPeople(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(people)
}

// Display a single data
func GetPerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range people {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Person{})
}

// create a new item
func CreatePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)
    person.ID = params["id"]
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}

// Delete an item
func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(people)
    }
}

// main function to boot up everything
func main() {
    router := mux.NewRouter()
    people = append(people, Person{ID: "1", Firstname: "Jarades", Lastname: "Monteiro", Address: &Address{City: "Ohaio", State: "Keopata"}})
    people = append(people, Person{ID: "2", Firstname: "Maria", Lastname: "Silva", Address: &Address{City: "Teófilo Otoni", State: "Minas Gerais"}})
    people = append(people, Person{ID: "3", Firstname: "João", Lastname: "José", Address: &Address{City: "Frei Gaspar", State: "Minas Gerais"}})
    people = append(people, Person{ID: "4", Firstname: "Daniel", Lastname: "Eykel", Address: &Address{City: "St. Louis", State: "Missouri"}})
    router.HandleFunc("/", handler)
    router.HandleFunc("/people", GetPeople).Methods("GET")
    router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
    router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":80", router))
}
