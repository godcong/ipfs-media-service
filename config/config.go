package config

import (
	"github.com/pelletier/go-toml"
	log "github.com/sirupsen/logrus"
	"os"
)

// Media ...
type Media struct {
	Upload      string `toml:"upload"`   //上传路径
	Download    string `toml:"download"` //下载路径
	Transfer    string `toml:"transfer"` //转换路径
	M3U8        string `toml:"m3u8"`     //m3u8文件名
	KeyURL      string `toml:"key_url"`  //default url
	KeyDest     string `toml:"key_dest"` //key 文件输出目录
	KeyFile     string `toml:"key_file"` //key文件名
	KeyInfoFile string `toml:"key_info_file"`

	//keyFile文件名
}

// IPFS ...
type IPFS struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

// GRPC ...
type GRPC struct {
	Enable bool   `toml:"enable"`
	Type   string `toml:"type"`
	Path   string `toml:"path"`
	Port   string `toml:"port"`
}

// REST ...
type REST struct {
	Enable bool   `toml:"enable"`
	Type   string `toml:"type"`
	Path   string `toml:"path"`
	Port   string `toml:"port"`
}

// Queue ...
type Queue struct {
	Type     string `json:"type"`
	HostPort string `json:"host_port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

// OSS ...
type OSS struct {
	Endpoint        string `toml:"endpoint"`
	AccessKeyID     string `toml:"access_key_id"`
	AccessKeySecret string `toml:"access_key_secret"`
	BucketName      string `toml:"bucket_name"`
}

type ManagerConfig struct {
	GRPC     GRPC   `toml:"grpc"`
	REST     REST   `toml:"rest"`
	CallType string `json:"call_type"`
}
type Node struct {
	ManagerName string `toml:"manager_name"`
	NodeName    string `toml:"node_name"`
	CensorName  string `toml:"censor_name"`
	EnableGRPC  bool   `toml:"enable_grpc"`
	EnableREST  bool   `toml:"enable_rest"`
	REST        REST   `toml:"rest"`
	RequestType string `toml:"request_type"`
	Clear       bool   `toml:"clear"`
}

// Callback ...
type Callback struct {
	Type     string `toml:"type"`
	BackType string `toml:"back_type"`
	BackAddr string `toml:"back_addr"`
	Version  string `toml:"version"`
}

// Configure ...
type Configure struct {
	Media    Media    `toml:"media"`
	Queue    Queue    `toml:"queue"`
	OSS      []OSS    `toml:"oss"`
	Node     Node     `toml:"node"`
	IPFS     IPFS     `toml:"ipfs"`
	Callback Callback `toml:"callback"`
}

var config *Configure

// NullString ...
func NullString(s string) bool {
	return s == ""
}

// Initialize ...
func Initialize(filePath ...string) error {
	if filePath == nil {
		filePath = []string{"config.toml"}
	}

	cfg := LoadConfig(filePath[0])
	cfg.Media.Download = DefaultString(cfg.Media.Download, "download")
	if !IsExists(cfg.Media.Download) {
		err := os.Mkdir(cfg.Media.Download, os.ModePerm)
		if err != nil {
			return err
		}
	}

	cfg.Media.Upload = DefaultString(cfg.Media.Upload, "upload")
	if !IsExists(cfg.Media.Upload) {
		err := os.Mkdir(cfg.Media.Upload, os.ModePerm)
		if err != nil {
			return err
		}
	}

	cfg.Media.Transfer = DefaultString(cfg.Media.Transfer, "transfer")
	if !IsExists(cfg.Media.Transfer) {
		err := os.Mkdir(cfg.Media.Transfer, os.ModePerm)
		if err != nil {
			return err
		}
	}
	cfg.Media.KeyDest = DefaultString(cfg.Media.KeyDest, "keydest")
	if !IsExists(cfg.Media.KeyDest) {
		err := os.Mkdir(cfg.Media.KeyDest, os.ModePerm)
		if err != nil {
			return err
		}
	}

	config = cfg

	return nil
}

// IsExists ...
func IsExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		log.Panicln(err)
	}
	return true
}

// LoadConfig ...
func LoadConfig(filePath string) *Configure {
	cfg := DefaultConfigure()
	openFile, err := os.OpenFile(filePath, os.O_RDONLY|os.O_SYNC, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	decoder := toml.NewDecoder(openFile)
	err = decoder.Decode(cfg)
	if err != nil {
		panic(err.Error())
	}
	log.Printf("config: %+v", cfg)
	return cfg
}

func DefaultConfigure() *Configure {
	return &Configure{
		Media: Media{
			//upload= "upload"
			//transfer= "transfer"
			//m3u8 = "media.m3u8"
			//key_url = "http://localhost:8080"
			//key_file = "key"
			//key_dest = "output_key"
			//key_info_file = "KeyInfoFile"

			Download:    "download",
			Upload:      "upload",
			Transfer:    "transfer",
			M3U8:        "media.m3u8",
			KeyURL:      "http://localhost:8080",
			KeyDest:     "output_key",
			KeyFile:     "key",
			KeyInfoFile: "KeyInfoFile",
		},
		Node: Node{
			ManagerName: "godcong.grpc.manager",
			NodeName:    "godcong.grpc.node",
			CensorName:  "godcong.grpc.censor",
			EnableGRPC:  true,
			EnableREST:  true,
			Clear:       false,
			REST: REST{
				Enable: true,
				Type:   "",
				Path:   "",
				Port:   ":7787",
			},
			RequestType: "grpc",
		},
		IPFS:     IPFS{},
		Callback: Callback{},
	}
}

// Config ...
func Config() *Configure {
	if config == nil {
		panic("nil config")
	}
	return config
}

// DefaultString ...
func DefaultString(v, def string) string {
	if v == "" {
		return def
	}
	return v
}
