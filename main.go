package main
import ("net/http";"os")
func main(){http.HandleFunc("/healthz",func(w http.ResponseWriter,_ *http.Request){w.WriteHeader(200);w.Write([]byte("ok"))});port:=os.Getenv("PORT");if port==""{port="8080"};http.ListenAndServe(":"+port,nil)}
