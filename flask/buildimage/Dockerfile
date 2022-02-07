FROM ubuntu:latest
RUN apt-get update -y
RUN apt-get install python3-pip -y
WORKDIR /home
COPY ./requirements.txt /home
COPY ./testflkafka.py /home
COPY ./train.py /home
RUN pip3 install -r requirements.txt
RUN python3 train.py
EXPOSE 5000

