binary:
	go build keyvol.go
	go install keyvol.go

deps:
	go get -v -u github.com/nsf/termbox-go
