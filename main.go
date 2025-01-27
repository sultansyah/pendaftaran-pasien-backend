package main

import (
	"log"
	"os"
	"pendaftaran-pasien-backend/internal/config"
	"pendaftaran-pasien-backend/internal/doctor"
	loginhistory "pendaftaran-pasien-backend/internal/login_history"
	"pendaftaran-pasien-backend/internal/middleware"
	"pendaftaran-pasien-backend/internal/patient"
	"pendaftaran-pasien-backend/internal/polyclinic"
	"pendaftaran-pasien-backend/internal/queue"
	"pendaftaran-pasien-backend/internal/register"
	"pendaftaran-pasien-backend/internal/token"
	"pendaftaran-pasien-backend/internal/transaction"
	"pendaftaran-pasien-backend/internal/user"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	dbConfig := config.DBConfig{
		User:     dbUser,
		Password: dbPassword,
		Host:     dbHost,
		Port:     dbPort,
		Name:     dbName,
	}

	db, err := config.InitDb(dbConfig)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api/v1")

	tokenService := token.NewTokenService([]byte(jwtSecretKey))

	loginHistoryRepository := loginhistory.NewLoginHistoryRepository()

	userRepository := user.NewUserRepository()
	userService := user.NewUserService(db, userRepository, tokenService, loginHistoryRepository)
	userHandler := user.NewUserHandler(userService)

	polyclinicRepository := polyclinic.NewPolyclinicRepository()
	polyclinicService := polyclinic.NewPolyclinicService(db, polyclinicRepository)
	polyclinicHandler := polyclinic.NewPolyclinicHandler(polyclinicService)

	doctorRepository := doctor.NewDoctorRepository()
	doctorService := doctor.NewDoctorService(db, doctorRepository, polyclinicRepository)
	doctorHandler := doctor.NewDoctorHandler(doctorService)

	patientRepository := patient.NewPatientRepository()
	patientService := patient.NewPatientService(db, patientRepository)
	patientHandler := patient.NewPatientHandler(patientService)

	transactionRepository := transaction.NewTransactionRepository()
	transactionService := transaction.NewTransactionService(db, transactionRepository)
	transactionHandler := transaction.NewTransactionHandler(transactionService)

	queueRepository := queue.NewQueueRepository()
	queueService := queue.NewQueueService(db, queueRepository)
	queueHandler := queue.NewQueueHandler(queueService)

	registerRepository := register.NewRegisterRepository()
	registerService := register.NewRegisterService(db, registerRepository, polyclinicRepository, doctorRepository, patientRepository, queueRepository, transactionRepository)
	registerHandler := register.NewRegisterHandler(registerService)

	api.POST("/auth/login", userHandler.Login)
	api.POST("/auth/forgot-password", userHandler.UpdatePassword)

	api.GET("/polyclinics", middleware.AuthMiddleware(tokenService), polyclinicHandler.GetAll)
	api.GET("/polyclinics/:clinic_id", middleware.AuthMiddleware(tokenService), polyclinicHandler.GetById)
	api.POST("/polyclinics", middleware.AuthMiddleware(tokenService), polyclinicHandler.Create)
	api.PUT("/polyclinics/:clinic_id", middleware.AuthMiddleware(tokenService), polyclinicHandler.Update)
	api.DELETE("/polyclinics/:clinic_id", middleware.AuthMiddleware(tokenService), polyclinicHandler.Delete)

	api.GET("/doctors", middleware.AuthMiddleware(tokenService), doctorHandler.GetAll)
	api.GET("/doctors/:doctor_id", middleware.AuthMiddleware(tokenService), doctorHandler.GetById)
	api.GET("/polyclinics/:clinic_id/doctors", middleware.AuthMiddleware(tokenService), doctorHandler.GetByClinicId)
	// retrieve doctor data based on clinic and available day
	api.GET("/polyclinics/:clinic_id/doctors/:day", middleware.AuthMiddleware(tokenService), doctorHandler.GetByDayAndClinicId)
	api.POST("/doctors", middleware.AuthMiddleware(tokenService), doctorHandler.Create)
	api.PUT("/doctors/:doctor_id", middleware.AuthMiddleware(tokenService), doctorHandler.Update)
	api.DELETE("/doctors/:doctor_id", middleware.AuthMiddleware(tokenService), doctorHandler.Delete)

	api.GET("/patients", middleware.AuthMiddleware(tokenService), patientHandler.GetAll)
	api.GET("/patients/:medical_record_no", middleware.AuthMiddleware(tokenService), patientHandler.GetByNoMR)
	api.POST("/patients", middleware.AuthMiddleware(tokenService), patientHandler.Create)
	api.PUT("/patients/:medical_record_no", middleware.AuthMiddleware(tokenService), patientHandler.Update)
	api.DELETE("/patients/:medical_record_no", middleware.AuthMiddleware(tokenService), patientHandler.Delete)

	api.GET("/registers", middleware.AuthMiddleware(tokenService), registerHandler.GetAll)
	api.GET("/registers/:register_id", middleware.AuthMiddleware(tokenService), registerHandler.GetById)
	api.GET("/registers/:medical_record_no/latest", middleware.AuthMiddleware(tokenService), registerHandler.GetLatestByMRNo)
	api.POST("/registers", middleware.AuthMiddleware(tokenService), registerHandler.Create)
	api.PUT("/registers/:register_id", middleware.AuthMiddleware(tokenService), registerHandler.Update)
	api.DELETE("/registers/:register_id", middleware.AuthMiddleware(tokenService), registerHandler.Delete)

	api.GET("/transactions", middleware.AuthMiddleware(tokenService), transactionHandler.GetAll)
	api.GET("/transactions/:transaction_id", middleware.AuthMiddleware(tokenService), transactionHandler.GetById)
	api.GET("/patients/:medical_record_no/transactions", middleware.AuthMiddleware(tokenService), transactionHandler.GetByMedicalRecordNo)
	api.PUT("/transactions/:transaction_id", middleware.AuthMiddleware(tokenService), transactionHandler.Update)

	api.GET("/queues", middleware.AuthMiddleware(tokenService), queueHandler.GetAll)
	api.GET("/queues/:queue_id", middleware.AuthMiddleware(tokenService), queueHandler.GetById)
	api.GET("/queues/:day/day", middleware.AuthMiddleware(tokenService), queueHandler.GetAllByDay)
	api.GET("/patients/:medical_record_no/queues", middleware.AuthMiddleware(tokenService), queueHandler.GetAllByMedicalRecordNo)
	api.PUT("/queues/:queue_id", middleware.AuthMiddleware(tokenService), queueHandler.Update)

	if err := router.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
