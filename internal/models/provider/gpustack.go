package provider

import (
	"fmt"

	"github.com/Tencent/WeKnora/internal/types"
)

const (
	// GPUStackBaseURL GPUStack API BaseURL (OpenAI compatible mode)
	GPUStackBaseURL = "http://your_gpustack_server_url/v1-openai"
	// GPUStackRerankBaseURL GPUStack Rerank API is OpenAI compatible but has different path (/v1/rerank instead of /v1-openai/rerank)
	GPUStackRerankBaseURL = "http://your_gpustack_server_url/v1"
)

// GPUStackProvider implements GPUStack Provider interface
type GPUStackProvider struct{}

func init() {
	Register(&GPUStackProvider{})
}

// Info returns GPUStack provider metadata
func (p *GPUStackProvider) Info() ProviderInfo {
	return ProviderInfo{
		Name:        ProviderGPUStack,
		DisplayName: "GPUStack",
		Description: "Choose your deployed model on GPUStack",
		DefaultURLs: map[types.ModelType]string{
			types.ModelTypeKnowledgeQA: GPUStackBaseURL,
			types.ModelTypeEmbedding:   GPUStackBaseURL,
			types.ModelTypeRerank:      GPUStackRerankBaseURL,
			types.ModelTypeVLLM:        GPUStackBaseURL,
		},
		ModelTypes: []types.ModelType{
			types.ModelTypeKnowledgeQA,
			types.ModelTypeEmbedding,
			types.ModelTypeRerank,
			types.ModelTypeVLLM,
		},
		RequiresAuth: true, // GPUStack requires API Key
	}
}

// ValidateConfig validates GPUStack provider configuration
func (p *GPUStackProvider) ValidateConfig(config *Config) error {
	if config.BaseURL == "" {
		return fmt.Errorf("base URL is required for GPUStack provider")
	}
	if config.APIKey == "" {
		return fmt.Errorf("API key is required for GPUStack provider")
	}
	if config.ModelName == "" {
		return fmt.Errorf("model name is required")
	}
	return nil
}
