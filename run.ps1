docker run --rm --name echo echo

#useful for getting the containers IP address
#docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' echo