package provider

import (
	"fmt"
	"strings"

	"github.com/Tencent/WeKnora/internal/types"
)

const (
	// LKEAPBaseURL Tencent Cloud Knowledge Engine Atomic Capability (LKEAP) OpenAI compatible protocol BaseURL
	LKEAPBaseURL = "https://api.lkeap.cloud.tencent.com/v1"
)

// LKEAPProvider implements Tencent Cloud LKEAP Provider interface
// Supports DeepSeek-R1, DeepSeek-V3 series models with thinking chain capability
type LKEAPProvider struct{}

func init() {
	Register(&LKEAPProvider{})
}

// Info returns LKEAP provider metadata
func (p *LKEAPProvider) Info() ProviderInfo {
	return ProviderInfo{
		Name:        ProviderLKEAP,
		DisplayName: "Tencent Cloud LKEAP",
		Description: "DeepSeek-R1, DeepSeek-V3 series models with thinking chain support",
		DefaultURLs: map[types.ModelType]string{
			types.ModelTypeKnowledgeQA: LKEAPBaseURL,
		},
		ModelTypes: []types.ModelType{
			types.ModelTypeKnowledgeQA,
		},
		RequiresAuth: true,
	}
}

// ValidateConfig validates LKEAP provider configuration
func (p *LKEAPProvider) ValidateConfig(config *Config) error {
	if config.APIKey == "" {
		return fmt.Errorf("API key is required for LKEAP provider")
	}
	if config.ModelName == "" {
		return fmt.Errorf("model name is required")
	}
	return nil
}

// IsLKEAPDeepSeekV3Model checks if it is a DeepSeek V3.x series model
// V3.x series supports controlling thinking chain switch through Thinking parameter
func IsLKEAPDeepSeekV3Model(modelName string) bool {
	return strings.Contains(strings.ToLower(modelName), "deepseek-v3")
}

// IsLKEAPDeepSeekR1Model checks if it is a DeepSeek R1 series model
// R1 series has thinking chain enabled by default
func IsLKEAPDeepSeekR1Model(modelName string) bool {
	return strings.Contains(strings.ToLower(modelName), "deepseek-r1")
}

// IsLKEAPThinkingModel checks if it is an LKEAP model that supports thinking chain
func IsLKEAPThinkingModel(modelName string) bool {
	return IsLKEAPDeepSeekR1Model(modelName) || IsLKEAPDeepSeekV3Model(modelName)
}
