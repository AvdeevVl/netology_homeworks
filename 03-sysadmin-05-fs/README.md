# Домашнее задание к занятию "3.5. Файловые системы"

1. Узнайте о [sparse](https://ru.wikipedia.org/wiki/%D0%A0%D0%B0%D0%B7%D1%80%D0%B5%D0%B6%D1%91%D0%BD%D0%BD%D1%8B%D0%B9_%D1%84%D0%B0%D0%B9%D0%BB) (разряженных) файлах.

Ответ:  
Разрежённый файл (англ. sparse file) — файл, в котором последовательности нулевых байтов заменены на информацию об этих последовательностях (список дыр).

---

2. Могут ли файлы, являющиеся жесткой ссылкой на один объект, иметь разные права доступа и владельца? Почему?

Ответ:  

Нет. Эти объекты файловой системы имеют одну и ту же inode

---

3. Сделайте `vagrant destroy` на имеющийся инстанс Ubuntu. Замените содержимое Vagrantfile следующим:

    ```bash
    Vagrant.configure("2") do |config|
      config.vm.box = "bento/ubuntu-20.04"
      config.vm.provider :virtualbox do |vb|
        lvm_experiments_disk0_path = "/tmp/lvm_experiments_disk0.vmdk"
        lvm_experiments_disk1_path = "/tmp/lvm_experiments_disk1.vmdk"
        vb.customize ['createmedium', '--filename', lvm_experiments_disk0_path, '--size', 2560]
        vb.customize ['createmedium', '--filename', lvm_experiments_disk1_path, '--size', 2560]
        vb.customize ['storageattach', :id, '--storagectl', 'SATA Controller', '--port', 1, '--device', 0, '--type', 'hdd', '--medium', lvm_experiments_disk0_path]
        vb.customize ['storageattach', :id, '--storagectl', 'SATA Controller', '--port', 2, '--device', 0, '--type', 'hdd', '--medium', lvm_experiments_disk1_path]
      end
    end
    ```

    Данная конфигурация создаст новую виртуальную машину с двумя дополнительными неразмеченными дисками по 2.5 Гб.

Ответ:  

---

4. Используя `fdisk`, разбейте первый диск на 2 раздела: 2 Гб, оставшееся пространство.

Ответ:  

`sudo fdisk /dev/sdb`  


---

5. Используя `sfdisk`, перенесите данную таблицу разделов на второй диск.

Ответ:  
`sudo sfdisk -d /dev/sdb | sudo sfdisk /dev/sdc`

---

6. Соберите `mdadm` RAID1 на паре разделов 2 Гб.

Ответ:  
`sudo mdadm --create --verbose /dev/md1 --level=1 --raid-devices=2 /dev/sdb1 /dev/sdc1`


---

7. Соберите `mdadm` RAID0 на второй паре маленьких разделов.

Ответ:  
`sudo mdadm --create --verbose /dev/md0 --level=0 --raid-devices=2 /dev/sdb2 /dev/sdc2`


---

8. Создайте 2 независимых PV на получившихся md-устройствах.

Ответ:  
`sudo pvcreate /dev/md1 /dev/md0`

---

9. Создайте общую volume-group на этих двух PV.

Ответ:  
`sudo vgcreate vg1 /dev/md1 /dev/md0`

---

10. Создайте LV размером 100 Мб, указав его расположение на PV с RAID0.

Ответ:  
`sudo lvcreate -L 100M vg1 /dev/md0`

---

11. Создайте `mkfs.ext4` ФС на получившемся LV.

Ответ:  
`sudo mkfs.ext4 /dev/vg1/lvol0`

---

12. Смонтируйте этот раздел в любую директорию, например, `/tmp/new`.

Ответ:  

`sudo mkdir /tmp/new`  
`sudo mount /dev/vg1/lvol0/ /tmp/new`
`sudo mount | grep "lvol0"`

---

13. Поместите туда тестовый файл, например `wget https://mirror.yandex.ru/ubuntu/ls-lR.gz -O /tmp/new/test.gz`.

Ответ:  
`sudo wget https://mirror.yandex.ru/ubuntu/ls-lR.gz -O /tmp/new/test.gz`  

---

14. Прикрепите вывод `lsblk`.

Ответ:  

---

15. Протестируйте целостность файла:

    ```bash
    root@vagrant:~# gzip -t /tmp/new/test.gz
    root@vagrant:~# echo $?
    0
    ```

Ответ:  
Вывод команды echo $? равен 0. Значит команда gzip -t отработала успешно и файл доступен.

---

16. Используя pvmove, переместите содержимое PV с RAID0 на RAID1.

Ответ:  
`sudo pvmove -n lvol0 /dev/md0 /dev/md1`

---

17. Сделайте `--fail` на устройство в вашем RAID1 md.

Ответ:  
`sudo mdadm --fail /dev/md1 /dev/sdb1`

---

18. Подтвердите выводом `dmesg`, что RAID1 работает в деградированном состоянии.

Ответ:  
`dmesg -T`

---

19. Протестируйте целостность файла, несмотря на "сбойный" диск он должен продолжать быть доступен:

    ```bash
    root@vagrant:~# gzip -t /tmp/new/test.gz
    root@vagrant:~# echo $?
    0
    ```

Ответ:  

Вывод команды echo $? равен 0. Значит команда gzip -t отработала успешно и файл остается доступным.

---

20. Погасите тестовый хост, `vagrant destroy`.

Ответ:  

 ---