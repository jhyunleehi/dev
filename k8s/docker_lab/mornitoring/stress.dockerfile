FROM ubuntu
RUN apt-get update ; apt-get install stress -y 
CMD [ "/bin/sh", "-c", "stress --cpu 8" ] 
