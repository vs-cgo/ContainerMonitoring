package server

import (
  "encoding/json"
  "net/http"
  "backend/db"
)

func init() {
  
}

type Container struct {
  ID   string `json:"id"`
  IP   string `json:"ip"`
  Ping string `json:"ping"`
  Time string `json:"time"`
}

type Info  struct {
  Containers []Container `json:"containers"`
}

func Start() {
  server := http.Server(Addr: 8888, Handler: nil) 
  http.HandleFunc("/get", getHandler)
  http.HandleFunc("/set", setHandler)

  log.Print("Server listen at %s", )
  log.Println(server.ListenAndServe())
}

func getHandler(w http.ResponseWriter, r *http.Request) {  
  res := db.Get()
  err := json.NewEncoder(w).Encode(res)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Header().Set("Content-Type", "application/json; charset=utf-8")
  w.WriteHeader(http.StatusOK)
}

func setHandler(w http.ResponseWriter, r *http.Request) {
  var c Container
  err := json.NewDecoder(r.Body).Decode(&c)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  defer r.Body.Close()
  w.Header().Set("Content-Type", "application/json; charset=utf-8")

  err = db.Set(c)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.WriteHeader(http.StatusCreated)
}
