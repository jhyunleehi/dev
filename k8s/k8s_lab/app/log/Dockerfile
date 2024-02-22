FROM ubuntu
ENTRYPOINT ["/bin/bash" , "-c", "for i in $(seq 1000) ; do echo $(date +%T) : service requests is being processed successfully >> /tmp/app.log ; sleep 3 ; done" ]
