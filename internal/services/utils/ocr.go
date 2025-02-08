package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

type OCRResponse struct {
	ParsedResults []struct {
		ParsedText string `json:"ParsedText"`
	} `json:"ParsedResults"`
	ErrorDetails          []string `json:"ErrorDetails"`
	IsErroredOnProcessing bool     `json:"IsErroredOnProcessing"`
}

func PerformOCR() (string, error) {
	file, err := os.Open("captcha.png")
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", "image.png")
	if err != nil {
		return "", fmt.Errorf("error creating form file: %v", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return "", fmt.Errorf("error copying file content: %v", err)
	}

	writer.WriteField("language", "eng")
	writer.WriteField("isOverlayRequired", "false")
	writer.WriteField("filetype", "PNG")
	writer.WriteField("detectOrientation", "true")
	writer.WriteField("scale", "true")
	writer.WriteField("OCREngine", "2")

	err = writer.Close()
	if err != nil {
		return "", fmt.Errorf("error closing writer: %v", err)
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	fmt.Println("client created")

	req, err := http.NewRequest("POST", "http://api.ocr.space/parse/image", body)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	fmt.Println("New request created")

	req.Header.Set("apikey", "K85521956488957")
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println("Request sent to OCR")

	responseBody, _ := io.ReadAll(resp.Body)
	fmt.Println("Raw response:", string(responseBody))

	var result OCRResponse
	if err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&result); err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	if result.IsErroredOnProcessing {
		return "", fmt.Errorf("API error: %v", result.ErrorDetails)
	}

	if len(result.ParsedResults) == 0 {
		return "", fmt.Errorf("no text found in image")
	}

	return result.ParsedResults[0].ParsedText, nil
}
