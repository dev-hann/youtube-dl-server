#!/bin/bash

echo "init install youtube-dl server..."
if [ $(id -u) -ne 0 ]; then exec sudo bash "$0" "$@"; exit; fi

apt-get update
apt-get install -y git
apt-get install -y curl
apt-get install -y systemctl

# install docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo usermod -aG  docker $USER

# install docker compose
#curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
#chmod +x /usr/local/bin/docker-compose
apt install -y docker-compose
chmod +x /usr/local/bin/docker-compose
docker-compose up
