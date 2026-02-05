# FAQ Management API

[Back to Index](./README.md)

| Method   | Path                                        | Description                    |
| -------- | ------------------------------------------- | ----------------------------- |
| GET      | `/knowledge-bases/:id/faq/entries`          | List FAQ entries              |
| POST     | `/knowledge-bases/:id/faq/entries`          | Batch import FAQ entries      |
| POST     | `/knowledge-bases/:id/faq/entry`            | Create single FAQ entry       |
| PUT      | `/knowledge-bases/:id/faq/entries/:entry_id`| Update single FAQ entry       |
| PUT      | `/knowledge-bases/:id/faq/entries/status`   | Batch update FAQ enabled status |
| PUT      | `/knowledge-bases/:id/faq/entries/tags`     | Batch update FAQ tags         |
| DELETE   | `/knowledge-bases/:id/faq/entries`          | Batch delete FAQ entries      |
| POST     | `/knowledge-bases/:id/faq/search`           | Hybrid search FAQ             |

## GET `/knowledge-bases/:id/faq/entries` - List FAQ Entries

**Query Parameters**:
- `page`: Page number (default 1)
- `page_size`: Items per page (default 20)
- `tag_id`: Filter by tag ID (optional)
- `keyword`: Keyword search (optional)
- `search_field`: Search field (optional), options:
  - `standard_question`: Search only standard questions
  - `similar_questions`: Search only similar questions
  - `answers`: Search only answers
  - Leave empty or omit: Search all fields
- `sort_order`: Sort order (optional), `asc` means ascending by update time, default is descending by update time

**Request**:

```curl
# Search all fields
curl --location 'http://localhost:8080/api/v1/knowledge-bases/kb-00000001/faq/entries?page=1&page_size=10&keyword=password' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json'

# Search only standard questions
curl --location 'http://localhost:8080/api/v1/knowledge-bases/kb-00000001/faq/entries?keyword=password&search_field=standard_question' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ'

# Search only similar questions
curl --location 'http://localhost:8080/api/v1/knowledge-bases/kb-00000001/faq/entries?keyword=forgot&search_field=similar_questions' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ'

# Search only answers
curl --location 'http://localhost:8080/api/v1/knowledge-bases/kb-00000001/faq/entries?keyword=click&search_field=answers' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ'
```

**Response**:

```json
{
    "data": {
        "total": 100,
        "page": 1,
        "page_size": 10,
        "data": [
            {
                "id": "faq-00000001",
                "chunk_id": "chunk-00000001",
                "knowledge_id": "knowledge-00000001",
                "knowledge_base_id": "kb-00000001",
                "tag_id": "tag-00000001",
                "is_enabled": true,
                "standard_question": "How to reset password?",
                "similar_questions": ["What to do if forgot password", "Password recovery"],
                "negative_questions": ["How to change username"],
                "answers": ["You can reset your password by clicking the 'Forgot Password' link on the login page."],
                "index_mode": "hybrid",
                "chunk_type": "faq",
                "created_at": "2025-08-12T10:00:00+08:00",
                "updated_at": "2025-08-12T10:00:00+08:00"
            }
        ]
    },
    "success": true
}
```

## POST `/knowledge-bases/:id/faq/entries` - Batch Import FAQ Entries

**Request Parameters**:
- `mode`: Import mode, `append` (append) or `replace` (replace)
- `entries`: FAQ entry array
- `knowledge_id`: Associated knowledge ID (optional)

**Request**:

```curl
curl --location 'http://localhost:8080/api/v1/knowledge-bases/kb-00000001/faq/entries' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json' \
--data '{
    "mode": "append",
    "entries": [
        {
            "standard_question": "How to contact customer service?",
            "similar_questions": ["Customer service phone", "Online customer service"],
            "answers": ["You can contact our customer service by calling 400-xxx-xxxx."],
            "tag_id": "tag-00000001"
        },
        {
            "standard_question": "What is the refund policy?",
            "answers": ["We offer a 7-day no-questions-asked refund service."]
        }
    ]
}'
```

**Response**:

```json
{
    "data": {
        "task_id": "task-00000001"
    },
    "success": true
}
```

Note: Batch import is an asynchronous operation, returns a task ID for tracking progress.

## POST `/knowledge-bases/:id/faq/entry` - Create Single FAQ Entry

Synchronously create a single FAQ entry, suitable for single entry scenarios. Automatically checks if standard questions and similar questions duplicate existing FAQs.

**Request Parameters**:
- `standard_question`: Standard question (required)
- `similar_questions`: Similar questions array (optional)
- `negative_questions`: Negative example questions array (optional)
- `answers`: Answers array (required)
- `tag_id`: Tag ID (optional)
- `is_enabled`: Whether enabled (optional, default true)

**Request**:

```curl
curl --location 'http://localhost:8080/api/v1/knowledge-bases/kb-00000001/faq/entry' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json' \
--data '{
    "standard_question": "How to contact customer service?",
    "similar_questions": ["Customer service phone", "Online customer service"],
    "answers": ["You can contact our customer service by calling 400-xxx-xxxx."],
    "tag_id": "tag-00000001",
    "is_enabled": true
}'
```

**Response**:

```json
{
    "data": {
        "id": "faq-00000001",
        "chunk_id": "chunk-00000001",
        "knowledge_id": "knowledge-00000001",
        "knowledge_base_id": "kb-00000001",
        "tag_id": "tag-00000001",
        "is_enabled": true,
        "standard_question": "How to contact customer service?",
        "similar_questions": ["Customer service phone", "Online customer service"],
        "negative_questions": [],
        "answers": ["You can contact our customer service by calling 400-xxx-xxxx."],
        "index_mode": "hybrid",
        "chunk_type": "faq",
        "created_at": "2025-08-12T10:00:00+08:00",
        "updated_at": "2025-08-12T10:00:00+08:00"
    },
    "success": true
}
```

**Error Response** (when standard question or similar question duplicates):

```json
{
    "success": false,
    "error": {
        "code": "BAD_REQUEST",
        "message": "Standard question duplicates existing FAQ"
    }
}
```

## PUT `/knowledge-bases/:id/faq/entries/:entry_id` - Update Single FAQ Entry

**Request**:

```curl
curl --location --request PUT 'http://localhost:8080/api/v1/knowledge-bases/kb-00000001/faq/entries/faq-00000001' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json' \
--data '{
    "standard_question": "How to reset account password?",
    "similar_questions": ["What to do if forgot password", "Password recovery", "Reset password"],
    "answers": ["You can reset your password by following these steps: 1. Click \"Forgot Password\" on the login page 2. Enter your registered email 3. Check your email for reset instructions"],
    "is_enabled": true
}'
```

**Response**:

```json
{
    "success": true
}
```

## PUT `/knowledge-bases/:id/faq/entries/status` - Batch Update FAQ Enabled Status

**Request**:

```curl
curl --location --request PUT 'http://localhost:8080/api/v1/knowledge-bases/kb-00000001/faq/entries/status' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json' \
--data '{
    "updates": {
        "faq-00000001": true,
        "faq-00000002": false,
        "faq-00000003": true
    }
}'
```

**Response**:

```json
{
    "success": true
}
```

## PUT `/knowledge-bases/:id/faq/entries/tags` - Batch Update FAQ Tags

**Request**:

```curl
curl --location --request PUT 'http://localhost:8080/api/v1/knowledge-bases/kb-00000001/faq/entries/tags' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json' \
--data '{
    "updates": {
        "faq-00000001": "tag-00000001",
        "faq-00000002": "tag-00000002",
        "faq-00000003": null
    }
}'
```

Note: Setting to `null` clears the tag association.

**Response**:

```json
{
    "success": true
}
```

## DELETE `/knowledge-bases/:id/faq/entries` - Batch Delete FAQ Entries

**Request**:

```curl
curl --location --request DELETE 'http://localhost:8080/api/v1/knowledge-bases/kb-00000001/faq/entries' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json' \
--data '{
    "ids": ["faq-00000001", "faq-00000002"]
}'
```

**Response**:

```json
{
    "success": true
}
```

## POST `/knowledge-bases/:id/faq/search` - Hybrid Search FAQ

**Request Parameters**:
- `query_text`: Search query text
- `vector_threshold`: Vector similarity threshold (0-1)
- `match_count`: Number of results to return (max 200)

**Request**:

```curl
curl --location 'http://localhost:8080/api/v1/knowledge-bases/kb-00000001/faq/search' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json' \
--data '{
    "query_text": "How to reset password",
    "vector_threshold": 0.5,
    "match_count": 10
}'
```

**Response**:

```json
{
    "data": [
        {
            "id": "faq-00000001",
            "chunk_id": "chunk-00000001",
            "knowledge_id": "knowledge-00000001",
            "knowledge_base_id": "kb-00000001",
            "tag_id": "tag-00000001",
            "is_enabled": true,
            "standard_question": "How to reset password?",
            "similar_questions": ["What to do if forgot password", "Password recovery"],
            "answers": ["You can reset your password by clicking the 'Forgot Password' link on the login page."],
            "chunk_type": "faq",
            "score": 0.95,
            "match_type": "vector",
            "created_at": "2025-08-12T10:00:00+08:00",
            "updated_at": "2025-08-12T10:00:00+08:00"
        }
    ],
    "success": true
}
```
