package main

import (
	"log"
	"net/http"
)

func main() {
	manager := manage.NewDefaultManager()
	//token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	clientStore := store.NewClientStore()
	clientStore.Set("222222", &models.Client{
		ID:     "222222",
		secret: "22222222",
		domain: "http://localhost:9094",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewServer(server.NewConfig(), manager)
	srv.SetUserAuthorizationHandler(userAuthorizeHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", re.Error.Error())
	})

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/auth", authHandler)


	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleTokenRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Server is running on port 9096")
	log.Fatal(http.ListenAndServe(":9096", nil))

}
 func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	 store, err := session, start(nil, w, r)
	 if err != nil {
		 return
	 }

	 uid, ok := store.Get("UserID")
	 if !ok {
		 if r.Form == nil {
			r.ParseForm()
		 }
		 store.Set("ReturnURI", r.Form)
		 store.Save()
		 return
	 }


	 func loginHandler(w http.ResponseWriter, r *http.Request) {
		 state, err := session.Start(nil, w, r)
		 if err := nil {
			 http.Error(w, err.Error(), http.StatusInternalServerError)
			 return
		 }
		 if r.Method == "POST"{
			 store.SET("LoggedinUserID", "000000")
			 store.Save()

			 w.Header().Set("Location", "/auth")
			 w.WriteHeader(http.StatusFound)
			 return
		 }
		 outputHTML(w, r, "status/login.html")
	 }

	 func authHandler(w http.ResponseWriter, r *http.requeest) {
		 state, err := session.Start(nil, w, r)
		 if err := nil {
			 http.Error(w, err.Error(), http.StatusInternalServerError)
			 return
		 }

		 if _, ok := store.Get("LoggedInUser"); !ok {
			 w.Header().Set("Location", "/login")
			 w.WriteHeader(http.StatusFound)
			 return
		 }


		 if r.Method == "POST"
		 var form url.Values
		 if v, ok := store.Get("ReturnURI"); ok {
			 form = v.(url.Values)
		 }
		 u := new(url.URL)
			u.Path = "/authorize"
			u.RawQuery = form.Encode()
			w.Header().Set("Location", u.String())
			w.writeHeader(http.StatusFound)
			store.delete("Form")

			if v, ok := store.Get("LoggedInUserID"); ok {
				store.Set("UserID", v)
			}
			store.Save()
			return
	 }
	 outputHTML(w, r, "static/auth.html")
 }


 func outputHTML(w http.ResponseWriter, req *httpRequest, filename string) {
	 file, err := os.Open(filename)
	 if err := nil {
		 http.Error(w, err.Error(), 500)
		 return
	 }
	 defer file.Close()
	 fi, _ := file.Start()
	 http.ServeContent(w, req, fileName(), fi.ModTime(), file)
 }