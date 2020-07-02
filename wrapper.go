package fasthttp2nethttp

import (
	"io/ioutil"
	"net"
	"net/http"

	"github.com/valyala/fasthttp"
)

func FastHTTPHandlerWrapper(h fasthttp.RequestHandler) http.Handler {
	return FastHTTPHandlerWrapperFunc(h)
}

func FastHTTPHandlerWrapperFunc(h fasthttp.RequestHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// New fasthttp request
		var req fasthttp.Request
		// Convert net/http -> fasthttp request
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}
		req.Header.SetMethod(r.Method)
		req.SetRequestURI(r.RequestURI)
		req.Header.SetContentLength(len(body))
		req.SetHost(r.Host)
		for key, val := range r.Header {
			for _, v := range val {
				req.Header.Add(key, v)
			}
		}
		req.BodyWriter().Write(body) //nolint
		remoteAddr, err := net.ResolveTCPAddr("tcp", r.RemoteAddr)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// New fasthttp Ctx
		var fctx fasthttp.RequestCtx
		fctx.Init(&req, remoteAddr, nil)
		// New fiber Ctx

		// Execute fiber Ctx
		h(&fctx)

		// Convert fasthttp Ctx > net/http
		fctx.Response.Header.VisitAll(func(k, v []byte) {
			sk := string(k)
			sv := string(v)
			w.Header().Set(sk, sv)
		})
		w.WriteHeader(fctx.Response.StatusCode())
		w.Write(fctx.Response.Body()) //nolint
	})
}
