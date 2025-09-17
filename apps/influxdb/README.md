<h3 align="center">InfluxDB - Configuration</h3>

  <p align="center">
    Every step in installing InfluxDB for the first time
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

