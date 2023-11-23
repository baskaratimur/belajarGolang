package testingapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func GetToken() (string, error) {
	url := "https://stag-digistar.visionplus.id/api/partnership/v1/cms/users/login"
	data := map[string]string{
		"username": "innocent@mncinnovation.id",
		"password": "!nn0c3nt",
	}
	headersAuth := map[string]string{
		"Accept-Langguage": "id",
	}
	payload, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}
	for key, value := range headersAuth {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	token := result["data"].(map[string]interface{})["access_token"].(string)
	return token, nil
}

func TestUpdateBanner(t *testing.T) {
	token, err := GetToken()
	if err != nil {
		t.Fatal(err)
	}
	url := "https://stag-digistar.visionplus.id/api/partnership/v1/contents/banner/digistar/id/1"
	data := map[string]interface{}{
		"title":           "Arab Maklum Automate",
		"title_en":        "Arab Maklum Automate",
		"image_url":       "https://static.mncnow.id/images/series/aeb12d16/6698_r83.jpg",
		"next_url":        "string1",
		"image_popup":     "automate",
		"image_popup_web": "automate",
		"position":        0,
		"content": map[string]interface{}{
			"id":       59000,
			"tags":     []string{"Drama", "Religi"},
			"type":     "Series",
			"synopsis": "Cerita tentang keluarga Arab yang tinggal di kota Jakarta, dimana sang Aba ingin mempertahankan tradisi Arab dalam keluarga dan pekerjaannya di era modernisasi. 1",
		},
		"content_en": map[string]interface{}{
			"id":       59000,
			"tags":     []string{"Drama", "Religi"},
			"type":     "Series",
			"synopsis": "Arab Maklum is a story about an Arab Family living in Jakarta, where the Aba wants to maintain Arabic traditions in his family and work in the modernization era 1",
		},
	}
	payload, _ := json.Marshal(data)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept-Language", "id")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Errorf("Expected status 200; got %d", resp.StatusCode)
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	title := result["data"].(map[string]interface{})["title"].(string)
	if title != "Arab Maklum Automate" {
		t.Errorf("Expected title: Arab Maklum Automate; got %s", title)
	} else {
		fmt.Println("berhasil assert title")
	}
}
func main() {

}
