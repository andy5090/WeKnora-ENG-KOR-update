# Session Management API

[Back to Index](./README.md)

| Method   | Path                                    | Description                  |
| -------- | --------------------------------------- | ---------------------------- |
| POST     | `/sessions`                             | Create session               |
| GET      | `/sessions/:id`                         | Get session details          |
| GET      | `/sessions`                             | Get tenant's session list   |
| PUT      | `/sessions/:id`                         | Update session               |
| DELETE   | `/sessions/:id`                         | Delete session               |
| POST     | `/sessions/:session_id/generate_title`  | Generate session title       |
| POST     | `/sessions/:session_id/stop`            | Stop session                 |
| GET      | `/sessions/continue-stream/:session_id` | Continue incomplete session  |


## POST `/sessions` - Create Session

**Request**:

```curl
curl --location 'http://localhost:8080/api/v1/sessions' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json' \
--data '{
    "knowledge_base_id": "kb-00000001",
    "session_strategy": {
        "max_rounds": 5,
        "enable_rewrite": true,
        "fallback_strategy": "FIXED_RESPONSE",
        "fallback_response": "Sorry, I cannot answer this question",
        "embedding_top_k": 10,
        "keyword_threshold": 0.5,
        "vector_threshold": 0.7,
        "rerank_model_id": "Rerank Model ID",
        "rerank_top_k": 3,
        "rerank_threshold": 0.7,
        "summary_model_id": "8aea788c-bb30-4898-809e-e40c14ffb48c",
        "summary_parameters": {
            "max_tokens": 0,
            "repeat_penalty": 1,
            "top_k": 0,
            "top_p": 0,
            "frequency_penalty": 0,
            "presence_penalty": 0,
            "prompt": "This is a conversation between user and assistant. xxx",
            "context_template": "You are a professional intelligent information retrieval assistant xxx",
            "no_match_prefix": "<think>\n</think>\nNO_MATCH",
            "temperature": 0.3,
            "seed": 0,
            "max_completion_tokens": 2048
        },
        "no_match_prefix": "<think>\n</think>\nNO_MATCH"
    }
}'
```

**Response**:

```json
{
    "data": {
        "id": "411d6b70-9a85-4d03-bb74-aab0fd8bd12f",
        "title": "",
        "description": "",
        "tenant_id": 1,
        "knowledge_base_id": "kb-00000001",
        "max_rounds": 5,
        "enable_rewrite": true,
        "fallback_strategy": "FIXED_RESPONSE",
        "fallback_response": "Sorry, I cannot answer this question",
        "embedding_top_k": 10,
        "keyword_threshold": 0.5,
        "vector_threshold": 0.7,
        "rerank_model_id": "Rerank Model ID",
        "rerank_top_k": 3,
        "rerank_threshold": 0.7,
        "summary_model_id": "8aea788c-bb30-4898-809e-e40c14ffb48c",
        "summary_parameters": {
            "max_tokens": 0,
            "repeat_penalty": 1,
            "top_k": 0,
            "top_p": 0,
            "frequency_penalty": 0,
            "presence_penalty": 0,
            "prompt": "This is a conversation between user and assistant. xxx",
            "context_template": "You are a professional intelligent information retrieval assistant xxx",
            "no_match_prefix": "<think>\n</think>\nNO_MATCH",
            "temperature": 0.3,
            "seed": 0,
            "max_completion_tokens": 2048
        },
        "agent_config": null,
        "context_config": null,
        "created_at": "2025-08-12T12:26:19.611616669+08:00",
        "updated_at": "2025-08-12T12:26:19.611616919+08:00",
        "deleted_at": null
    },
    "success": true
}
```

## GET `/sessions/:id` - Get Session Details

**Request**:

```curl
curl --location 'http://localhost:8080/api/v1/sessions/ceb9babb-1e30-41d7-817d-fd584954304b' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json'
```

**Response**:

```json
{
    "data": {
        "id": "ceb9babb-1e30-41d7-817d-fd584954304b",
        "title": "Model Optimization Strategy",
        "description": "",
        "tenant_id": 1,
        "knowledge_base_id": "kb-00000001",
        "max_rounds": 5,
        "enable_rewrite": true,
        "fallback_strategy": "fixed",
        "fallback_response": "Sorry, I cannot answer this question.",
        "embedding_top_k": 10,
        "keyword_threshold": 0.3,
        "vector_threshold": 0.5,
        "rerank_model_id": "",
        "rerank_top_k": 5,
        "rerank_threshold": 0.7,
        "summary_model_id": "8aea788c-bb30-4898-809e-e40c14ffb48c",
        "summary_parameters": {
            "max_tokens": 0,
            "repeat_penalty": 1,
            "top_k": 0,
            "top_p": 0,
            "frequency_penalty": 0,
            "presence_penalty": 0,
            "prompt": "This is a conversation between user and assistant",
            "context_template": "You are a professional intelligent information retrieval assistant",
            "no_match_prefix": "<think>\n</think>\nNO_MATCH",
            "temperature": 0.3,
            "seed": 0,
            "max_completion_tokens": 2048
        },
        "agent_config": null,
        "context_config": null,
        "created_at": "2025-08-12T10:24:38.308596+08:00",
        "updated_at": "2025-08-12T10:25:41.317761+08:00",
        "deleted_at": null
    },
    "success": true
}
```

## GET `/sessions?page=&page_size=` - Get Tenant's Session List

**Request**:

```curl
curl --location 'http://localhost:8080/api/v1/sessions?page=1&page_size=1' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json'
```

**Response**:

```json
{
    "data": [
        {
            "id": "411d6b70-9a85-4d03-bb74-aab0fd8bd12f",
            "title": "",
            "description": "",
            "tenant_id": 1,
            "knowledge_base_id": "kb-00000001",
            "max_rounds": 5,
            "enable_rewrite": true,
            "fallback_strategy": "FIXED_RESPONSE",
            "fallback_response": "Sorry, I cannot answer this question",
            "embedding_top_k": 10,
            "keyword_threshold": 0.5,
            "vector_threshold": 0.7,
            "rerank_model_id": "Rerank Model ID",
            "rerank_top_k": 3,
            "rerank_threshold": 0.7,
            "summary_model_id": "8aea788c-bb30-4898-809e-e40c14ffb48c",
            "summary_parameters": {
                "max_tokens": 0,
                "repeat_penalty": 1,
                "top_k": 0,
                "top_p": 0,
                "frequency_penalty": 0,
                "presence_penalty": 0,
                "prompt": "This is a conversation between user and assistant. xxx",
                "context_template": "You are a professional intelligent information retrieval assistant xxx",
                "no_match_prefix": "<think>\n</think>\nNO_MATCH",
                "temperature": 0.3,
                "seed": 0,
                "max_completion_tokens": 2048
            },
            "created_at": "2025-08-12T12:26:19.611616+08:00",
            "updated_at": "2025-08-12T12:26:19.611616+08:00",
            "deleted_at": null
        }
    ],
    "page": 1,
    "page_size": 1,
    "success": true,
    "total": 2
}
```

## PUT `/sessions/:id` - Update Session

**Request**:

```curl
curl --location --request PUT 'http://localhost:8080/api/v1/sessions/411d6b70-9a85-4d03-bb74-aab0fd8bd12f' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json' \
--data '{
    "title": "weknora",
    "description": "weknora description",
    "knowledge_base_id": "kb-00000001",
    "max_rounds": 5,
    "enable_rewrite": true,
    "fallback_strategy": "FIXED_RESPONSE",
    "fallback_response": "Sorry, I cannot answer this question",
    "embedding_top_k": 10,
    "keyword_threshold": 0.5,
    "vector_threshold": 0.7,
    "rerank_model_id": "Rerank Model ID",
    "rerank_top_k": 3,
    "rerank_threshold": 0.7,
    "summary_model_id": "8aea788c-bb30-4898-809e-e40c14ffb48c",
    "summary_parameters": {
        "max_tokens": 0,
        "repeat_penalty": 1,
        "top_k": 0,
        "top_p": 0,
        "frequency_penalty": 0,
        "presence_penalty": 0,
        "prompt": "This is a conversation between user and assistant. xxx",
        "context_template": "You are a professional intelligent information retrieval assistant xxx",
        "no_match_prefix": "<think>\n</think>\nNO_MATCH",
        "temperature": 0.3,
        "seed": 0,
        "max_completion_tokens": 2048
    }
}'
```

**Response**:

```json
{
    "data": {
        "id": "411d6b70-9a85-4d03-bb74-aab0fd8bd12f",
        "title": "weknora",
        "description": "weknora description",
        "tenant_id": 1,
        "knowledge_base_id": "kb-00000001",
        "max_rounds": 5,
        "enable_rewrite": true,
        "fallback_strategy": "FIXED_RESPONSE",
        "fallback_response": "Sorry, I cannot answer this question",
        "embedding_top_k": 10,
        "keyword_threshold": 0.5,
        "vector_threshold": 0.7,
        "rerank_model_id": "Rerank Model ID",
        "rerank_top_k": 3,
        "rerank_threshold": 0.7,
        "summary_model_id": "8aea788c-bb30-4898-809e-e40c14ffb48c",
        "summary_parameters": {
            "max_tokens": 0,
            "repeat_penalty": 1,
            "top_k": 0,
            "top_p": 0,
            "frequency_penalty": 0,
            "presence_penalty": 0,
            "prompt": "This is a conversation between user and assistant. xxx",
            "context_template": "You are a professional intelligent information retrieval assistant xxx",
            "no_match_prefix": "<think>\n</think>\nNO_MATCH",
            "temperature": 0.3,
            "seed": 0,
            "max_completion_tokens": 2048
        },
        "created_at": "0001-01-01T00:00:00Z",
        "updated_at": "2025-08-12T14:20:56.738424351+08:00",
        "deleted_at": null
    },
    "success": true
}
```

## DELETE `/sessions/:id` - Delete Session

**Request**:

```curl
curl --location --request DELETE 'http://localhost:8080/api/v1/sessions/411d6b70-9a85-4d03-bb74-aab0fd8bd12f' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json'
```

**Response**:

```json
{
    "message": "Session deleted successfully",
    "success": true
}
```

## POST `/sessions/:session_id/generate_title` - Generate Session Title

**Request**:

```curl
curl --location 'http://localhost:8080/api/v1/sessions/ceb9babb-1e30-41d7-817d-fd584954304b/generate_title' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json' \
--data '{
  "messages": [
    {
      "role": "user",
      "content": "Hello, I would like to learn about artificial intelligence"
    },
    {
      "role": "assistant",
      "content": "Artificial intelligence is a branch of computer science..."
    }
  ]
}'
```

**Response**:

```json
{
    "data": "Model Optimization Strategy",
    "success": true
}
```

## POST `/sessions/:session_id/stop` - Stop Session

**Request**:

```curl
curl --location 'http://localhost:8080/api/v1/sessions/7c966c74-610e-4516-8d5b-05e14b2e4ee0/stop' \
--header 'X-API-Key: sk-An-W8T4tdZDbWKJgfwgdea5rS8ue_mRhCTZ8Smhnvku-bWEE' \
--header 'Content-Type: application/json' \
--data '{"message_id":"ebbf7e53-dfe6-44d5-882f-36a4104910b5"}'
```

**Response**:

```json
{
    "message": "Session stopped successfully",
    "success": true
}
```

## GET `/sessions/continue-stream/:session_id` - Continue Incomplete Session

**Query Parameters**:
- `message_id`: Message ID from `/messages/:session_id/load` endpoint where `is_completed` is `false`

**Request**:

```curl
curl --location 'http://localhost:8080/api/v1/sessions/continue-stream/ceb9babb-1e30-41d7-817d-fd584954304b?message_id=b8b90eeb-7dd5-4cf9-81c6-5ebcbd759451' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json'
```

**Response Format**:
Server-Sent Events, consistent with `/knowledge-chat/:session_id` response
