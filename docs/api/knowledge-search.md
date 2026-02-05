# Knowledge Search API

[Back to Index](./README.md)

| Method | Path               | Description     |
| ------ | ------------------ | --------------- |
| POST   | `/knowledge-search` | Knowledge search |

## POST `/knowledge-search` - Knowledge Search

Search for related content in knowledge bases (without LLM summarization), directly returning retrieval results.

**Request Parameters**:
- `query`: Search query text (required)
- `knowledge_base_id`: Single knowledge base ID (backward compatible)
- `knowledge_base_ids`: Knowledge base ID list (supports multi-knowledge base search)
- `knowledge_ids`: Specified knowledge (file) ID list

**Request**:

```curl
# Search single knowledge base
curl --location 'http://localhost:8080/api/v1/knowledge-search' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json' \
--data '{
    "query": "How to use knowledge base",
    "knowledge_base_id": "kb-00000001"
}'

# Search multiple knowledge bases
curl --location 'http://localhost:8080/api/v1/knowledge-search' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json' \
--data '{
    "query": "How to use knowledge base",
    "knowledge_base_ids": ["kb-00000001", "kb-00000002"]
}'

# Search specified files
curl --location 'http://localhost:8080/api/v1/knowledge-search' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json' \
--data '{
    "query": "How to use knowledge base",
    "knowledge_ids": ["4c4e7c1a-09cf-485b-a7b5-24b8cdc5acf5"]
}'
```

**Response**:

```json
{
    "data": [
        {
            "id": "chunk-00000001",
            "content": "Knowledge base is a system for storing and retrieving knowledge...",
            "knowledge_id": "knowledge-00000001",
            "chunk_index": 0,
            "knowledge_title": "Knowledge Base Usage Guide",
            "start_at": 0,
            "end_at": 500,
            "seq": 1,
            "score": 0.95,
            "chunk_type": "text",
            "image_info": "",
            "metadata": {},
            "knowledge_filename": "guide.pdf",
            "knowledge_source": "file"
        }
    ],
    "success": true
}
```
