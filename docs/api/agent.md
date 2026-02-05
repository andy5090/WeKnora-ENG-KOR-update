# Agent Management API

[Back to Index](./README.md)

## Overview

The Agent API is used to manage custom agents. The system provides built-in agents and also supports users creating custom agents to meet different business scenario requirements.

### Built-in Agents

The system provides the following built-in agents by default:

| ID | Name | Description | Mode |
|----|------|-------------|------|
| `builtin-quick-answer` | Quick Answer | RAG-based Q&A using knowledge base, answering questions quickly and accurately | quick-answer |
| `builtin-smart-reasoning` | Smart Reasoning | ReAct reasoning framework supporting multi-step thinking and tool calling | smart-reasoning |
| `builtin-data-analyst` | Data Analyst | Professional data analysis agent supporting SQL queries and statistical analysis for CSV/Excel files | smart-reasoning |

### Agent Modes

| Mode | Description |
|------|-------------|
| `quick-answer` | RAG mode for quick Q&A, directly generating answers based on knowledge base retrieval results |
| `smart-reasoning` | ReAct mode supporting multi-step reasoning and tool calling |

## API List

| Method | Path | Description |
|--------|------|-------------|
| POST | `/agents` | Create agent |
| GET | `/agents` | List agents |
| GET | `/agents/:id` | Get agent details |
| PUT | `/agents/:id` | Update agent |
| DELETE | `/agents/:id` | Delete agent |
| POST | `/agents/:id/copy` | Copy agent |
| GET | `/agents/placeholders` | Get placeholder definitions |

---

## POST `/agents` - Create Agent

Create a new custom agent.

**Request**:

```curl
curl --location 'http://localhost:8080/api/v1/agents' \
--header 'X-API-Key: your_api_key' \
--header 'Content-Type: application/json' \
--data '{
    "name": "My Agent",
    "description": "Custom agent description",
    "avatar": "🤖",
    "config": {
        "agent_mode": "smart-reasoning",
        "system_prompt": "You are a professional assistant...",
        "temperature": 0.7,
        "max_iterations": 10,
        "kb_selection_mode": "all",
        "web_search_enabled": true,
        "multi_turn_enabled": true,
        "history_turns": 5
    }
}'
```

**Request Parameters**:

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `name` | string | Yes | Agent name |
| `description` | string | No | Agent description |
| `avatar` | string | No | Agent avatar (emoji or icon name) |
| `config` | object | No | Agent configuration, see [Configuration Parameters](#configuration-parameters) |

**Response**:

```json
{
    "success": true,
    "data": {
        "id": "550e8400-e29b-41d4-a716-446655440000",
        "name": "My Agent",
        "description": "Custom agent description",
        "avatar": "🤖",
        "is_builtin": false,
        "tenant_id": 1,
        "created_by": "user-123",
        "config": {
            "agent_mode": "smart-reasoning",
            "system_prompt": "You are a professional assistant...",
            "temperature": 0.7,
            "max_iterations": 10
        },
        "created_at": "2025-01-19T10:00:00Z",
        "updated_at": "2025-01-19T10:00:00Z"
    }
}
```

**Error Response**:

| Status Code | Error Code | Error | Description |
|-------------|------------|-------|-------------|
| 400 | 1000 | Bad Request | Invalid request parameters or agent name is empty |
| 500 | 1007 | Internal Server Error | Internal server error |

---

## GET `/agents` - List Agents

Get all agents for the current tenant, including built-in agents and custom agents.

**Request**:

```curl
curl --location 'http://localhost:8080/api/v1/agents' \
--header 'X-API-Key: your_api_key'
```

**Response**:

```json
{
    "success": true,
    "data": [
        {
            "id": "builtin-quick-answer",
            "name": "Quick Answer",
            "description": "RAG-based Q&A using knowledge base, answering questions quickly and accurately",
            "avatar": "💬",
            "is_builtin": true,
            "tenant_id": 10000,
            "created_by": "",
            "config": {
                "agent_mode": "quick-answer",
                "system_prompt": "You are a professional intelligent information retrieval assistant named WeKnora. You are like a professional senior secretary, answering user questions based on retrieved information, without using any prior knowledge.\nWhen users ask questions, the assistant will answer based on specific information. The assistant first thinks through the reasoning process, then provides answers to users.\n",
                "context_template": "...",
                "model_id": "...",
                "rerank_model_id": "",
                "temperature": 0.3,
                "max_completion_tokens": 2048,
                "max_iterations": 10,
                "allowed_tools": [],
                "reflection_enabled": false,
                "mcp_selection_mode": "",
                "mcp_services": null,
                "kb_selection_mode": "all",
                "knowledge_bases": [],
                "supported_file_types": null,
                "faq_priority_enabled": false,
                "faq_direct_answer_threshold": 0,
                "faq_score_boost": 0,
                "web_search_enabled": false,
                "web_search_max_results": 5,
                "multi_turn_enabled": true,
                "history_turns": 5,
                "embedding_top_k": 10,
                "keyword_threshold": 0.3,
                "vector_threshold": 0.5,
                "rerank_top_k": 5,
                "rerank_threshold": 0.5,
                "enable_query_expansion": true,
                "enable_rewrite": true,
                "rewrite_prompt_system": "...",
                "rewrite_prompt_user": "...",
                "fallback_strategy": "fixed",
                "fallback_response": "...",
                "fallback_prompt": "..."
            },
            "created_at": "2025-12-29T20:06:01.696308+08:00",
            "updated_at": "2025-12-29T20:06:01.696308+08:00",
            "deleted_at": null
        },
        {
            "id": "builtin-smart-reasoning",
            "name": "Smart Reasoning",
            "description": "ReAct reasoning framework supporting multi-step thinking and tool calling",
            "is_builtin": true,
            "config": {
                "agent_mode": "smart-reasoning"
  
            }
        },
        {
            "id": "550e8400-e29b-41d4-a716-446655440000",
            "name": "My Agent",
            "description": "Custom agent description",
            "is_builtin": false,
            "config": {
                "agent_mode": "smart-reasoning"
            }
        }
    ]
}
```

---

## GET `/agents/:id` - Get Agent Details

Get detailed information about an agent by ID.

**Request**:

```curl
curl --location 'http://localhost:8080/api/v1/agents/builtin-quick-answer' \
--header 'X-API-Key: your_api_key'
```

**Response**:

```json
{
    "success": true,
    "data": {
        "id": "builtin-quick-answer",
        "name": "Quick Answer",
        "description": "RAG-based Q&A using knowledge base, answering questions quickly and accurately",
        "is_builtin": true,
        "tenant_id": 1,
        "config": {
            "agent_mode": "quick-answer",
            "system_prompt": "",
            "context_template": "Please answer the user's question based on the following reference materials...",
            "temperature": 0.7,
            "max_completion_tokens": 2048,
            "kb_selection_mode": "all",
            "web_search_enabled": true,
            "multi_turn_enabled": true,
            "history_turns": 5
        },
        "created_at": "2025-01-01T00:00:00Z",
        "updated_at": "2025-01-01T00:00:00Z"
    }
}
```

**Error Response**:

| Status Code | Error Code | Error | Description |
|-------------|------------|-------|-------------|
| 400 | 1000 | Bad Request | Agent ID is empty |
| 404 | 1003 | Not Found | Agent not found |
| 500 | 1007 | Internal Server Error | Internal server error |

---

## PUT `/agents/:id` - Update Agent

Update agent name, description, and configuration. Built-in agents cannot be modified.

**Request**:

```curl
curl --location --request PUT 'http://localhost:8080/api/v1/agents/550e8400-e29b-41d4-a716-446655440000' \
--header 'X-API-Key: your_api_key' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Updated Agent",
    "description": "Updated description",
    "config": {
        "agent_mode": "smart-reasoning",
        "temperature": 0.8,
        "max_iterations": 20
    }
}'
```

**Request Parameters**:

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `name` | string | No | Agent name |
| `description` | string | No | Agent description |
| `avatar` | string | No | Agent avatar |
| `config` | object | No | Agent configuration |

**Response**:

```json
{
    "success": true,
    "data": {
        "id": "550e8400-e29b-41d4-a716-446655440000",
        "name": "Updated Agent",
        "description": "Updated description",
        "config": {
            "agent_mode": "smart-reasoning",
            "temperature": 0.8,
            "max_iterations": 20
        },
        "updated_at": "2025-01-19T11:00:00Z"
    }
}
```

**Error Response**:

| Status Code | Error Code | Error | Description |
|-------------|------------|-------|-------------|
| 400 | 1000 | Bad Request | Invalid request parameters or agent name is empty |
| 403 | 1002 | Forbidden | Cannot modify basic information of built-in agent |
| 404 | 1003 | Not Found | Agent not found |
| 500 | 1007 | Internal Server Error | Internal server error |

---

## DELETE `/agents/:id` - Delete Agent

Delete the specified custom agent. Built-in agents cannot be deleted.

**Request**:

```curl
curl --location --request DELETE 'http://localhost:8080/api/v1/agents/550e8400-e29b-41d4-a716-446655440000' \
--header 'X-API-Key: your_api_key'
```

**Response**:

```json
{
    "success": true,
    "message": "Agent deleted successfully"
}
```

**Error Response**:

| Status Code | Error Code | Error | Description |
|-------------|------------|-------|-------------|
| 400 | 1000 | Bad Request | Agent ID is empty |
| 403 | 1002 | Forbidden | Cannot delete built-in agent |
| 404 | 1003 | Not Found | Agent not found |
| 500 | 1007 | Internal Server Error | Internal server error |

---

## POST `/agents/:id/copy` - Copy Agent

Copy the specified agent to create a new copy. Supports copying built-in agents.

**Request**:

```curl
curl --location --request POST 'http://localhost:8080/api/v1/agents/builtin-smart-reasoning/copy' \
--header 'X-API-Key: your_api_key'
```

**Response**:

```json
{
    "success": true,
    "data": {
        "id": "660e8400-e29b-41d4-a716-446655440001",
        "name": "Smart Reasoning (Copy)",
        "description": "ReAct reasoning framework supporting multi-step thinking and tool calling",
        "is_builtin": false,
        "config": {
            "agent_mode": "smart-reasoning",
            "max_iterations": 50
        },
        "created_at": "2025-01-19T12:00:00Z",
        "updated_at": "2025-01-19T12:00:00Z"
    }
}
```

**Error Response**:

| Status Code | Error Code | Error | Description |
|-------------|------------|-------|-------------|
| 400 | 1000 | Bad Request | Agent ID is empty |
| 404 | 1003 | Not Found | Agent not found |
| 500 | 1007 | Internal Server Error | Internal server error |

---

## GET `/agents/placeholders` - Get Placeholder Definitions

Get all available prompt placeholder definitions, grouped by field type. These placeholders can be used in system prompts and context templates.

**Request**:

```curl
curl --location 'http://localhost:8080/api/v1/agents/placeholders' \
--header 'X-API-Key: your_api_key'
```

**Response**:

```json
{
    "success": true,
    "data": {
        "all": [...],
        "system_prompt": [...],
        "agent_system_prompt": [...],
        "context_template": [...],
        "rewrite_system_prompt": [...],
        "rewrite_prompt": [...],
        "fallback_prompt": [...]
    }
}
```

---

## Configuration Parameters

The agent's `config` object supports the following configuration items:

### Basic Settings

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `agent_mode` | string | - | Agent mode: `quick-answer` (RAG) or `smart-reasoning` (ReAct) |
| `system_prompt` | string | - | System prompt, supports placeholders |
| `context_template` | string | - | Context template (only used in quick-answer mode) |

### Model Settings

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `model_id` | string | - | Conversation model ID |
| `rerank_model_id` | string | - | Rerank model ID |
| `temperature` | float | 0.7 | Temperature parameter (0-1) |
| `max_completion_tokens` | int | 2048 | Maximum completion tokens |

### Agent Mode Settings

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `max_iterations` | int | 10 | Maximum ReAct iterations |
| `allowed_tools` | []string | - | List of allowed tools |
| `reflection_enabled` | bool | false | Whether reflection is enabled |
| `mcp_selection_mode` | string | - | MCP service selection mode: `all`/`selected`/`none` |
| `mcp_services` | []string | - | Selected MCP service ID list |

### Knowledge Base Settings

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `kb_selection_mode` | string | - | Knowledge base selection mode: `all`/`selected`/`none` |
| `knowledge_bases` | []string | - | Associated knowledge base ID list |
| `supported_file_types` | []string | - | Supported file types (e.g., `["csv", "xlsx"]`) |

### FAQ Strategy Settings

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `faq_priority_enabled` | bool | true | FAQ priority strategy switch |
| `faq_direct_answer_threshold` | float | 0.9 | FAQ direct answer threshold |
| `faq_score_boost` | float | 1.2 | FAQ score boost multiplier |

### Web Search Settings

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `web_search_enabled` | bool | true | Whether web search is enabled |
| `web_search_max_results` | int | 5 | Maximum web search results |

### Multi-turn Conversation Settings

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `multi_turn_enabled` | bool | true | Whether multi-turn conversation is enabled |
| `history_turns` | int | 5 | Number of history turns to keep |

### Retrieval Strategy Settings

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `embedding_top_k` | int | 10 | Vector retrieval TopK |
| `keyword_threshold` | float | 0.3 | Keyword retrieval threshold |
| `vector_threshold` | float | 0.5 | Vector retrieval threshold |
| `rerank_top_k` | int | 5 | Rerank TopK |
| `rerank_threshold` | float | 0.5 | Rerank threshold |

### Advanced Settings

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `enable_query_expansion` | bool | true | Whether query expansion is enabled |
| `enable_rewrite` | bool | true | Whether multi-turn conversation query rewriting is enabled |
| `rewrite_prompt_system` | string | - | Rewrite system prompt |
| `rewrite_prompt_user` | string | - | Rewrite user prompt template |
| `fallback_strategy` | string | model | Fallback strategy: `fixed` (fixed response) or `model` (model generation) |
| `fallback_response` | string | - | Fixed fallback response (used when `fallback_strategy` is `fixed`) |
| `fallback_prompt` | string | - | Fallback prompt (used when `fallback_strategy` is `model`) |

---

## Using Agent for Q&A

After creating or obtaining an agent, you can use the agent for Q&A through the `/agent-chat/:session_id` endpoint. For details, please refer to [Chat API](./chat.md).

Use the `agent_id` parameter in the Q&A request to specify the agent to use:

```curl
curl --location 'http://localhost:8080/api/v1/agent-chat/session-123' \
--header 'X-API-Key: your_api_key' \
--header 'Content-Type: application/json' \
--data '{
    "query": "Help me analyze this data",
    "agent_enabled": true,
    "agent_id": "builtin-data-analyst"
}'
```
