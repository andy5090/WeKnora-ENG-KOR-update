# Chat API

[Back to Index](./README.md)

| Method | Path                          | Description                    |
| ------ | ----------------------------- | ----------------------------- |
| POST   | `/knowledge-chat/:session_id` | Knowledge base Q&A             |
| POST   | `/agent-chat/:session_id`     | Agent-based intelligent Q&A    |
| POST   | `/knowledge-search`           | Knowledge base search           |

## POST `/knowledge-chat/:session_id` - Knowledge Base Q&A

**Request**:

```curl
curl --location 'http://localhost:8080/api/v1/knowledge-chat/ceb9babb-1e30-41d7-817d-fd584954304b' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json' \
--data '{
    "query": "Comet tail shape"
}'
```

**Response Format**:
Server-Sent Events (Content-Type: text/event-stream)

**Response**:

```
event: message
data: {"id":"3475c004-0ada-4306-9d30-d7f5efce50d2","response_type":"references","content":"","done":false,"knowledge_references":[{"id":"c8347bef-127f-4a22-b962-edf5a75386ec","content":"Comet xxx.","knowledge_id":"a6790b93-4700-4676-bd48-0d4804e1456b","chunk_index":0,"knowledge_title":"Comet.txt","start_at":0,"end_at":2760,"seq":0,"score":4.038836479187012,"match_type":3,"sub_chunk_id":["688821f0-40bf-428e-8cb6-541531ebeb76","c1e9903e-2b4d-4281-be15-0149288d45c2","7d955251-3f79-4fd5-a6aa-02f81e044091"],"metadata":{},"chunk_type":"text","parent_chunk_id":"","image_info":"","knowledge_filename":"Comet.txt","knowledge_source":""},{"id":"fa3aadee-cadb-4a84-9941-c839edc3e626","content":"# Document Name\nComet.txt\n\n# Summary\nComets are small solar system bodies composed of ice and dust. When approaching the sun, they release gas forming a coma and tail. Their orbital periods vary greatly, with sources including the Kuiper Belt and Oort Cloud. The distinction between comets and asteroids is gradually blurring, with some comets having lost volatile material, similar to asteroids. Currently, numerous comets are known, and exocomets exist. Comets were considered omens in ancient times, and modern research reveals their complex structure and origin.","knowledge_id":"a6790b93-4700-4676-bd48-0d4804e1456b","chunk_index":6,"knowledge_title":"Comet.txt","start_at":0,"end_at":0,"seq":6,"score":0.6131043121858466,"match_type":3,"sub_chunk_id":null,"metadata":{},"chunk_type":"summary","parent_chunk_id":"c8347bef-127f-4a22-b962-edf5a75386ec","image_info":"","knowledge_filename":"Comet.txt","knowledge_source":""}]}

event: message
data: {"id":"3475c004-0ada-4306-9d30-d7f5efce50d2","response_type":"answer","content":"Manifests as","done":false,"knowledge_references":null}

event: message
data: {"id":"3475c004-0ada-4306-9d30-d7f5efce50d2","response_type":"answer","content":" structure","done":false,"knowledge_references":null}

event: message
data: {"id":"3475c004-0ada-4306-9d30-d7f5efce50d2","response_type":"answer","content":".","done":false,"knowledge_references":null}

event: message
data: {"id":"3475c004-0ada-4306-9d30-d7f5efce50d2","response_type":"answer","content":"","done":true,"knowledge_references":null}
```

## POST `/agent-chat/:session_id` - Agent-based Intelligent Q&A

Agent mode supports more intelligent Q&A, including tool calling, web search, multi-knowledge base retrieval, and other capabilities.

**Request Parameters**:
- `query`: Query text (required)
- `knowledge_base_ids`: Knowledge base ID array, can dynamically specify knowledge bases to use for this query (optional)
- `knowledge_ids`: Knowledge file ID array, can dynamically specify specific knowledge files to use for this query (optional)
- `agent_enabled`: Whether to enable Agent mode (optional, default false)
- `agent_id`: Custom Agent ID, specifies the custom agent to use (optional)
- `web_search_enabled`: Whether to enable web search (optional, default false)
- `summary_model_id`: Override the session's default summary model ID (optional)
- `mentioned_items`: @ Mentioned knowledge bases and file list (optional)
- `disable_title`: Whether to disable automatic title generation (optional, default false)
- `mcp_service_ids`: MCP service whitelist (optional, deprecated)

**Request**:

```curl
curl --location 'http://localhost:8080/api/v1/agent-chat/ceb9babb-1e30-41d7-817d-fd584954304b' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json' \
--data '{
    "query": "Help me check today'\''s weather",
    "agent_enabled": true,
    "web_search_enabled": true,
    "knowledge_base_ids": ["kb-00000001"],
    "agent_id": "agent-001",
    "mentioned_items": [
        {
            "id": "kb-00000001",
            "name": "Weather Knowledge Base",
            "type": "kb",
            "kb_type": "document"
        }
    ]
}'
```

**Response Format**:
Server-Sent Events (Content-Type: text/event-stream)

**Response Type Description**:

| response_type | Description          |
|---------------|----------------------|
| `thinking`    | Agent thinking process |
| `tool_call`   | Tool call information |
| `tool_result`| Tool call result      |
| `references` | Knowledge base retrieval references |
| `answer`      | Final answer content |
| `reflection`  | Agent reflection content |
| `error`       | Error information     |

**Response Example**:

```
event: message
data: {"id":"agent-001","response_type":"thinking","content":"User wants to check weather, I need to use web search tool...","done":false,"knowledge_references":null}

event: message
data: {"id":"agent-001","response_type":"tool_call","content":"","done":false,"knowledge_references":null,"data":{"tool_name":"web_search","arguments":{"query":"Today weather"}}}

event: message
data: {"id":"agent-001","response_type":"tool_result","content":"Search results: Today sunny, temperature 25°C...","done":false,"knowledge_references":null}

event: message
data: {"id":"agent-001","response_type":"answer","content":"According to the search results, today's weather is sunny with a temperature of about 25°C.","done":false,"knowledge_references":null}

event: message
data: {"id":"agent-001","response_type":"answer","content":"","done":true,"knowledge_references":null}
```
