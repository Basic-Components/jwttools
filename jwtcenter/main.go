package main //import "github.com/Basic-Components/jwttools/jwtcenter"

import (
	context "context"
	"os"

	"net"

	errs "github.com/Basic-Components/jwttools/jwtcenter/errs"
	pb "github.com/Basic-Components/jwttools/jwtcenter/jwtrpcdeclare"
	log "github.com/Basic-Components/jwttools/jwtcenter/logger"
	script "github.com/Basic-Components/jwttools/jwtcenter/script"
	"github.com/Basic-Components/jwttools/jwtsigner"
	"github.com/Basic-Components/jwttools/jwtverifier"
	jsoniter "github.com/json-iterator/go"
	grpc "google.golang.org/grpc"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type rpcservice struct {
	AsymmetricSigner   *jwtsigner.Asymmetric
	AsymmetricVerifier *jwtverifier.Asymmetric
	SymmetricSigner    *jwtsigner.Symmetric
	SymmetricVerifier  *jwtverifier.Symmetric
}

// NewService 创建一个新的服务
func NewService(conf script.ConfigType) (*rpcservice, error) {
	s := &rpcservice{}
	asymmetricSigner, err := jwtsigner.AsymmetricFromPEMFile("RS256", conf.PrivateKeyPath)
	if err != nil {
		return s, err
	}
	asymmetricVerifier, err := jwtverifier.AsymmetricFromPEMFile("RS256", conf.PublicKeyPath)
	if err != nil {
		return s, err
	}
	symmetricSigner, err := jwtsigner.SymmetricNew("HS256", conf.Hashkey)
	if err != nil {
		return s, err
	}
	symmetricVerifier := jwtverifier.SymmetricNew(conf.Hashkey)
	if err != nil {
		return s, err
	}
	s.AsymmetricSigner = asymmetricSigner
	s.AsymmetricVerifier = asymmetricVerifier
	s.SymmetricSigner = symmetricSigner
	s.SymmetricVerifier = symmetricVerifier
	return s, nil
}

func (s *rpcservice) SignJSON(ctx context.Context, in *pb.SignJSONRequest) (*pb.SignResponse, error) {
	log.Debug(map[string]interface{}{
		"in":     in,
		"Method": "SignJSON",
	}, "get request")
	var res *pb.SignResponse
	switch in.Algo {
	case pb.Algo_HS256:
		{
			if in.Exp < 0 {
				status := pb.StatusData{
					Status: pb.StatusData_ERROR,
					Msg:    errs.ErrExpOutOfRange.Error()}
				res = &pb.SignResponse{Status: &status}
			} else {
				if in.Exp > 0 {
					token, err := s.SymmetricSigner.ExpSignJSON(in.Payload, in.Aud, in.Iss, in.Exp)
					if err != nil {
						status := pb.StatusData{
							Status: pb.StatusData_ERROR,
							Msg:    err.Error()}
						res = &pb.SignResponse{Status: &status}
					} else {
						status := pb.StatusData{
							Status: pb.StatusData_SUCCEED,
						}
						res = &pb.SignResponse{Status: &status, Token: token}
					}
				} else {
					token, err := s.SymmetricSigner.SignJSON(in.Payload, in.Aud, in.Iss)
					if err != nil {
						status := pb.StatusData{
							Status: pb.StatusData_ERROR,
							Msg:    err.Error()}
						res = &pb.SignResponse{Status: &status}
					} else {
						status := pb.StatusData{
							Status: pb.StatusData_SUCCEED,
						}
						res = &pb.SignResponse{Status: &status, Token: token}
					}
				}
			}
		}
	case pb.Algo_RS256:
		{
			if in.Exp < 0 {
				status := pb.StatusData{
					Status: pb.StatusData_ERROR,
					Msg:    errs.ErrExpOutOfRange.Error()}
				res = &pb.SignResponse{Status: &status}
			} else {
				if in.Exp > 0 {
					token, err := s.AsymmetricSigner.ExpSignJSON(in.Payload, in.Aud, in.Iss, in.Exp)
					if err != nil {
						status := pb.StatusData{
							Status: pb.StatusData_ERROR,
							Msg:    err.Error()}
						res = &pb.SignResponse{Status: &status}
					} else {
						status := pb.StatusData{
							Status: pb.StatusData_SUCCEED,
						}
						res = &pb.SignResponse{Status: &status, Token: token}
					}
				} else {
					token, err := s.AsymmetricSigner.SignJSON(in.Payload, in.Aud, in.Iss)
					if err != nil {
						status := pb.StatusData{
							Status: pb.StatusData_ERROR,
							Msg:    err.Error()}
						res = &pb.SignResponse{Status: &status}
					} else {
						status := pb.StatusData{
							Status: pb.StatusData_SUCCEED,
						}
						res = &pb.SignResponse{Status: &status, Token: token}
					}
				}
			}
		}
	default:
		{
			status := pb.StatusData{
				Status: pb.StatusData_ERROR,
				Msg:    errs.ErrAlgoType.Error()}
			res = &pb.SignResponse{Status: &status}
		}
	}
	log.Debug(map[string]interface{}{
		"return": res,
		"Method": "SignJSON",
	}, "response")
	return res, nil
}
func (s *rpcservice) VerifyJSON(ctx context.Context, in *pb.VerifyRequest) (*pb.VerifyJSONResponse, error) {
	log.Debug(map[string]interface{}{"in": in, "Method": "VerifyJSON"}, "get request")
	var res *pb.VerifyJSONResponse
	switch in.Algo {
	case pb.Algo_HS256:
		{
			payload, err := s.SymmetricVerifier.Verify(in.Token)
			if err != nil {
				status := pb.StatusData{
					Status: pb.StatusData_ERROR,
					Msg:    err.Error()}
				res = &pb.VerifyJSONResponse{Status: &status}
			} else {
				payloadJSON, err := json.Marshal(payload)
				if err != nil {
					status := pb.StatusData{
						Status: pb.StatusData_ERROR,
						Msg:    err.Error()}
					res = &pb.VerifyJSONResponse{Status: &status}
				} else {
					status := pb.StatusData{
						Status: pb.StatusData_SUCCEED,
					}
					res = &pb.VerifyJSONResponse{Status: &status, Payload: payloadJSON}
				}
			}
		}
	case pb.Algo_RS256:
		{
			payload, err := s.AsymmetricVerifier.Verify(in.Token)
			if err != nil {
				status := pb.StatusData{
					Status: pb.StatusData_ERROR,
					Msg:    err.Error()}
				res = &pb.VerifyJSONResponse{Status: &status}
			} else {
				payloadJSON, err := json.Marshal(payload)
				if err != nil {
					status := pb.StatusData{
						Status: pb.StatusData_ERROR,
						Msg:    err.Error()}
					res = &pb.VerifyJSONResponse{Status: &status}
				} else {
					status := pb.StatusData{
						Status: pb.StatusData_SUCCEED,
					}
					res = &pb.VerifyJSONResponse{Status: &status, Payload: payloadJSON}
				}
			}
		}
	default:
		{
			status := pb.StatusData{
				Status: pb.StatusData_ERROR,
				Msg:    errs.ErrAlgoType.Error()}
			res = &pb.VerifyJSONResponse{Status: &status}
		}
	}
	log.Debug(map[string]interface{}{
		"return": res,
		"Method": "VerifyJSON",
	}, "response")
	return res, nil
}

// Run 执行签名验签服务
func (s *rpcservice) Run() {
	listener, err := net.Listen("tcp", script.Config.Address)
	if err != nil {
		log.Logger.Fatalf("failed to listen: %v", err)
		return
	}
	log.Info(map[string]interface{}{"address": script.Config.Address}, "server started")
	server := grpc.NewServer()
	pb.RegisterJwtServiceServer(server, s)
	if err := server.Serve(listener); err != nil {
		log.Logger.Fatalf("failed to serve: %v", err)
		return
	}
}

func main() {
	err := script.Init()
	if err != nil {
		log.Warn(map[string]interface{}{"error": err}, "init config error")
		os.Exit(1)
	}
	log.Init(script.Config.LogLevel, map[string]interface{}{"component_name": script.Config.ComponentName})
	log.Info(map[string]interface{}{"conf": script.Config}, "setted server config")
	service, err := NewService(script.Config)
	if err != nil {
		log.Error(map[string]interface{}{"error": err}, "service not inited")
		os.Exit(1)
	}
	service.Run()
}
