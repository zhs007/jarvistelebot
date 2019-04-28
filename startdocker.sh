docker container stop jarvistelebot
docker container rm jarvistelebot
docker run -d --name jarvistelebot -v $PWD/cfg:/app/jarvistelebot/cfg -v $PWD/logs:/app/jarvistelebot/logs -v $PWD/dat:/app/jarvistelebot/dat jarvistelebot