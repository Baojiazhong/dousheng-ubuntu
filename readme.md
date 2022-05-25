# user service

## version 1 
- register and auto login
- login

### start
#### 1. Setup Basic Dependence
```shell
sudo docker compose up
```

#### 2.Run User RPC Server
```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

#### 3.Run API Server
```shell
cd cmd/api
chmod +x run.sh
sh run.sh
```

### API test
postman desktop
#### 1.register

<img src="images/shot_register.jpg" width="2850"  alt=""/>

#### 2.login
<img src="images/shot_login.jpg" width="2850"  alt=""/>
