
all: blog currtime auth

blog:
	go build -o bin/blog cmd/blog/*go

auth:
	go build -o bin/auth cmd/auth/*go

currtime:
	go build -o bin/currtime cmd/currtime/*go


clean:
	@@rm bin -r 2> /dev/null || true


.PHONY: all blog auth currtime clean
