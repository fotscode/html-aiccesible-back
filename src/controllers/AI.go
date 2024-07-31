package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	ct "html-aiccesible/constants"
	"html-aiccesible/httputil"
	m "html-aiccesible/models"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (b *Controller) ListModels(c *gin.Context) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s:%s/api/tags", ct.OLLAMA_HOST, ct.OLLAMA_PORT), nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		httputil.InternalServerError(c, err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		httputil.InternalServerError(c, fmt.Errorf("failed to get response from model"))
		return
	}
	defer resp.Body.Close()
	var jsonResp map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&jsonResp)
	var respArr []string
	for _, value := range jsonResp["models"].([]any) {
		value := value.(map[string]interface{})
		respArr = append(respArr, strings.Split(value["name"].(string), ":")[0])
	}
	c.JSON(http.StatusOK, respArr)
}

func (b *Controller) Accesibilize(c *gin.Context) {
	body := c.MustGet(gin.BindKey).(*m.AccesibilizeBody)
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)
	chanStream := make(chan string)
	go func() {
		defer close(chanStream)
		req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s:%s/api/generate", ct.OLLAMA_HOST, ct.OLLAMA_PORT), buf)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			httputil.InternalServerError(c, err)
			return
		}
		if resp.StatusCode != http.StatusOK {
			httputil.InternalServerError(c, fmt.Errorf("failed to get response from model"))
			return
		}
		defer resp.Body.Close()
		reader := io.Reader(resp.Body)
		for {
			buf := make([]byte, 8192)
			n, err := reader.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				httputil.InternalServerError(c, err)
				return
			}
			chanStream <- string(buf[:n])
		}
	}()

	c.Stream(func(_ io.Writer) bool {
		if msg, ok := <-chanStream; ok {
			c.SSEvent("message", msg)
			return true
		}
		return false

	})

}
