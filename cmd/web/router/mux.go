package router

import (
	"github.com/dnagikh/website/pkg"
	"github.com/dnagikh/website/pkg/dig"
	"github.com/dnagikh/website/pkg/dkim"
	"github.com/dnagikh/website/pkg/myip"
	"net/http"
	"path/filepath"
)

// NewRouter provides http handlers for website.
func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	files := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})
	mux.Handle("/assets/", http.StripPrefix("/assets", files))

	mux.HandleFunc("/", pkg.Index)
	mux.HandleFunc("/my_ip", myip.Index)

	mux.HandleFunc("/dkim", dkim.Index)
	mux.HandleFunc("/dkim/result", dkim.Result)

	mux.HandleFunc("/dig", dig.Index)

	return mux
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
