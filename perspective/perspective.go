package perspective

import (
	"bytes"
	"encoding/json"
	"net/http"
)

//GetToxicity calculates the toxicity of a given string, returning a float64 of the decimal percentage
func GetToxicity(comment string) float64 {
	var jsonStr = []byte(`{"comment":{"text":"` + comment + `"},"requested_attributes":{"TOXICITY":{}},"client_token":"issecretshh"}`)
	resp, _ := http.Post("http://www.perspectiveapi.com/check", "application/json", bytes.NewBuffer(jsonStr))
	var toxicity struct {
		AttributeScores struct {
			TOXICITY struct {
				SpanScores []struct {
					Score struct {
						Value float64 `json:"value"`
						Type  string  `json:"type"`
					} `json:"score"`
				} `json:"spanScores"`
				SummaryScore struct {
					Value float64 `json:"value"`
					Type  string  `json:"type"`
				} `json:"summaryScore"`
			} `json:"TOXICITY"`
		} `json:"attributeScores"`
		Languages   []string `json:"languages"`
		ClientToken string   `json:"clientToken"`
	}
	json.NewDecoder(resp.Body).Decode(&toxicity)
	return toxicity.AttributeScores.TOXICITY.SummaryScore.Value
}
