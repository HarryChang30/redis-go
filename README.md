# redis-go

Redis go implementation by using package library go-redis dependency. Simulating the time process for each operations in redis instance / clusters.

## Pre-requisites
The integration is required some installation redis local instance

- Redis v5.0.3 or latest version
- Go v1.12 or latest version

## Feature Implemented for Simulation
- Set Operation
- Get Operation
- Pipeline Single Command Execution
- Pipeline Multiple Command Execution

## Running Project Steps
1. Clone repo from the repository \
```git clone https://github.com/HarryChang30/redis-go.git```

2. To build the project and dependencies \
```go build .```

3. Run the program \
```go run main.go```
