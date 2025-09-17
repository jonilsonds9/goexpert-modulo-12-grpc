package grpcServer

import (
	"database/sql"
	"net"

	"github.com/jonilsonds9/goexpert-modulo-12-grpc/internal/database"
	"github.com/jonilsonds9/goexpert-modulo-12-grpc/internal/pb"
	"github.com/jonilsonds9/goexpert-modulo-12-grpc/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := services.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
