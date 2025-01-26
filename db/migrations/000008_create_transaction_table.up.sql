CREATE TABLE transactions (
    transaction_id INT AUTO_INCREMENT PRIMARY KEY,
    register_id VARCHAR(10),
    registration_fee DECIMAL(10, 2) NOT NULL,
    examination_fee DECIMAL(10, 2) NOT NULL,
    total_fee DECIMAL(10, 2) NOT NULL,
    discount DECIMAL(10, 2) DEFAULT 0.00,
    total_payment DECIMAL(10, 2) NOT NULL,
    payment_type ENUM('cash', 'qris', 'bank_transfer', 'credit_card') DEFAULT 'cash',
    payment_status ENUM('PAID', 'UNPAID') DEFAULT 'UNPAID',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (register_id) REFERENCES register(register_id)
);
