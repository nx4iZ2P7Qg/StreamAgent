package ssh

import (
	"StreamAgent/internal/cus/config"
	"fmt"
	"golang.org/x/crypto/ssh"
	"time"
)

func Cmd(cmd string) string {
	sshHost := config.C.Ssh.Host
	sshPort := config.C.Ssh.Port
	sshUser := config.C.Ssh.User
	sshPassword := config.C.Ssh.Password
	sshType := "password" //password 或者 key
	//sshKeyPath := ""//ssh id_rsa.id 路径"

	c := &ssh.ClientConfig{
		Timeout:         time.Second,
		User:            sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	if sshType == "password" {
		c.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}
	} else {
		//config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(sshKeyPath)}
	}

	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	sshClient, err := ssh.Dial("tcp", addr, c)
	if err != nil {
		panic(err)
	}
	defer sshClient.Close()

	session, err := sshClient.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	combo, err := session.CombinedOutput(cmd)
	fmt.Printf("cmd = %v\n", cmd)
	fmt.Printf("combo = %v\n", string(combo))
	if err != nil {
		panic(err)
	}
	return string(combo)
}