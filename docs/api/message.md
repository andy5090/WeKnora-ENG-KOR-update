# Message Management API

[Back to Index](./README.md)

| Method   | Path                         | Description                    |
| -------- | ---------------------------- | ------------------------------ |
| GET      | `/messages/:session_id/load` | Get recent session message list |
| DELETE   | `/messages/:session_id/:id`  | Delete message                 |

## GET `/messages/:session_id/load` - Get Recent Session Message List

**Query Parameters**:

- `before_time`: The `created_at` field of the earliest message from the last pull, empty to pull recent messages
- `limit`: Items per page (default 20)

**Request**:

```curl
curl --location --request GET 'http://localhost:8080/api/v1/messages/ceb9babb-1e30-41d7-817d-fd584954304b/load?limit=3&before_time=2030-08-12T14%3A35%3A42.123456789Z' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json' \
--data '{
    "query": "Comet tail shape"
}'
```

**Response**:

```json
{
    "data": [
        {
            "id": "b8b90eeb-7dd5-4cf9-81c6-5ebcbd759451",
            "session_id": "ceb9babb-1e30-41d7-817d-fd584954304b",
            "request_id": "hCA8SDjxcAvv",
            "content": "<think>\nOkay",
            "role": "assistant",
            "knowledge_references": [
                {
                    "id": "c8347bef-127f-4a22-b962-edf5a75386ec",
                    "content": "Comet xxx",
                    "knowledge_id": "a6790b93-4700-4676-bd48-0d4804e1456b",
                    "chunk_index": 0,
                    "knowledge_title": "Comet.txt",
                    "start_at": 0,
                    "end_at": 2760,
                    "seq": 0,
                    "score": 4.038836479187012,
                    "match_type": 4,
                    "sub_chunk_id": [
                        "688821f0-40bf-428e-8cb6-541531ebeb76",
                        "c1e9903e-2b4d-4281-be15-0149288d45c2",
                        "7d955251-3f79-4fd5-a6aa-02f81e044091"
                    ],
                    "metadata": {},
                    "chunk_type": "text",
                    "parent_chunk_id": "",
                    "image_info": "",
                    "knowledge_filename": "Comet.txt",
                    "knowledge_source": ""
                },
                {
                    "id": "fa3aadee-cadb-4a84-9941-c839edc3e626",
                    "content": "# Document Name\nComet.txt\n\n# Summary\nComets are small solar system bodies composed of ice and dust. When approaching the sun, they release gas forming a coma and tail. Their orbital periods vary greatly, with sources including the Kuiper Belt and Oort Cloud. The distinction between comets and asteroids is gradually blurring, with some comets having lost volatile material, similar to asteroids. Currently, numerous comets are known, and exocomets exist. Comets were considered omens in ancient times, and modern research reveals their complex structure and origin.",
                    "knowledge_id": "a6790b93-4700-4676-bd48-0d4804e1456b",
                    "chunk_index": 6,
                    "knowledge_title": "Comet.txt",
                    "start_at": 0,
                    "end_at": 0,
                    "seq": 6,
                    "score": 0.6131043121858466,
                    "match_type": 0,
                    "sub_chunk_id": null,
                    "metadata": {},
                    "chunk_type": "summary",
                    "parent_chunk_id": "c8347bef-127f-4a22-b962-edf5a75386ec",
                    "image_info": "",
                    "knowledge_filename": "Comet.txt",
                    "knowledge_source": ""
                }
            ],
            "agent_steps": [],
            "is_completed": true,
            "created_at": "2025-08-12T10:24:38.370548+08:00",
            "updated_at": "2025-08-12T10:25:40.416382+08:00",
            "deleted_at": null
        },
        {
            "id": "7fa136ae-a045-424e-baac-52113d92ae94",
            "session_id": "ceb9babb-1e30-41d7-817d-fd584954304b",
            "request_id": "3475c004-0ada-4306-9d30-d7f5efce50d2",
            "content": "Comet tail shape",
            "role": "user",
            "knowledge_references": [],
            "agent_steps": [],
            "is_completed": true,
            "created_at": "2025-08-12T14:30:39.732246+08:00",
            "updated_at": "2025-08-12T14:30:39.733277+08:00",
            "deleted_at": null
        },
        {
            "id": "9bcafbcf-a758-40af-a9a3-c4d8e0f49439",
            "session_id": "ceb9babb-1e30-41d7-817d-fd584954304b",
            "request_id": "3475c004-0ada-4306-9d30-d7f5efce50d2",
            "content": "<think>\nOkay",
            "role": "assistant",
            "knowledge_references": [
                {
                    "id": "c8347bef-127f-4a22-b962-edf5a75386ec",
                    "content": "Comet xxx",
                    "knowledge_id": "a6790b93-4700-4676-bd48-0d4804e1456b",
                    "chunk_index": 0,
                    "knowledge_title": "Comet.txt",
                    "start_at": 0,
                    "end_at": 2760,
                    "seq": 0,
                    "score": 4.038836479187012,
                    "match_type": 3,
                    "sub_chunk_id": [
                        "688821f0-40bf-428e-8cb6-541531ebeb76",
                        "c1e9903e-2b4d-4281-be15-0149288d45c2",
                        "7d955251-3f79-4fd5-a6aa-02f81e044091"
                    ],
                    "metadata": {},
                    "chunk_type": "text",
                    "parent_chunk_id": "",
                    "image_info": "",
                    "knowledge_filename": "Comet.txt",
                    "knowledge_source": ""
                },
                {
                    "id": "fa3aadee-cadb-4a84-9941-c839edc3e626",
                    "content": "# Document Name\nComet.txt\n\n# Summary\nComets are small solar system bodies composed of ice and dust. When approaching the sun, they release gas forming a coma and tail. Their orbital periods vary greatly, with sources including the Kuiper Belt and Oort Cloud. The distinction between comets and asteroids is gradually blurring, with some comets having lost volatile material, similar to asteroids. Currently, numerous comets are known, and exocomets exist. Comets were considered omens in ancient times, and modern research reveals their complex structure and origin.",
                    "knowledge_id": "a6790b93-4700-4676-bd48-0d4804e1456b",
                    "chunk_index": 6,
                    "knowledge_title": "Comet.txt",
                    "start_at": 0,
                    "end_at": 0,
                    "seq": 6,
                    "score": 0.6131043121858466,
                    "match_type": 3,
                    "sub_chunk_id": null,
                    "metadata": {},
                    "chunk_type": "summary",
                    "parent_chunk_id": "c8347bef-127f-4a22-b962-edf5a75386ec",
                    "image_info": "",
                    "knowledge_filename": "Comet.txt",
                    "knowledge_source": ""
                }
            ],
            "agent_steps": [],
            "is_completed": true,
            "created_at": "2025-08-12T14:30:39.735108+08:00",
            "updated_at": "2025-08-12T14:31:17.829926+08:00",
            "deleted_at": null
        }
    ],
    "success": true
}
```

## DELETE `/messages/:session_id/:id` - Delete Message

**Request**:

```curl
curl --location --request DELETE 'http://localhost:8080/api/v1/messages/ceb9babb-1e30-41d7-817d-fd584954304b/9bcafbcf-a758-40af-a9a3-c4d8e0f49439' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json'
```

**Response**:

```json
{
    "message": "Message deleted successfully",
    "success": true
}
```
