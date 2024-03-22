# Simple random redirect

## Description

This project is a simple web server written in Go that demonstrates dynamic URL redirection. It uses a JSON configuration file to determine the server's listening port and the list of target URLs for redirection. The application supports hot reloading of the configuration file to update the list of target URLs without restarting the server.

## Features

- Dynamic redirection to randomly chosen URLs from a predefined list.
- Configurable listening port via a JSON file.
- Hot reloading of configuration to update target URLs and port without server restart.
- Implementation of a basic intermediate redirect to manipulate HTTP headers.

## Getting Started

### Prerequisites

- Go 1.22 or later

### Installation

1. Clone the repository:
```sh
git clone https://github.com/barantaran/simple-redirect.git
```

2. Navigate to the project directory:
```sh
cd simple-redirect
```

3. Build the project:
```sh
go build
```

### Configuration

Edit the `config.json` file in the project's root directory to set up your desired port and target URLs:

```json
{
  "port": 8080,
  "targets": [
    "http://example1.com",
    "http://example2.com"
  ]
}
```

### Running the Application

Execute the compiled binary to start the server, e.g.:

```sh
./m.exe
```

The server will start and listen on the port specified in `config.json`. It will redirect incoming requests to one of the configured target URLs.

### Updating Configuration

To update the list of target URLs or the listening port, modify the `config.json` file and save your changes. The application will automatically reload the configuration without needing a restart.

## Usage

To test the redirection functionality, navigate to `http://localhost:port/` in your web browser, replacing `port` with the actual port number you configured. The server will redirect you to one of the target URLs.

Below is a template for a README file for your project, including a section on recommendations for running your Go application as a daemon. This template assumes your project involves the web application you've been developing, with dynamic configuration reloading and systemd for auto-restart. Adjust the content as needed to better fit the specifics of your project.

---

## Running as a Daemon (Service)

To ensure that the project runs continuously, even after a reboot or crash, it's recommended to run it as a daemon. Here are the steps to set it up as a systemd service on a Linux system:

1. **Create a systemd service file** at `/etc/systemd/system/simple-redirect.service`:

```ini
[Unit]
Description=Your Project Name
After=network.target

[Service]
User=username
Group=usergroup
WorkingDirectory=/path/to/your/project
ExecStart=/path/to/your/project/simple-redirect
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
```

Replace `username`, `usergroup`, `/path/to/your/project`, and `simple-redirect` with your actual user name, group, project path, and executable name.

2. **Reload systemd** to recognize your new service:

```bash
sudo systemctl daemon-reload
```

3. **Enable the service** to start on boot:

```bash
sudo systemctl enable simple-redirect.service
```

4. **Start the service**:

```bash
sudo systemctl start simple-redirect.service
```

### Recommendations for Daemon Operation

- **Logging**: Ensure your application logs important events and errors. Configure systemd to manage and rotate logs as necessary.
- **Monitoring**: Use tools like `systemctl status`, `journalctl`, or third-party monitoring solutions to keep an eye on your service's health.
- **Security**: Run your service as a non-root user with minimal privileges to reduce security risks.
- **Configuration Changes**: When updating `config.json`, simply restart the service for changes to take effect:

```bash
sudo systemctl restart simple-redirect.service
```
