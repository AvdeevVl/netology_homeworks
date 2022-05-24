# Домашнее задание к занятию "3.4. Операционные системы, лекция 2"

1. На лекции мы познакомились с [node_exporter](https://github.com/prometheus/node_exporter/releases). В демонстрации его исполняемый файл запускался в background. Этого достаточно для демо, но не для настоящей production-системы, где процессы должны находиться под внешним управлением. Используя знания из лекции по systemd, создайте самостоятельно простой [unit-файл](https://www.freedesktop.org/software/systemd/man/systemd.service.html) для node_exporter:

    * поместите его в автозагрузку,
    * предусмотрите возможность добавления опций к запускаемому процессу через внешний файл (посмотрите, например, на `systemctl cat cron`),
    * удостоверьтесь, что с помощью systemctl процесс корректно стартует, завершается, а после перезагрузки автоматически поднимается.

Ответ: 

`wget https://github.com/prometheus/node_exporter/releases/download/v1.3.1/node_exporter-1.3.1.linux-amd64.tar.gz`  
  Распакуем скачанный архив:  
`tar xvfz node_exporter-1.3.1.linux-amd64.tar.gz`  
  перейдем в каталог с распакованными файлами:  
`cd node_exporter-1.3.1.linux-amd64`  
Копируем исполняемый файл в bin:  
`cp node_exporter /usr/local/bin/`   
Создаем пользователя nodeusr:  
`useradd --no-create-home --shell /bin/false nodeusr`  
Задаем владельца для исполняемого файла:  
`chown -R nodeusr:nodeusr /usr/local/bin/node_exporter`  

Автозапуск  
Создаем файл `node_exporter.service в systemd`  
`nano /etc/systemd/system/node_exporter.service ` 

```
[Unit]
Description=Node Exporter Service
After=network.target

[Service]
User=nodeusr
Group=nodeusr
Type=simple
ExecStart=/usr/local/bin/node_exporter
ExecReload=/bin/kill -HUP $MAINPID
Restart=on-failure

[Install]
WantedBy=multi-user.target
```
Cоздаем парамерты возможность добавления опций к запускаемому процессу через внешний файл  

`echo "EXTRA_OPTS=\"--log.level=info\"" > /opt/node_exporter.env`  

Перечитываем конфигурацию systemd:  
systemctl daemon-reload  
Разрешаем автозапуск:  
systemctl enable node_exporter  
Запускаем службу:  
systemctl start node_exporter  

systemctl status node_exporter  

![2022-05-24_22-21-06](https://user-images.githubusercontent.com/93952387/170128738-b0928180-c10a-4fa9-af60-0f26fc18c153.png)
  
  
![Снимок экрана от 2022-05-24 22-02-05](https://user-images.githubusercontent.com/93952387/170128632-a4f0c17c-12ff-4304-bcba-be29f9027023.png)  


---

2. Ознакомьтесь с опциями node_exporter и выводом `/metrics` по-умолчанию. Приведите несколько опций, которые вы бы выбрали для базового мониторинга хоста по CPU, памяти, диску и сети.

Ответ:  
CPU

      `node_cpu_seconds_total{cpu="0",mode="idle"} 5026.45
      node_cpu_seconds_total{cpu="0",mode="system"} 8.13
      node_cpu_seconds_total{cpu="0",mode="user"} 4.84
      node_cpu_seconds_total{cpu="1",mode="idle"} 5025.01
      node_cpu_seconds_total{cpu="1",mode="system"} 13.46
      node_cpu_seconds_total{cpu="1",mode="user"} 1.61
      process_cpu_seconds_total 0.58`
      
Memory

      `node_memory_MemTotal_bytes 1.028694016e+09
      node_memory_MemFree_bytes 5.68066048e+08
      node_memory_MemAvailable_bytes 7.6163072e+08
      node_memory_Buffers_bytes 2.3498752e+07
      node_memory_Cached_bytes 2.90205696e+08`
      
Disk

      `node_disk_io_time_seconds_total{device="sda"} 12.012
      node_disk_read_bytes_total{device="sda"} 3.16904448e+08
      node_disk_read_time_seconds_total{device="sda"} 7.303
      node_disk_written_bytes_total{device="sda"} 7.4236928e+07
      node_disk_write_time_seconds_total{device="sda"} 7.746`
      
Network

      `node_network_receive_bytes_total{device="eth0"} 545686
      node_network_receive_errs_total{device="eth0"} 0
      node_network_transmit_bytes_total{device="eth0"} 434942
      node_network_transmit_errs_total{device="eth0"} 0`

---

3. Установите в свою виртуальную машину [Netdata](https://github.com/netdata/netdata). Воспользуйтесь [готовыми пакетами](https://packagecloud.io/netdata/netdata/install) для установки (`sudo apt install -y netdata`). После успешной установки:
    * в конфигурационном файле `/etc/netdata/netdata.conf` в секции [web] замените значение с localhost на `bind to = 0.0.0.0`,
    * добавьте в Vagrantfile проброс порта Netdata на свой локальный компьютер и сделайте `vagrant reload`:

    ```bash
    config.vm.network "forwarded_port", guest: 19999, host: 19999
    ```

    После успешной перезагрузки в браузере *на своем ПК* (не в виртуальной машине) вы должны суметь зайти на `localhost:19999`. Ознакомьтесь с метриками, которые по умолчанию собираются Netdata и с комментариями, которые даны к этим метрикам.
   
Ответ:  

![Снимок экрана от 2022-05-24 22-25-16](https://user-images.githubusercontent.com/93952387/170128964-ee16bed7-8aed-4972-9ea2-92b3a9d14547.png)


---

4. Можно ли по выводу `dmesg` понять, осознает ли ОС, что загружена не на настоящем оборудовании, а на системе виртуализации?

Ответ:  
`sudo dmesg | grep -in virtual`  

![2022-05-24_22-38-04](https://user-images.githubusercontent.com/93952387/170129102-ded42456-4894-4cfd-8ee3-f9b4480d441e.png)  


`Booting paravirtualized kernel on KVM`  - программное решение, обеспечивающее виртуализацию в среде Linux.  
`systemd[1]: Detected virtualization oracle.` - определение системы виртуализации.

---

5. Как настроен sysctl `fs.nr_open` на системе по-умолчанию? Узнайте, что означает этот параметр. Какой другой существующий лимит не позволит достичь такого числа (`ulimit --help`)?

Ответ:   
`fs.nr_open` -  лимит на количество открытых дескрипторов   
 
 ![2022-05-24_22-43-04](https://user-images.githubusercontent.com/93952387/170129186-d367e274-152a-403e-9d79-d4f6a228cada.png)  
 
 
`ulimit -Hn` - жесткое ограничение после установки превосходить нельзя;  
`ulimit -Sn` - мягкое ограничение можно превосходить вплоть до значения соответствующего жесткого ограничения.
   
---

6. Запустите любой долгоживущий процесс (не `ls`, который отработает мгновенно, а, например, `sleep 1h`) в отдельном неймспейсе процессов; покажите, что ваш процесс работает под PID 1 через `nsenter`. Для простоты работайте в данном задании под root (`sudo -i`). Под обычным пользователем требуются дополнительные опции (`--map-root-user`) и т.д.

Ответ:  
pts/1  


![2022-05-24_23-25-23](https://user-images.githubusercontent.com/93952387/170129585-1d00e59a-7136-4d11-b23b-d29548f2e1d9.png)  

pts/0  
`ps aux`

![Снимок экрана от 2022-05-24 23-26-18](https://user-images.githubusercontent.com/93952387/170129701-fecc04fc-e094-45c4-a1d0-94247875da49.png)   
![Снимок экрана от 2022-05-24 23-27-47](https://user-images.githubusercontent.com/93952387/170129905-a65c64ce-418a-40e7-aafb-018b15ebfab4.png)


---

7. Найдите информацию о том, что такое `:(){ :|:& };:`. Запустите эту команду в своей виртуальной машине Vagrant с Ubuntu 20.04 (**это важно, поведение в других ОС не проверялось**). Некоторое время все будет "плохо", после чего (минуты) – ОС должна стабилизироваться. Вызов `dmesg` расскажет, какой механизм помог автоматической стабилизации. Как настроен этот механизм по-умолчанию, и как изменить число процессов, которое можно создать в сессии?

Ответ:  

![Снимок экрана от 2022-05-24 22-51-29](https://user-images.githubusercontent.com/93952387/170130130-dfed310c-3ff4-475e-9890-e3a8a0bee143.png)  




функция : рекурсивно вызывает сама себя до тех пор пока не забьёт все ресурсы системы.
  
`[Tue May 24 19:53:20 2022] cgroup: fork rejected by pids controller in /user.slice/user-1000.slice/session-3.scope`  

 Cработал механизм сgroups - это способ ограничить ресурсы внутри конкретной cgroup(контрольной группы процессов).
Параметры по умолчанию можно просмотреть командой `ulimit -a`  


![2022-05-24_23-05-32](https://user-images.githubusercontent.com/93952387/170130203-1f44cfae-d497-407c-92fb-b1a5e18ecc88.png)

 ---


