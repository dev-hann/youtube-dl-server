FROM golang
ENV port 80
EXPOSE 80

# install ngrok
#RUN apt-get update && apt-get upgrade
#RUN apt-get install -y unzip wget
#RUN wget https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-linux-amd64.zip --no-check-certificate
#RUN unzip ngrok-stable-linux-amd64.zip
#RUN mv ./ngrok /usr/bin/ngrok

# install youtube-dl
#RUN apt-get update
#RUN curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/local/bin/youtube-dl
#RUN chmod a+rx /usr/local/bin/youtube-dl
#RUN apt install -y python3-pip
#RUN pip3 install --upgrade youtube-dl



#WORKDIR /go/src/app
#COPY . .
#RUN go build -o app
#RUN cp ./app /
#
#CMD ["./app"]