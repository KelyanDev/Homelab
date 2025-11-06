<h3 align="center">DNS & Proxy setting</h3>

  <p align="center">
    Regarding my Proxy setup, I'm running a Nginx Proxy Manager LXC with Docker and Docker compose. This repository is my own <strong>Work in Progress</strong> installation guide, featuring the knowledge linked to my reverse proxy setup and how I access my services from outside, but also my Pi-hole config     
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

The installation steps of each service assume that your LXC is already configured, and that you're connected to it as the root user (no sudo in here, so double check before following this)     
Regarding these 3 services, here is my recommandation for the OS to run them:
- **Pi-Hole / Nginx Proxy Manager** - Latest Debian version, for its compatibility and how simple it is to use
- **CloudFlare DDNS** - Latest Alpine version, for its lightness and its efficiency

The original documentations might as well be a better source of information, depending on your tech stack.

### Nginx Proxy Manager

> **Specs**
> - 2 vCpu
> - 2 GiB memory
> - 10 GiB storage

These are the initial steps regarding the installation of Nginx Proxy Manager.    
The installation is made using this [Docker compose file](https://github.com/KelyanDev/Homelab/blob/main/proxy/npm/docker-compose.yml). You can modify it to create your own, or find an alternative on Internet.   
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

<hr />

### Pi-Hole

> **Specs**
> - 1 vCpu
> - 512 MiB memory
> - 2 GiB storage

These are the initial steps regarding the creation of a Pi-Hole instance                   
   
```
# Update packages:
apt update -y && apt upgrade -y

# CInstall dependencies
apt install curl gnupg lsb-release ca-certificates -y

# Use the official installation script to install Pi-Hole - MAKE SURE THAT THE SCRIPT IS SAFE
curl -sSL https://install.pi-hole.net | bash
```
Once you installed Pi-Hole, the service should be up and running on the default http port on your host; Just add /admin to the link to access the admin panel (for example: http://10.0.0.1/admin)         

The installation script will generate a default admin password for you - You should probably change it once you connected for the first time     
It also provides a default blocklist, that you can remove if you want     

Play with the configuration, make sure that it can access Internet and upstream DNS servers, and you can configure it as your network DNS server to let equipments use it ! 

You can find a lot of blocklist online, add them whenever you want to block more things - you can also block your own domains without the need of a list     

<hr />

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

