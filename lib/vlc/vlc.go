package vlc

import (
	"context"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"github.com/YusufOzmen01/otaku-cli/constants"
	"github.com/YusufOzmen01/otaku-cli/lib/network"
	"math/rand"
	"os/exec"
	"runtime"
	"strconv"
)

type vlc struct {
	file   string
	port   int
	auth   string
	vlcPID int
	params []string
}

type VLC interface {
	GetVLCData() (*VLCData, error)
}

func NewVLC(file string, parameters []string) (VLC, error) {
	constants.KillProcessByNameWindows("vlc.exe")

	port := rand.Intn(10000) + 50000
	password := constants.RandomString(32)
	baseParams := []string{file, "--intf", "qt", "--extraintf", "http"}

	if runtime.GOOS == "windows" {
		baseParams = append(baseParams, []string{"--http-password=" + password, "--http-port=" + strconv.Itoa(port)}...)
	} else {
		baseParams = append(baseParams, []string{"--http-password ", "--http-port ", strconv.Itoa(port)}...)
	}

	finalParams := append(baseParams, parameters...)

	proc := exec.Command("vlc", finalParams...)
	err := proc.Start()
	if err != nil {
		return nil, err
	}

	authString := ":" + password
	encodedString := base64.StdEncoding.EncodeToString([]byte(authString))

	return &vlc{
		file:   file,
		port:   port,
		vlcPID: proc.Process.Pid,
		auth:   encodedString,
		params: finalParams,
	}, nil
}

func (v *vlc) GetVLCData() (*VLCData, error) {
	body, status, err := network.ProcessGet(context.Background(), fmt.Sprintf("http://localhost:%d/requests/status.xml", v.port), map[string]string{"Authorization": "Basic " + v.auth})
	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, fmt.Errorf("an error occured")
	}

	data := new(VLCData)

	if err := xml.Unmarshal(body, data); err != nil {
		return nil, err
	}

	return data, nil
}
