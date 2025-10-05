<h3 align="center">Cloud</h3>

  <p align="center">
    Regarding my cloud stack, I'm running Nextcloud in an LXC Docker compose container. This repository is my own <strong>WIP</strong> installation guide.
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
    <a href="https://github.com/KelyanDev/Homelab/blob/main/proxy/README.md">Proxy</a>
    ·
    <a href="https://github.com/KelyanDev/Homelab/blob/main/cloud/README.md"><strong>Cloud</strong></a>
    ·
    <a href="">WIP</a>
  </p>
</div>

## Installation

These steps assume that you have already configured the debian LXC container, and that you are connected as the root user in them.

The original documentations might as well be a better source of information, depending on your tech stack.


### NextCloud

> **Specs**
> - 2 vCpu
> - 4 GiB memory
> - 40 GiB storage

These are the initial steps regarding the installation of Nextcloud. 
The installation is done by using the [Docker compose file](https://github.com/KelyanDev/Homelab/blob/main/cloud/docker-compose.yml) of this repository
```
# Update packages:
apt update && apt upgrade -y

# Install the dependencies
apt install ca-certificates curl gnupg lsb-release apt-transport-https -y

# Add the Docker key to verify downloads and add the repository:
curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] \
  https://download.docker.com/linux/debian $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null

# Finally, update your packages and install Docker & Docker compose
apt update && apt install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin -y
```

Once all these steps are done and the installation proceeded, we'll finally be able to start configure Nextcloud   

First, we'll have to prepare the folder in which we'll put the docker-compose.yml file.  
```
# Create the directory
mkdir -p /opt/nextcloud
cd /opt/nextcloud

# Create the docker-compose.yml file - my config file can be found at the beginning
nano docker-compose.yml

# Create and start the Nextcloud stack
docker compose up -d
```
Then, you can start to configure and play with your Nextcloud. The initial steps needs you to configure the Admin user and the database user.
