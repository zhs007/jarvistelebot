docker container stop jarvistelebot
docker container rm jarvistelebot
docker run -d --name jarvistelebot -v $PWD/cfg:/home/jarvistelebot/cfg jarvistelebot