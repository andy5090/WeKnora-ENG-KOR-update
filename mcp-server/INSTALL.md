# WeKnora MCP Server Installation and Usage Guide

## Quick Start

### 1. Install Dependencies
```bash
pip install -r requirements.txt
```

### 2. Set Environment Variables
```bash
# Linux/macOS
export WEKNORA_BASE_URL="http://localhost:8080/api/v1"
export WEKNORA_API_KEY="your_api_key_here"

# Windows PowerShell
$env:WEKNORA_BASE_URL="http://localhost:8080/api/v1"
$env:WEKNORA_API_KEY="your_api_key_here"

# Windows CMD
set WEKNORA_BASE_URL=http://localhost:8080/api/v1
set WEKNORA_API_KEY=your_api_key_here
```

### 3. Run the Server

There are multiple ways to run the server:

#### Method 1: Using the main entry point (Recommended)
```bash
python main.py
```

#### Method 2: Using the original startup script
```bash
python run_server.py
```

#### Method 3: Run the server module directly
```bash
python weknora_mcp_server.py
```

#### Method 4: Run as a Python module
```bash
python -m weknora_mcp_server
```

## Install as Python Package

### Development Mode Installation
```bash
pip install -e .
```

After installation, you can use the command line tools:
```bash
weknora-mcp-server
# or
weknora-server
```

### Production Mode Installation
```bash
pip install .
```

### Build Distribution Package
```bash
# Build source distribution and wheel
python setup.py sdist bdist_wheel

# Or use build tool
pip install build
python -m build
```

## Command Line Options

The main entry point `main.py` supports the following options:

```bash
python main.py --help                 # Show help information
python main.py --check-only           # Only check environment configuration
python main.py --verbose              # Enable verbose logging
python main.py --version              # Show version information
```

## Environment Check

Run the following command to check environment configuration:
```bash
python main.py --check-only
```

This will display:
- WeKnora API base URL configuration
- API key setting status
- Dependency package installation status

## Troubleshooting

### 1. Import Error
If you encounter `ImportError`, please ensure:
- All dependencies are installed: `pip install -r requirements.txt`
- Python version is compatible (recommended 3.10+)
- No filename conflicts

### 2. Connection Error
If you cannot connect to the WeKnora API:
- Check if `WEKNORA_BASE_URL` is correct
- Confirm that the WeKnora service is running
- Verify network connection

### 3. Authentication Error
If you encounter authentication issues:
- Check if `WEKNORA_API_KEY` is set
- Confirm that the API key is valid
- Verify permission settings

## Development Mode

### Project Structure
```
WeKnoraMCP/
├── __init__.py              # Package initialization file
├── main.py                  # Main entry point
├── run_server.py           # Original startup script
├── weknora_mcp_server.py   # MCP server implementation
├── requirements.txt        # Dependency list
├── setup.py               # Installation script
├── MANIFEST.in            # Include file manifest
├── LICENSE                # License
├── README.md              # Project description
└── INSTALL.md             # Installation guide
```

### Adding New Features
1. Add new API methods to the `WeKnoraClient` class
2. Register new tools in `handle_list_tools()`
3. Implement tool logic in `handle_call_tool()`
4. Update documentation and tests

### Testing
```bash
# Run basic tests
python test_imports.py

# Test environment configuration
python main.py --check-only

# Test server startup
python main.py --verbose
```

## Deployment

### Docker Deployment
Create a `Dockerfile`:
```dockerfile
FROM python:3.11-slim

WORKDIR /app
COPY requirements.txt .
RUN pip install -r requirements.txt

COPY . .
RUN pip install -e .

ENV WEKNORA_BASE_URL=http://localhost:8080/api/v1
EXPOSE 8000

CMD ["weknora-mcp-server"]
```

### System Service
Create a systemd service file `/etc/systemd/system/weknora-mcp.service`:
```ini
[Unit]
Description=WeKnora MCP Server
After=network.target

[Service]
Type=simple
User=weknora
WorkingDirectory=/opt/weknora-mcp
Environment=WEKNORA_BASE_URL=http://localhost:8080/api/v1
Environment=WEKNORA_API_KEY=your_api_key
ExecStart=/usr/local/bin/weknora-mcp-server
Restart=always

[Install]
WantedBy=multi-user.target
```

Enable the service:
```bash
sudo systemctl enable weknora-mcp
sudo systemctl start weknora-mcp
```

## Support

If you encounter issues, please:
1. Check log output
2. Check environment configuration
3. Refer to the troubleshooting section
4. Submit an Issue to the project repository: https://github.com/NannaOlympicBroadcast/WeKnoraMCP/issues
