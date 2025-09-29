<h3 align="center">Monitoring</h3>

  <p align="center">
    In this page, you'll find avery general informations regarding all of the different services that I'm currently running on my Homelab. Each service will (probably) have its own set of scripts and general informations on how to run them, even if I'm not sure how to organize it yet
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
    <a href="">WIP</a>
  </p>
</div>

## Grafana

There is multiple steps in installing Grafana. These steps assume that you already have a running Debian 12 LXC   
The Grafana documentations can also provide great amount of informations regarding the installation steps, but they assume that you're not using the root user   

Also, it is not possible to run a simple Grafana service without going with Docker if the LXC container is unprivilegied, because Grafana will have some trouble when trying to start.

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
Once Docker is finally installed, you just need to create a Grafana container, using the parameters you desire. Here's what I used to create my Grafana's container:
```
# Command to create a Grafana container named "grafana", always restarting unless manually stopped.
# Grafana configs / dashboards are stored in /var/lib/grafana, in case the container is deleted.
docker run -d --name=grafana --restart=unless-stopped -p 3000:3000 -v /opt/grafana/data:/var/lib/grafana grafana/grafana:latest

# If you run Grafana the same way I did, you'll also need to do the following:
mkdir -p /opt/grafana/data
chown -R grafana:grafana /opt/grafana/data
chmod -R 755 /opt/grafana/data
```
Then, you can start to configure your Grafana. It usually runs on the port 3000, so if you configured yours correctly and applied the port, you should be able to access it directly from here

## Dashboard

You can find my custom dashboard here:
[Dashboard](https://github.com/KelyanDev/Homelab/blob/main/apps/grafana/proxmox-ve.json)

Here's the original Dashboard I used; It is almost the same, I just tweaked some things here and there:
[Original](https://grafana.com/grafana/dashboards/23164-proxmox-ve/)


## InfluxDB

There is multiple steps in installing InfluxDB. These steps assume that you already have a running Debian 12 LXC
The InfluxDB documentations can also provide great amount of informations regarding the installation steps, but they assume that you're not using the root user

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

Once all these steps are done and the installation proceeded, you are safe to start the service. Don't forget to enable it if you want it to start when the LXC containers start
```
# Enable the service:
systemctl enable influxdb

# Start the service
systemctl start influxdb
```
Then, you can start to configure your InfluxDB. InfluxDB generally runs on the port 8086 of your container, so you'll need to access the web interface using this port to configure InfluxDB directly   
The first steps during the initial configuration of InfluxDB consists of creating your admin user, as well as your first Organization and its first bucket
