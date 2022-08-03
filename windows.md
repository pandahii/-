## 磁盘格式化 覆写

### D 盘用随机数覆盖 8 次
format D:  /P:8
### D 盘未使用空间会被覆盖三次，一次 0x00、一次 0xFF，一次随机数
cipher /w:D:\Private
