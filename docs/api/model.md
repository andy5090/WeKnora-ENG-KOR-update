# Model Management API

[Back to Index](./README.md)

| Method   | Path                    | Description              |
| -------- | ----------------------- | ------------------------ |
| POST     | `/models`               | Create model             |
| GET      | `/models`               | List models              |
| GET      | `/models/:id`           | Get model details        |
| PUT      | `/models/:id`           | Update model             |
| DELETE   | `/models/:id`           | Delete model             |
| GET      | `/models/providers`     | List model providers     |

## Provider Support

WeKnora supports multiple mainstream AI model providers. When creating models, you can specify the provider type through the `provider` field for better compatibility.

### Supported Provider List

| Provider ID    | Name                         | Supported Model Types            |
| -------------- | ---------------------------- | -------------------------------- |
| `generic`      | Custom (OpenAI compatible)   | Chat, Embedding, Rerank, VLLM   |
| `openai`       | OpenAI                       | Chat, Embedding, Rerank, VLLM   |
| `aliyun`       | Alibaba Cloud DashScope      | Chat, Embedding, Rerank, VLLM   |
| `zhipu`        | Zhipu BigModel               | Chat, Embedding, Rerank, VLLM   |
| `volcengine`   | ByteDance Volcengine          | Chat, Embedding, VLLM           |
| `hunyuan`      | Tencent Hunyuan               | Chat, Embedding                 |
| `deepseek`     | DeepSeek                      | Chat                            |
| `minimax`      | MiniMax                       | Chat                            |
| `mimo`         | Xiaomi MiMo                   | Chat                            |
| `siliconflow`  | SiliconFlow                   | Chat, Embedding, Rerank, VLLM   |
| `jina`         | Jina                          | Embedding, Rerank               |
| `openrouter`   | OpenRouter                    | Chat, VLLM                      |
| `gemini`       | Google Gemini                 | Chat                            |
| `modelscope`   | ModelScope                    | Chat, Embedding, VLLM           |
| `moonshot`     | Moonshot                      | Chat, VLLM                      |
| `qianfan`      | Baidu Cloud Qianfan           | Chat, Embedding, Rerank, VLLM   |
| `qiniu`        | Qiniu Cloud                   | Chat                            |
| `longcat`      | LongCat AI                    | Chat                            |
| `gpustack`     | GPUStack                      | Chat, Embedding, Rerank, VLLM   |

## GET `/models/providers` - List Model Providers

Get the list of supported providers and configuration information based on model type.

**Request Parameters**:

| Parameter   | Type   | Required | Description                                    |
| ----------- | ------ | -------- | ---------------------------------------------- |
| model_type  | string | No       | Model type: `chat`, `embedding`, `rerank`, `vllm` |

**Request**:

```curl
# Get all providers
curl --location 'http://localhost:8080/api/v1/models/providers' \
--header 'X-API-Key: your_api_key'

# Get providers supporting Embedding type
curl --location 'http://localhost:8080/api/v1/models/providers?model_type=embedding' \
--header 'X-API-Key: your_api_key'
```

**Response**:

```json
{
    "success": true,
    "data": [
        {
            "value": "aliyun",
            "label": "Alibaba Cloud DashScope",
            "description": "qwen-plus, tongyi-embedding-vision-plus, qwen3-rerank, etc.",
            "defaultUrls": {
                "chat": "https://dashscope.aliyuncs.com/compatible-mode/v1",
                "embedding": "https://dashscope.aliyuncs.com/compatible-mode/v1",
                "rerank": "https://dashscope.aliyuncs.com/api/v1/services/rerank/text-rerank/text-rerank"
            },
            "modelTypes": ["chat", "embedding", "rerank", "vllm"]
        },
        {
            "value": "zhipu",
            "label": "Zhipu BigModel",
            "description": "glm-4.7, embedding-3, rerank, etc.",
            "defaultUrls": {
                "chat": "https://open.bigmodel.cn/api/paas/v4",
                "embedding": "https://open.bigmodel.cn/api/paas/v4/embeddings",
                "rerank": "https://open.bigmodel.cn/api/paas/v4/rerank"
            },
            "modelTypes": ["chat", "embedding", "rerank", "vllm"]
        }
    ]
}
```

## POST `/models` - Create Model

### Create Chat Model (KnowledgeQA)

**Local Ollama Model**:

```curl
curl --location 'http://localhost:8080/api/v1/models' \
--header 'Content-Type: application/json' \
--header 'X-API-Key: your_api_key' \
--data '{
    "name": "qwen3:8b",
    "type": "KnowledgeQA",
    "source": "local",
    "description": "LLM Model for Knowledge QA",
    "parameters": {
        "base_url": "",
        "api_key": ""
    }
}'
```

**Remote API Model (Specify Provider)**:

```curl
curl --location 'http://localhost:8080/api/v1/models' \
--header 'Content-Type: application/json' \
--header 'X-API-Key: your_api_key' \
--data '{
    "name": "qwen-plus",
    "type": "KnowledgeQA",
    "source": "remote",
    "description": "Alibaba Cloud Qwen Large Model",
    "parameters": {
        "base_url": "https://dashscope.aliyuncs.com/compatible-mode/v1",
        "api_key": "sk-your-dashscope-api-key",
        "provider": "aliyun"
    }
}'
```

### Create Embedding Model

**Local Ollama Model**:

```curl
curl --location 'http://localhost:8080/api/v1/models' \
--header 'Content-Type: application/json' \
--header 'X-API-Key: your_api_key' \
--data '{
    "name": "nomic-embed-text:latest",
    "type": "Embedding",
    "source": "local",
    "description": "Embedding Model",
    "parameters": {
        "base_url": "",
        "api_key": "",
        "embedding_parameters": {
            "dimension": 768,
            "truncate_prompt_tokens": 0
        }
    }
}'
```

**Remote API Model (Alibaba Cloud DashScope)**:

```curl
curl --location 'http://localhost:8080/api/v1/models' \
--header 'Content-Type: application/json' \
--header 'X-API-Key: your_api_key' \
--data '{
    "name": "text-embedding-v3",
    "type": "Embedding",
    "source": "remote",
    "description": "Alibaba Cloud Tongyi Qianwen Embedding Model",
    "parameters": {
        "base_url": "https://dashscope.aliyuncs.com/compatible-mode/v1",
        "api_key": "sk-your-dashscope-api-key",
        "provider": "aliyun",
        "embedding_parameters": {
            "dimension": 1024,
            "truncate_prompt_tokens": 0
        }
    }
}'
```

**Remote API Model (Jina AI)**:

```curl
curl --location 'http://localhost:8080/api/v1/models' \
--header 'Content-Type: application/json' \
--header 'X-API-Key: your_api_key' \
--data '{
    "name": "jina-embeddings-v3",
    "type": "Embedding",
    "source": "remote",
    "description": "Jina AI Embedding Model",
    "parameters": {
        "base_url": "https://api.jina.ai/v1",
        "api_key": "jina_your_api_key",
        "provider": "jina",
        "embedding_parameters": {
            "dimension": 1024,
            "truncate_prompt_tokens": 0
        }
    }
}'
```

### Create Rerank Model

**Remote API Model (Alibaba Cloud DashScope)**:

```curl
curl --location 'http://localhost:8080/api/v1/models' \
--header 'Content-Type: application/json' \
--header 'X-API-Key: your_api_key' \
--data '{
    "name": "gte-rerank",
    "type": "Rerank",
    "source": "remote",
    "description": "Alibaba Cloud GTE Rerank Model",
    "parameters": {
        "base_url": "https://dashscope.aliyuncs.com/api/v1/services/rerank/text-rerank/text-rerank",
        "api_key": "sk-your-dashscope-api-key",
        "provider": "aliyun"
    }
}'
```

**Remote API Model (Jina AI)**:

```curl
curl --location 'http://localhost:8080/api/v1/models' \
--header 'Content-Type: application/json' \
--header 'X-API-Key: your_api_key' \
--data '{
    "name": "jina-reranker-v2-base-multilingual",
    "type": "Rerank",
    "source": "remote",
    "description": "Jina AI Rerank Model",
    "parameters": {
        "base_url": "https://api.jina.ai/v1",
        "api_key": "jina_your_api_key",
        "provider": "jina"
    }
}'
```

### Create Vision Model (VLLM)

```curl
curl --location 'http://localhost:8080/api/v1/models' \
--header 'Content-Type: application/json' \
--header 'X-API-Key: your_api_key' \
--data '{
    "name": "qwen-vl-plus",
    "type": "VLLM",
    "source": "remote",
    "description": "Alibaba Cloud Tongyi Qianwen Vision Model",
    "parameters": {
        "base_url": "https://dashscope.aliyuncs.com/compatible-mode/v1",
        "api_key": "sk-your-dashscope-api-key",
        "provider": "aliyun"
    }
}'
```

**Response**:

```json
{
    "success": true,
    "data": {
        "id": "09c5a1d6-ee8b-4657-9a17-d3dcbd5c70cb",
        "tenant_id": 1,
        "name": "text-embedding-v3",
        "type": "Embedding",
        "source": "remote",
        "description": "Alibaba Cloud Tongyi Qianwen Embedding Model",
        "parameters": {
            "base_url": "https://dashscope.aliyuncs.com/compatible-mode/v1",
            "api_key": "sk-***",
            "provider": "aliyun",
            "embedding_parameters": {
                "dimension": 1024,
                "truncate_prompt_tokens": 0
            }
        },
        "is_default": false,
        "status": "active",
        "created_at": "2025-08-12T10:39:01.454591766+08:00",
        "updated_at": "2025-08-12T10:39:01.454591766+08:00",
        "deleted_at": null
    }
}
```

## GET `/models` - List Models

**Request**:

```curl
curl --location 'http://localhost:8080/api/v1/models' \
--header 'Content-Type: application/json' \
--header 'X-API-Key: your_api_key'
```

**Response**:

```json
{
    "success": true,
    "data": [
        {
            "id": "dff7bc94-7885-4dd1-bfd5-bd96e4df2fc3",
            "tenant_id": 1,
            "name": "text-embedding-v3",
            "type": "Embedding",
            "source": "remote",
            "description": "Alibaba Cloud Tongyi Qianwen Embedding Model",
            "parameters": {
                "base_url": "https://dashscope.aliyuncs.com/compatible-mode/v1",
                "api_key": "sk-***",
                "provider": "aliyun",
                "embedding_parameters": {
                    "dimension": 1024,
                    "truncate_prompt_tokens": 0
                }
            },
            "is_default": true,
            "status": "active",
            "created_at": "2025-08-11T20:10:41.813832+08:00",
            "updated_at": "2025-08-11T20:10:41.822354+08:00",
            "deleted_at": null
        },
        {
            "id": "8aea788c-bb30-4898-809e-e40c14ffb48c",
            "tenant_id": 1,
            "name": "qwen-plus",
            "type": "KnowledgeQA",
            "source": "remote",
            "description": "Alibaba Cloud Qwen Large Model",
            "parameters": {
                "base_url": "https://dashscope.aliyuncs.com/compatible-mode/v1",
                "api_key": "sk-***",
                "provider": "aliyun",
                "embedding_parameters": {
                    "dimension": 0,
                    "truncate_prompt_tokens": 0
                }
            },
            "is_default": true,
            "status": "active",
            "created_at": "2025-08-11T20:10:41.811761+08:00",
            "updated_at": "2025-08-11T20:10:41.825381+08:00",
            "deleted_at": null
        }
    ]
}
```

## GET `/models/:id` - Get Model Details

**Request**:

```curl
curl --location 'http://localhost:8080/api/v1/models/dff7bc94-7885-4dd1-bfd5-bd96e4df2fc3' \
--header 'Content-Type: application/json' \
--header 'X-API-Key: your_api_key'
```

**Response**:

```json
{
    "success": true,
    "data": {
        "id": "dff7bc94-7885-4dd1-bfd5-bd96e4df2fc3",
        "tenant_id": 1,
        "name": "text-embedding-v3",
        "type": "Embedding",
        "source": "remote",
        "description": "Alibaba Cloud Tongyi Qianwen Embedding Model",
        "parameters": {
            "base_url": "https://dashscope.aliyuncs.com/compatible-mode/v1",
            "api_key": "sk-***",
            "provider": "aliyun",
            "embedding_parameters": {
                "dimension": 1024,
                "truncate_prompt_tokens": 0
            }
        },
        "is_default": true,
        "status": "active",
        "created_at": "2025-08-11T20:10:41.813832+08:00",
        "updated_at": "2025-08-11T20:10:41.822354+08:00",
        "deleted_at": null
    }
}
```

## PUT `/models/:id` - Update Model

**Request**:

```curl
curl --location --request PUT 'http://localhost:8080/api/v1/models/8fdc464d-8eaa-44d4-a85b-094b28af5330' \
--header 'Content-Type: application/json' \
--header 'X-API-Key: your_api_key' \
--data '{
    "name": "gte-rerank-v2",
    "description": "Alibaba Cloud GTE Rerank Model V2",
    "parameters": {
        "base_url": "https://dashscope.aliyuncs.com/api/v1/services/rerank/text-rerank/text-rerank",
        "api_key": "sk-your-new-api-key",
        "provider": "aliyun"
    }
}'
```

**Response**:

```json
{
    "success": true,
    "data": {
        "id": "8fdc464d-8eaa-44d4-a85b-094b28af5330",
        "tenant_id": 1,
        "name": "gte-rerank-v2",
        "type": "Rerank",
        "source": "remote",
        "description": "Alibaba Cloud GTE Rerank Model V2",
        "parameters": {
            "base_url": "https://dashscope.aliyuncs.com/api/v1/services/rerank/text-rerank/text-rerank",
            "api_key": "sk-***",
            "provider": "aliyun",
            "embedding_parameters": {
                "dimension": 0,
                "truncate_prompt_tokens": 0
            }
        },
        "is_default": false,
        "status": "active",
        "created_at": "2025-08-12T10:57:39.512681+08:00",
        "updated_at": "2025-08-12T11:00:27.271678+08:00",
        "deleted_at": null
    }
}
```

## DELETE `/models/:id` - Delete Model

**Request**:

```curl
curl --location --request DELETE 'http://localhost:8080/api/v1/models/8fdc464d-8eaa-44d4-a85b-094b28af5330' \
--header 'Content-Type: application/json' \
--header 'X-API-Key: your_api_key'
```

**Response**:

```json
{
    "success": true,
    "message": "Model deleted"
}
```

## Parameter Description

### ModelType

| Value        | Description        | Usage                                    |
| ------------ | ------------------ | ---------------------------------------- |
| KnowledgeQA  | Chat Model         | Knowledge base Q&A, conversation generation |
| Embedding    | Embedding Model    | Text vectorization, knowledge base retrieval |
| Rerank       | Rerank Model       | Retrieval result reranking, relevance optimization |
| VLLM         | Vision Language Model | Multimodal analysis, image-text understanding |

### ModelSource

| Value   | Description    | Configuration Requirements                    |
| ------- | -------------- | --------------------------------------------- |
| local   | Local Model    | Requires Ollama installed and model pulled   |
| remote  | Remote API     | Requires `base_url` and `api_key`             |

### Parameters

| Field                 | Type   | Description                                    |
| --------------------- | ------ | ---------------------------------------------- |
| base_url              | string | API service address (required for remote models) |
| api_key               | string | API key (required for remote models)           |
| provider              | string | Provider identifier (optional, for selecting specific API adapter) |
| embedding_parameters  | object | Embedding model specific parameters            |
| extra_config          | object | Provider-specific extra configuration          |

### EmbeddingParameters

| Field                  | Type | Description                              |
| ---------------------- | ---- | ---------------------------------------- |
| dimension              | int  | Vector dimension (e.g., 768, 1024)      |
| truncate_prompt_tokens | int  | Truncate token count (0 means no truncation) |
