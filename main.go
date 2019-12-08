package main
import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"src/app"
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
	auth.HandleFunc("/delete", controllers.DeleteAccount).Methods(http.MethodPost)
	auth.HandleFunc("/login", controllers.Authenticate).Methods(http.MethodPost)
	auth.HandleFunc("/logout", controllers.Logout).Methods(http.MethodGet)
	// Configure user
	api.HandleFunc("/{user_id}", controllers.GetUserProfile).Methods(http.MethodGet)
	api.HandleFunc("/update_profile", controllers.UpdateProfile).Methods(http.MethodPost)
	api.HandleFunc("/{user_id}/courses", controllers.GetUserCourses).Methods(http.MethodGet)

	// Configure courses
	courses := api.PathPrefix("/courses").Subrouter()
	courses.HandleFunc("/list", controllers.GetAllCourses).Methods(http.MethodGet)
	courses.HandleFunc("/institution/{institution_id}", controllers.GetInstitutionCourses).Methods(http.MethodGet)
	courses.HandleFunc("/{course_id}", controllers.GetCourse).Methods(http.MethodGet)
	courses.HandleFunc("/{course_id}/comments", controllers.GetCourseComments).Methods(http.MethodGet)

	r.Use(app.JwtAuthentication) // attach JWT auth middleware

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
