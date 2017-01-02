// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package setting

import (
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Unknwon/com"
	"gopkg.in/ini.v1"

	"github.com/rodkranz/fakeApi/modules/log"
)

type Scheme string

const (
	HTTP  Scheme = "http"
	HTTPS Scheme = "https"
)

var (
	// Fake Api
	SeedFolder    string
	SeedExtension string

	// App Serrings
	AppVer         string
	AppName        string
	AppDesc        string
	AppUrl         string
	AppPath        string
	AppSubUrl      string
	AppSubUrlDepth int // Number of slashes
	AppDataPath    string

	// Server setting
	Protocol             Scheme
	Domain               string
	HTTPAddr, HTTPPort   string
	LocalURL             string
	DisableRouterLog     bool
	CertFile, KeyFile    string
	StaticRootPath       string
	EnableGzip           bool
	UnixSocketPermission uint32

	// Global setting objects
	Cfg          *ini.File
	CustomPath   string // Custom directory path
	CustomConf   string
	ProdMode     bool
	IsWindows    bool
	HasRobotsTxt bool

	// Log setting
	LogRootPath string
	LogModes    []string
	LogConfigs  []string

	// Api
	AllowCrossDomain bool
)

// execPath returns the executable path.
func execPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Abs(file)
}

func init() {
	IsWindows = runtime.GOOS == "windows"
	//log.NewLogger(0, "console", `{"level": 0}`)

	var err error
	if AppPath, err = execPath(); err != nil {
		//		log.Fatal(4, "fail to get app path: %v\n", err)
	}

	// Note: we don't use path.Dir here because it does not handle case
	//      which path starts with two "/" in Windows: "//psf/Home/..."
	AppPath = strings.Replace(AppPath, "\\", "/", -1)
}

// WorkDir returns absolute path of work directory.
func WorkDir() (string, error) {
	wd := os.Getenv("WWW_DIR")
	if len(wd) > 0 {
		return wd, nil
	}

	i := strings.LastIndex(AppPath, "/")
	if i == -1 {
		return AppPath, nil
	}
	return AppPath[:i], nil
}

func forcePathSeparator(path string) {
	if strings.Contains(path, "\\") {
		log.Fatal(4, "Do not use '\\' or '\\\\' in paths, instead, please use '/' in all places")
	}
}

func NewContext() {
	workDir, err := WorkDir()
	if err != nil {
		log.Fatal(4, "Fail to get work directory: %v", err)
	}

	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal(4, "Fail to parse 'conf/app.ini': %v", err)
	}

	CustomPath = os.Getenv("FAKE_API_CUSTOM")
	if len(CustomPath) == 0 {
		CustomPath = workDir + "/custom"
	}

	if len(CustomConf) == 0 {
		CustomConf = CustomPath + "/conf/app.ini"
	}

	if com.IsFile(CustomConf) {
		if err = Cfg.Append(CustomConf); err != nil {
			log.Fatal(4, "Fail to load custom conf '%s': %v", CustomConf, err)
		}
	} else {
		log.Warn("Custom config '%s' not found, ignore this if you're running first time", CustomConf)
	}
	Cfg.NameMapper = ini.AllCapsUnderscore

	sec := Cfg.Section("app")
	sec = Cfg.Section("server")
	AppUrl = sec.Key("ROOT_URL").MustString("http://localhost:9090/")
	if AppUrl[len(AppUrl)-1] != '/' {
		AppUrl += "/"
	}

	// Check if has app suburl.
	surl, err := url.Parse(AppUrl)
	if err != nil {
		log.Fatal(4, "Invalid ROOT_URL '%s': %s", AppUrl, err)
	}

	// Suburl should start with '/' and end without '/', such as '/{subpath}'.
	AppSubUrl = strings.TrimSuffix(surl.Path, "/")
	AppSubUrlDepth = strings.Count(AppSubUrl, "/")

	Protocol = HTTP
	if sec.Key("PROTOCOL").String() == "https" {
		Protocol = HTTPS
		CertFile = sec.Key("CERT_FILE").String()
		KeyFile = sec.Key("KEY_FILE").String()
	}

	Domain = sec.Key("DOMAIN").MustString("localhost")
	HTTPAddr = sec.Key("HTTP_ADDR").MustString("0.0.0.0")
	HTTPPort = sec.Key("HTTP_PORT").MustString("3000")
	LocalURL = sec.Key("LOCAL_ROOT_URL").MustString("http://localhost:" + HTTPPort + "/")
	DisableRouterLog = sec.Key("DISABLE_ROUTER_LOG").MustBool()
	AppDataPath = sec.Key("APP_DATA_PATH").MustString("data")
	StaticRootPath = sec.Key("STATIC_ROOT_PATH").MustString(workDir)
	EnableGzip = sec.Key("ENABLE_GZIP").MustBool()

	AllowCrossDomain = Cfg.Section("api").Key("ALLOW_CROSS_DOMAIN").MustBool(true)

	sec = Cfg.Section("fakeApi")
	SeedExtension = sec.Key("SEED_EXTENSION").MustString(".json")
	SeedFolder = sec.Key("SEED_FOLDER").MustString("fakes")
}
