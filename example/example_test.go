package example

import "testing"
import "net/http"
import "fmt"
import "github.com/hydrogen18/memlistener"
import "io"
import "bytes"
import "os"

func helloHttp(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(os.Stdout, "Request from: %v\n", req.RemoteAddr)
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, "Hello http\n")
}

func TestMemoryServer(t *testing.T) {
	//Start the HTTP server using an in memory Listener
	server := memlistener.NewInMemoryServer(http.HandlerFunc(helloHttp))

	//Make an HTTP request. Any domain works
	response, err := server.NewClient().Get("http://test.server/bar")
	if err != nil {
		t.Fatal(err)
	}

	//Validate the results
	if response.StatusCode != http.StatusOK {
		t.Fatalf("Status code is %q", response.Status)
	}

	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, response.Body)
	if err != nil {
		t.Fatal(err)
	}

	if buf.String() != "Hello http\n" {
		t.Fatalf("Output is %q", buf.String())
	}

	//Close the in memory listener, stopping the server
	server.Listener.Close()

}
