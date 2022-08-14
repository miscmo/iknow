all: proto backend fronted

backend:
	cd ./backend
	go build
	cd ..

fronted:


.PHONY: proto
proto:
	protoc --go_out=backend proto/*.proto

.PHONY: clean
clean:
	rm -fr backend/proto
