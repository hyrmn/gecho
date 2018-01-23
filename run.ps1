#After this, use putty and connect to localhost:7000
docker run --rm -p 7000:7 echo

#useful for getting the containers IP address
#docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' echo