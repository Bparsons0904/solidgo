# SolidGo CLI Tool

SolidGo is a command-line interface tool designed to simplify managing SolidGo projects. It provides functionalities for creating and managing routes in your SolidGo application.

## Installation

To install SolidGo, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/your-repo/solidgo.git
   ```

2. Navigate to the cloned directory:

   ```bash
   cd solidgo
   ```

3. Build the application:

   ```bash
   go build -o solidgo
   ```

4. Install the application:

   ```bash
   go install
   ```

5. (Optional) Move the binary to a directory in your PATH for global access:
   ```bash
   sudo mv solidgo /usr/local/bin/
   ```

## Usage

The SolidGo CLI provides commands for managing routes in your application.

### Adding a New Route

To add a new route:

```bash
solidgo route add <route-name> -m <HTTP-method> -f <function-name> -p <path>
```

- `<route-name>`: Name of the route.
- `<HTTP-method>`: HTTP method for the route (GET, POST, PUT, PATCH, DELETE).
- `<function-name>`: Function name to handle the route.
- `<path>`: API path name

Example:

```bash
solidgo route add user -m POST -f CreateUserLogin -p :userId/login
```

This command adds a new POST route to the user routes with the handler function `CreateUserLogin`.

### Creating a New Route File

To create a new route file with a group and initial route:

```bash
solidgo route new <group-name> -m <HTTP-method> -f <function-name> -p <path>
```

- `<group-name>`: Name of the route group.
- `-m`, `-f`, `-p`: Same as above.

Example:

```bash
solidgo route new user -m GET -f GetUser
```

This command creates a new route file for user routes with a default GET route handled by `GetUser`.
