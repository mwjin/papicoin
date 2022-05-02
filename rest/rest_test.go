package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mwjjeong/papicoin/blockchain"
	"github.com/stretchr/testify/assert"
)

func TestDocumentHandler(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := createHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var data []urlDescription
	err := json.Unmarshal(res.Body.Bytes(), &data)
	assert.Nil(err)

	url, _ := url("/").MarshalText()
	assert.Equal(url, []byte(data[0].URL))
	assert.Equal("GET", data[0].Method)
	assert.Equal("See Documentation", data[0].Description)
	assert.Equal("", data[0].Payload)
}

func TestGetBlockList(t *testing.T) {
	blockchain.GetBlockchain().AddBlock("Test Block 1")
	blockchain.GetBlockchain().AddBlock("Test Block 2")

	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/blocks", nil)

	mux := createHandler()
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)

	var testBlocks []blockchain.Block
	err := json.Unmarshal(res.Body.Bytes(), &testBlocks)
	assert.Nil(err)

	assert.Equal("Test Block 1", testBlocks[len(testBlocks)-2].Data)
	assert.Equal("Test Block 2", testBlocks[len(testBlocks)-1].Data)
	assert.Equal(len(testBlocks), testBlocks[len(testBlocks)-1].Height)
}

func TestPostBlock(t *testing.T) {
	testReqBody, _ := json.Marshal(blocksPostReqBody{"My block"})

	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/blocks", bytes.NewReader(testReqBody))

	mux := createHandler()
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusCreated, res.Code)

	blocks := blockchain.GetBlockchain().GetAllBlocks()
	assert.Equal("My block", blocks[len(blocks)-1].Data)
}

func TestGetBlock(t *testing.T) {
	blockchain.GetBlockchain().AddBlock("New Test Block")
	height := len(blockchain.GetBlockchain().GetAllBlocks())

	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", fmt.Sprintf("/blocks/%d", height), nil)

	mux := createHandler()
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)

	var block blockchain.Block
	err := json.Unmarshal(res.Body.Bytes(), &block)
	assert.Nil(err)

	assert.Equal(height, block.Height)
	assert.Equal("New Test Block", block.Data)
}
