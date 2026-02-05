package types

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"sort"
	"strings"
	"time"
)

// FAQChunkMetadata defines the structure of FAQ entries in Chunk.Metadata
type FAQChunkMetadata struct {
	StandardQuestion  string         `json:"standard_question"`
	SimilarQuestions  []string       `json:"similar_questions,omitempty"`
	NegativeQuestions []string       `json:"negative_questions,omitempty"`
	Answers           []string       `json:"answers,omitempty"`
	AnswerStrategy    AnswerStrategy `json:"answer_strategy,omitempty"`
	Version           int            `json:"version,omitempty"`
	Source            string         `json:"source,omitempty"`
}

// GeneratedQuestion represents a single AI-generated question
type GeneratedQuestion struct {
	ID       string `json:"id"`       // Unique identifier, used to construct source_id
	Question string `json:"question"` // Question content
}

// DocumentChunkMetadata defines the metadata structure for document chunks
// Used to store enhanced information such as AI-generated questions
type DocumentChunkMetadata struct {
	// GeneratedQuestions stores AI-generated related questions for this chunk
	// These questions are independently indexed to improve recall rate
	GeneratedQuestions []GeneratedQuestion `json:"generated_questions,omitempty"`
}

// GetQuestionStrings returns a list of question content strings (for backward compatibility)
func (m *DocumentChunkMetadata) GetQuestionStrings() []string {
	if m == nil || len(m.GeneratedQuestions) == 0 {
		return nil
	}
	result := make([]string, len(m.GeneratedQuestions))
	for i, q := range m.GeneratedQuestions {
		result[i] = q.Question
	}
	return result
}

// DocumentMetadata parses document metadata from Chunk
func (c *Chunk) DocumentMetadata() (*DocumentChunkMetadata, error) {
	if c == nil || len(c.Metadata) == 0 {
		return nil, nil
	}
	var meta DocumentChunkMetadata
	if err := json.Unmarshal(c.Metadata, &meta); err != nil {
		return nil, err
	}
	return &meta, nil
}

// SetDocumentMetadata sets document metadata for Chunk
func (c *Chunk) SetDocumentMetadata(meta *DocumentChunkMetadata) error {
	if c == nil {
		return nil
	}
	if meta == nil {
		c.Metadata = nil
		return nil
	}
	bytes, err := json.Marshal(meta)
	if err != nil {
		return err
	}
	c.Metadata = JSON(bytes)
	return nil
}

// Normalize cleans whitespace and removes duplicates
func (m *FAQChunkMetadata) Normalize() {
	if m == nil {
		return
	}
	m.StandardQuestion = strings.TrimSpace(m.StandardQuestion)
	m.SimilarQuestions = normalizeStrings(m.SimilarQuestions)
	m.NegativeQuestions = normalizeStrings(m.NegativeQuestions)
	m.Answers = normalizeStrings(m.Answers)
	if m.Version <= 0 {
		m.Version = 1
	}
}

// FAQMetadata 解析 Chunk 中的 FAQ 元数据
func (c *Chunk) FAQMetadata() (*FAQChunkMetadata, error) {
	if c == nil || len(c.Metadata) == 0 {
		return nil, nil
	}
	var meta FAQChunkMetadata
	if err := json.Unmarshal(c.Metadata, &meta); err != nil {
		return nil, err
	}
	meta.Normalize()
	return &meta, nil
}

// SetFAQMetadata 设置 Chunk 的 FAQ 元数据
func (c *Chunk) SetFAQMetadata(meta *FAQChunkMetadata) error {
	if c == nil {
		return nil
	}
	if meta == nil {
		c.Metadata = nil
		c.ContentHash = ""
		return nil
	}
	meta.Normalize()
	bytes, err := json.Marshal(meta)
	if err != nil {
		return err
	}
	c.Metadata = JSON(bytes)
	// Calculate and set ContentHash
	c.ContentHash = CalculateFAQContentHash(meta)
	return nil
}

// CalculateFAQContentHash calculates the hash value of FAQ content
// Hash is based on: standard question + similar questions (sorted) + negative examples (sorted) + answers (sorted)
// Used for fast matching and deduplication
func CalculateFAQContentHash(meta *FAQChunkMetadata) string {
	if meta == nil {
		return ""
	}

	// Create a copy and normalize
	normalized := *meta
	normalized.Normalize()

	// Sort arrays (ensures same content produces same hash)
	similarQuestions := make([]string, len(normalized.SimilarQuestions))
	copy(similarQuestions, normalized.SimilarQuestions)
	sort.Strings(similarQuestions)

	negativeQuestions := make([]string, len(normalized.NegativeQuestions))
	copy(negativeQuestions, normalized.NegativeQuestions)
	sort.Strings(negativeQuestions)

	answers := make([]string, len(normalized.Answers))
	copy(answers, normalized.Answers)
	sort.Strings(answers)

	// Build string for hashing: standard question + similar questions + negative examples + answers
	var builder strings.Builder
	builder.WriteString(normalized.StandardQuestion)
	builder.WriteString("|")
	builder.WriteString(strings.Join(similarQuestions, ","))
	builder.WriteString("|")
	builder.WriteString(strings.Join(negativeQuestions, ","))
	builder.WriteString("|")
	builder.WriteString(strings.Join(answers, ","))

	// Calculate SHA256 hash
	hash := sha256.Sum256([]byte(builder.String()))
	return hex.EncodeToString(hash[:])
}

// AnswerStrategy defines the answer return strategy
type AnswerStrategy string

const (
	// AnswerStrategyAll returns all answers
	AnswerStrategyAll AnswerStrategy = "all"
	// AnswerStrategyRandom randomly returns one answer
	AnswerStrategyRandom AnswerStrategy = "random"
)

// FAQEntry represents an FAQ entry returned to the frontend
type FAQEntry struct {
	ID                int64          `json:"id"`
	ChunkID           string         `json:"chunk_id"`
	KnowledgeID       string         `json:"knowledge_id"`
	KnowledgeBaseID   string         `json:"knowledge_base_id"`
	TagID             int64          `json:"tag_id"`
	TagName           string         `json:"tag_name"`
	IsEnabled         bool           `json:"is_enabled"`
	IsRecommended     bool           `json:"is_recommended"`
	StandardQuestion  string         `json:"standard_question"`
	SimilarQuestions  []string       `json:"similar_questions"`
	NegativeQuestions []string       `json:"negative_questions"`
	Answers           []string       `json:"answers"`
	AnswerStrategy    AnswerStrategy `json:"answer_strategy"`
	IndexMode         FAQIndexMode   `json:"index_mode"`
	UpdatedAt         time.Time      `json:"updated_at"`
	CreatedAt         time.Time      `json:"created_at"`
	Score             float64        `json:"score,omitempty"`
	MatchType         MatchType      `json:"match_type,omitempty"`
	ChunkType         ChunkType      `json:"chunk_type"`
	// MatchedQuestion is the actual question text that was matched in FAQ search
	// Could be the standard question or one of the similar questions
	MatchedQuestion string `json:"matched_question,omitempty"`
}

// FAQEntryPayload is the payload for creating/updating FAQ entries
type FAQEntryPayload struct {
	// ID is optional, used to specify seq_id during data migration (must be less than auto-increment starting value 100000000)
	ID                *int64          `json:"id,omitempty"`
	StandardQuestion  string          `json:"standard_question"    binding:"required"`
	SimilarQuestions  []string        `json:"similar_questions"`
	NegativeQuestions []string        `json:"negative_questions"`
	Answers           []string        `json:"answers"              binding:"required"`
	AnswerStrategy    *AnswerStrategy `json:"answer_strategy,omitempty"`
	TagID             int64           `json:"tag_id"`
	TagName           string          `json:"tag_name"`
	IsEnabled         *bool           `json:"is_enabled,omitempty"`
	IsRecommended     *bool           `json:"is_recommended,omitempty"`
}

const (
	FAQBatchModeAppend  = "append"
	FAQBatchModeReplace = "replace"
)

// FAQBatchUpsertPayload represents batch import of FAQ entries
type FAQBatchUpsertPayload struct {
	Entries     []FAQEntryPayload `json:"entries"      binding:"required"`
	Mode        string            `json:"mode"         binding:"oneof=append replace"`
	KnowledgeID string            `json:"knowledge_id"`
	TaskID      string            `json:"task_id"` // Optional, auto-generates UUID if not provided
	DryRun      bool              `json:"dry_run"` // Only validate, do not actually import
}

// FAQFailedEntry represents an entry that failed to import/validate
type FAQFailedEntry struct {
	Index             int      `json:"index"`                        // Entry index in batch (0-based)
	Reason            string   `json:"reason"`                       // Failure reason
	TagName           string   `json:"tag_name,omitempty"`           // Category/tag
	StandardQuestion  string   `json:"standard_question"`            // Standard question
	SimilarQuestions  []string `json:"similar_questions,omitempty"`  // Similar questions
	NegativeQuestions []string `json:"negative_questions,omitempty"` // Negative examples
	Answers           []string `json:"answers,omitempty"`            // Answers
	AnswerAll         bool     `json:"answer_all,omitempty"`         // Whether to reply to all
	IsDisabled        bool     `json:"is_disabled,omitempty"`        // Whether disabled
}

// FAQSuccessEntry represents simple information about a successfully imported entry
type FAQSuccessEntry struct {
	Index            int    `json:"index"`              // Entry index in batch (0-based)
	SeqID            int64  `json:"seq_id"`             // Entry sequence ID after import
	TagID            int64  `json:"tag_id,omitempty"`   // Category ID (seq_id)
	TagName          string `json:"tag_name,omitempty"` // Category name
	StandardQuestion string `json:"standard_question"`  // Standard question
}

// FAQDryRunResult represents the validation result in dry_run mode
type FAQDryRunResult struct {
	TaskID        string           `json:"task_id,omitempty"` // Async task ID (returned in async mode)
	Total         int              `json:"total"`             // Total number of entries
	SuccessCount  int              `json:"success_count"`     // Number of entries that passed validation
	FailedCount   int              `json:"failed_count"`      // Number of entries that failed validation
	FailedEntries []FAQFailedEntry `json:"failed_entries"`    // Failed entry details
}

// FAQSearchRequest represents FAQ search request parameters
type FAQSearchRequest struct {
	QueryText            string  `json:"query_text"             binding:"required"`
	VectorThreshold      float64 `json:"vector_threshold"`
	MatchCount           int     `json:"match_count"`
	FirstPriorityTagIDs  []int64 `json:"first_priority_tag_ids"`  // First priority tag ID list, limits match scope, highest priority
	SecondPriorityTagIDs []int64 `json:"second_priority_tag_ids"` // Second priority tag ID list, limits match scope, lower priority than first
	OnlyRecommended      bool    `json:"only_recommended"`        // Whether to return only recommended entries
}

// UntaggedTagName is the default tag name for entries without a tag
const UntaggedTagName = "Untagged"

// FAQEntryFieldsUpdate represents field updates for a single FAQ entry
type FAQEntryFieldsUpdate struct {
	IsEnabled     *bool  `json:"is_enabled,omitempty"`
	IsRecommended *bool  `json:"is_recommended,omitempty"`
	TagID         *int64 `json:"tag_id,omitempty"`
	// More fields can be extended in the future
}

// FAQEntryFieldsBatchUpdate represents a request to batch update FAQ entry fields
// Supports two modes:
// 1. Update by entry ID: use ByID field
// 2. Update by Tag: use ByTag field to apply the same update to all entries under that tag
type FAQEntryFieldsBatchUpdate struct {
	// ByID updates by entry ID, key is entry ID (seq_id)
	ByID map[int64]FAQEntryFieldsUpdate `json:"by_id,omitempty"`
	// ByTag batch updates by tag, key is TagID (seq_id)
	ByTag map[int64]FAQEntryFieldsUpdate `json:"by_tag,omitempty"`
	// ExcludeIDs list of IDs to exclude in ByTag operation (seq_id)
	ExcludeIDs []int64 `json:"exclude_ids,omitempty"`
}

// FAQImportTaskStatus represents the import task status
type FAQImportTaskStatus string

const (
	// FAQImportStatusPending represents the pending status of the FAQ import task
	FAQImportStatusPending FAQImportTaskStatus = "pending"
	// FAQImportStatusProcessing represents the processing status of the FAQ import task
	FAQImportStatusProcessing FAQImportTaskStatus = "processing"
	// FAQImportStatusCompleted represents the completed status of the FAQ import task
	FAQImportStatusCompleted FAQImportTaskStatus = "completed"
	// FAQImportStatusFailed represents the failed status of the FAQ import task
	FAQImportStatusFailed FAQImportTaskStatus = "failed"
)

// FAQImportProgress represents the progress of an FAQ import task stored in Redis
// When Status is "completed", the result fields (SkippedCount, ImportMode, ImportedAt, DisplayStatus, ProcessingTime) are populated.
type FAQImportProgress struct {
	TaskID            string              `json:"task_id"`                       // UUID for the import task
	KBID              string              `json:"kb_id"`                         // Knowledge Base ID
	KnowledgeID       string              `json:"knowledge_id"`                  // FAQ Knowledge ID
	Status            FAQImportTaskStatus `json:"status"`                        // Task status
	Progress          int                 `json:"progress"`                      // 0-100 percentage
	Total             int                 `json:"total"`                         // Total entries to import
	Processed         int                 `json:"processed"`                     // Entries processed so far
	SuccessCount      int                 `json:"success_count"`                 // Number of successfully imported/validated entries
	FailedCount       int                 `json:"failed_count"`                  // Number of failed entries
	SkippedCount      int                 `json:"skipped_count,omitempty"`       // Number of skipped entries (e.g., duplicates)
	FailedEntries     []FAQFailedEntry    `json:"failed_entries,omitempty"`      // Failed entry details (returned directly when small amount)
	FailedEntriesURL  string              `json:"failed_entries_url,omitempty"`  // Failed entries CSV download URL (returned when large amount)
	SuccessEntries    []FAQSuccessEntry   `json:"success_entries,omitempty"`     // Success entry simple info (returned directly when small amount)
	ValidEntryIndices []int               `json:"valid_entry_indices,omitempty"` // Validated entry indices (used to skip validation on retry)
	Message           string              `json:"message"`                       // Status message
	Error             string              `json:"error"`                         // Error message if failed
	CreatedAt         int64               `json:"created_at"`                    // Task creation timestamp
	UpdatedAt         int64               `json:"updated_at"`                    // Last update timestamp
	DryRun            bool                `json:"dry_run,omitempty"`             // Whether it's dry run mode

	// Result fields (populated when Status == "completed")
	ImportMode     string    `json:"import_mode,omitempty"`     // Import mode: append or replace
	ImportedAt     time.Time `json:"imported_at,omitempty"`     // Import completion time
	DisplayStatus  string    `json:"display_status,omitempty"`  // Display status: open or close
	ProcessingTime int64     `json:"processing_time,omitempty"` // Processing time (milliseconds)
}

// FAQImportMetadata stores FAQ import task information in Knowledge.Metadata
// Deprecated: Use FAQImportProgress with Redis storage instead
type FAQImportMetadata struct {
	ImportProgress  int `json:"import_progress"` // 0-100
	ImportTotal     int `json:"import_total"`
	ImportProcessed int `json:"import_processed"`
}

// FAQImportResult stores statistical results after FAQ import completion
// This information is persistent and does not follow progress status until replaced by next import
type FAQImportResult struct {
	// Import statistics
	TotalEntries int `json:"total_entries"` // Total number of entries
	SuccessCount int `json:"success_count"` // Number of successfully imported entries
	FailedCount  int `json:"failed_count"`  // Number of failed entries
	SkippedCount int `json:"skipped_count"` // Number of skipped entries (e.g., duplicates)

	// Import mode and time information
	ImportMode string    `json:"import_mode"` // Import mode: append or replace
	ImportedAt time.Time `json:"imported_at"` // Import completion time
	TaskID     string    `json:"task_id"`     // Import task ID

	// Failed details URL (download link provided when many failed entries)
	FailedEntriesURL string `json:"failed_entries_url,omitempty"` // Failed entries CSV download URL

	// Display control
	DisplayStatus string `json:"display_status"` // Display status: open or close

	// Additional statistics
	ProcessingTime int64 `json:"processing_time"` // Processing time (milliseconds)
}

// ToJSON converts the metadata to JSON type.
func (m *FAQImportMetadata) ToJSON() (JSON, error) {
	if m == nil {
		return nil, nil
	}
	bytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return JSON(bytes), nil
}

// ToJSON converts the import result to JSON type.
func (r *FAQImportResult) ToJSON() (JSON, error) {
	if r == nil {
		return nil, nil
	}
	bytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return JSON(bytes), nil
}

// ParseFAQImportMetadata parses FAQ import metadata from Knowledge.
func ParseFAQImportMetadata(k *Knowledge) (*FAQImportMetadata, error) {
	if k == nil || len(k.Metadata) == 0 {
		return nil, nil
	}
	var metadata FAQImportMetadata
	if err := json.Unmarshal(k.Metadata, &metadata); err != nil {
		return nil, err
	}
	return &metadata, nil
}

func normalizeStrings(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	dedup := make([]string, 0, len(values))
	seen := make(map[string]struct{}, len(values))
	for _, v := range values {
		trimmed := strings.TrimSpace(v)
		if trimmed == "" {
			continue
		}
		if _, exists := seen[trimmed]; exists {
			continue
		}
		seen[trimmed] = struct{}{}
		dedup = append(dedup, trimmed)
	}
	if len(dedup) == 0 {
		return nil
	}
	return dedup
}
