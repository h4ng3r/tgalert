BINARY = tgalert

all: linux darwin windows

linux:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}-linux-amd64 .

darwin:
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}-darwin-amd64 .

windows:
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}-windows-amd64.exe .
