# USB-NAS-UPDATE-SERVER

A minimal Go HTTP server for serving files from a `public/` directory and providing directory browsing. Ideal for hosting update files, version manifests, or any other static assets on a local network.

## Features

- **Simple Setup**: Just place your files in the `public/` folder; the server handles the rest.
- **Automatic Directory Browsing**: Lists files and subdirectories when a directory path is requested.
- **File Serving**: Serves individual files directly when a file path is requested.
- **Configurable Port**: Change the port via a command-line flag (`-port`).
- **Lightweight & Fast**: Uses Go’s built-in HTTP server, requiring no dependencies.

## Installation & Usage

### 1. Clone or Download  
   ```bash
   git clone https://github.com/MattInnovates/USB-NAS-UPDATE-SERVER.git
   cd USB-NAS-UPDATE-SERVER
   ```

### 2. Build  
   ```bash
   go build -o usb-nas-update-server.exe server.go
   ```
   This produces an executable named `usb-nas-update-server`.

### 3. Create a `public` Directory  
   Ensure you have a folder named `public` in the same directory as your executable. Place all the files you want to serve (for example, `.exe` files, `latestVersion.txt`, etc.) inside `public`.

### 4. Run  
   ```bash
   ./usb-nas-update-server -port 8080
   ```
   By default, it listens on port 8080, but you can specify another port if needed:
   ```bash
   ./usb-nas-update-server -port 9090
   ```

### 5. Access Your Files  
   Open a browser or use a tool like `curl` to navigate to:
   ```
   http://127.0.0.1:8080/
   ```
   You should see a simple directory listing of the `public/` folder. Subfolders and files are listed, and clicking on any file downloads or displays it.

---

## Directory Browsing

When you open a directory path in your browser, the server automatically generates a basic HTML page listing its contents. This includes:

- A link to the parent directory (`..`), unless you’re at the root (`/`).
- A link for each file or subdirectory in the current directory.

---

## Example Folder Structure

```
USB-NAS-UPDATE-SERVER/
├─ usb-nas-update-server (executable)
├─ public/
│  ├─ latestVersion.txt
│  ├─ usb-nas-cli-v1.0.0.exe
│  ├─ usb-nas-cli-v1.0.1.exe
│  └─ ...
└─ server.go
```

- **`server.go`**: The source code for this HTTP server.
- **`public/`**: All files you want to serve (update binaries, text files, etc.).

---

## Customization

- **Port Configuration**: Adjust the default port by modifying the `flag.Int("port", 8080, ...)` line in `server.go` or by passing a different port number at runtime.
- **Public Directory**: Currently set to `./public`. You can change the directory name in `handler()` if you want a different folder structure.
- **Security**: This code is meant for simple local or LAN-based file serving. For production use, consider adding TLS/HTTPS, authentication, or other security measures.

---

## Contributing

Contributions are welcome. If you want to contribute:

1. **Report Issues**: Open an issue if you encounter bugs, need enhancements, or have feature suggestions.
2. **Improve Documentation**: If you find missing or unclear documentation, consider submitting improvements.
3. **Submit a Pull Request**:
   - Fork the repo and clone it to your local machine.
   - Create a new branch for your changes (`git checkout -b fix/bug-name` or `feature/new-feature`).
   - Make your changes and commit them (`git commit -m "Fix bug: description"`).
   - Push your changes to your fork and submit a pull request.

Please ensure your contributions align with the project’s goals and are well-tested before submission.

---

## License

*(Add your preferred license here, such as MIT, Apache 2.0, etc.)*

---

## Acknowledgments

- **Go** for providing a simple and efficient HTTP server in the standard library.

