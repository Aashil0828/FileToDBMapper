package main

import (
	"FiletoDBMapper/models"
	"FiletoDBMapper/pb/pb"
	"FiletoDBMapper/service"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("could not load env: %v", err)
	}
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.FieldMapping{}, &models.MapDetail{}, &models.Mapping{}, &models.MasterProduct{}, &models.MasterModule{}, &models.MasterTemplate{}, &models.TenantMasterMapping{})
	server := &service.MapperServer{Db: db}
	grpcServer := grpc.NewServer()
	pb.RegisterFileToDBMapperServer(grpcServer, server)
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
	log.Print("server started on port 50051")
	grpcServer.Serve(listener)
}
