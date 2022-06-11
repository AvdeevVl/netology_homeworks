# Домашнее задание к занятию "3.8. Компьютерные сети, лекция 3"

1. Подключитесь к публичному маршрутизатору в интернет. Найдите маршрут к вашему публичному IP
```
telnet route-views.routeviews.org
Username: rviews
show ip route x.x.x.x/32
show bgp x.x.x.x/32
```

Ответ:  
![2022-06-11_21-22-06](https://user-images.githubusercontent.com/93952387/173201128-d7aad33c-8dbb-4ea3-b873-c6074bdec392.png)  

![2022-06-11_21-22-45](https://user-images.githubusercontent.com/93952387/173201147-b4584993-66fb-45b6-9242-c9cf962d78ff.png)  

![2022-06-11_21-24-53](https://user-images.githubusercontent.com/93952387/173201155-07b0ea41-ae84-4942-99d0-70e47e20c506.png)  


---

2. Создайте dummy0 интерфейс в Ubuntu. Добавьте несколько статических маршрутов. Проверьте таблицу маршрутизации.

Ответ:  

Добавляем `dummy` интерфейс:  
`sudo ip link add dummy0 type dummy`  
Назначаем адрес интерфейсу:  
`sudo ip addr add 192.168.9.0/24 dev dummy0`   
Включаем интерфей:  
`sudo ip link set dummy0 up`  
Проверяем:  
`sudo ip address`  

![2022-06-11_21-37-35](https://user-images.githubusercontent.com/93952387/173201185-761b4226-b904-4de0-8604-573a4611f805.png)

Добавляем маршруты:  
`sudo ip route add 8.8.8.0/24 via 192.168.0.1`  
`sudo ip route add 8.16.28.0/24 via 192.168.9.0`  
Проверяем:  
`sudo ip route`  

![2022-06-11_21-38-58](https://user-images.githubusercontent.com/93952387/173201195-125cabcb-1051-4cb4-a490-f8693bb45aae.png)


---

3. Проверьте открытые TCP порты в Ubuntu, какие протоколы и приложения используют эти порты? Приведите несколько примеров.

Ответ:  
`ss -tnlp`  

![2022-06-11_19-31-26](https://user-images.githubusercontent.com/93952387/173201242-00add586-3198-48d8-8a4d-b57375babe77.png)  


Используему порты приложениями:  
53 - DNS  
22 - SSH  

4. Проверьте используемые UDP сокеты в Ubuntu, какие протоколы и приложения используют эти порты?

Ответ:  
`ss -unap`  

![2022-06-11_19-33-23](https://user-images.githubusercontent.com/93952387/173201252-5f35330e-e872-4c27-a046-78dd1f88092f.png)  

Используему порты приложениями:  
53 - DNS  
68 - DHCP  

---

5. Используя diagrams.net, создайте L3 диаграмму вашей домашней сети или любой другой сети, с которой вы работали. 

Ответ:  

![2022-06-11_21-12-53](https://user-images.githubusercontent.com/93952387/173201273-7c6072fa-0c04-4a32-a27d-fc806ea9606a.png)


---
