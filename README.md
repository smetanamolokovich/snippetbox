# snippetbox

Snippetbox lets people paste
and share snippets of text — a bit like Pastebin or GitHub’s Gists.

![Screenshot](https://lets-go.alexedwards.net/sample/assets/img/01.00-01.png)

## App routes:

| Method |   URL Pattern   |      Handler      | Action                          |
| ------ | :-------------: | :---------------: | ------------------------------- |
| GET    |        /        |       home        | Display the home page           |
| GET    |  /snippet?id=1  |    showSnippet    | Display a specific snippet      |
| GET    | /snippet/create | createSnippetForm | Display the new snippet form    |
| POST   | /snippet/create |   createSnippet   | Create a new snippet            |
| GET    |  /user/signup   |  signupUserForm   | Display the user signup form    |
| POST   |  /user/signup   |    signupUser     | Create a new user               |
| GET    |   /user/login   |   loginUserForm   | Display the user login form     |
| POST   |   /user/login   |     loginUser     | Authenticate and login the user |
| POST   |  /user/logout   |    logoutUser     | Logout the user                 |
| GET    |    /static/     |  http.FileServer  | Serve a specific static file    |

## App structure:

```bash
├── cmd
│   └── web
│       ├── handlers.go
│       ├── handlers_test.go
│       ├── helpers.go
│       ├── main.go
│       ├── middleware.go
│       ├── middleware_test.go
│       ├── routes.go
│       ├── templates.go
│       ├── templates_test.go
│       └── testutils_test.go
├── go.mod
├── go.sum
├── pkg
│   ├── forms
│   │   ├── errors.go
│   │   └── form.go
│   └── models
│       ├── mock
│       │   ├── snippets.go
│       │   └── users.go
│       ├── models.go
│       └── mysql
│           ├── snippets.go
│           ├── testdata
│           │   ├── setup.sql
│           │   └── teardown.sql
│           ├── testutils_test.go
│           ├── users.go
│           └── users_test.go
├── README.md
├── tls
│   ├── cert.pem
│   └── key.pem
└── ui
    ├── html
    │   ├── about.page.tmpl
    │   ├── base.layout.tmpl
    │   ├── create.page.tmpl
    │   ├── footer.partial.tmpl
    │   ├── home.page.tmpl
    │   ├── login.page.tmpl
    │   ├── password.page.tmpl
    │   ├── profile.page.tmpl
    │   ├── show.page.tmpl
    │   └── signup.page.tmpl
    └── static
        ├── css
        │   └── main.css
        ├── img
        │   ├── favicon.ico
        │   └── logo.png
        └── js
            └── main.js

15 directories, 40 files
```

## Development

```bash
$ go mod download
```

```bash
$ go run ./cmd/web/
```
