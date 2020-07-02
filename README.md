# Fasthttp handlers wrapper

You can adapt fasthttp handlers for serving with default net/http handlers or rewrap them to over frameworks

Based on [https://github.com/gofiber/adaptor/](https://github.com/gofiber/adaptor/).

## Usage 


```go
import (
 "github.com/mr-tron/fasthttp2nethttp"
)
func MyOldLegacyFasthttpHandler(ctx *fasthttp.RequestCtx) {
		body, err := json.Marshal(struct{Status string}{"Ok"})
    	if err != nil {
    		return err
    	}
    	ctx.SetContentType("application/json")
    	ctx.SetStatusCode(fasthttp.StatusOK)
    	ctx.SetBody(body)
}




func main() {
handler := fasthttp2nethttp.FastHTTPHandlerWrapper(MyOldLegacyFasthttpHandler)
// now you can serve handler with any framework like default net/http or wrap for another framework like echo.

}








```