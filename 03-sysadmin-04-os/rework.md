# Доработка к домашнему заданию по занятию "3.4. Операционные системы, лекция 2"

## Задание 1
Уточнить как именно в службу будут передаваться дополнительные опции.  
Речь идёт не о переменных окружения, а об опциях (параметрах) запуска службы.  

Ответ:  

Параметры запуска можно указать в отдельном файле.  
Дополним файл `/etc/systemd/system/node_exporter.service` :  

```
[Service]
EnvironmentFile=/etc/node_exporter.conf
ExecStart=/usr/local/bin/node_exporter $EXT_OPTS
```

Теперь `systemd` запустит службу c переменной EXT_OPTS, которая определенна в файле `/etc/node_exporter.conf `

Содержимое файла `/etc/node_exporter.conf `  
```
EXT_OPTS="--log.level=info --version"
```