# Testing API Golang with Postman
## Struktur Project
```
catalog_week_6
│
├── main.go
├── go.mod
├── .env
├── .gitignore
│
├── config
│   ├── database.go
│   └── firebase.go
│
├── models
│   ├── user.go
│   └── product.go
│
├── middleware
│   └── auth_middleware.go
│
├── handlers
│   ├── auth_handler.go
│   └── product_handler.go
│
├── services
│   ├── auth_service.go
│   └── product_service.go
│
├── repositories
│   ├── user_repository.go
│   └── product_repository.go
│
└── routes
    └── router.go
```


### Penjelasan
| Folder       | Fungsi                                         |
| ------------ | ---------------------------------------------- |
| config       | konfigurasi database dan firebase              |
| models       | struktur data yang terhubung ke tabel database |
| repositories | layer akses database                           |
| services     | business logic aplikasi                        |
| handlers     | menerima HTTP request                          |
| middleware   | autentikasi JWT                                |
| routes       | definisi endpoint API                          |
| images       | image untuk di readme                          |

---

## API Endpoint
Berikut Endpoint yang tersedia pada sistem.
### Health Check
```
GET 
/v1/health
```
<img width="1388" height="880" alt="image" src="https://github.com/user-attachments/assets/eb0c5ee1-ca36-4b4b-aa60-c931aa127386" />
<img width="952" height="967" alt="image" src="https://github.com/user-attachments/assets/f27f0809-a74f-4b0d-976c-b6ebf40aecff" />

---

### Authentication
#### Login
```
POST 
https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key={{FIREBASE_API_KEY}}
```
<img width="1377" height="882" alt="image" src="https://github.com/user-attachments/assets/d490039a-a9cf-4a7a-ab03-32f59cbe0797" />

#### Verify
```
POST 
/v1/auth/verify-token
```
<img width="1386" height="878" alt="image" src="https://github.com/user-attachments/assets/56fd79ed-1f93-470e-8fba-af0bb8344413" />


---
---

### Products
```
GET 
/v1/products
```
<img width="1383" height="875" alt="image" src="https://github.com/user-attachments/assets/ec98dffb-ef2f-42f5-9ab0-50cd8ae58a1b" />

```
GET /v1/products/:id
```
<img width="1392" height="880" alt="image" src="https://github.com/user-attachments/assets/65c7425c-ffa9-4003-ae2f-52735cae1e2f" />

---


## Code
```
sequenceDiagram
participant Client
participant Middleware
participant Handler
participant Service
participant Repository
participant DB
Client->>Middleware: HTTP Request
Middleware->>Handler: forward request
Handler->>Service: CreateUser()
Service->>Repository: SaveUser()
Repository->>DB: INSERT USER
DB-->>Repository: OK
Repository-->>Service: user saved
Service-->>Handler: user response
Handler-->>Client: JSON Response
```

---


## Code
```
sequenceDiagram
participant Client
participant Middleware
participant Handler
participant Service
participant Repository
participant DB

Client->>Middleware: HTTP Request (GET /products)
Middleware->>Handler: forward request
Handler->>Service: GetAllProducts()
Service->>Repository: FindAll()
Repository->>DB: SELECT * FROM products
DB-->>Repository: product list
Repository-->>Service: products data
Service-->>Handler: products response
Handler-->>Client: JSON Response
```
