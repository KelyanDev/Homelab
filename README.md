<a id="readme-top"></a>

<!--[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url] -->

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/KelyanDev/Homelab">
    <img src="images/logo.png" alt="Logo" width="160" height="160">
  </a>

<h3 align="center">Homelab</h3>

  <p align="center">
    My own homelab project, as a way to self-host services and applications, as well as learn and experiment with technologies. This project allows me to both use my network skills and my programming skills to build my own infrastructure.
    <br />
    <a href="https://github.com/KelyanDev/Homelab"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/KelyanDev/Homelab">View Demo</a>
    ·
    <a href="https://github.com/KelyanDev/Homelab/issues/new?labels=bug&template=bug-report---.md">Report Bug</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#infrastructure">Infrastructure</a></li>
        <li><a href="#services">Services</a></li>
      </ul>
    </li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

In this github repository, I'll centralize all of the informations regarding my current homelab setup (from hardware to software), my goals regarding it, aswell as some reusable configuration files.

> What is a Homelab ?
>
> A Homelab is a laboratory at home where you can self-host, experiment with new technologies, practice for certifications, and so on.
> For more information, please check the [r/homelab introduction](https://www.reddit.com/r/homelab/wiki/introduction/)

This project is, of course, still **Work In Progress**

<p align="right">(<a href="#readme-top">back to top</a>)</p>


### Infrastructure

#### SER8 (Proxmox)  
This machine is running my Proxmox server. Handling my windows servers, as well as my different services running on containers.
* Ryzen 7 8845HS - 8 cores, 16 threads
* 32Go DDR5 5600 MHz
* 1To NVMe SSD


<p align="right">(<a href="#readme-top">back to top</a>)</p>


### Services

#### Monitoring
As my monitoring service, I intend to use Grafana with either Prometheus or InfluxDB (I'm still not sure what I want to use)


### Experimental
For the most experimental part of this Homelab, I intend to run at least 2 windows servers:
* AD DS/DNS server - To experiment a more "enterprise" infrastructure, and also used to experiment with my Veeam server
* Veeam server - To experiment and increase my skills towards Veeam infrastructures and backups, mainly focusing on Veeam Backup and Replication while I experiment with Veeam Backup for Office 365 at work?


<!-- ROADMAP -->
## Roadmap

### Services roadmap

- [X] Proxmox VE configuration
- [ ] Metric collector / Monitoring
    - [ ] Monitoring
    - [ ] Logging
    - [ ] Alerting
- [ ] Backup solutions
    - [ ] Veeam server configuration
    - [ ] Proxmox Backup Server
- [ ] Dashboard
- [ ] Proxy manager
    - [ ] Certificates manager
- [ ] DNS Sinkhole
- [ ] VPN
- [ ] Self hosted cloud

### Physical infrastructure roadmap

- [X] Proxmox server
- [ ] Switch
- [ ] NAS

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

My Portfolio - [Click here](https://kelyandev.github.io/)

Project Link: [https://github.com/KelyanDev/Homelab](https://github.com/KelyanDev/Homelab)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* [Proxmox Documentation](https://pve.proxmox.com/wiki/Main_Page)
* [Grafana's Documentation](https://grafana.com/docs/)
* [TechHutTV's Homelab & Documentations](https://github.com/TechHutTV/homelab)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/KelyanDev/Homelab.svg?style=for-the-badge
[contributors-url]: https://github.com/KelyanDev/Homelab/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/KelyanDev/Homelab.svg?style=for-the-badge
[forks-url]: https://github.com/KelyanDev/Homelab/network/members
[stars-shield]: https://img.shields.io/github/stars/KelyanDev/Homelab.svg?style=for-the-badge
[stars-url]: https://github.com/KelyanDev/Homelab/stargazers
[issues-shield]: https://img.shields.io/github/issues/KelyanDev/Homelab.svg?style=for-the-badge
[issues-url]: https://github.com/KelyanDev/Homelab/issues
[license-shield]: https://img.shields.io/github/license/KelyanDev/Homelab.svg?style=for-the-badge
[license-url]: https://github.com/KelyanDev/Homelab/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/linkedin_username

[Java.com]: https://img.shields.io/badge/java-%23ED8B00.svg?style=for-the-badge&logo=openjdk&logoColor=white
[Java-url]: [https://nextjs.org/](https://www.w3schools.com/java/)
[Firebase.com]: https://img.shields.io/badge/firebase-ffca28?style=for-the-badge&logo=firebase&logoColor=black
[Firebase-url]: [https://reactjs.org/](https://firebase.google.com/)
