build : clean setup
	docker build -t spew .

setup :
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dist/spew .

clean :
	rm -rf dist
