<h3 align="center">Monitoring</h3>

  <p align="center">
    Regarding my monitoring stack, I'm running 2 Debian LXC - One running InfluxDB2 without Docker, while the other one is running Grafana in Docker. This repository is my own <strong>WIP</strong> installation guide. Some important aspects such as securing access using SSL/TLS connexions are yet to come.
    <br />
    <a href="https://github.com/KelyanDev/Homelab"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    Navigation <br />
    <a href="https://github.com/KelyanDev/Homelab">General</a>
    ·
    <a href="https://github.com/KelyanDev/Homelab/blob/main/apps/README.md">Apps</a>
    ·
    <a href="https://github.com/KelyanDev/Homelab/blob/main/monitoring/README.md"><strong>Monitoring</strong></a>
    ·
    <a href="https://github.com/KelyanDev/Homelab/blob/main/proxy/README.md">Proxy</a>
    ·
    <a href="">WIP</a>
    ·
    <a href="">WIP</a>
  </p>
</div>

## Installation

These steps assume that you have already configured 2 debian LXC containers, and that you are connected as the root user in them.

The original documentations might as well be a better source of information, depending on your tech stack.


### InfluxDB

> **Specs**
> - 2 vCpu
> - 2 GiB memory
> - 10 GiB storage

These are the initial steps regarding the installation of InfluxDB2 natively.   
```
# Update packages:
apt update && apt upgrade -y

# Install curl and gpg for the next command
apt install curl gpg -y

# Add the InfluxData key to verify downloads and add the repository:
curl --silent --location -O https://repos.influxdata.com/influxdata-archive.key
gpg --show-keys --with-fingerprint --with-colons ./influxdata-archive.key 2>&1 \
| grep -q '^fpr:\+24C975CBA61A024EE1B631787C3D57159FC2F927:$' \
&& cat influxdata-archive.key \
| gpg --dearmor \
| tee /etc/apt/keyrings/influxdata-archive.gpg > /dev/null \
&& echo 'deb [signed-by=/etc/apt/keyrings/influxdata-archive.gpg] https://repos.influxdata.com/debian stable main' \
| tee /etc/apt/sources.list.d/influxdata.list

# Finally, update your packages and install influxdb
apt update && apt install influxdb2 -y
```

Once all these steps are done and the installation proceeded, you are safe to start the service.    

Don't forget to enable it if you want it to start when the LXC containers start   
```
# Enable the service:
systemctl enable influxdb

# Start the service
systemctl start influxdb
```
Then, you can start to configure your InfluxDB, which should run on the port 8086 of your LXC, so you'll need to access the web interface using this port to configure InfluxDB directly  

The first steps during the initial configuration of InfluxDB consists of creating your admin user, as well as your first Organization and its first bucket    


### Grafana

> **Specs**
> - 1 vCpu
> - 512 MiB memory
> - 3 GiB storage     

As it is impossible to run Grafana natively if the LXC container is unprivilegied, we'll use docker to properly install it.     
```
# Update packages:
apt update && apt upgrade -y

# Install the dependencies
apt install -y apt-transport-https ca-certificates curl gnupg lsb-release

# Add the Docker key to verify downloads:
curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

# Add the Docker depot

# Finally, update your packages and install docker and its components
apt update && apt install -y docker-ce docker-ce-cli containerd.io
```
Once all these steps are done and the installation proceeded, you can verify if docker is successfully installed by running ``docker --version``     

Then, you just need to create a Grafana container, using the parameters you desire. Here's what I used to create my Grafana's container:
```
# Command to create a Grafana container named "grafana", always restarting unless manually stopped.
# Grafana configs / dashboards are stored in /var/lib/grafana, in case the container is deleted.
docker run -d --name=grafana --restart=unless-stopped -p 3000:3000 -v /opt/grafana/data:/var/lib/grafana grafana/grafana:latest

# If you run Grafana the same way I did, you'll also need to do the following:
mkdir -p /opt/grafana/data
chown -R grafana:grafana /opt/grafana/data
chmod -R 755 /opt/grafana/data
```
Then, you can start to configure your Grafana, which should run on the port 3000, but can vary depending on your configuration.     

## Dashboards

The entire point of this is to have beautiful dashboards. There are plenty online, but here's the one I'm currently using. I've redone some aspects of the original one, so you'll have the option of mine, or the original     

[Proxmox VE - InfluxDB2](https://grafana.com/grafana/dashboards/23164-proxmox-ve/) (Flux)    
[Custom version - Proxmox cluster](https://github.com/KelyanDev/Homelab/blob/main/apps/grafana/proxmox-ve.json) (Flux)     



