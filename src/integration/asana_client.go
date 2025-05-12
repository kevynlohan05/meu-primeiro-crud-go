package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
)

var (
	ASANA_PROJECT_ID   = "ASANA_PROJECT_ID"
	ASANA_ACCESS_TOKEN = "ASANA_ACCESS_TOKEN"
)

// Estrutura para o payload da task
type AsanaTaskRequest struct {
	Data AsanaTaskData `json:"data"`
}

type AsanaTaskData struct {
	Name     string   `json:"name"`
	Notes    string   `json:"notes"`
	Projects []string `json:"projects"`
}

// Criação da tarefa no Asana
func CreateAsanaTask(ticket ticketModel.TicketDomainInterface) (string, error) {
	projectID := os.Getenv(ASANA_PROJECT_ID)
	token := os.Getenv(ASANA_ACCESS_TOKEN)

	// Monta a descrição completa com todas as informações
	notes := fmt.Sprintf(
		"Solicitante: %s\n\nDescrição: \n\n%s\n\nDepartamento: %s\n\nPrioridade: %s\n\nAnexo: %s",
		ticket.GetUserEmail(),
		ticket.GetDescription(),
		ticket.GetRequestType(),
		ticket.GetPriority(),
		ticket.GetAttachmentURL(),
	)

	task := AsanaTaskRequest{
		Data: AsanaTaskData{
			Name:     ticket.GetTitle(),
			Notes:    notes,
			Projects: []string{projectID},
		},
	}

	payload, err := json.Marshal(task)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://app.asana.com/api/1.0/tasks", bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	// Se deu erro
	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("erro ao criar tarefa: %s", string(body))
	}

	// Parse do ID da task criada
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	data := result["data"].(map[string]interface{})
	taskID := data["gid"].(string)

	return taskID, nil
}
