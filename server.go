package todo 

import(
    "time"
    "net/http"
    "context"
)

type HTTPServer struct {
    server *http.Server
}

func (this *HTTPServer) Run(port string, handler http.Handler) error {
    this.server = &http.Server {
	Addr: 		":" + port,
	Handler: 	handler,
	MaxHeaderBytes: 1 << 20, // 1 mb
	ReadTimeout: 	10 * time.Second, 
	WriteTimeout: 	10 * time.Second,
    }

    return this.server.ListenAndServe()
}

func (this *HTTPServer) Stop(ctx context.Context) error {
    return this.server.Shutdown(ctx)
}











