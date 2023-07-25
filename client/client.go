package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var endpoint = "http://localhost:8080/cotacao"

type Quotation struct {
	Bid string `json:"bid"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func GetQuotation() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		fmt.Println("Request error: ", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Response error: ", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read body error: ", err)
		return
	}

	fmt.Println("Status code: ", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		err = json.Unmarshal(body, &errorResponse)
		if err != nil {
			fmt.Println("Unmarshal error: ", err)
			return
		}

		fmt.Println("Error response: ", errorResponse)
		return
	}

	var quotation Quotation
	err = json.Unmarshal(body, &quotation)
	if err != nil {
		fmt.Println("Unmarshal error: ", err)
		return
	}

	fmt.Println("Quotation: ", quotation)
}

func SaveTxtQuotation(quotation Quotation) error {
	file, err := os.OpenFile("cotacao.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	text := fmt.Sprintf("Cotação do dólar: %s", quotation.Bid)
	_, err = file.WriteString(text)
	if err != nil {
		return err
	}

	return nil
}
