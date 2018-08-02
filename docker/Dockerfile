FROM centos:7

USER root

COPY jarvistelebot /usr/bin/

COPY config.yaml /home/jarvistelebot/

CMD ["jarvistelebot", "--run=/home/jarvistelebot"]