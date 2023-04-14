package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "app/dict"
)

func main() {

	conn, err := grpc.Dial("localhost:9001",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("error to connect: %+v", err)
	}

	c := pb.NewTranslateServiceClient(conn)

	res, err := c.Dictionary(
		context.Background(),
		&pb.DictionaryRequest{
			Word: "Apple",
		},
	)
	if err != nil {
		log.Println("error to Dictionary:", err)
		return
	}

	fmt.Println(res.Definition)


	data  := map[string]string{
		"Onion": "Piyoz",
		"Tomato": "Pamidor",
		"Cucumber": "Bodring",
	}


	for key, val := range data{
		
		_, err  := c.CreateDictionary(
			context.Background(),
			&pb.CreateDictionaryRequest{
				Word: key,
				Definition: val,
			},	
		)
		if err != nil {
			log.Fatalf("error to create: %+v", err)
		}

	}
	


	resp, err  := c.GetDictionary(
		context.Background(),
		&pb.GetDictionaryRequest{
			Key: "Onion",
		},	
	)
	if err != nil {
		log.Fatalf("error to create: %+v", err)
	}
	fmt.Println(resp.Value)


	resSum, err := c.SumArr(
		context.Background(),
		&pb.SumArrRequest{Data: []int32{ 1, 2, 3, 4, 5}},
	)
	fmt.Println(resSum.Sum)


	curencyres, err := c.Currency(
		context.Background(),
		&pb.CurrencyRequest{ValuteName: "USD"},	
	)
	if err != nil {
		log.Fatal("errors when getting currency")
	}
	fmt.Println(curencyres)

	
	max, err := c.Max(
		context.Background(),
		&pb.MaxRequest{Data: []int32{1,2,43,5,98}},	
	)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(max.Max)

	
	converter, err := c.Pdf(
		context.Background(),
		&pb.PdfReq{RawPath: "/home/muhiddin/univer_tasks/1-2_amaliy.docx"},
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(converter.ReadyPath)
}
