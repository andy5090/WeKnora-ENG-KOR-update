# WeKnora API Documentation

## Table of Contents

- [Overview](#overview)
- [Basic Information](#basic-information)
- [Authentication](#authentication)
- [Error Handling](#error-handling)
- [API Overview](#api-overview)

## Overview

WeKnora provides a series of RESTful APIs for creating and managing knowledge bases, retrieving knowledge, and conducting knowledge-based Q&A. This documentation describes how to use these APIs in detail.

## Basic Information

- **Base URL**: `/api/v1`
- **Response Format**: JSON
- **Authentication**: API Key

## Authentication

All API requests require authentication by including `X-API-Key` in the HTTP request headers:

```
X-API-Key: your_api_key
```

For easier issue tracking and debugging, it is recommended to add `X-Request-ID` to each request's HTTP headers:

```
X-Request-ID: unique_request_id
```

### Obtaining API Key

After completing account registration on the web page, please go to the account information page to obtain your API Key.

Please keep your API Key secure and avoid disclosure. The API Key represents your account identity and has full API access permissions.

## Error Handling

All APIs use standard HTTP status codes to indicate request status and return a unified error response format:

```json
{
  "success": false,
  "error": {
    "code": "Error Code",
    "message": "Error Message",
    "details": "Error Details"
  }
}
```

## API Overview

WeKnora APIs are categorized by functionality as follows:

| Category | Description | Documentation Link |
|----------|-------------|---------------------|
| Tenant Management | Create and manage tenant accounts | [tenant.md](./tenant.md) |
| Knowledge Base Management | Create, query and manage knowledge bases | [knowledge-base.md](./knowledge-base.md) |
| Knowledge Management | Upload, retrieve and manage knowledge content | [knowledge.md](./knowledge.md) |
| Model Management | Configure and manage various AI models | [model.md](./model.md) |
| Chunk Management | Manage knowledge chunks | [chunk.md](./chunk.md) |
| Tag Management | Manage knowledge base tag classifications | [tag.md](./tag.md) |
| FAQ Management | Manage FAQ Q&A pairs | [faq.md](./faq.md) |
| Agent Management | Create and manage custom agents | [agent.md](./agent.md) |
| Session Management | Create and manage conversation sessions | [session.md](./session.md) |
| Knowledge Search | Search content in knowledge bases | [knowledge-search.md](./knowledge-search.md) |
| Chat Functionality | Q&A based on knowledge base and Agent | [chat.md](./chat.md) |
| Message Management | Get and manage conversation messages | [message.md](./message.md) |
| Evaluation Functionality | Evaluate model performance | [evaluation.md](./evaluation.md) |
