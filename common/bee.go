package common

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	baseURL = "http://localhost:1633"
)

type BeeClient struct {
	baseURL   string
	postageId string
}

func NewBeeClient(baseURL, postageId string) *BeeClient {
	return &BeeClient{
		baseURL:   baseURL,
		postageId: postageId,
	}
}

type BytesUploadResponse struct {
	Reference string `json:"reference"`
}

func (c *BeeClient) UploadBytes(data []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/bytes", c.baseURL)
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to upload to Swarm bytes EP: %v", err)
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("swarm-postage-batch-id", c.postageId)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to upload to Swarm bytes EP: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected response code: %d, body: %s", resp.StatusCode, body)
	}

	var uploadResponse BytesUploadResponse
	if err := json.NewDecoder(resp.Body).Decode(&uploadResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	swarmRef, err := hex.DecodeString(uploadResponse.Reference)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return swarmRef, nil
}

func (c *BeeClient) UploadSoc(data []byte) ([]byte, error) {
	// Parse the base URL
	parsedURL, err := url.Parse(fmt.Sprintf("%s/soc", c.baseURL))
	if err != nil {
		return nil, fmt.Errorf("Error parsing URL: %v", err)
	}
	// Create a url.Values object and add query parameters
	socSigOffset := 32
	socSigLen := 65
	socSig := data[socSigOffset : socSigOffset+socSigLen]
	params := url.Values{}
	params.Add("sig", hex.EncodeToString(socSig))
	parsedURL.RawQuery = params.Encode()

	req, err := http.NewRequest("POST", parsedURL.String(), bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to upload to Swarm SOC EP: %v", err)
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("swarm-postage-batch-id", c.postageId)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to upload to Swarm SOC EP: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected response code: %d, body: %s", resp.StatusCode, body)
	}

	var uploadResponse BytesUploadResponse
	if err := json.NewDecoder(resp.Body).Decode(&uploadResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	swarmRef, err := hex.DecodeString(uploadResponse.Reference)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return swarmRef, nil
}
