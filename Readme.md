# FastMux - A Simple HTTP Router for Go

**FastMux** is a lightweight HTTP router built with Go's `net/http` package.  
The name "FastMux" is a playful nod to its simplicity and ease of use, but its primary purpose is **educational** rather than achieving maximum performance.

This library was created to deepen understanding of Go's `net/http` package and to explore the principles of library design in Go.

---

## Purpose

FastMux is not intended to be the fastest or most performant HTTP router available. Instead, its main goals are:

- **Learning**: To serve as a practical example for studying how to use Go's `net/http` package effectively.
- **Library Design**: To demonstrate clean and modular design principles for building reusable Go libraries.
- **Simplicity**: To provide a minimal, easy-to-understand router implementation that supports basic routing functionality.

This makes FastMux an excellent tool for developers who want to learn more about Go's HTTP handling, routing mechanisms, and how to structure a Go library.

---

## Features

- Supports basic HTTP methods (`GET`, `POST`, etc.)
- Simple and intuitive API for defining routes
- Matches routes based on HTTP method and path
- Lightweight with minimal dependencies (only uses Go’s standard library)
- Easy to extend for custom routing logic

---

## Installation

```bash
go get github.com/hermangoncalves/fastmux
````

---

## Usage

Below is an example of how to use FastMux to create a simple HTTP server with a few routes.

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/hermangoncalves/fastmux"
)

func main() {
	port := "8007"
	mux := fastmux.New()

	// Register a route with a dynamic parameter :name
	mux.GET("/hello/:name", func(w http.ResponseWriter, r *http.Request, params fastmux.Params) {
		name := params.ByName("name")
		fmt.Fprintf(w, "Hello, %s!", name)
	})

	fmt.Println("Server starting on :" + port + "...")
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}

```

---

## Explanation of the Code

* **Creating a Router**: `fastmux.New()` initializes a new FastMux instance.
* **Registering Routes**: Use `mux.GET()` or `mux.POST()` to define routes with their respective handlers. You can also use `mux.Handle(method, path, handler)` for other HTTP methods.
* **Starting the Server**: Pass the FastMux instance to `http.ListenAndServe` to start the server.

## Limitations

Since FastMux is for educational purposes, it has some limitations:
* ❌ Not Production-Optimized: For production use, prefer robust routers like `chi` or `gin`.

---

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository
2. Create a new branch

   ```bash
   git checkout -b feature/your-feature
   ```
3. Make your changes and commit

   ```bash
   git commit -m "Add your feature"
   ```
4. Push to your branch

   ```bash
   git push origin feature/your-feature
   ```
5. Open a pull request

Please ensure your code follows Go best practices and includes appropriate tests.

---

## Acknowledgments

FastMux was inspired by the desire to learn more about Go's `net/http` package and to create a teaching tool for others interested in Go library design.

Thanks to the Go community for providing excellent resources and inspiration.
