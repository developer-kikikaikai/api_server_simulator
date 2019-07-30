IMAGENAME=api_server_simurator

all:
	sudo docker build -t ${IMAGENAME} .
	sudo docker run -d -p 8080:8080 ${IMAGENAME}
