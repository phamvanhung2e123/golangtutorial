deps:
	@echo "====> Install dependencies..."

clean:
	@echo "====> Remove installed binary"
	rm -f bin/tail

install:
	@echo "====> Build hget in ./bin "
	go build -o bin/tail

test:
	@echo "====> Remove installed binary"
	go test