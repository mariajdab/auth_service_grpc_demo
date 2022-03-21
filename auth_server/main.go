package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	pb "github.com/mariajdab/auth-service/auth"
	"github.com/mariajdab/auth-service/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port       = flag.Int("port", 50055, "The server port")
	configFile = flag.String("conf", "./config.json", "configuration file")
	BaseAPIURL = "https://verify.twilio.com/v2/Services/"
	users      = make(map[string]AccountInfo)
)

type server struct {
	pb.UnimplementedAuthServiceServer
	pb.UnimplementedOTPServiceServer
}

// VerificationCheckOutput is received when a verification request has been processed.
type VerificationCheck struct {
	Valid bool `json:"valid"`
}

type AccountInfo struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Username    string `json:"username"`
	Verify      bool   `json:"verify"`
}

var (
	grpcClient pb.OTPServiceClient = nil
	conn       *grpc.ClientConn    = nil
)

func CallVerifyAPI(query url.Values, customPath string) (*http.Response, error) {
	URL := BaseAPIURL + config.Configuration.ServicesID + customPath

	msgDataR := *strings.NewReader(query.Encode())
	client := &http.Client{}
	request, _ := http.NewRequest("POST", URL, &msgDataR)
	request.SetBasicAuth(
		config.Configuration.AccountSSID,
		config.Configuration.AuthToken)

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return response, fmt.Errorf("an error happend sending the request: %w", err)
	}

	return response, nil
}

func init() {
	flag.Parse()
	var err error

	addr := fmt.Sprintf("localhost:%d", *port)
	opts := insecure.NewCredentials()

	conn, err = grpc.Dial(addr, grpc.WithTransportCredentials(opts))
	if err != nil {
		log.Fatalf("did not connect: %v\\n", err)
	}

	err = config.LoadConfig(*configFile)
	if err != nil {
		log.Fatalln("could not load config:", err)
	}

	grpcClient = pb.NewOTPServiceClient(conn)
}

func SendOTPRequest(phoneNumber string) error {
	ctx, canel := context.WithTimeout(context.Background(), 20*time.Second)
	defer canel()

	_, err := grpcClient.CreateTwillioOTP(ctx, &pb.Phone{Number: phoneNumber})
	if err != nil {
		return fmt.Errorf("an error createOTP: %v", err)
	}
	return nil
}

func (s *server) SignupWithPhoneNumber(ctx context.Context, req *pb.SignupWithPhoneNumberRequest) (*pb.GenericResponse, error) {
	fmt.Println("Signup With Phone Number")
	account := req.GetAccount()

	// Check request
	if req == nil {
		return nil, errors.New("request must not be nil")
	}

	if req.Account.Name == "" {
		return nil, errors.New("name could not be empty")
	}

	if req.Account.Username == "" {
		return nil, errors.New("username could not be empty")
	}

	if req.Account.PhoneNumber == "" {
		return nil, errors.New("phone number could not be empty")
	}

	err := SendOTPRequest(req.Account.PhoneNumber)
	if err != nil {
		return nil, err
	}

	users[req.Account.PhoneNumber] = AccountInfo{
		Name:        account.Name,
		Username:    account.Username,
		PhoneNumber: account.PhoneNumber,
	}

	return &pb.GenericResponse{
		StatusCode: "Codigo enviado",
	}, nil
}

func (s *server) CreateTwillioOTP(ctx context.Context, req *pb.Phone) (*pb.GenericResponse, error) {
	res := &pb.GenericResponse{}

	query := make(url.Values)
	query.Set("To", req.Number)
	query.Set("Channel", "sms")

	r, err := CallVerifyAPI(query, "/Verifications")
	if err != nil {
		return res, fmt.Errorf("could not call this api %v", err)
	}

	if r.StatusCode != http.StatusCreated {
		return res, errors.New("an error happend could not send the OTP code")
	}

	return res, nil
}

func (s *server) VerifyPhoneNumber(ctx context.Context, req *pb.VerifyPhoneNumberRequest) (*pb.GenericResponse, error) {

	query := make(url.Values)
	query.Set("To", req.Otp.PhoneNumber)
	query.Set("Code", req.Otp.Code)

	r, err := CallVerifyAPI(query, "/VerificationCheck")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("an error happend reading the the body")
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.New("an error happend the code could no be verified")
	}

	checkCode := VerificationCheck{}
	result := &pb.GenericResponse{}

	if err := json.Unmarshal(body, &checkCode); err != nil {
		return nil, err
	}

	if checkCode.Valid {
		result.StatusCode = "Valid code"
		users[req.Otp.PhoneNumber] = AccountInfo{
			Verify: true,
		}
	} else {
		result.StatusCode = "Invalid code "
	}

	return result, nil
}

func (s *server) LoginWithPhoneNumber(ctx context.Context, req *pb.Phone) (*pb.GenericResponse, error) {
	fmt.Println("Login with Phone")

	// Check request
	if req == nil {
		return nil, errors.New("request must not be nil")
	}

	if req.Number == "" {
		return nil, errors.New("phone number could not be empty")
	}

	if _, ok := users[req.GetNumber()]; ok {
		err := SendOTPRequest(req.GetNumber())
		if err != nil {
			log.Println(err)
			return nil, err
		}

		return &pb.GenericResponse{
			StatusCode: "Codigo enviado",
		}, nil
	}

	return nil, errors.New("phone number is not registred")
}

func (s *server) ValidatePhoneNumberLogin(ctx context.Context, req *pb.OTP) (*pb.GenericResponse, error) {

	query := make(url.Values)
	query.Set("To", req.PhoneNumber)
	query.Set("Code", req.Code)

	r, err := CallVerifyAPI(query, "/VerificationCheck")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("an error happend reading the the body")
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.New("an error happend the code could no be verified")
	}

	checkCode := VerificationCheck{}
	if err := json.Unmarshal(body, &checkCode); err != nil {
		return nil, err
	}

	result := &pb.GenericResponse{}
	if checkCode.Valid {
		result.StatusCode = "Valid code"

	} else {
		result.StatusCode = "Invalid code "
	}

	return result, nil

}

func main() {

	defer conn.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &server{})
	pb.RegisterOTPServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
