build:
	@if not exist out mkdir out
	set CGO_ENABLED=1
	go build -o .\out .

run: build
	./out/bongserver.exe
