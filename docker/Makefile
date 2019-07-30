IMAGENAME=api_server_simurator

all:
	sudo docker build -t ${IMAGENAME} .
	sudo docker run -d -p 8080:8080 ${IMAGENAME} 
	ruby import.rb
clean:
	sudo docker ps --filter ancestor=${IMAGENAME} | grep -v CONTAINER | cut -c1-12 | xargs sudo docker kill
clean-image:
	sudo docker rmi -f ${IMAGENAME}
