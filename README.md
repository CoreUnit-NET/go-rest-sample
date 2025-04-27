# 🔥 go-rest-sample — REST API Boilerplate Using Fuego

## 📌 Purpose

This repo is a **Go boilerplate** using the **Fuego** web framework to build REST APIs.

It's a **template project** — the resource types are *illustrative* and not tied to actual infrastructure.
The example implementation is for managing resources such as **servers**, **containers**, and **services**.

---

## 🧱 Tech Stack

- **Language**: Go
- **Framework**: [Fuego](https://github.com/alexflint/fuego)
- **Database**: JSON file or SQLite
- **Auth**: None (initially); Later: OIDC

---

## 📘 Why Fuego?

- 🧪 **OpenAPI generation** from code automatically
- ✅ **Validation**: Fast + go-playground/validator-based
- 🔧 **Built on `net/http`**: No vendor lock-in
- 🔄 **Serialization/Deserialization**: Auto JSON/XML/HTML responses

---

## 🧩 Example Resources

Each resource (server, container, service) follows a **RESTful schema** with 5 standard endpoints:

### Endpoints for each resource

| Method | Path                | Description        |
|--------|---------------------|--------------------|
| GET    | `/resource/:id`     | Retrieve one       |
| GET    | `/resource`         | Retrieve all       |
| POST   | `/resource`         | Create new         |
| PATCH  | `/resource/:id`     | Update existing    |
| DELETE | `/resource/:id`     | Delete existing    |

Replace `/resource` with:
- `/servers`
- `/containers`
- `/services`

---

## 🧱 Sample Structs

### `Server`

```go
type Server struct {
  ID       string `json:"id" validate:"required,uuid4"`
  Name     string `json:"name" validate:"required"`
  OS       string `json:"os" validate:"required"`
  Hostname string `json:"hostname" validate:"required,hostname_rfc1123"`
  SSHPort  int    `json:"sshPort" validate:"required,min=1,max=65535"`
  Owner    string `json:"owner" validate:"required"`
}
```

### `Container`/`Service`
```go
type Container struct {
  ID          string         `json:"id" validate:"required,uuid4"`
  ServerID    string         `json:"serverId" validate:"required,uuid4"`
  Name        string         `json:"name" validate:"required"`
  Description string         `json:"description"`
  Ports       map[string]int `json:"ports"`
}
```

## 🗃️ Persistence
Abstracted via interface:

```go
type ResourceStore[T any] interface {
  Get(id string) (*T, error)
  GetAll() ([]T, error)
  Create(obj *T) error
  Update(id string, update *T) error
  Delete(id string) error
}
```
Initial implementation can use:
- JSON file on fs
- SQLite

## 🔐 Future Auth: OIDC via CoreUnit.NET SSO
Planned login flow:
- Use Cunet SSO as OIDC provider
- Add middleware to Fuego routes for token validation
- Use github.com/coreos/go-oidc or similar

## 🧪 Future: React + TypeScript Dashboard
React frontend (Vite-based) with:
- Visual management for servers, containers & services
- REST API communication via generated code from a openapi v3 code generator
- UI built using tailwind
- Clean layout with create/edit/delete support

