TARGET=api_server_simulator

all:
	test -e vendor || dep ensure
	go build -o ${TARGET}
clean:
	rm ${TARGET}
ut:
	make -C test
clean-cache:
	rm -rf vendor
container:
	make -C docker
clean-container:
	make -C docker clean
clean-container-image:
	make -C docker clean-image
