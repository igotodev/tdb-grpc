package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
	"tdb-grpc/pb"
	"time"
)

const (
	address = "0.0.0.0:50051"
)

var collection *mongo.Collection

type server struct{}

type noteItem struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Title string             `bson:"title"`
	Note  string             `bson:"note"`
	Time  string             `bson:"time"`
}

func (*server) CreateNotes(ctx context.Context, req *pb.CreateNoteRequest) (*pb.CreateNoteResponse, error) {
	log.Printf("starting to create a blog...")
	noteReq := req.GetNote()
	note := noteItem{
		Title: noteReq.GetTitle(),
		Note:  noteReq.GetNote(),
		Time:  time.Now().Format("2006/01/02 15:04:05"),
	}

	result, err := collection.InsertOne(ctx, note)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("internal error from mongodb: %v", err))
	}
	// get object id
	oid := result.InsertedID.(primitive.ObjectID)

	log.Printf("blog is created with id=%v", oid.Hex())
	resp := &pb.CreateNoteResponse{Note: &pb.Note{
		Id:    oid.Hex(),
		Title: note.Title,
		Note:  note.Note,
		Time:  note.Time,
	}}
	return resp, nil
}

func main() {
	log.Printf("the server is running...")
	// init mongoDB
	log.Printf("connecting to mongodb...")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("error while connect to mongodb : %v", err)
	}
	collection = client.Database("mynotesdb").Collection("notes")
	log.Printf("connection to mongodb is complete!")

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("error while listen tcp: %v", err)
	}
	defer lis.Close()

	// empty options (for use https need to edit options)
	var opts []grpc.ServerOption
	/*
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if sslErr != nil {
			log.Fatalf("error while loading certificates: %v", err)
			return
		}
		opts = append(opts, grpc.Creds(creds))
	*/
	s := grpc.NewServer(opts...)
	pb.RegisterNotesServiceServer(s, &server{})
	// Register reflection service
	reflection.Register(s)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("error while serve: %v", err)
		}
	}()

	// wait ctrl-c for stop server
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	log.Printf("stopping server...")
	s.Stop()
	log.Print("closing mongodb connection...")
	if err = client.Disconnect(ctx); err != nil {
		log.Fatalf("error on disconnection with mongodb : %v", err)
	}
	log.Print("mongodb connection closed!")
	log.Printf("end of program")
}
