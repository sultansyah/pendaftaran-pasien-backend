# pendaftaran-pasien-backend

## Prerequisites
- Go

## Setup Instructions

### 1. Clone Repository
```bash
git clone https://github.com/sultansyah/pendaftaran-pasien-backend.git
cd pendaftaran-pasien-backend
```

### 2. Environment Variables
Copy `.env.example` to `.env`:
```bash
cp .env.example .env
```

Update the values in `.env`:
```env
JWT_SECRET_KEY=your_jwt_secret
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
DB_HOST=localhost
DB_PORT=3306
```

### 3. Install Dependencies
```bash
go mod tidy
```

### 4. Run Server
```bash
go run main.go
```