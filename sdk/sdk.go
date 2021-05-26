package sdk

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	declare "github.com/Basic-Components/jwttools/jwtrpcdeclare"
	"github.com/Basic-Components/jwttools/options"
	log "github.com/Golang-Tools/loggerhelper"
	jsoniter "github.com/json-iterator/go"
	"github.com/liyue201/grpc-lb/balancer"
	registry "github.com/liyue201/grpc-lb/registry/zookeeper"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	resolver "google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
	_ "google.golang.org/grpc/xds"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

//SDKConfig 的客户端类型
type SDKConfig struct {
	Address                                 []string `json:"address" jsonschema:"description=连接服务的主机和端口"`
	ServiceName                             string   `json:"service_name,omitempty" jsonschema:"description=服务器域名"`
	AppName                                 string   `json:"app_name,omitempty" jsonschema:"description=服务名"`
	AppVersion                              string   `json:"app_version,omitempty" jsonschema:"description=服务版本"`
	BalanceWithZookeeper                    bool     `json:"balance_with_zookeeper,omitempty" jsonschema:"description=是否使用zookeeper做本地负载均衡"`
	CaCert                                  string   `json:"ca_cert,omitempty" jsonschema:"description=如果要使用tls则需要指定根证书位置"`
	InitialWindowSize                       int      `json:"initial_window_size,omitempty" jsonschema:"description=基于Stream的滑动窗口大小"`
	InitialConnWindowSize                   int      `json:"initial_conn_window_size,omitempty" jsonschema:"description=基于Connection的滑动窗口大小"`
	KeepaliveTime                           int      `json:"keepalive_time,omitempty" jsonschema:"description=空闲连接每隔n秒ping一次客户端已确保连接存活"`
	KeepaliveTimeout                        int      `json:"keepalive_timeout,omitempty" jsonschema:"description=ping时长超过n则认为连接已死"`
	KeepaliveEnforcementPermitWithoutStream bool     `json:"keepalive_enforement_permit_without_stream,omitempty" jsonschema:"description=是否当连接空闲时仍然发送PING帧监测"`
	ConnWithBlock                           bool     `json:"conn_with_block,omitempty" jsonschema:"description=同步的连接建立"`
	MaxRecvMsgSize                          int      `json:"max_rec_msg_size,omitempty" jsonschema:"description=允许接收的最大消息长度"`
	MaxSendMsgSize                          int      `json:"max_send_msg_size,omitempty" jsonschema:"description=允许发送的最大消息长度"`
	QueryTimeout                            int      `json:"query_timeout,omitempty" jsonschema:"description=请求的超时时长,单位毫秒"`
}

//NewSDK 创建客户端对象
func (c *SDKConfig) NewSDK() *SDK {
	sdk := New()
	sdk.Init(c)
	return sdk
}

//SDK 的客户端类型
type SDK struct {
	*SDKConfig
	opts          []grpc.DialOption
	serviceconfig map[string]interface{}
	addr          string
	conn          *Conn
}

//New 创建客户端对象
func New() *SDK {
	c := new(SDK)
	c.opts = make([]grpc.DialOption, 0, 10)
	return c
}

//InitCallOpts 初始化连接选项
func (c *SDK) InitCallOpts() {
	callopts := []grpc.CallOption{}
	if c.MaxRecvMsgSize != 0 {
		callopts = append(callopts, grpc.MaxCallRecvMsgSize(c.MaxRecvMsgSize))
	}
	if c.MaxSendMsgSize != 0 {
		callopts = append(callopts, grpc.MaxCallSendMsgSize(c.MaxSendMsgSize))
	}
	if len(callopts) > 0 {
		c.opts = append(c.opts, grpc.WithDefaultCallOptions(callopts...))
	}
}

//InitOpts 初始化连接选项
func (c *SDK) InitOpts() error {
	c.opts = append(c.opts)
	if c.CaCert != "" {
		creds, err := credentials.NewClientTLSFromFile(c.CaCert, "")
		if err != nil {
			log.Error("failed to load credentials", log.Dict{"err": err.Error()})
			return err
		}
		c.opts = append(c.opts, grpc.WithTransportCredentials(creds))
	} else {
		c.opts = append(c.opts, grpc.WithInsecure())
	}
	if c.KeepaliveTime != 0 || c.KeepaliveTimeout != 0 || c.KeepaliveEnforcementPermitWithoutStream == true {
		kacp := keepalive.ClientParameters{
			Time:                time.Duration(c.KeepaliveTime) * time.Second,
			Timeout:             time.Duration(c.KeepaliveTimeout) * time.Second,
			PermitWithoutStream: c.KeepaliveEnforcementPermitWithoutStream, // send pings even without active streams
		}
		c.opts = append(c.opts, grpc.WithKeepaliveParams(kacp))
	}
	if c.ConnWithBlock == true {
		c.opts = append(c.opts, grpc.WithBlock())
	}
	if c.InitialWindowSize != 0 {
		c.opts = append(c.opts, grpc.WithInitialWindowSize(int32(c.InitialWindowSize)))
	}
	if c.InitialConnWindowSize != 0 {
		c.opts = append(c.opts, grpc.WithInitialConnWindowSize(int32(c.InitialConnWindowSize)))
	}

	return nil
}

//Init 初始化sdk客户端的连接信息
func (c *SDK) Init(conf *SDKConfig) error {
	c.SDKConfig = conf
	if conf.Address == nil {
		return errors.New("必须至少有一个地址")
	}
	switch len(conf.Address) {
	case 0:
		{
			return errors.New("必须至少有一个地址")
		}
	case 1:
		{
			if conf.BalanceWithZookeeper {
				c.initWithZooKeeperBalance()
			} else {
				c.initStandalone()
			}
		}
	default:
		{
			if conf.BalanceWithZookeeper {
				c.initWithZooKeeperBalance()
			} else {
				c.initWithLocalBalance()
			}
		}
	}
	err := c.InitOpts()
	if err != nil {
		return err
	}
	c.InitCallOpts()
	if c.serviceconfig != nil {
		serviceconfig, err := json.MarshalToString(c.serviceconfig)
		if err != nil {
			return err
		}
		c.opts = append(c.opts, grpc.WithDefaultServiceConfig(serviceconfig))
	}
	return nil
}

//InitStandalone 初始化单机服务的连接配置
func (c *SDK) initStandalone() error {
	c.addr = c.Address[0]
	return nil
}

//InitWithLocalBalance 初始化本地负载均衡的连接配置
func (c *SDK) initWithLocalBalance() error {
	serverName := ""
	if c.AppName != "" {
		if c.AppVersion != "" {
			serverName = fmt.Sprintf("%s-%s", c.AppName, strings.ReplaceAll(c.AppVersion, ".", "_"))
		} else {
			serverName = c.AppName
		}

	}
	if c.serviceconfig == nil {
		c.serviceconfig = map[string]interface{}{
			"loadBalancingPolicy": "round_robin",
			"healthCheckConfig":   map[string]string{"serviceName": c.ServiceName},
		}
	} else {
		c.serviceconfig["loadBalancingPolicy"] = "round_robin"
		c.serviceconfig["healthCheckConfig"] = map[string]string{"serviceName": serverName}
	}
	r := manual.NewBuilderWithScheme("localbalancer")
	addresses := []resolver.Address{}
	for _, addr := range c.Address {
		addresses = append(addresses, resolver.Address{Addr: addr})
	}
	r.InitialState(resolver.State{
		Addresses: addresses,
	})
	c.addr = fmt.Sprintf("%s:///%s", r.Scheme(), serverName)
	c.opts = append(c.opts, grpc.WithResolvers(r))
	return nil
}

//InitWithZooKeeperBalance 初始zokeeper负载均衡的连接配置
func (c *SDK) initWithZooKeeperBalance() error {
	registry.RegisterResolver("zk", c.Address, "/backend/services", c.AppName, c.AppVersion)
	c.addr = "zk:///"
	c.opts = append(c.opts, grpc.WithBalancerName(balancer.RoundRobin))
	return nil
}

func (c *SDK) GetConn() (*Conn, error) {
	if c.conn == nil {
		return c.NewConn()
	}
	return c.conn, nil
}

func (c *SDK) Close() error {
	if c.conn != nil {
		err := c.conn.Close()
		c.conn = nil
		return err
	}
	return nil
}

//NewConn 建立一个新的连接
func (c *SDK) NewConn() (*Conn, error) {
	conn, err := newConn(c, c.addr, c.opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
func (c *SDK) queryCtx() (context.Context, context.CancelFunc) {
	var ctx context.Context
	var cancel context.CancelFunc
	if c.QueryTimeout > 0 {
		timeout := time.Duration(c.QueryTimeout) * time.Millisecond
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	return ctx, cancel
}

//Algo 查看指向的服务使用的算法
func (c *SDK) Meta() (declare.EncryptionAlgorithm, error) {
	conn, err := c.GetConn()
	if err != nil {
		return declare.EncryptionAlgorithm_UNKNOWN, err
	}
	ctx, cancel := c.queryCtx()
	defer cancel()
	res, err := conn.Meta(ctx, &declare.MetaRequest{})
	if err != nil {
		return declare.EncryptionAlgorithm_UNKNOWN, err
	}
	return res.Algo, nil
}

//Sign 查看指向的服务使用的算法
func (c *SDK) Sign(payload []byte, opts ...options.SignOption) (string, error) {
	conn, err := c.GetConn()
	if err != nil {
		return "", err
	}
	ctx, cancel := c.queryCtx()
	defer cancel()
	query := declare.SignRequest{
		Payload: payload,
	}

	defaultopt := &options.SignOptions{}
	for _, opt := range opts {
		opt.Apply(defaultopt)
	}
	if defaultopt.Sub != "" {
		query.Sub = defaultopt.Sub
	}
	if defaultopt.Aud != nil {
		query.Aud = defaultopt.Aud
	}
	if defaultopt.Nbf > 0 {
		query.Nbf = defaultopt.Nbf
	}
	if defaultopt.Exp > 0 {
		query.Exp = defaultopt.Exp
	}

	res, err := conn.Sign(ctx, &query)
	if err != nil {
		return "", err
	}
	return res.Token, nil
}

//Verify
func (c *SDK) Verify(tokenstring string) (map[string]interface{}, error) {
	conn, err := c.GetConn()
	if err != nil {
		return nil, err
	}
	ctx, cancel := c.queryCtx()
	defer cancel()
	res, err := conn.Verify(ctx, &declare.VerifyRequest{Token: tokenstring})
	if err != nil {
		return nil, err
	}
	result := map[string]interface{}{}
	err = json.Unmarshal(res.Payload, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//Conn 客户端类
type Conn struct {
	declare.JwtServiceClient
	conn *grpc.ClientConn
	sdk  *SDK
	once bool
}

func newConn(sdk *SDK, addr string, opts ...grpc.DialOption) (*Conn, error) {
	c := new(Conn)
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}
	c.conn = conn
	c.JwtServiceClient = declare.NewJwtServiceClient(conn)
	return c, nil
}

//Close 断开连接
func (c *Conn) Close() error {
	return c.conn.Close()
}

//DefaultSDK 默认的sdk客户端
var DefaultSDK = New()