import "github.com/gorilla/sessions"

var store = sessions.NewCookieStore([]byte("33446a9dcf9ea060a0a6532b166da32f304af0de"))

func Handler(w http.ResponseWriter, r *http.Request){
    session, _ := store.Get(r, "session-name")

    session.Values["foo"] = "bar"
    session.Values[42] = 43
    session.Save(r, w)

    fmt.Fprint(w, "Hello world :)")
}

func main(){
    store.Options = &sessions.Options{
        Domain:     "localhost",
        Path:       "/",
        MaxAge:     60 * 15,
        Secure:     false,
        HttpOnly:   true,
    }
}