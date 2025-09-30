<h3 align="center">Services</h3>

  <p align="center">
    In this page, you'll find avery general informations regarding all of the different services that I'm currently running on my Homelab. Each service will (probably) have its own set of scripts and general informations on how to run them, even if I'm not sure how to organize it yet
    <br />
    <a href="https://github.com/KelyanDev/Homelab"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    Navigation <br />
    <a href="https://github.com/KelyanDev/Homelab">General</a>
    ·
    <a href="https://github.com/KelyanDev/Homelab/blob/main/apps/README.md"><strong>Apps</strong></a>
    ·
    <a href="https://github.com/KelyanDev/Homelab/blob/main/monitoring/README.md">Monitoring</a>
    ·
    <a href="https://github.com/KelyanDev/Homelab/blob/main/proxy/README.md">Proxy</a>
    ·
    <a href="">WIP</a>
    ·
    <a href="">WIP</a>
  </p>
</div>

## DNS

### Pi-Hole

Probably the most known DNS sinkhole (or network-wide ad-blocker), designed to protect devices on a private network from unwanted content by blocking certain DNS inquiries, using blacklists full of domains. Pi-Hole also has other multiple use cases, as I have personally configured mine as a classic DNS server. I haven't checked all of its functionalities yet, but it is surely a wonderful service to run and play with.

<div align="center">
  <br />
  <img src="images/pihole-dashboard.png" alt="Logo" width="900"/>
  <br />
</div>

<div align="center">
  <a href="https://github.com/KelyanDev/Homelab/blob/main/apps/pi-hole/README.md">Configuration</a>
  ·
  <a href="https://github.com/pi-hole/pi-hole">GitHub</a>
</div>

<br />

## Metrics & Monitoring

### Grafana

Grafana is a stunning dashboard tool that can translate various databases and showcase it into beautiful graphs, charts, or even presentations. There is a ton of custom dashboard created by the community for different use cases, and it can integrates with a lot of tech stacks. You can easily add a connection and import a dashboard to display the data.

<div align="center">
  <br />
  <img src="images/grafana-dashboard.png" alt="Logo" width="900"/>
  <br />
</div>

<div align="center">
  <a href="https://github.com/KelyanDev/Homelab/blob/main/monitoring/README.md">Configuration</a>
  ·
  <a href="https://github.com/grafana/grafana">GitHub</a>
</div>

<br />

### InfluxDB2

InfluxDB is an open-source time series database designed for high-speed ingestion and efficient querying of time-stamped data. This database tool is also capable of displaying the data into charts and graphs, but I'm mainly using it to gather metrics from my Proxmox server, before pulling it to Grafana dashboards

<div align="center">
  <br />
  <img src="images/influxdb-dashboard.png" alt="Logo" width="900"/>
  <br />
</div>

<div align="center">
  <a href="https://github.com/KelyanDev/Homelab/blob/main/monitoring/README.md">Configuration</a>
  ·
  <a href="https://github.com/influxdata/influxdb">GitHub</a>
</div>
