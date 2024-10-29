# Installation 

```bash
docker build -f Dockerfile.dev -t go-dev .

docker run --name go-dev -v .:/app go-dev
```
Or use VsCode devcontainer

# Start

```bash 
go run main.go

# other script
cd scripts
go run script_file.go

```