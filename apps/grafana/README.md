<h3 align="center">Grafana - Configuration</h3>

  <p align="center">
    Every step in installing Grafana for the first time
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

# You can also remove the grafana 
```
Then, you can start to configure your Grafana. It usually runs on the port 3000, so if you configured yours correctly and applied the port, you should be able to access it directly from here

