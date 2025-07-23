package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRetrievegOwnershipsByID(t *testing.T) {
	mockResp := Ownership{
		ID:         "ownership123",
		Blockchain: "ETHEREUM",
	}
	body, _ := json.Marshal(mockResp)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}))
	defer server.Close()

	client := server.Client()
	result, err := RetrieveOwnershipByID(client, "test-id")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if result.ID != "ownership123" {
		t.Errorf("Expected ID 'ownership123', got '%s'", result.ID)
	}
}

func TestRetrieveTraitsRarity(t *testing.T) {
	mockResp := TraitsRarityResponse{
		Traits: []ExtendedTraitProperty{
			{Key: "Hat", Value: "Halo", Rarity: 0.15},
		},
	}
	body, _ := json.Marshal(mockResp)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}))
	defer server.Close()

	client := server.Client()

	reqData := TraitRarityRequest{
		CollectionID: "ETHEREUM:0x123",
		Properties: []TraitProperty{
			{Key: "Hat", Value: "Halo"},
		},
	}

	result, err := RetrieveTraitsRarity(client, reqData)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if len(result.Traits) != 1 || result.Traits[0].Key != "Hat" {
		t.Errorf("Unexpected trait result: %+v", result.Traits)
	}
}
