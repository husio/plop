
all: auth blog

blog:
	go build -o bin/blog cmd/blog/*go

auth:
	go build -o bin/auth cmd/auth/*go


clean:
	@@rm bin -r 2> /dev/null || true
