build:
	mkdir -p bin
	go build -o bin/roctl

clean:
	rm -rf bin