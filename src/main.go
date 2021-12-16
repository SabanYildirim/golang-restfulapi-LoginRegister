package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func IsEmpty(data string) bool {
	if len(data) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	
    r := mux.NewRouter()

	r.HandleFunc("/api/signup",signup).Methods("POST")
	r.HandleFunc("/api/login",login).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func login(w http.ResponseWriter, r *http.Request)  {		
		w.Header().Set("Content-Type", "application/json")
		r.ParseForm()

		email := r.FormValue("email")
		pwd := r.FormValue("password")
 
		fmt.Printf(email)
		fmt.Printf(pwd)
		
		emailCheck := IsEmpty(email)
		pwdCheck := IsEmpty(pwd)

		if emailCheck || pwdCheck {
			fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
		}

		dbEmail := "admin@admin"
        dbPwd := "admin"

		if email == dbEmail && pwd == dbPwd {
			fmt.Fprintf(w,"Login succesful.")
		} else {
			fmt.Fprintf(w,"Login failed.")
		}
}


func signup(w http.ResponseWriter, r *http.Request){
		    
			uName, email, pwd, pwdConfirm := "", "", "", ""

			w.Header().Set("Content-Type", "application/json")
			r.ParseForm()
	
			uName = r.FormValue("username")
			email = r.FormValue("email")
			pwd = r.FormValue("password")
			pwdConfirm = r.FormValue("Confirm")
	
			uNameCheck := IsEmpty(uName)
			emailCheck := IsEmpty(email)
			pwdCheck := IsEmpty(pwd)
			pwdConfirmCheck := IsEmpty(pwdConfirm)
	
			if uNameCheck || emailCheck || pwdCheck || pwdConfirmCheck {
				
				fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
				return
			}
	
			if pwd == pwdConfirm {
				//veritabanına kullanıcı kaydet
				fmt.Fprintf(w, "Registration succesful.")
			} else {
				fmt.Fprintf(w, "Paswword information must be the same")
			}

			// for key, value := range r.Form {
			// 	fmt.Printf("%s = %s\n", key, value)
			// }
}
