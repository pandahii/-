# docker
### docker clean build container
```
docker rmi -f $(docker images --filter dangling=true -qa)
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
### wifi
```
[root@localhost ~]# nmcli dev show
[root@localhost ~]# nmtui
[root@localhost ~]# ip link [dev name] set up
[root@localhost ~]# chkconfig NetworkManaget on
```
