cookbook/
├── cmd/
│   └── main.go              # Entry point for the application
├── config/
│   └── config.go            # Configuration handling
├── internal/
│   ├── handlers/            # HTTP handlers for routes
│   │   ├── ocr.go
│   │   ├── recipes.go
│   │   └── google_docs.go
│   ├── services/            # Business logic
│   │   ├── ocr_service.go
│   │   ├── recipe_service.go
│   │   └── google_docs_service.go
│   ├── repositories/        # Database interactions
│   │   ├── recipe_repo.go
│   │   └── user_repo.go
│   ├── models/              # Data models
│   │   ├── recipe.go
│   │   └── user.go
│   ├── middlewares/         # Middleware for request handling
│   │   └── auth_middleware.go
│   └── utils/               # Utility functions
│       ├── http_response.go
│       └── image_processing.go
├── pkg/
│   └── external/            # External integrations
│       ├── ocr_client.go
│       ├── google_auth.go
│       └── google_docs_client.go
├── go.mod                   # Dependency management
├── go.sum
└── README.md
