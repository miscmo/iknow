all: proto backend fronted

backend:
	cd ./backend
	go build
	cd ..

fronted:


.PHONY: clean
clean:
