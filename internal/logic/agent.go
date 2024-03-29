package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/tools"
	"github.com/tmc/langchaingo/tools/scraper"
)

func SpawnAgent(query string, website string) (output map[string]any, err error) {
	llm, err := ollama.New(ollama.WithModel("gemma:2b"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	scraper, err := scraper.New()
	if err != nil {
		log.Fatal(err)
	}
	agent := agents.NewOpenAIFunctionsAgent(llm, []tools.Tool{scraper})
	executor := agents.NewExecutor(agent, []tools.Tool{scraper})

	input := fmt.Sprintf(`
	Utilise le scraper tool pour résumer la position du parti à propos de %s, 
	en te basant sur les informations disponibles sur leur site web à l'adresse %s
	Ne décris pas les étapes à suivre pour trouver cette information, mais donne directement le résumé de la position du parti.`, website, query)

	response, err := executor.Call(ctx, map[string]any{
		"input": input})

	if err != nil {
		log.Fatal(err)
	}
	return response, nil
}

func CallAgentHook(query string, website string) (string, error) {
	var body struct {
		Query   string `json:"query"`
		Website string `json:"website"`
	}

	json.Unmarshal([]byte(fmt.Sprintf(`{"query": "%s", "website": "%s"}`, query, website)), &body)
	payload, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(os.Getenv("N8N_HOOK"), "application/json", bytes.NewReader(payload))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	respBody := string(bodyBytes)

	return respBody, nil
}
