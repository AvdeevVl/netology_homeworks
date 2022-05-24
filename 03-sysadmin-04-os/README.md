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

---

4. Можно ли по выводу `dmesg` понять, осознает ли ОС, что загружена не на настоящем оборудовании, а на системе виртуализации?

Ответ:  
`sudo dmesg | grep -in virtual`



`Booting paravirtualized kernel on KVM`  - программное решение, обеспечивающее виртуализацию в среде Linux.  
`systemd[1]: Detected virtualization oracle.` - определение системы виртуализации.

---

5. Как настроен sysctl `fs.nr_open` на системе по-умолчанию? Узнайте, что означает этот параметр. Какой другой существующий лимит не позволит достичь такого числа (`ulimit --help`)?

Ответ: 
`fs.nr_open` -  лимит на количество открытых дескрипторов  

 
`ulimit -Hn` - жесткое ограничение после установки превосходить нельзя;
`ulimit -Sn` - мягкое ограничение можно превосходить вплоть до значения соответствующего жесткого ограничения.
   
---

6. Запустите любой долгоживущий процесс (не `ls`, который отработает мгновенно, а, например, `sleep 1h`) в отдельном неймспейсе процессов; покажите, что ваш процесс работает под PID 1 через `nsenter`. Для простоты работайте в данном задании под root (`sudo -i`). Под обычным пользователем требуются дополнительные опции (`--map-root-user`) и т.д.

Ответ: 




---

7. Найдите информацию о том, что такое `:(){ :|:& };:`. Запустите эту команду в своей виртуальной машине Vagrant с Ubuntu 20.04 (**это важно, поведение в других ОС не проверялось**). Некоторое время все будет "плохо", после чего (минуты) – ОС должна стабилизироваться. Вызов `dmesg` расскажет, какой механизм помог автоматической стабилизации. Как настроен этот механизм по-умолчанию, и как изменить число процессов, которое можно создать в сессии?

Ответ:  

функция : рекурсивно вызывает сама себя до тех пор пока не забьёт все ресурсы системы.
  
`[Tue May 24 19:53:20 2022] cgroup: fork rejected by pids controller in /user.slice/user-1000.slice/session-3.scope`  

 Cработал механизм сgroups - это способ ограничить ресурсы внутри конкретной cgroup(контрольной группы процессов).
Параметры по умолчанию можно просмотреть командой `ulimit -a`

 ---


