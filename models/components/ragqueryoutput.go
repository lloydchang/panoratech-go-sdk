// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type RagQueryOutput struct {
	// The chunk which matches the embed query
	Chunk string `json:"chunk"`
	// The metadata tied to the chunk
	Metadata map[string]any `json:"metadata"`
	// The score
	Score *float64 `json:"score"`
	// The embedding of the relevant chunk
	Embedding []float64 `json:"embedding"`
}

func (o *RagQueryOutput) GetChunk() string {
	if o == nil {
		return ""
	}
	return o.Chunk
}

func (o *RagQueryOutput) GetMetadata() map[string]any {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *RagQueryOutput) GetScore() *float64 {
	if o == nil {
		return nil
	}
	return o.Score
}

func (o *RagQueryOutput) GetEmbedding() []float64 {
	if o == nil {
		return nil
	}
	return o.Embedding
}
