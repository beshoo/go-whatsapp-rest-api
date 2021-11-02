# go-whatsapp-rest-API
go-whatsapp-rest-API is a **Go library** for the WhatsApp web which use Swagger as api interface

# Multi-devices (MD) Support.
This version does NOT support multi-devices (MD)

# Multi-Session with singile go-whatsapp processes support.
Yes, you can run multi-sessions

# Sessions 
All sessions are stored in MYSQL

# The API support the following
```
​/profile​/scanqr [Scan the QR code via the browser]
​/profile​/phone​/isconnected
​/profile​/phone​/connect
​/profile​/phone​/disconnect
​/profile​/logout
​/profile​/me
​/profile​/hook​/set
​/profile​/contacts
​/send​/ack​/read
​/send​/text
​/send​/location
​/send​/image
​/send​/video
​/send​/audio
​/send​/audiorecord
​/send​/vcard
​/send​/link
​/send​/doc
```
# Web Hooks

```
/notify​/logout
/notify​/receive
/notify​/connectivity
/message​/text
/message​/image
/message​/video
/message​/audio
/message​/doc
/message​/contact
/power​/battery
/message​/location
/message​/livelocation
```

The API is built on [go-whatsapp], but it does not support the MD right now, any PL is welcome

[go-whatsapp]: https://github.com/Rhymen/go-whatsapp

# Prerequisites

  - [GoLang](https://golang.org/doc/install) Runtime environment 
  - [MYSQL](https://www.mysql.com/downloads/) 
  - [PM2](https://pm2.keymetrics.io/) install using command 
  wget -qO- https://getpm2.com/install.sh | bash


Build:
  - Build it for x86 architecture and linux os
  
```
sh ./build.sh
```

# Run
 - Copy wa-api executable to your server
 - Copy static folder for docs
 - Open static/qrcode/index.html and update the [ip_address] to match your deployement server ip address or domian name.
 - Change the url to point to deployement server in static/qrcode/index.html for qr to work
 - Create a MySql DB eg

 ```
mysql -u root -p
create database wapidb;
```
Run the server using pm2 with yout db user and password
```
pm2 start ./wa-api -- --port {Port Name} --dburl {DBUSER:DBPASSWORD@/DBNAME} --fromMe  --timeout 60
Eg :
pm2 start ./wa-api --  --dburl beshoo_wapidb:_5,v3NaE+0Os@/beshoo_wapidb --fromMe  --timeout 60
```

### Available Flags

| Flag | Function | Default|
| ------ | ------ |--------|
| port | The port on which the server will run | 3000 |
| dburl | Database access url |The sytem will not persist any session|
| fromMe | If the hook should recieve message from your own number | false|
| timeout | The time in seconds, API should wait before droping the connection while QR scan and session restored | 5 |
