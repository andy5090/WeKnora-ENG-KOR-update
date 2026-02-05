# Tag Management API

[Back to Index](./README.md)

| Method   | Path                                  | Description                    |
| -------- | ------------------------------------- | ------------------------------ |
| GET      | `/knowledge-bases/:id/tags`           | List knowledge base tags       |
| POST     | `/knowledge-bases/:id/tags`           | Create tag                     |
| PUT      | `/knowledge-bases/:id/tags/:tag_id`   | Update tag                     |
| DELETE   | `/knowledge-bases/:id/tags/:tag_id`   | Delete tag                     |

## GET `/knowledge-bases/:id/tags` - List Knowledge Base Tags

**Query Parameters**:
- `page`: Page number (default 1)
- `page_size`: Items per page (default 20)
- `keyword`: Tag name keyword search (optional)

**Request**:

```curl
curl --location 'http://localhost:8080/api/v1/knowledge-bases/kb-00000001/tags?page=1&page_size=10' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json'
```

**Response**:

```json
{
    "data": {
        "total": 2,
        "page": 1,
        "page_size": 10,
        "data": [
            {
                "id": "tag-00000001",
                "tenant_id": 1,
                "knowledge_base_id": "kb-00000001",
                "name": "Technical Documentation",
                "color": "#1890ff",
                "sort_order": 1,
                "created_at": "2025-08-12T10:00:00+08:00",
                "updated_at": "2025-08-12T10:00:00+08:00",
                "knowledge_count": 5,
                "chunk_count": 120
            },
            {
                "id": "tag-00000002",
                "tenant_id": 1,
                "knowledge_base_id": "kb-00000001",
                "name": "FAQ",
                "color": "#52c41a",
                "sort_order": 2,
                "created_at": "2025-08-12T10:00:00+08:00",
                "updated_at": "2025-08-12T10:00:00+08:00",
                "knowledge_count": 3,
                "chunk_count": 45
            }
        ]
    },
    "success": true
}
```

## POST `/knowledge-bases/:id/tags` - Create Tag

**Request**:

```curl
curl --location 'http://localhost:8080/api/v1/knowledge-bases/kb-00000001/tags' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Product Manual",
    "color": "#faad14",
    "sort_order": 3
}'
```

**Response**:

```json
{
    "data": {
        "id": "tag-00000003",
        "tenant_id": 1,
        "knowledge_base_id": "kb-00000001",
        "name": "Product Manual",
        "color": "#faad14",
        "sort_order": 3,
        "created_at": "2025-08-12T11:00:00+08:00",
        "updated_at": "2025-08-12T11:00:00+08:00"
    },
    "success": true
}
```

## PUT `/knowledge-bases/:id/tags/:tag_id` - Update Tag

**Request**:

```curl
curl --location --request PUT 'http://localhost:8080/api/v1/knowledge-bases/kb-00000001/tags/tag-00000003' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Product Manual Updated",
    "color": "#ff4d4f"
}'
```

**Response**:

```json
{
    "data": {
        "id": "tag-00000003",
        "tenant_id": 1,
        "knowledge_base_id": "kb-00000001",
        "name": "Product Manual Updated",
        "color": "#ff4d4f",
        "sort_order": 3,
        "created_at": "2025-08-12T11:00:00+08:00",
        "updated_at": "2025-08-12T11:30:00+08:00"
    },
    "success": true
}
```

## DELETE `/knowledge-bases/:id/tags/:tag_id` - Delete Tag

**Query Parameters**:
- `force`: Set to `true` to force delete (even if tag is referenced)

**Request**:

```curl
curl --location --request DELETE 'http://localhost:8080/api/v1/knowledge-bases/kb-00000001/tags/tag-00000003?force=true' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json'
```

**Response**:

```json
{
    "success": true
}
```
