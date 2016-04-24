package example

import (
	"io"
	"net/http"

	m "github.com/otiai10/marmoset"
)

func init() {

	m.LoadViews("./views")

	r := m.NewRouter()

	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		m.Render(w).HTML("index", nil)
	})

	r.GET("/users/(?P<name>[a-zA-Z0-9]+)/hello", func(w http.ResponseWriter, r *http.Request) {
		m.Render(w).HTML("hello", map[string]interface{}{"name": r.FormValue("name")})
	})

	r.POST("/upload", func(w http.ResponseWriter, r *http.Request) {
		f, h, err := r.FormFile("upload")
		if err != nil {
			m.RenderJSON(w, http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}
		w.Header().Set("Content-Type", h.Header.Get("Content-Type"))
		io.Copy(w, f)
	})

	r.GET("/api", func(w http.ResponseWriter, r *http.Request) {
		m.Render(w).JSON(http.StatusOK, map[string]interface{}{
			"message":     "Hello, this is pygmy marmoset API!",
			"remote_addr": r.RemoteAddr,
		})
	})
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		m.Render(w).HTML("notfound", nil)
	})

	r.StaticRelative("/public", "./assets")

	http.Handle("/", r)
	// ListenAndServe will be invoked by GAE SDK.
	// http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
