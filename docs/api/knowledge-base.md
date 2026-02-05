# Knowledge Base Management API

[Back to Index](./README.md)

| Method   | Path                                 | Description                    |
| -------- | ------------------------------------ | ------------------------------ |
| POST     | `/knowledge-bases`                   | Create knowledge base          |
| GET      | `/knowledge-bases`                   | List knowledge bases           |
| GET      | `/knowledge-bases/:id`               | Get knowledge base details     |
| PUT      | `/knowledge-bases/:id`               | Update knowledge base          |
| DELETE   | `/knowledge-bases/:id`               | Delete knowledge base          |
| POST     | `/knowledge-bases/copy`              | Copy knowledge base            |
| GET      | `/knowledge-bases/:id/hybrid-search` | Hybrid search (vector + keyword) |

## POST `/knowledge-bases` - Create Knowledge Base

**Request**:

```curl
curl --location 'http://localhost:8080/api/v1/knowledge-bases' \
--header 'Content-Type: application/json' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--data '{
    "name": "weknora",
    "description": "weknora description",
    "chunking_config": {
        "chunk_size": 1000,
        "chunk_overlap": 200,
        "separators": [
            "."
        ],
        "enable_multimodal": true
    },
    "image_processing_config": {
        "model_id": "f2083ad7-63e3-486d-a610-e6c56e58d72e"
    },
    "embedding_model_id": "dff7bc94-7885-4dd1-bfd5-bd96e4df2fc3",
    "summary_model_id": "8aea788c-bb30-4898-809e-e40c14ffb48c",
    "rerank_model_id": "b30171a1-787b-426e-a293-735cd5ac16c0",
    "vlm_config": {
        "enabled": true,
        "model_id": "f2083ad7-63e3-486d-a610-e6c56e58d72e"
    },
    "cos_config": {
        "secret_id": "",
        "secret_key": "",
        "region": "",
        "bucket_name": "",
        "app_id": "",
        "path_prefix": ""
    }
}'
```

**Response**:

```json
{
    "data": {
        "id": "b5829e4a-3845-4624-a7fb-ea3b35e843b0",
        "name": "weknora",
        "description": "weknora description",
        "tenant_id": 1,
        "chunking_config": {
            "chunk_size": 1000,
            "chunk_overlap": 200,
            "separators": [
                "."
            ],
            "enable_multimodal": true
        },
        "image_processing_config": {
            "model_id": "f2083ad7-63e3-486d-a610-e6c56e58d72e"
        },
        "embedding_model_id": "dff7bc94-7885-4dd1-bfd5-bd96e4df2fc3",
        "summary_model_id": "8aea788c-bb30-4898-809e-e40c14ffb48c",
        "rerank_model_id": "b30171a1-787b-426e-a293-735cd5ac16c0",
        "vlm_config": {
            "enabled": true,
            "model_id": "f2083ad7-63e3-486d-a610-e6c56e58d72e"
        },
        "cos_config": {
            "secret_id": "",
            "secret_key": "",
            "region": "",
            "bucket_name": "",
            "app_id": "",
            "path_prefix": ""
        },
        "created_at": "2025-08-12T11:30:09.206238645+08:00",
        "updated_at": "2025-08-12T11:30:09.206238854+08:00",
        "deleted_at": null
    },
    "success": true
}
```

## GET `/knowledge-bases` - List Knowledge Bases

**Request**:

```curl
curl --location 'http://localhost:8080/api/v1/knowledge-bases' \
--header 'Content-Type: application/json' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ'
```

**Response**:

```json
{
    "data": [
        {
            "id": "kb-00000001",
            "name": "Default Knowledge Base",
            "description": "System Default Knowledge Base",
            "tenant_id": 1,
            "chunking_config": {
                "chunk_size": 1000,
                "chunk_overlap": 200,
                "separators": [
                    "\n\n",
                    "\n",
                    "。",
                    "！",
                    "？",
                    ";",
                    "；"
                ],
                "enable_multimodal": true
            },
            "image_processing_config": {
                "model_id": ""
            },
            "embedding_model_id": "dff7bc94-7885-4dd1-bfd5-bd96e4df2fc3",
            "summary_model_id": "8aea788c-bb30-4898-809e-e40c14ffb48c",
            "rerank_model_id": "b30171a1-787b-426e-a293-735cd5ac16c0",
            "vlm_config": {
                "enabled": true,
                "model_id": "f2083ad7-63e3-486d-a610-e6c56e58d72e"
            },
            "cos_config": {
                "secret_id": "",
                "secret_key": "",
                "region": "",
                "bucket_name": "",
                "app_id": "",
                "path_prefix": ""
            },
            "created_at": "2025-08-11T20:10:41.817794+08:00",
            "updated_at": "2025-08-12T11:23:00.593097+08:00",
            "deleted_at": null
        }
    ],
    "success": true
}
```

## GET `/knowledge-bases/:id` - Get Knowledge Base Details

**Request**:

```curl
curl --location 'http://localhost:8080/api/v1/knowledge-bases/kb-00000001' \
--header 'Content-Type: application/json' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ'
```

**Response**:

```json
{
    "data": {
        "id": "kb-00000001",
        "name": "Default Knowledge Base",
        "description": "System Default Knowledge Base",
        "tenant_id": 1,
        "chunking_config": {
            "chunk_size": 1000,
            "chunk_overlap": 200,
            "separators": [
                "\n\n",
                "\n",
                "。",
                "！",
                "？",
                ";",
                "；"
            ],
            "enable_multimodal": true
        },
        "image_processing_config": {
            "model_id": ""
        },
        "embedding_model_id": "dff7bc94-7885-4dd1-bfd5-bd96e4df2fc3",
        "summary_model_id": "8aea788c-bb30-4898-809e-e40c14ffb48c",
        "rerank_model_id": "b30171a1-787b-426e-a293-735cd5ac16c0",
        "vlm_config": {
            "enabled": true,
            "model_id": "f2083ad7-63e3-486d-a610-e6c56e58d72e"
        },
        "cos_config": {
            "secret_id": "",
            "secret_key": "",
            "region": "",
            "bucket_name": "",
            "app_id": "",
            "path_prefix": ""
        },
        "created_at": "2025-08-11T20:10:41.817794+08:00",
        "updated_at": "2025-08-12T11:23:00.593097+08:00",
        "deleted_at": null
    },
    "success": true
}
```

## PUT `/knowledge-bases/:id` - Update Knowledge Base

**Request**:

```curl
curl --location --request PUT 'http://localhost:8080/api/v1/knowledge-bases/b5829e4a-3845-4624-a7fb-ea3b35e843b0' \
--header 'Content-Type: application/json' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--data '{
    "name": "weknora new",
    "description": "weknora description new",
    "config": {
        "chunking_config": {
            "chunk_size": 1000,
            "chunk_overlap": 200,
            "separators": [
                "\n\n",
                "\n",
                "。",
                "！",
                "？",
                ";",
                "；"
            ],
            "enable_multimodal": true
        },
        "image_processing_config": {
            "model_id": ""
        }
    }
}'
```

**Response**:

```json
{
    "data": {
        "id": "b5829e4a-3845-4624-a7fb-ea3b35e843b0",
        "name": "weknora new",
        "description": "weknora description new",
        "tenant_id": 1,
        "chunking_config": {
            "chunk_size": 1000,
            "chunk_overlap": 200,
            "separators": [
                "\n\n",
                "\n",
                "。",
                "！",
                "？",
                ";",
                "；"
            ],
            "enable_multimodal": true
        },
        "image_processing_config": {
            "model_id": ""
        },
        "embedding_model_id": "dff7bc94-7885-4dd1-bfd5-bd96e4df2fc3",
        "summary_model_id": "8aea788c-bb30-4898-809e-e40c14ffb48c",
        "rerank_model_id": "b30171a1-787b-426e-a293-735cd5ac16c0",
        "vlm_config": {
            "enabled": true,
            "model_id": "f2083ad7-63e3-486d-a610-e6c56e58d72e"
        },
        "cos_config": {
            "secret_id": "",
            "secret_key": "",
            "region": "",
            "bucket_name": "",
            "app_id": "",
            "path_prefix": ""
        },
        "created_at": "2025-08-12T11:30:09.206238+08:00",
        "updated_at": "2025-08-12T11:36:09.083577609+08:00",
        "deleted_at": null
    },
    "success": true
}
```

## DELETE `/knowledge-bases/:id` - Delete Knowledge Base

**Request**:

```curl
curl --location --request DELETE 'http://localhost:8080/api/v1/knowledge-bases/b5829e4a-3845-4624-a7fb-ea3b35e843b0' \
--header 'Content-Type: application/json' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ'
```

**Response**:

```json
{
    "message": "Knowledge base deleted successfully",
    "success": true
}
```

## GET `/knowledge-bases/:id/hybrid-search` - Hybrid Search

Perform hybrid retrieval combining vector search and keyword search.

**Note**: This endpoint uses GET method but requires a JSON request body.

**Request Parameters**:
- `query_text`: Search query text (required)
- `vector_threshold`: Vector similarity threshold (0-1, optional)
- `keyword_threshold`: Keyword matching threshold (optional)
- `match_count`: Number of results to return (optional)
- `disable_keywords_match`: Whether to disable keyword matching (optional)
- `disable_vector_match`: Whether to disable vector matching (optional)

**Request**:

```curl
curl --location --request GET 'http://localhost:8080/api/v1/knowledge-bases/kb-00000001/hybrid-search' \
--header 'X-API-Key: sk-vQHV2NZI_LK5W7wHQvH3yGYExX8YnhaHwZipUYbiZKCYJbBQ' \
--header 'Content-Type: application/json' \
--data '{
    "query_text": "How to use knowledge base",
    "vector_threshold": 0.5,
    "match_count": 10
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
