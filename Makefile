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
