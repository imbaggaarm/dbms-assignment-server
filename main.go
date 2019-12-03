package main
import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"src/controllers"
)
func main() {
	r := mux.NewRouter().StrictSlash(true)

	// Configure api route
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", HandleIndex)

	// Configure auth route
	auth := api.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register", controllers.CreateAccount).Methods(http.MethodPost)
	auth.HandleFunc("/login", controllers.Authenticate).Methods(http.MethodPost)

	// Configure user
	api.HandleFunc("/{user_id}", controllers.GetUserProfile).Methods(http.MethodGet)
	api.HandleFunc("/update_profile/", controllers.UpdateProfile).Methods(http.MethodPost)

	// Configure courses
	courses := api.PathPrefix("/courses").Subrouter()
	courses.HandleFunc("/", controllers.GetCourses).Methods(http.MethodGet)
	courses.HandleFunc("/{institution_id}", controllers.GetInstitutionCourses).Methods(http.MethodGet)
	courses.HandleFunc("/{user_id}", controllers.GetUserCourses).Methods(http.MethodGet)
	courses.HandleFunc("/{courses_id}", controllers.GetCourseDetail).Methods(http.MethodGet)
	courses.HandleFunc("/{courses_id}/comments", controllers.GetCourseComments).Methods(http.MethodGet)
	// Starting server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

 	log.Println("Starting server at :" + port)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
