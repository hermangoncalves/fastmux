# FastMux - A Simple HTTP Router for Go

**FastMux** is a lightweight HTTP router built with Go's `net/http` package.  

This library was created to deepen understanding of Go's `net/http` package and to explore the principles of library design in Go.

---

## Purpose

FastMux is not intended to be the fastest or most performant HTTP router available. Instead, its main goals are:

- **Learning**: To serve as a practical example for studying how to use Go's `net/http` package effectively.
- **Library Design**: To demonstrate clean and modular design principles for building reusable Go libraries.
- **Simplicity**: To provide a minimal, easy-to-understand router implementation that supports basic routing functionality.

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
	addr := ":8008"
	r := fastmux.New()

	r.GET("/hello/:name", func(ctx *fastmux.Context) {
		name, _ := ctx.Param("name")
		ctx.JSON(http.StatusOK, fastmux.H{
			"message": name,
		})
	})

	log.Fatal(r.Run(addr))
}

```



## Limitations

Since FastMux is for study purposes, it has some limitations:
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

## Acknowledgments

FastMux was inspired by the desire to learn more about Go's `net/http` package.

Thanks to the Go community for providing excellent resources and inspiration.
