package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os/exec"

	pb "app/dict"

	"google.golang.org/grpc"
)

  var maps = map[string]string{}

type server struct {
	pb.UnimplementedTranslateServiceServer
}


func (s *server) Dictionary(ctx context.Context, req *pb.DictionaryRequest) (*pb.DictionaryResponse, error) {

	log.Println("Dictionary >>>>>>>> ", req)
	if req.Word == "Apple" {
		return &pb.DictionaryResponse{Definition: "Olma"}, nil
	}

	return &pb.DictionaryResponse{}, nil
}


func (s *server) CreateDictionary(ctx context.Context, req *pb.CreateDictionaryRequest)  (*pb.CreateDictionaryResponse, error) {

	log.Println("createDictionary >>>>>>>> ", req)
	maps[req.Word] = req.Definition

	return &pb.CreateDictionaryResponse{}, nil
}


func (s *server) GetDictionary(ctx context.Context, req *pb.GetDictionaryRequest)  (*pb.GetDictionaryResponse, error) {

	log.Println( req)
	
	for k, v := range maps{
		if req.Key == k {
			return &pb.GetDictionaryResponse{Value: v}, nil
		}
}
			return &pb.GetDictionaryResponse{}, nil
}


func (s *server) SumArr(ctx context.Context, req *pb.SumArrRequest)  (*pb.SumArrResponse, error) {
	
	var sum int32 

	for _, v := range req.Data{
		sum += v
	}
	
	return &pb.SumArrResponse{Sum: sum}, nil	
}

func (s *server) Currency(ctx context.Context, req *pb.CurrencyRequest) (*pb.CurrencyResponse, error){
	
	conc := []pb.CurrencyResponse{}
	data, err := ioutil.ReadFile("/home/muhiddin/Documents/03_third_month/lesson6/dictionary/concData.json")
	if err != nil {
		fmt.Println(err)
	}

	erro := json.Unmarshal(data, &conc)
	if erro != nil {
		fmt.Println(erro)
	}

	for _, v := range conc {
		if req.ValuteName == v.Ccy {
		return &v, nil	
		}
	}

	return nil, fmt.Errorf("Currency not found")
}


func (s *server) Pdf(ctx context.Context, req *pb.PdfReq) (*pb.PdfRes, error){
	
	cmd := exec.Command("libreoffice","--headless", "--convert-to", "pdf", req.RawPath , "--outdir", "/home/muhiddin")
	out, err := cmd.Output()

	if err != nil {
  		fmt.Println("could not run command: ", err)
	}
	return &pb.PdfRes{ReadyPath: string(out)}, nil 
}

func (s *server) Max(ctx context.Context, req *pb.MaxRequest) (*pb.MaxResponse, error) {
	var max int32 

	// var ints []int

	// for i, v := range req.Data {
	// 	ints[i] = int(v)
	// }

	// sort.Ints(ints)

	// max = int32(ints[0])

	for i := 0; i < len(req.Data)-1; i++ {
		if req.Data[i] > req.Data[i+1]{
			max = req.Data[i]
		} else {
			max = req.Data[i+1]
		}
	}
	
	return &pb.MaxResponse{Max: max}, nil
}

func main() {

	lis, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		log.Fatalf("failed to listen: %+v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTranslateServiceServer(s, &server{})

	fmt.Println("Listen RPC server...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %+v", err)
	}

}
