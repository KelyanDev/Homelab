<h3 align="center">Proxy</h3>

  <p align="center">
    Regarding my Proxy setup, I'm running a Nginx Proxy Manager LXC with Docker and Docker compose. This repository is my own <strong>Work in Progress</strong> installation guide, featuring the knowledge linked to my reverse proxy setup and how I access my services from outside
    <br />
    <a href="https://github.com/KelyanDev/Homelab"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    Navigation <br />
    <a href="https://github.com/KelyanDev/Homelab">General</a>
    ·
    <a href="https://github.com/KelyanDev/Homelab/blob/main/apps/README.md">Apps</a>
    ·
    <a href="https://github.com/KelyanDev/Homelab/blob/main/monitoring/README.md">Monitoring</a>
    ·
    <a href="https://github.com/KelyanDev/Homelab/blob/main/proxy/README.md"><strong>Proxy</strong></a>
    ·
    <a href="https://github.com/KelyanDev/Homelab/blob/main/cloud/README.md">Cloud</a>
    ·
    <a href="">WIP</a>
  </p>
</div>

## Installation

These steps assume that you already have a running Debian 13 LXC   

The original documentations might as well be a better source of information, depending on your tech stack.

### Nginx Proxy Manager

> **Specs**
> - 2 vCpu
> - 2 GiB memory
> - 10 GiB storage

These are the initial steps regarding the installation of Nginx Proxy Manager.    
The installation is made using this [Docker compose file](https://github.com/KelyanDev/Homelab/blob/main/proxy/docker-compose.yml). You can modify it to create your own, or find an alternative on Internet.   
```
# Update packages:
apt update && apt upgrade -y

# Install the dependencies
apt install -y curl gnupg2 ca-certificates lsb-release apt-transport-https

# Install Docker and Docker-compose using their online script
curl -fsSL https://get.docker.com | sh
```
Once all these steps are done and the installation proceeded, you can verify if docker is successfully installed by running ``docker --version``, and ``docker compose version`` for docker compose   

Then, we'll need to create Nginx Proxy Manager docker-compose.yml file.
```
# Create a repository for the container
mkdir -p /opt/npm
cd /opt/npm

# Create your "docker-compose.yml" file to run the container:
nano docker-compose.yml ## check the compose.yaml file in this repository if you need the configuration file

# Start the container
docker compose up -d
```
Then, you can start to configure your Nginx Proxy Manager. The admin interface is running on the port 81, so if you configured yours correctly and applied the port, you should be able to access it directly from here

The default user / password is ``admin@example.com`` / ``changeme``
