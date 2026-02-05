package provider

import (
	"fmt"

	"github.com/Tencent/WeKnora/internal/types"
)

const (
	// QiniuBaseURL Qiniu API BaseURL (OpenAI compatible mode)
	QiniuBaseURL = "https://api.qnaigc.com/v1"
)

// QiniuProvider implements Qiniu Provider interface
type QiniuProvider struct{}

func init() {
	Register(&QiniuProvider{})
}

// Info returns Qiniu provider metadata
func (p *QiniuProvider) Info() ProviderInfo {
	return ProviderInfo{
		Name:        ProviderQiniu,
		DisplayName: "Qiniu",
		Description: "deepseek/deepseek-v3.2-251201, z-ai/glm-4.7, etc.",
		DefaultURLs: map[types.ModelType]string{
			types.ModelTypeKnowledgeQA: QiniuBaseURL,
		},
		ModelTypes: []types.ModelType{
			types.ModelTypeKnowledgeQA,
		},
		RequiresAuth: true,
	}
}

// ValidateConfig validates Qiniu provider configuration
func (p *QiniuProvider) ValidateConfig(config *Config) error {
	if config.BaseURL == "" {
		return fmt.Errorf("base URL is required for Qiniu provider")
	}
	if config.APIKey == "" {
		return fmt.Errorf("API key is required for Qiniu provider")
	}
	if config.ModelName == "" {
		return fmt.Errorf("model name is required")
	}
	return nil
}
