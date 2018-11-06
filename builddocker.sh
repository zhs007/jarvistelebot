docker build -t jarvistelebot .

if [ ! -d "logs" ]; then
    mkdir logs
fi

if [ ! -d "dat" ]; then
    mkdir dat
fi