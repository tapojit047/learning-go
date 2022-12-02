# API server using GO

### RESTful API using [go](https://github.com/golang), [cobra CLI](https://github.com/spf13/cobra), [go-chi/chi](https://github.com/go-chi/chi), Basic Auth, [JWT Auth](https://github.com/dgrijalva/jwt-go)

--- 
## API Endpoints
| Endpoint            | Function | Method | StatusCode | Auth  |
|---------------------|---| ------ | ---------- |-------|
| `/api/login`        | Login | POST | Success - 200, Failure - 401 | Basic |
| `/api/books`        | GetBooks | GET | Success - 200, Failure - 401 | Basic |
| `/api/books`        | AddBook | POST | Success - 200, Failure - 401, 409 | JWT   |
| `/api/books/{id}`   | UpdateBook | PUT | Success - 200, Failure - 401, 404 | JWT   |
| `/api/books/{id}`   | DeleteBook | DELETE | Success - 200, Failure - 401, 404 | JWT   |
| `/api/books/{id}`   | GetBookById  | GET | Success - 200, Failure - 401 | JWT   |

---
## Data Model
```
package model

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

```
```
package model

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastName"`
}

```
```
package model

type Credential struct {
	Username string `json:"username"'`
	Password string `json:"password"`
}
```
---
## Installation
* `go install github.com/tapojit047/learning-go/api-server/`
---
## Authentication Method
* Basic Authentication
* JWT Authentication
---
Resources:
* [sysdevbd learn GO](https://github.com/sysdevbd/sysdevbd.github.io/tree/master/go)
* [A Beginner’s Guide to HTTP and REST](https://code.tutsplus.com/tutorials/a-beginners-guide-to-http-and-rest--net-16340)
* [HTTP Response Status Codes](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)