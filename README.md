# Arlo 🚀

> Transform your Go backend into a beautiful web application with modern web frontends

## Why Arlo?

- **Go-powered backend** for system operations, performance, and reliability
- **Modern web frontend** using your favorite JavaScript/TypeScript framework
- **Hot-reload development** for instant feedback during development
- **Single executable** distribution for easy deployment

## Requirements

Before getting started, ensure you have these tools installed:

### Core Dependencies
- **Go** (version 1.21 or higher)
- **Node.js** (version 14 or higher)
- **Air** - Go hot-reloading tool

### Package Manager (choose one)
- npm
- yarn
- pnpm
- bun
- deno

## Installation

Install Arlo using Go's package manager:

```bash
go install github.com/JasnRathore/arlo@latest
```

## Getting Started

### Creating Your First Project

Initialize a new Arlo project with the interactive setup:

```bash
arlo init
```

This command will guide you through:
1. **Project naming** - Choose a descriptive name for your application
2. **Package manager selection** - Pick your preferred JavaScript package manager
3. **Dependency verification** - Automatically check that all required tools are installed
4. **Project scaffolding** - Generate the complete project structure
5. **Dependency installation** - Set up all necessary packages

### Project Structure

After initialization, your project will be organized as follows:

```
your-project/
├── src/                    # Frontend source code (Vite project)
├── src-backend/           # Go backend code
│   ├── app/               # Application logic
│   │   └── app.go         # HTTP handlers and business logic
│   ├── main.go           # Development entry point
│   ├── build.go          # Production build configuration
│   └── .air.toml         # Hot-reload configuration
└── arlo.config.json      # Project configuration
```

## Development Workflow

### Running in Development Mode

Start your development environment:

```bash
arlo dev
```

This single command orchestrates your entire development setup:
- Launches your frontend development server
- Starts the Go backend with automatic reloading
- Establishes proxy connections between frontend and backend
- Watches for file changes and reloads automatically

Your application will be available at `http://localhost:port` with the frontend proxied through Vite's development server.

### Building Go APIs

Create HTTP endpoints in your `app/app.go` file:

```go
func App() http.Handler {
    mux := http.NewServeMux()
    
    // Simple API endpoint
    mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        users := []string{"Alice", "Bob", "Charlie"}
        json.NewEncoder(w).Encode(users)
    })
    
    // Handle different HTTP methods
    mux.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            // Handle GET request
        case "POST":
            // Handle POST request
        }
    })
    
    return mux
}
```

### Frontend Integration

Call your Go APIs from the frontend using standard HTTP requests:

```javascript
// Fetch data from your Go backend
async function fetchUsers() {
    const response = await fetch('/api/users');
    const users = await response.json();
    return users;
}

// Post data to your Go backend
async function saveUser(userData) {
    const response = await fetch('/api/users', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(userData)
    });
    return response.json();
}
```

### TypeScript Support

Arlo includes TypeScript utilities for type-safe development:

```typescript
// Define your API response types
interface User {
    id: number;
    name: string;
    email: string;
}

// Type-safe API calls
async function getUsers(): Promise<User[]> {
    const response = await fetch('/api/users');
    return response.json();
}
```

## Production Builds

### Creating Distribution Builds

When ready to distribute your application:

```bash
arlo build
```

This command performs several optimizations:
1. **Frontend optimization** - Builds your frontend for production with minification and bundling
2. **Asset embedding** - Embeds frontend assets directly into the Go binary
3. **Binary compilation** - Creates a single executable file
4. **Output organization** - Places the final executable in `src-backend/target/`

The resulting binary is completely self-contained and can be distributed without any dependencies.

### Deployment

Your built application is a single executable file that includes:
- All frontend assets (HTML, CSS, JavaScript)
- Go backend compiled for the target platform
- Static file serving capabilities
- API routing and handlers

Simply copy the executable to any compatible system and run it - no installation required.

## Configuration

### Application Settings

Customize your application in the `app.go` file:

```go
func setupServer() {
    port := findAvailablePort(8080)
    
    mux := http.NewServeMux()
    mux.Handle("/api/", http.StripPrefix("/api", App()))
    
    log.Printf("Server running on port %d", port)
    http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
```

### Environment Variables

Configure your frontend to communicate with the backend using environment variables:

```bash
# .env file in your frontend
VITE_API_URL=http://localhost:8080/api
```

Arlo automatically manages these configurations during development and production builds.

## Examples

### File Processing Application

```go
// Backend: Handle file uploads and processing
mux.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
    file, header, err := r.FormFile("document")
    if err != nil {
        http.Error(w, "Failed to read file", http.StatusBadRequest)
        return
    }
    defer file.Close()
    
    // Process the file with Go's powerful standard library
    result := processDocument(file)
    
    json.NewEncoder(w).Encode(map[string]interface{}{
        "filename": header.Filename,
        "result":   result,
    })
})
```

```javascript
// Frontend: Upload and display results
function uploadFile(file) {
    const formData = new FormData();
    formData.append('document', file);
    
    return fetch('/api/process', {
        method: 'POST',
        body: formData
    }).then(response => response.json());
}
```

### System Information Dashboard

```go
// Backend: System monitoring endpoints
mux.HandleFunc("/system/stats", func(w http.ResponseWriter, r *http.Request) {
    stats := map[string]interface{}{
        "memory":    getMemoryUsage(),
        "cpu":       getCPUUsage(),
        "processes": getRunningProcesses(),
    }
    json.NewEncoder(w).Encode(stats)
})
```

```javascript
// Frontend: Real-time dashboard updates
setInterval(async () => {
    const stats = await fetch('/api/system/stats').then(r => r.json());
    updateDashboard(stats);
}, 1000);
```

## Advanced Features

### Custom Build Configuration

Modify the build process in `build.go`:

```go
//go:embed dist/*
var content embed.FS

func serveStatic(w http.ResponseWriter, r *http.Request) {
    // Custom static file serving logic
    file := r.URL.Path
    if file == "/" {
        file = "/index.html"
    }
    
    data, err := content.ReadFile("dist" + file)
    if err != nil {
        http.NotFound(w, r)
        return
    }
    
    w.Header().Set("Content-Type", detectContentType(file))
    w.Write(data)
}
```

### Hot Reload Configuration

Customize hot reloading in `.air.toml`:

```toml
[build]
  cmd = "go build -o ./tmp/main.exe ."
  bin = "tmp\\main.exe"
  exclude_dir = ["assets", "tmp", "vendor", "dist"]
  include_ext = ["go", "html", "css", "js"]
```

## Architecture Benefits

Arlo's architecture provides several key advantages:

**Performance**: Go's compiled nature and efficient runtime provide excellent performance for system operations, file processing, and concurrent tasks.

**Developer Experience**: Modern JavaScript tooling (Vite, hot-reload, TypeScript) combined with Go's simple deployment model creates an optimal development workflow.

**Maintainability**: Clear separation between frontend and backend concerns, with well-defined API boundaries, makes applications easier to maintain and extend.

## Conclusion


Ready to build your next web application? Start with `arlo init` and experience the power of Go-powered desktop development.

**Get involved**: Found a bug or have a feature request? Contributions are welcome! The simplicity of Arlo's architecture makes it easy to understand and extend.