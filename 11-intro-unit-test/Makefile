run:
	go run main.go

test: 
	go test ./calculate/nilai/... -cover

test-html:
	go test ./calculate/nilai/... -coverprofile=cover.out && go tool cover -html=cover.out