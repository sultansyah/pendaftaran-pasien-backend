CREATE TABLE register (
    register_id VARCHAR(10) PRIMARY KEY,
    medical_record_no VARCHAR(50),
    session_polyclinic VARCHAR(20),
    clinic_id VARCHAR(10),
    doctor_id VARCHAR(10),
    department VARCHAR(20),
    class VARCHAR(20),
    entry_method VARCHAR(30),
    admission_fee VARCHAR(30),
    medical_procedure VARCHAR(30),
    is_deleted TINYINT(1) DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (medical_record_no) REFERENCES patient(medical_record_no),
    FOREIGN KEY (clinic_id) REFERENCES polyclinic(clinic_id),
    FOREIGN KEY (doctor_id) REFERENCES doctor(doctor_id)
);