package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Ownership struct {
	ID            string `json:"id"`
	Blockchain    string `json:"blockchain"`
	ItemID        string `json:"itemId"`
	Contract      string `json:"contract"`
	Collection    string `json:"collection"`
	TokenID       string `json:"tokenId"`
	Owner         string `json:"owner"`
	Value         string `json:"value"`
	LazyValue     string `json:"lazyValue"`
	CreatedAt     string `json:"createdAt"`
	LastUpdatedAt string `json:"lastUpdatedAt"`

	BestSellOrder *Order `json:"bestSellOrder,omitempty"`
	BestBidOrder  *Order `json:"bestBidOrder,omitempty"`
}

type Order struct {
	ID            string              `json:"id"`
	Fill          string              `json:"fill"`
	Platform      string              `json:"platform"`
	Status        string              `json:"status"`
	MakeStock     string              `json:"makeStock"`
	Cancelled     bool                `json:"cancelled"`
	CreatedAt     string              `json:"createdAt"`
	LastUpdatedAt string              `json:"lastUpdatedAt"`
	DbUpdatedAt   string              `json:"dbUpdatedAt"`
	AuctionId     string              `json:"auctionId,omitempty"`
	MakePrice     string              `json:"makePrice,omitempty"`
	TakePrice     string              `json:"takePrice,omitempty"`
	MakePriceUsd  string              `json:"makePriceUsd,omitempty"`
	TakePriceUsd  string              `json:"takePriceUsd,omitempty"`
	PriceHistory  []PriceHistoryEntry `json:"priceHistory,omitempty"`
	Make          AssetType           `json:"make"`
	Take          AssetType           `json:"take"`
	Maker         string              `json:"maker"`
	Taker         string              `json:"taker,omitempty"`
	Salt          string              `json:"salt"`
	Signature     string              `json:"signature,omitempty"`
	Start         int64               `json:"start,omitempty"`
	End           int64               `json:"end,omitempty"`
	Data          OrderData           `json:"data"`
}

type AssetType struct {
	Type  AssetClass `json:"type"`
	Value string     `json:"value"`
}

type AssetClass struct {
	AssetClass string `json:"assetClass"`
	Contract   string `json:"contract,omitempty"`
	TokenId    string `json:"tokenId,omitempty"`
}

type OrderData struct {
	DataType   string   `json:"dataType"`
	PayOuts    []PayOut `json:"payOuts,omitempty"`
	OriginFees []PayOut `json:"originFees,omitempty"`
}

type PayOut struct {
	Account string `json:"account"`
	Value   int    `json:"value"`
}

type PriceHistoryEntry struct {
	Date  string `json:"date"`
	Value string `json:"value"`
}

type TraitProperty struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type TraitRarityRequest struct {
	CollectionID string          `json:"collectionID"`
	Properties   []TraitProperty `json:"properties"`
}

type ExtendedTraitProperty struct {
	Key    string  `json:"key"`
	Value  string  `json:"value"`
	Rarity float64 `json:"rarity"`
}

type TraitsRarityResponse struct {
	Continuation *string                 `json:"continuation,omitempty"`
	Traits       []ExtendedTraitProperty `json:"traits"`
}

func (o *Ownership) GetCreatedAt() (time.Time, error) {
	return time.Parse(time.RFC3339, o.CreatedAt)
}

func (o *Ownership) GetLastUpdatedAt() (time.Time, error) {
	return time.Parse(time.RFC3339, o.LastUpdatedAt)
}

func RetrieveOwnershipByID(client *http.Client, ownershipID string) (*Ownership, error) {
	url := fmt.Sprintf("https://api.rarible.org/v0.1/ownerships/%s", ownershipID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("X-API-KEY", os.Getenv("RARIBLE_API_KEY"))
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: status %d, body: %s", resp.StatusCode, string(body))
	}

	var ownership Ownership
	if err := json.Unmarshal(body, &ownership); err != nil {
		return nil, fmt.Errorf("JSON decode failed: %w", err)
	}

	return &ownership, nil
}

func RetrieveTraitsRarity(client *http.Client, reqData TraitRarityRequest) (*TraitsRarityResponse, error) {
	url := "https://api.rarible.org/v0.1/items/traits/rarity"

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request data: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-API-KEY", os.Getenv("RARIBLE_API_KEY"))

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: status %d, body: %s", resp.StatusCode, string(body))
	}

	var traitsResp TraitsRarityResponse
	if err := json.Unmarshal(body, &traitsResp); err != nil {
		return nil, fmt.Errorf("JSON decode failed: %w", err)
	}

	return &traitsResp, nil
}
