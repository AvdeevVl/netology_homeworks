
# Домашнее задание к занятию "5.2. Применение принципов IaaC в работе с виртуальными машинами"

## Задача 1

- Опишите своими словами основные преимущества применения на практике IaaC паттернов.
- Какой из принципов IaaC является основополагающим?


Ответ:  

- Основные преимущества применения на практике IaaC паттернов:
    -  Ускорение производства и вывода продукта на рынок.
    -  Стабильность среды, устранение дрейфа конфигураций.
    -  Более быстрая и эффективная разработка.

- Основопологающий принцип IaaC являеться: Идемпоте́нтность - это свойство объекта или операции, при повторном выполнении которой мы получаем результат идентичный предыдущему и всем последующим выполнениям.

## Задача 2

- Чем Ansible выгодно отличается от других систем управление конфигурациями?
- Какой, на ваш взгляд, метод работы систем конфигурации более надёжный push или pull?


Ответ:  
Преимущества Ansible:
- Работает без агента на клиентах, использует ssh для доступа на клиент
- Для описания конфигурационных файлов используется удобный для чтения формат YAML.
- При неуспешной доставке конфигурации на сервер, оповестит об этом.

Метод push не требует развертывания дополнительных агентов, которые выступают дополнительным звеном, снижающим надежность а так же централизованное управляет конфигурацией исключает ситуации, когда прямое изменение
на сервере не отразится в репозитории - что может привести к непредсказуемым ситуациям.


## Задача 3

Установить на личный компьютер:

- VirtualBox
- Vagrant
- Ansible

*Приложить вывод команд установленных версий каждой из программ, оформленный в markdown.*


Ответ:  

```sh
gips@UD22:/$ VBoxManage --version
6.1.34_Ubuntur150636
```

```sh
gips@UD22:/$ vagrant --version
Vagrant 2.3.0
```

```sh
gips@UD22:/$ ansible --version
ansible 2.10.8
  config file = None
  configured module search path = ['/home/gips/.ansible/plugins/modules', '/usr/share/ansible/plugins/modules']
  ansible python module location = /usr/lib/python3/dist-packages/ansible
  executable location = /usr/bin/ansible
  python version = 3.10.4 (main, Jun 29 2022, 12:14:53) [GCC 11.2.0]
  ```


## Задача 4 (*)

Воспроизвести практическую часть лекции самостоятельно.

- Создать виртуальную машину.
- Зайти внутрь ВМ, убедиться, что Docker установлен с помощью команды
```
docker ps
```


Ответ:  

```sh
gips@UD22:~/src/vagrant$ vagrant global-status
id       name             provider   state   directory                           
---------------------------------------------------------------------------------
434ba8f  server1.netology virtualbox running /home/gips/src/vagrant              
 
The above shows information about all known Vagrant environments
on this machine. This data is cached and may not be completely
up-to-date (use "vagrant global-status --prune" to prune invalid
entries). To interact with any of the machines, you can go to that
directory and run Vagrant, or you can use the ID directly with
Vagrant commands from any directory. For example:
"vagrant destroy 1a2b3c4d"
```
```sh
vagrant@server1:~$ docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
```
```sh
vagrant@server1:~$ docker run hello-world
Unable to find image 'hello-world:latest' locally
latest: Pulling from library/hello-world
2db29710123e: Pull complete 
Digest: sha256:7d246653d0511db2a6b2e0436cfd0e52ac8c066000264b3ce63331ac66dca625
Status: Downloaded newer image for hello-world:latest

Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker ID:
 https://hub.docker.com/

For more examples and ideas, visit:
 https://docs.docker.com/get-started/
```

