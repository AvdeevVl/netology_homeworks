# Домашнее задание к занятию "3.9. Элементы безопасности информационных систем"

1. Установите Bitwarden плагин для браузера. Зарегестрируйтесь и сохраните несколько паролей.

Ответ: 

![2022-06-19_21-38-40](https://user-images.githubusercontent.com/93952387/174499097-d466efb7-c39b-4f40-a82f-8093a3f5dc62.png)


---

2. Установите Google authenticator на мобильный телефон. Настройте вход в Bitwarden акаунт через Google authenticator OTP.

Ответ: 

![2022-06-19_21-45-27](https://user-images.githubusercontent.com/93952387/174499117-a183b0f6-6603-4697-95ce-8612f4179006.png)


---

3. Установите apache2, сгенерируйте самоподписанный сертификат, настройте тестовый сайт для работы по HTTPS.

Ответ:  

`sudo apt install apache2`  
`sudo openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout /etc/ssl/private/apache-selfsigned.key -out /etc/ssl/certs/apache-selfsigned.crt`  

![2022-06-19_22-45-31](https://user-images.githubusercontent.com/93952387/174499254-0a52bba0-7ec0-4c74-8f3f-9fe269253f55.png)


`sudo nano /etc/apache2/conf-available/ssl-params.conf`  

![2022-06-19_23-20-17](https://user-images.githubusercontent.com/93952387/174499280-91c5ae3e-bf34-4988-a16f-358105a8dd36.png)


`sudo nano /etc/apache2/sites-available/default-ssl.conf`  

![2022-06-19_23-21-16](https://user-images.githubusercontent.com/93952387/174499288-a17e5a9a-3de8-41c1-81e1-8870c42777de.png)

`sudo nano /etc/apache2/sites-available/000-default.conf`  

![2022-06-19_23-21-50](https://user-images.githubusercontent.com/93952387/174499294-c383ad9f-1fb4-47f2-92b6-c5dd8fd9a094.png)


`sudo a2enmod ssl`  
`sudo a2enmod headers`  
`sudo a2ensite default-ssl`  
`sudo a2enconf ssl-params`  
`sudo apache2ctl configtest`  
`sudo systemctl restart apache2`  

![2022-06-19_23-16-58](https://user-images.githubusercontent.com/93952387/174499299-a27018e4-028c-411a-9550-a864f674493d.png)

---

4. Проверьте на TLS уязвимости произвольный сайт в интернете (кроме сайтов МВД, ФСБ, МинОбр, НацБанк, РосКосмос, РосАтом, РосНАНО и любых госкомпаний, объектов КИИ, ВПК ... и тому подобное).

Ответ: 

`docker run --rm -ti  drwetter/testssl.sh -U --sneaky https://ya.ru`

![2022-06-19_21-52-52](https://user-images.githubusercontent.com/93952387/174499136-0337c205-b7f2-4911-b105-526ba86f2546.png)


---

5. Установите на Ubuntu ssh сервер, сгенерируйте новый приватный ключ. Скопируйте свой публичный ключ на другой сервер. Подключитесь к серверу по SSH-ключу.

Ответ:  

* Установка ssh сервера на Ubuntu:  
`sudo apt install openssh-server`

* Генерируем открытый ssh ключ:  
`ssh-keygen`  

![2022-06-19_22-13-13](https://user-images.githubusercontent.com/93952387/174499174-1ac70f70-15b5-4a20-8451-f2ed22f995bb.png)


* Копируем ключ на сервер:  
`ssh-copy-id -i .ssh/id_rsa gips@192.168.0.12`  
* Проверяем подключение:  
`ssh gips@192.168.0.12` 

![2022-06-19_22-15-06](https://user-images.githubusercontent.com/93952387/174499184-af68accd-5644-4750-9998-71314dc01be4.png)


---

6. Переименуйте файлы ключей из задания 5. Настройте файл конфигурации SSH клиента, так чтобы вход на удаленный сервер осуществлялся по имени сервера.

Ответ:  

* Переименовываем файлы ключей из задания 5:  
`sudo mv ~/.ssh/id_rsa ~/.ssh/id_rsa_HW39`  

* Настраиваем файл конфигурации SSH клиента:  
`sudo nano ~/.ssh/config`  

![2022-06-19_22-26-04](https://user-images.githubusercontent.com/93952387/174499218-0481f2be-1c7d-40cf-ac1a-b26fb9ba7059.png)  

* Добовляем содержимое:  
```
Host HW39
        HostName 192.168.0.12
        User gips
        Port 22
        IdentityFile ~/.ssh/id_rsa_HW39
```

* Проверяем подключение:  
`ssh HW39`  

![2022-06-19_22-28-10](https://user-images.githubusercontent.com/93952387/174499230-41f4230f-b31a-433f-a56c-cc0e2b0e5ca7.png)

---

7. Соберите дамп трафика утилитой tcpdump в формате pcap, 100 пакетов. Откройте файл pcap в Wireshark.

Ответ:  

`tcpdump -nnei any -c 100 -w 100packets.pcap`  

![2022-06-19_22-39-16](https://user-images.githubusercontent.com/93952387/174499392-5b4105db-2b39-493a-9b60-69538544be3b.png)  

![2022-06-19_22-40-48](https://user-images.githubusercontent.com/93952387/174499400-fb7c4015-a903-448c-ac84-5a3a695a294e.png)




---
