package router

import (
	"github.com/dnagikh/website/pkg"
	"github.com/dnagikh/website/pkg/dkim"
	"github.com/dnagikh/website/pkg/dns"
	"github.com/dnagikh/website/pkg/domain_ip"
	"github.com/dnagikh/website/pkg/my_ip"
	"github.com/dnagikh/website/pkg/ports"
	"net/http"
	"path/filepath"
)

// NewRouter provides http handlers for website.
func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	files := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})
	mux.Handle("/assets/", http.StripPrefix("/assets", files))

	mux.HandleFunc("/", pkg.Index)
	mux.HandleFunc("/my_ip", my_ip.Index)
	mux.HandleFunc("/dkim", dkim.Index)
	mux.HandleFunc("/dkim/result", dkim.Result)
	mux.HandleFunc("/dns", dns.Index)
	mux.HandleFunc("/domain_ip", domain_ip.Index)
	mux.HandleFunc("/ports", ports.Index)

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
