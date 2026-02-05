# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2024-01-XX

### Added
- Initial version release
- WeKnora MCP Server core functionality
- Complete WeKnora API integration
- Tenant management tools
- Knowledge base management tools
- Knowledge management tools
- Model management tools
- Session management tools
- Chat functionality tools
- Chunk management tools
- Multiple startup method support
- Command line argument support
- Environment variable configuration
- Complete package installation support
- Development and production modes
- Detailed documentation and installation guide

### Tool List
- `create_tenant` - Create a new tenant
- `list_tenants` - List all tenants
- `create_knowledge_base` - Create a knowledge base
- `list_knowledge_bases` - List knowledge bases
- `get_knowledge_base` - Get knowledge base details
- `delete_knowledge_base` - Delete a knowledge base
- `hybrid_search` - Hybrid search
- `create_knowledge_from_url` - Create knowledge from URL
- `list_knowledge` - List knowledge
- `get_knowledge` - Get knowledge details
- `delete_knowledge` - Delete knowledge
- `create_model` - Create a model
- `list_models` - List models
- `get_model` - Get model details
- `create_session` - Create a chat session
- `get_session` - Get session details
- `list_sessions` - List sessions
- `delete_session` - Delete a session
- `chat` - Send a chat message
- `list_chunks` - List knowledge chunks
- `delete_chunk` - Delete a knowledge chunk

### File Structure
```
WeKnoraMCP/
├── __init__.py              # Package initialization file
├── main.py                  # Main entry point (Recommended)
├── run.py                   # Convenience startup script
├── run_server.py           # Original startup script
├── weknora_mcp_server.py   # MCP server implementation
├── test_module.py          # Module test script
├── requirements.txt        # Dependency list
├── setup.py               # Installation script (Traditional)
├── pyproject.toml         # Modern project configuration
├── MANIFEST.in            # Include file manifest
├── LICENSE                # MIT License
├── README.md              # Project description
├── INSTALL.md             # Detailed installation guide
└── CHANGELOG.md           # Changelog
```

### Startup Methods
1. `python main.py` - Main entry point (Recommended)
2. `python run_server.py` - Original startup script
3. `python run.py` - Convenience startup script
4. `python weknora_mcp_server.py` - Run directly
5. `python -m weknora_mcp_server` - Run as module
6. `weknora-mcp-server` - Command line tool after installation
7. `weknora-server` - Command line tool after installation (alias)

### Technical Features
- Based on Model Context Protocol (MCP) 1.0.0+
- Async I/O support
- Complete error handling
- Detailed logging
- Environment variable configuration
- Command line argument support
- Multiple installation methods
- Development and production modes
- Complete test coverage

### Dependencies
- Python 3.10+
- mcp >= 1.0.0
- requests >= 2.31.0

### Compatibility
- Supports Windows, macOS, Linux
- Supports Python 3.10-3.12
- Compatible with modern Python package management tools
