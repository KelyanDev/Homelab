<h3 align="center">DNS & Proxy setting</h3>

  <p align="center">
    Regarding my Proxy setup, I'm running a Nginx Proxy Manager LXC with Docker and Docker compose. This repository is my own <strong>Work in Progress</strong> installation guide, featuring the knowledge linked to my reverse proxy setup and how I access my services from outside
    For the DNS part, as my public IP is dynamic, I had to setup a cloudflare ddns. I setup my own dns updater using a self made script in Go, which can be found here
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

For Nginx Proxy manager, this guide assumes that you already have a Debian 13 LXC running. For the DDNS, I used an Alpine linux because of its lightness.  

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


### DDNS

> **Specs**
> - 1 vCpu
> - 128 MiB memory
> - 800 MiB storage

These are the initial steps regarding the creation of a DDNS to update cloudflare.      
We'll need to install multiple things to be able to run our Go script       
   
```
# Install Go, Curl and Nano:
apk add go curl nano

# Create a repository for our ddns
mkdir /opt/ddns
cd /opt/ddns

# Create and modify the Go script
touch main.go && nano main.go
```
Once you installed Go, created your repository and your script file, you can check my [custom script](https://github.com/KelyanDev/Homelab/blob/main/proxy/ddns/main.go)     

The particularity of this script is that you can add a discord webhook link to send you a message when your public ip changes. This part can be removed from the script        

Once the script has been created, we simply need to compile it and use cron to automatically executes it every 5 minutes      
```
# CCompile / Build our Go script to make it runnable by Alpine
go build -o ddns

# Verify if crond is already installed:
crond --version

# Activates crond to make it starts when the LXC starts
rc-update add crond
rc-status

# Starts crond
rc-service crond start

# Edit the cron tasks
crontab -e

# Finally, add the following line
*/5 * * * * /opt/ddns/ddns >> /opt/ddns/ddns.log 2>&1

# To make sure that your cron has been successfully configured, use the following command. If you see the line you added, it should be good.
crontab -l

# Restart your cron (just in case)
rc-service crond restart
```
Now, your ddns should be correctly configured if you followed every step     

To check if your ddns is up and running, just wait 5 minutes and then check your log file. You should see a line specifying that your IP has changed

