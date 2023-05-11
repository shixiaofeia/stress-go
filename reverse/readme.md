## http压测结果

### 服务器配置
> 代理端, 源站, 压测机均为同样配置: Centos 8C16G

压测机 192.168.3.28
反向代理 192.168.3.29
源站 192.168.3.30

### Fast => Fast

![fast_fast.png](fast_fast.png)

max: 208092

min: 192725

### Fast => Net

![fast_net.png](fast_net.png)

max: 174800

min: 161833

### Hertz => Fast

![hertz_fast.png](hertz_fast.png)

max: 165763

min: 131759

### Hertz => Net

![hertz_net.png](hertz_net.png)

max: 126939

min: 125850


### Net => Fast

![net_fast.png](net_fast.png)

max: 40280

min: 40103


### Net => Net

![net_net.png](net_net.png)

max: 49547

min: 47532
