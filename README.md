# scp-client 

## Usage

### Server

reference: https://docs.docker.com/engine/examples/running_ssh_service/

1. build image
`docker build -t ssh-server`

2. run container
`docker run -d -P --name test-scp ssh-server` 

3. check your port
`docker port test-scp 22`

### Client

#### command help
```
Usage of client:
  -host string
        default: localhost (default "localhost")
  -local string
        local/text.txt
  -port int
        default: 22 (default 22)
  -private-key string
        designate private SSH key
  -remote string
        remote/text.txt
```
