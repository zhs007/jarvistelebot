docker container stop jarvistelebot
docker container rm jarvistelebot
docker run -d --name jarvistelebot -v $PWD/cfg:/home/jarvistelebot/cfg -v $PWD/logs:/home/jarvistelebot/logs -v $PWD/dat:/home/jarvistelebot/dat jarvistelebot