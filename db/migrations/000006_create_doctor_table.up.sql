CREATE TABLE doctor (
    doctor_id VARCHAR(10) PRIMARY KEY,
    clinic_id VARCHAR(10),
    doctor_name VARCHAR(255) NOT NULL,
    specialization VARCHAR(100) NOT NULL,
    days VARCHAR(255) NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    phone_number VARCHAR(20) NULL,
    FOREIGN KEY (clinic_id) REFERENCES polyclinic(clinic_id)
);
