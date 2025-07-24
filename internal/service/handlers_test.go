package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/casual-user-asm/raribleClient/internal/client"
	"github.com/gin-gonic/gin"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestRetrieveOwnershipByID(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.rarible.org/v0.1/ownerships/test-id",
		httpmock.NewStringResponder(200, `{
            "id": "test-id",
            "blockchain": "ETHEREUM",
            "owner": "0xowner"
        }`),
	)

	clientForFunc := &http.Client{Timeout: 10 * time.Second}
	ownership, err := client.RetrieveOwnershipByID(clientForFunc, "test-id")
	require.NoError(t, err)
	require.Equal(t, "test-id", ownership.ID)
	require.Equal(t, "ETHEREUM", ownership.Blockchain)
	require.Equal(t, "0xowner", ownership.Owner)
}

func TestTraitsHandler(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockResponse := `{
		"traits": [
			{
				"key": "color",
				"value": "red",
				"rarity": 0.05
			}
		],
		"continuation": ""
	}`

	httpmock.RegisterResponder("POST", "https://api.rarible.org/v0.1/items/traits/rarity",
		httpmock.NewStringResponder(200, mockResponse),
	)

	reqBody := client.TraitRarityRequest{
		CollectionID: "test-collection",
		Properties: []client.TraitProperty{
			{Key: "color", Value: "red"},
		},
	}
	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/traits", TraitsHandler)

	req := httptest.NewRequest("POST", "/traits", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp client.TraitsRarityResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Len(t, resp.Traits, 1)
	require.Equal(t, "color", resp.Traits[0].Key)
	require.Equal(t, "red", resp.Traits[0].Value)
	require.InDelta(t, 0.05, resp.Traits[0].Rarity, 0.001)
}
