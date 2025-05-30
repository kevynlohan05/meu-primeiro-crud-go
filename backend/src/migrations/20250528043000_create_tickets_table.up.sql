CREATE TABLE IF NOT EXISTS tickets (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  request_user VARCHAR(100) NOT NULL,
  sector VARCHAR(100) NOT NULL,
  description TEXT,
  request_type VARCHAR(100),
  priority VARCHAR(50),
  attachment_urls TEXT,
  status VARCHAR(50),
  asana_task_id VARCHAR(100),
  project_id INT, 
  comments TEXT, 
  CONSTRAINT fk_project
    FOREIGN KEY (project_id)
    REFERENCES projects(id)
    ON DELETE SET NULL
    ON UPDATE CASCADE
);
