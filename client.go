package main

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tmc/scp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"net"
	"os"
	"strconv"
)

var (
	keyfile string
	localpath string
	remotepath string
	host string
	port int
)

func getAgent() (agent.Agent, error) {
	agentConn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	return agent.NewClient(agentConn), err
}

func main() {
	flag.StringVar(&host, "host", "localhost", "default: localhost")
	flag.StringVar(&keyfile, "private-key", "", "designate private SSH key")
	flag.StringVar(&localpath, "local", "", "local/text.txt")
	flag.IntVar(&port, "port", 22, "default: 22")
	flag.StringVar(&remotepath, "remote", "", "remote/text.txt")
	flag.Parse()

	f, err := os.Open(localpath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	agent, err := getAgent()
	if err != nil {
		panic(errors.Wrap(err, "Failed to connect to SSH_AUTH_SOCK"))
	}

	client, err := ssh.Dial("tcp", host + ":" + strconv.Itoa(port), &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeysCallback(agent.Signers),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		panic(errors.Wrap(err, "Failed to dial"))
	}

	session, err := client.NewSession()
	if err != nil {
		panic(errors.Wrap(err, "Failed to create session"))
	}

	err = scp.CopyPath(localpath, remotepath, session)
	if err != nil {
		panic(err)
	}
	fmt.Println("success!")
}
