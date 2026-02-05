package provider

import (
	"fmt"

	"github.com/Tencent/WeKnora/internal/types"
)

const (
	MoonshotBaseURL = "https://api.moonshot.ai/v1"
)

// MoonshotProvider implements Moonshot AI (Kimi) Provider interface
type MoonshotProvider struct{}

func init() {
	Register(&MoonshotProvider{})
}

// Info returns Moonshot provider metadata
func (p *MoonshotProvider) Info() ProviderInfo {
	return ProviderInfo{
		Name:        ProviderMoonshot,
		DisplayName: "Moonshot AI",
		Description: "kimi-k2-turbo-preview, moonshot-v1-8k-vision-preview, etc.",
		DefaultURLs: map[types.ModelType]string{
			types.ModelTypeKnowledgeQA: MoonshotBaseURL,
			types.ModelTypeVLLM:        MoonshotBaseURL,
		},
		ModelTypes: []types.ModelType{
			types.ModelTypeKnowledgeQA,
			types.ModelTypeVLLM,
		},
		RequiresAuth: true,
	}
}

// ValidateConfig validates Moonshot provider configuration
func (p *MoonshotProvider) ValidateConfig(config *Config) error {
	if config.BaseURL == "" {
		return fmt.Errorf("base URL is required for Moonshot provider")
	}
	if config.APIKey == "" {
		return fmt.Errorf("API key is required for Moonshot provider")
	}
	if config.ModelName == "" {
		return fmt.Errorf("model name is required")
	}
	return nil
}
