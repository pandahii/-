# docker
### docker clean build container
```
docker rmi -f $(docker images --filter dangling=true -qa)
```

### docker install by yum
```
[root@localhost ~]# yum install -y yum-utils
[root@localhost ~]# yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
[root@localhost ~]# yum install -y docker-ce docker-ce-cli containerd.io docker-compose-plugin 
[root@localhost ~]# mkdir /etc/docker
[root@localhost ~]# echo "{"data-root": "/home/docker-root"}" >> /etc/docker/daemon.json
[root@localhost ~]# systemctl enable docker
[root@localhost ~]# systemctl start docker
[root@localhost ~]# docker info
```


# centos 7
### 修改history 格式
1. 内容
```
[root@localhost ~]# echo "HISTFILESIZE=100000\n \
HISTSIZE=100000\n \
USER_IP=`who -u am i 2>/dev/null| awk '{print $NF}'|sed -e 's/[()]//g'` \n \
if [ -z $USER_IP ] \n \
then \n \
  USER_IP=`hostname` \n \
fi \n \
HISTTIMEFORMAT="%F %T $USER_IP [`whoami`] " \n \
export HISTTIMEFORMAT" >> /etc/profile
[root@localhost ~]# source /etc/profile
[root@localhost ~]# history
```
### contos7 connect wifi
```
[root@localhost ~]# nmcli dev show
[root@localhost ~]# nmtui
[root@localhost ~]# ip link [dev name] set up
[root@localhost ~]# chkconfig NetworkManaget on
[root@localhost ~]# wpa_supplicant -B -i wlp3s0 -c <(wpa_passphrase "[SSID]" "[PIN]")
```
