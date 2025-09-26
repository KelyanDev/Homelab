<h3 align="center">Nginx Proxy Manager - Configuration</h3>

  <p align="center">
    Every step in installing Nginx Proxy Manager for the first time
    <br />
    <a href="https://github.com/KelyanDev/Homelab"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    Navigation <br />
    <a href="https://github.com/KelyanDev/Homelab">General</a>
    ·
    <a href="https://github.com/KelyanDev/Homelab/blob/main/apps/README.md">Apps</a>
    ·
    <a href="">WIP</a>
    ·
    <a href="">WIP</a>
  </p>
</div>

## Installation

There is multiple steps in installing Nginx Proxy Manager. These steps assume that you already have a running Debian 13 LXC
The Nginx Proxy Manager documentations can also provide great amount of informations regarding the installation steps, but they assume that you're not using the root user

This tutorial will allow you to install NPM using Docker and Docker-compose.

```
# Update packages:
apt update && apt upgrade -y

# Install the dependencies
apt install -y curl gnupg2 ca-certificates lsb-release apt-transport-https

# Install Docker and Docker-compose using their online script
curl -fsSL https://get.docker.com | sh
```
Once all these steps are done and the installation proceeded, you can verify if docker is successfully installed by running ``docker --version``, and ``docker compose version`` for docker compose
Once Docker is finally installed, you just need to create a Nginx Proxy Manager. Here's the steps I used to install it with docker compose:
```
# Create a repository for the container
mkdir -p /opt/npm
cd /opt/npm

# Create your "compose.yaml" file to run the container:
nano compose.yaml ## check the compose.yaml file in this repository if you need the configuration file

# Start the container
docker compose up -d
```
Then, you can start to configure your Nginx Proxy Manager. It usually runs on the port 3000, so if you configured yours correctly and applied the port, you should be able to access it directly from here

## Configuration file

You can find the config file here (or really anywhere on Internet):
[Config](https://github.com/KelyanDev/Homelab/blob/main/apps/Nginx Proxy Manager/compose.yaml)

