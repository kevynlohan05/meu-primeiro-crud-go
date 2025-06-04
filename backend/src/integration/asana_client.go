package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"

	ticketModel "github.com/kevynlohan05/meu-primeiro-crud-go/src/model/ticket"
)

var (
	ASANA_ACCESS_TOKEN = "ASANA_ACCESS_TOKEN"
)

// Structure for task payload
type AsanaTaskRequest struct {
	Data AsanaTaskData `json:"data"`
}

type AsanaTaskData struct {
	Name     string   `json:"name"`
	Notes    string   `json:"notes"`
	Projects []string `json:"projects"`
}

// Create a task in Asana
func CreateAsanaTask(ticket ticketModel.TicketDomainInterface) (string, error) {
	projectID := ticket.GetAsanaProjectID()
	if projectID == "" {
		return "", fmt.Errorf("asanaProjectID is empty")
	}

	token := os.Getenv(ASANA_ACCESS_TOKEN)

	// Compose the full description with all information
	notes := fmt.Sprintf(
		"Request title: \n%s\n\nRequester email: \n%s\n\nRequester sector: \n%s\n\nRequest details: \n%s\n\nRequest type: \n%s\n\nPriority: \n%s",
		ticket.GetTitle(),
		ticket.GetRequestUser(),
		ticket.GetSector(),
		ticket.GetDescription(),
		ticket.GetRequestType(),
		ticket.GetPriority(),
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

	// If error occurred
	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("error creating task: %s", string(body))
	}

	// Parse the created task ID
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	data := result["data"].(map[string]interface{})
	taskID := data["gid"].(string)

	return taskID, nil
}

type AsanaTaskResponse struct {
	Data struct {
		Gid      string `json:"gid"`
		Name     string `json:"name"`
		Notes    string `json:"notes"`
		Projects []struct {
			Gid  string `json:"gid"`
			Name string `json:"name"`
		} `json:"projects"`
		Memberships []struct {
			Section struct {
				Gid  string `json:"gid"`
				Name string `json:"name"`
			} `json:"section"`
		} `json:"memberships"`
	} `json:"data"`
}

func GetAsanaTaskDetails(taskID string) (string, []string, error) {
	token := os.Getenv("ASANA_ACCESS_TOKEN")

	req, _ := http.NewRequest("GET", "https://app.asana.com/api/1.0/tasks/"+taskID, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", nil, fmt.Errorf("error fetching task: %s", string(body))
	}

	var result AsanaTaskResponse
	json.Unmarshal(body, &result)

	// Current section name
	status := "Undefined"
	if len(result.Data.Memberships) > 0 {
		status = result.Data.Memberships[0].Section.Name
	}

	return status, nil, nil
}

func UploadAttachmentToAsana(taskID string, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Detect Content-Type based on file extension
	contentType := mime.TypeByExtension(filepath.Ext(file.Name()))
	if contentType == "" {
		contentType = "application/octet-stream" // generic fallback
	}

	// Create part header with appropriate Content-Type
	header := textproto.MIMEHeader{}
	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, filepath.Base(file.Name())))
	header.Set("Content-Type", contentType)

	part, err := writer.CreatePart(header)
	if err != nil {
		return fmt.Errorf("error creating file part: %v", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return fmt.Errorf("error copying file content: %v", err)
	}

	writer.Close()

	req, err := http.NewRequest("POST", "https://app.asana.com/api/1.0/tasks/"+taskID+"/attachments", body)
	if err != nil {
		return fmt.Errorf("error creating upload request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv(ASANA_ACCESS_TOKEN))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("upload failed: %s", string(respBody))
	}

	return nil
}
