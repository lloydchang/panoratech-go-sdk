// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type QueryBody struct {
	// The query you want to received embeddings and chunks for
	Query string `json:"query"`
	// The number of most appropriate documents for your query.
	TopK *float64 `json:"topK,omitempty"`
}

func (o *QueryBody) GetQuery() string {
	if o == nil {
		return ""
	}
	return o.Query
}

func (o *QueryBody) GetTopK() *float64 {
	if o == nil {
		return nil
	}
	return o.TopK
}
