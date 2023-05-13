package cache

import (
	"net/http"
	"strings"
)

type Net struct {
	addr   string
	prefix string
}

func (n *Net) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	if of := strings.HasPrefix(path, n.prefix); !of {
		panic("unexpected path" + path)
	}

	separate := strings.SplitN(path[len(n.prefix):], "/", 100)
	if len(separate) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	groupName := separate[0]
	key := separate[1]
	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, "no such group", http.StatusNotFound)
	}
	v, err := group.Get(key)
	if err != nil {
		http.Error(w, "fail to get element", http.StatusInternalServerError)
	}

	//w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Type", "octet-stream")
	w.Write(v.Copy())
}

func NewNetService(prefix string) *Net {
	return &Net{
		prefix: prefix,
	}
}
