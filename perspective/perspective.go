package perspective

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Request struct {
	Comment             `json:"comment"`
	RequestedAttributes struct {
		TOXICITY struct {
		} `json:"TOXICITY"`
	} `json:"requested_attributes"`
	ClientToken string `json:"client_token"`
}

type Comment struct {
	Text string `json:"text"`
}

type RequestedAttributes struct {
	TOXICITY struct{} `json:"TOXICITY"`
}

//GetToxicity calculates the toxicity of a given string, returning a float64 of the decimal percentage
func GetToxicity(comment string) float64 {
	var request struct {
		Comment struct {
			Text string `json:"text"`
		} `json:"comment"`
		RequestedAttributes struct {
			TOXICITY struct {
			} `json:"TOXICITY"`
		} `json:"requested_attributes"`
		ClientToken string `json:"client_token"`
	}
	request.Comment.Text = comment
	request.ClientToken = "issecretshh"
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(request)
	resp, _ := http.Post("http://www.perspectiveapi.com/check", "application/json", b)
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
