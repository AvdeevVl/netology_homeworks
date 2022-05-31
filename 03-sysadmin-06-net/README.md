# Домашнее задание к занятию "3.6. Компьютерные сети, лекция 1"

1. Работа c HTTP через телнет.
- Подключитесь утилитой телнет к сайту stackoverflow.com
`telnet stackoverflow.com 80`
- отправьте HTTP запрос
```bash
GET /questions HTTP/1.0
HOST: stackoverflow.com
[press enter]
[press enter]
```
- В ответе укажите полученный HTTP код, что он означает?  

Ответ:  

`HTTP/1.1 301 Moved Permanently` - стандартный код ответа HTTP сервера запрошенный ресурс был на постоянной основе перемещён в новое месторасположение.  
Новое месторасположение ресуса https://stackexchange.com  

![2022-05-31_16-27-04](https://user-images.githubusercontent.com/93952387/171193218-90844a89-efb8-44ae-91b0-bf172c576d2f.png)

---

2. Повторите задание 1 в браузере, используя консоль разработчика F12.
- откройте вкладку `Network`
- отправьте запрос http://stackoverflow.com
- найдите первый ответ HTTP сервера, откройте вкладку `Headers`
- укажите в ответе полученный HTTP код.
- проверьте время загрузки страницы, какой запрос обрабатывался дольше всего?
- приложите скриншот консоли браузера в ответ.

Ответ:  
Первый код ответа HTTP сервера  
`Status Code: 301 Moved Permanently`  

![2022-05-31_16-52-49](https://user-images.githubusercontent.com/93952387/171193526-03b6b1db-4af9-4b00-a132-879b63c61926.png)  

Время загрузки страницы  
`Finish: 1.19 s`

Дольше всего обрабатывался запрос:  
Документ stackoverflow.com 541 ms  

![2022-05-31_16-52-00](https://user-images.githubusercontent.com/93952387/171193595-eb386f47-78e4-4aba-9ce0-c8243c3b3869.png)

---

3. Какой IP адрес у вас в интернете?

Ответ:  
`wget -qO- eth0.me`  

![2022-05-31_15-20-55](https://user-images.githubusercontent.com/93952387/171193833-46bed061-7dfb-4b0c-be0f-abff7c930de7.png)

---

4. Какому провайдеру принадлежит ваш IP адрес? Какой автономной системе AS? Воспользуйтесь утилитой `whois`

Ответ:  
`whois 89.109.14.129`  
```
descr:          Network for OJSC VolgaTelecom
origin:         AS25405
```

![2022-05-31_15-27-33](https://user-images.githubusercontent.com/93952387/171195398-a6c18c6c-d3d2-4476-977b-2969b19645ae.png)

---

5. Через какие сети проходит пакет, отправленный с вашего компьютера на адрес 8.8.8.8? Через какие AS? Воспользуйтесь утилитой `traceroute`

Ответ:  
`traceroute -An 8.8.8.8`

```
AS12389
AS15169
```

![2022-05-31_15-44-10](https://user-images.githubusercontent.com/93952387/171195453-b586655b-ae72-4d84-bb5e-cb2ca66940ff.png)

---

6. Повторите задание 5 в утилите `mtr`. На каком участке наибольшая задержка - delay?

Ответ:  
`mtr 8.8.8.8 -znrc 1`  
Наибольшая задержка:
 ```
  6. AS15169  216.239.49.19        0.0%     1  102.3 102.3 102.3 102.3   0.0
```

![2022-05-31_15-52-24](https://user-images.githubusercontent.com/93952387/171195571-fb49bdf0-c65f-4365-8fe0-70ba63f104c1.png)

---

7. Какие DNS сервера отвечают за доменное имя dns.google? Какие A записи? воспользуйтесь утилитой `dig`

Ответ:  

```
gips@ud22:~$ dig +short NS dns.google
ns4.zdns.google.
ns1.zdns.google.
ns3.zdns.google.
ns2.zdns.google.
```

```
gips@ud22:~$ dig +short A dns.google
8.8.8.8
8.8.4.4
gips@ud22:~$ 
```

---

8. Проверьте PTR записи для IP адресов из задания 7. Какое доменное имя привязано к IP? воспользуйтесь утилитой `dig`

Ответ:  
`dig +short -x 8.8.8.8`  
`dig +short -x 8.8.4.4`  

![2022-05-31_16-19-59](https://user-images.githubusercontent.com/93952387/171196404-2bf53ce5-786c-4394-be91-7b22f947d5ca.png)  

