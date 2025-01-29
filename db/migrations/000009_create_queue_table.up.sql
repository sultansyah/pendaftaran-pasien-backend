CREATE TABLE queue (
    queue_id INT AUTO_INCREMENT PRIMARY KEY,
    register_id VARCHAR(10),
    queue_number INT NOT NULL,
    is_completed TINYINT(1) DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (register_id) REFERENCES register(register_id)
);
