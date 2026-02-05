package provider

import (
	"fmt"

	"github.com/Tencent/WeKnora/internal/types"
)

// GenericProvider implements generic OpenAI compatible Provider interface
type GenericProvider struct{}

func init() {
	Register(&GenericProvider{})
}

// Info returns generic provider metadata
func (p *GenericProvider) Info() ProviderInfo {
	return ProviderInfo{
		Name:        ProviderGeneric,
		DisplayName: "Custom (OpenAI Compatible)",
		Description: "Generic API endpoint (OpenAI-compatible)",
		DefaultURLs: map[types.ModelType]string{}, // Users need to configure themselves
		ModelTypes: []types.ModelType{
			types.ModelTypeKnowledgeQA,
			types.ModelTypeEmbedding,
			types.ModelTypeRerank,
			types.ModelTypeVLLM,
		},
		RequiresAuth: false, // May or may not be required
	}
}

// ValidateConfig validates generic provider configuration
func (p *GenericProvider) ValidateConfig(config *Config) error {
	if config.BaseURL == "" {
		return fmt.Errorf("base URL is required for generic provider")
	}
	if config.ModelName == "" {
		return fmt.Errorf("model name is required")
	}
	return nil
}
