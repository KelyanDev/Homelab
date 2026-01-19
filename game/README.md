<h3 align="center">Game</h3>

  <p align="center">
    Regarding my game servers, I'm running a Pterodactyl stack featuring 2 LXC, one for the Pterodactyl Panel and one for Pterodactyl Wings - This setup allows me to easily manage game servers. This repository is my own <strong>light</strong> guide to help me remember the setup process.
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
    <a href="https://github.com/KelyanDev/Homelab/blob/main/proxy/README.md">Network</a>
    ·
    <a href="https://github.com/KelyanDev/Homelab/blob/main/cloud/README.md">Cloud</a>
    <br />
    <a href="https://github.com/KelyanDev/Homelab/blob/main/media/README.md">Media</a>
    ·
    <a href="https://github.com/KelyanDev/Homelab/blob/main/storage/README.md">Storage</a>
    ·
    <a href="https://github.com/KelyanDev/Homelab/blob/main/game/README.md"><strong>Game<strong/></a>
    ·
    <a href="">WIP</a>
  </p>
</div>

## Installation

These steps assume that you have already configured the debian LXC container, and that you are connected as the root user in them.

I recommend checking out the [official documentation](https://pterodactyl.io/project/terms.html) of Pterodactyl, in which you'll find plenty of informations you might need.     


### Pterodactyl Panel

> **Specs**
> - 1 vCpu
> - 2 GiB memory
> - 4 GiB storage

These are the initial steps regarding the installation of Pterodactyl Panel. 

Multiple steps are required to configure this panel, so make sure to do all of these steps.      

#### First part - Panel configuration

First, we'll install the required packages (you might need to add Redis and MariaDB official repo for Debian 12 or lower):
```
# Install necessary packages
apt install -y curl ca-certificates gnupg2 sudo lsb-release

# Add additional repositories for PHP
echo "deb https://packages.sury.org/php/ $(lsb_release -sc) main" | sudo tee /etc/apt/sources.list.d/sury-php.list
curl -fsSL https://packages.sury.org/php/apt.gpg | sudo gpg --dearmor -o /etc/apt/trusted.gpg.d/sury-keyring.gpg

# Add Redis official APT repository (Debian 11 & 12 only)
curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/redis.list

# MariaDB repo setup script (Debian 11 & 12 only)
curl -LsS https://r.mariadb.com/downloads/mariadb_repo_setup | sudo bash

# Update repositories list & finish installing dependencies
apt update && apt install -y php8.3 php8.3-{common,cli,gd,mysql,mbstring,bcmath,xml,fpm,curl,zip} mariadb-server nginx tar unzip git redis-server

# Install Composer (dependency manager for php)
curl -sS https://getcomposer.org/installer | sudo php -- --install-dir=/usr/local/bin --filename=composer
```
Now, we should have every dependency we need.        
Let's start the installation & configuration of the actual panel:
```
# Create the directory
mkdir -p /var/www/pterodactyl
cd /var/www/pterodactyl

# Download the panel files, unpack the archive and set the correct permissions
curl -Lo panel.tar.gz https://github.com/pterodactyl/panel/releases/latest/download/panel.tar.gz
tar -xzvf panel.tar.gz
chmod -R 755 storage/* bootstrap/cache/
```
Now that we installed the panel, we'll start the actual configuration of our different components.       
First, let's start with MariaDB - We'll create the pterodactyl user and the database, and grant all privileges to our user on it:
```
mariadb -u root -p

#PHP
CREATE USER 'pterodactyl'@'127.0.0.1' IDENTIFIED BY 'YourSecret';
CREATE DATABASE panel;
GRANT ALL PRIVILEGES ON panel.* TO 'pterodactyl'@'127.0.0.1' WITH OPTION GRANT;
exit
```

Now, we'll setup our environment file:
```
# Copy the example environment file
cp .env.example .env

# Install core dependencies
COMPOSER_ALLOW_SUPERUSER=1 composer install --no-dev --optimize-autoloader

# Generate a new key for our panel - RUN ONLY IN CASE OF FIRST INSTALL / NO DATA IN DATABASE
php artisan key:generate --force
```

Then, we'll configure our .env - You can use the commands to configure it, like this:
```
### Configuration commands:
php artisan p:environment:setup
php artisan p:environment:database
php artisan p:environment:mail
```
Or you can also just change manually the required parameters:
```
### Configure it by hand
nano .env

# Make sure your .env file has these
APP_ENV=production
APP_DEBUG=false
APP_URL=http://PANEL_IP

DB_CONNECTION=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=panel
DB_USERNAME=pterodactyl
DB_PASSWORD=YourPassword

CACHE_DRIVER=file
QUEUE_CONNECTION=sync
SESSION_DRIVER=file
```

Then we need to setup the database for Pterodactyl, and configure our first user.
```
# Seed the MariaDB database
php artisan migrate --seed --force

# Create the user
php artisan p:user:make

# Set the correct ownership
chown -R www-data:www-data /var/www/pterodactyl/*
```

Then, finally, we'll just configure a cron to run every minute to run specific pterodactyl tasks (necessary for schedules to run)
```
# Edit crontab
crontab -e

# Crontab - add this line:
* * * * * php /var/www/pterodactyl/artisan schedule:run >> /dev/null 2>&1
```

#### Second part - Webserver Configuration

Now that we've setup the Pterodactyl Panel, we need to setup Nginx to correctly be able to access our panel.     
First, we'll remove the default nginx panel and create a new file for our panel:
```
# Remove default config
rm /etc/nginx/sites-enabled/default

# Create our new file and edit it
nano /etc/nginx/sites-available/pterodactyl
```
Now, we'll create the web interface for our panel. For this, add the following to the file we just created. You just need to modify the IP of the server:
```
server {
    listen 80;
    server_name PANEL_IP;

    root /var/www/pterodactyl/public;
    index index.html index.htm index.php;
    charset utf-8;

    location / {
        try_files $uri $uri/ /index.php?$query_string;
    }

    location = /favicon.ico { access_log off; log_not_found off; }
    location = /robots.txt  { access_log off; log_not_found off; }

    access_log off;
    error_log  /var/log/nginx/pterodactyl.app-error.log error;

    # allow larger file uploads and longer script runtimes
    client_max_body_size 100m;
    client_body_timeout 120s;

    sendfile off;

    location ~ \.php$ {
        fastcgi_split_path_info ^(.+\.php)(/.+)$;
        fastcgi_pass unix:/run/php/php8.3-fpm.sock;
        fastcgi_index index.php;
        include fastcgi_params;
        fastcgi_param PHP_VALUE "upload_max_filesize = 100M \n post_max_size=100M";
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        fastcgi_param HTTP_PROXY "";
        fastcgi_intercept_errors off;
        fastcgi_buffer_size 16k;
        fastcgi_buffers 4 16k;
        fastcgi_connect_timeout 300;
        fastcgi_send_timeout 300;
        fastcgi_read_timeout 300;
    }

    location ~ /\.ht {
        deny all;
    }
}
```
Finally, we just need to enable the configuration:
```
# Create a symbolic link to our pterodactyl
ln -s /etc/nginx/sites-available/pterodactyl.conf /etc/nginx/sites-enabled/pterodactyl.conf

# Restart Nginx
systemctl restart nginx
```

### Pterodactyl Wings

> **Specs**
> - 4 vCpu
> - 8 GiB memory
> - 50 GiB storage

These steps are necessary to configure Pterodactyl Wings on your LXC - they are a lot quicker to execute, and your Wings instance should be running pretty fast.

```
#WIP
```
