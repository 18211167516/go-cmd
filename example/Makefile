.PHONY: all build clean

CMD = "gocmd"

all: check cmd run

cmd:
	go build -o ${CMD} 

linux-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${CMD}

windows-build:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${CMD}.exe

mac-build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${CMD}
check:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${CMD} ] ; then rm ${CMD} ; fi
