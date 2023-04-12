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


	addResult, err :=  c.Add(
		context.Background(),
		&pb.AddRequest{FirsNum: 2, SecondNum: 2},
	)

	fmt.Println("add: ", addResult.Result)

	divResult, err :=  c.Div(
		context.Background(),
		&pb.DivRequest{FirsNum: 2, SecondNum: 2},
	)

	fmt.Println("div: ", divResult.Result)

	multResult, err :=  c.Mult(
		context.Background(),
		&pb.MultRequest{FirsNum: 2, SecondNum: 2},
	)

	fmt.Println("mult: ", multResult.Result)

	subResult, err :=  c.Sub(
		context.Background(),
		&pb.SubRequest{FirsNum: 2, SecondNum: 2},
	)

	fmt.Println("sub: ", subResult.Result)


	sqrtResult, err :=  c.Sqrt(
		context.Background(),
		&pb.SqrtRequest{FirsNum: 64},
	)

	fmt.Println("sqrt: ", sqrtResult.Result)

	powResult, err :=  c.Pow(
		context.Background(),
		&pb.PowRequest{FirsNum: 2, SecondNum: 3},
	)

	fmt.Println("pow: ", powResult.Result)


	minResult, err :=  c.Min(
		context.Background(),
		&pb.MinRequest{FirsNum: 5, SecondNum: 1},
	)

	fmt.Println("min: ", minResult.Result)	

}
