# scp-client 

## Usage

### Docker

reference: https://docs.docker.com/engine/examples/running_ssh_service/

1. build image
`docker build -t ssh-server`

2. run container
`docker run -d -P --name test-scp ssh-server` 

3. check your port
`docker port test-scp 22`
