package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/mariajdab/auth-service/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr        = flag.String("addr", "localhost:50055", "the address to connect to")
	name        = flag.String("name", "maria", "name")
	username    = flag.String("username", "mariaj", "username")
	phoneNumber = flag.String("phoneNumber", "+584127650341", "phoneNumber")
)

func main() {
	opts := insecure.NewCredentials()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(opts))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c1 := pb.NewAuthServiceClient(conn)

	fmt.Println("Creating the account")
	acc := &pb.Account{
		Name:        *name,
		Username:    *username,
		PhoneNumber: *phoneNumber,
	}
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	SignupWithPhoneNumberRes, err := c1.SignupWithPhoneNumber(context.Background(), &pb.SignupWithPhoneNumberRequest{Account: acc})
	if err != nil {
		log.Fatalf("could not created the account: %v", err)
	}

	fmt.Println(SignupWithPhoneNumberRes)

	fmt.Print("Please, enter the OTP code: \n")
	var code string
	_, err = fmt.Scanln(&code)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(code)

	otp := &pb.OTP{
		PhoneNumber: *phoneNumber,
		Code:        code,
	}

	VerifyPhoneNumberRes, err := c1.VerifyPhoneNumber(ctx, &pb.VerifyPhoneNumberRequest{Otp: otp})
	if err != nil {
		log.Fatalf("could not verify the number: %v", err)
	}

	fmt.Printf("Account has been verifited %v \n", VerifyPhoneNumberRes)

	fmt.Println("Login with phone number")

	fmt.Print("Enter a phone number: \n")
	var phoneN string
	_, err = fmt.Scanln(&phoneN)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	LoginWithPhoneNumberRes, err := c1.LoginWithPhoneNumber(ctx, &pb.Phone{Number: phoneN})
	if err != nil {
		log.Fatalf("The number is not registed: %v", err)
	}

	fmt.Print(LoginWithPhoneNumberRes)

	fmt.Print("Enter code: \n")
	var codeLogin string
	_, err = fmt.Scanln(&codeLogin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	ValidatePhoneNumberLoginRes, err := c1.ValidatePhoneNumberLogin(ctx, &pb.OTP{Code: codeLogin,
		PhoneNumber: phoneN})

	if err != nil {
		log.Fatalf("The code is not valit: %v", err)
	}

	fmt.Print(ValidatePhoneNumberLoginRes)

}
