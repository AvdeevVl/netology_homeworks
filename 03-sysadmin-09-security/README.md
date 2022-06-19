# Домашнее задание к занятию "3.9. Элементы безопасности информационных систем"

1. Установите Bitwarden плагин для браузера. Зарегестрируйтесь и сохраните несколько паролей.

Ответ: 



---

2. Установите Google authenticator на мобильный телефон. Настройте вход в Bitwarden акаунт через Google authenticator OTP.

Ответ: 



---

3. Установите apache2, сгенерируйте самоподписанный сертификат, настройте тестовый сайт для работы по HTTPS.

Ответ:  

`sudo apt install apache2`  
`sudo openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout /etc/ssl/private/apache-selfsigned.key -out /etc/ssl/certs/apache-selfsigned.crt`  

`sudo nano /etc/apache2/conf-available/ssl-params.conf`  

`sudo nano /etc/apache2/sites-available/default-ssl.conf`  

`sudo nano /etc/apache2/sites-available/000-default.conf`  

`sudo a2enmod ssl`  
`sudo a2enmod headers`  
`sudo a2ensite default-ssl`  
`sudo a2enconf ssl-params`  
`sudo apache2ctl configtest`  
`sudo systemctl restart apache2`  



---

1. Проверьте на TLS уязвимости произвольный сайт в интернете (кроме сайтов МВД, ФСБ, МинОбр, НацБанк, РосКосмос, РосАтом, РосНАНО и любых госкомпаний, объектов КИИ, ВПК ... и тому подобное).

Ответ: 

`docker run --rm -ti  drwetter/testssl.sh -U --sneaky https://ya.ru`



---

5. Установите на Ubuntu ssh сервер, сгенерируйте новый приватный ключ. Скопируйте свой публичный ключ на другой сервер. Подключитесь к серверу по SSH-ключу.

Ответ:  

* Установка ssh сервера на Ubuntu:  
`sudo apt install openssh-server`

* Генерируем открытый ssh ключ:  
`ssh-keygen`  



* Копируем ключ на сервер:  
`ssh-copy-id -i .ssh/id_rsa gips@192.168.0.12`  
* Проверяем подключение:  
`ssh gips@192.168.0.12` 



---

6. Переименуйте файлы ключей из задания 5. Настройте файл конфигурации SSH клиента, так чтобы вход на удаленный сервер осуществлялся по имени сервера.

Ответ:  

* Переименовываем файлы ключей из задания 5:  
`sudo mv ~/.ssh/id_rsa ~/.ssh/id_rsa_HW39`  

* Настраиваем файл конфигурации SSH клиента:  
`sudo nano ~/.ssh/config`  

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

---

7. Соберите дамп трафика утилитой tcpdump в формате pcap, 100 пакетов. Откройте файл pcap в Wireshark.

Ответ:  

`tcpdump -nnei any -c 100 -w 100packets.pcap`



---