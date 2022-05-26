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
![2022-05-25_10-56-19](https://user-images.githubusercontent.com/93952387/170354911-4732d58c-1cf4-450a-a551-fc61d0209645.png)


---

4. Используя `fdisk`, разбейте первый диск на 2 раздела: 2 Гб, оставшееся пространство.

Ответ:  

`sudo fdisk /dev/sdb`  

![2022-05-25_11-25-08](https://user-images.githubusercontent.com/93952387/170355464-d41ef46e-d989-4615-a568-ca670ad9c1c4.png)


---

5. Используя `sfdisk`, перенесите данную таблицу разделов на второй диск.

Ответ:  
`sudo sfdisk -d /dev/sdb | sudo sfdisk /dev/sdc`  

![2022-05-25_11-31-11](https://user-images.githubusercontent.com/93952387/170449005-2be0b4fb-f0b3-4f45-8189-fcdcfe9e1e68.png)

---

6. Соберите `mdadm` RAID1 на паре разделов 2 Гб.

Ответ:  
`sudo mdadm --create --verbose /dev/md1 --level=1 --raid-devices=2 /dev/sdb1 /dev/sdc1`  

![2022-05-25_21-04-50](https://user-images.githubusercontent.com/93952387/170449304-4ddc59d7-da77-4865-b9de-72286e8d1185.png)


---

7. Соберите `mdadm` RAID0 на второй паре маленьких разделов.

Ответ:  
`sudo mdadm --create --verbose /dev/md0 --level=0 --raid-devices=2 /dev/sdb2 /dev/sdc2`  

![2022-05-25_21-05-25](https://user-images.githubusercontent.com/93952387/170449713-44171276-410c-4b6e-b0e8-f32e9399386d.png)

---

8. Создайте 2 независимых PV на получившихся md-устройствах.

Ответ:  
`sudo pvcreate /dev/md1 /dev/md0`  

![2022-05-25_21-16-25](https://user-images.githubusercontent.com/93952387/170450085-a1606275-772a-492c-8979-f8915ee95d8b.png)  

![2022-05-25_21-19-33](https://user-images.githubusercontent.com/93952387/170450141-f70af97f-322a-4108-a622-c9e91f53ceeb.png)



---

9. Создайте общую volume-group на этих двух PV.

Ответ:  
`sudo vgcreate vg1 /dev/md1 /dev/md0`  

![2022-05-25_21-22-08](https://user-images.githubusercontent.com/93952387/170450328-7b5c2683-36fa-4821-8659-870662c958b5.png)


---

10. Создайте LV размером 100 Мб, указав его расположение на PV с RAID0.

Ответ:  
`sudo lvcreate -L 100M vg1 /dev/md0`  

![2022-05-25_21-25-57](https://user-images.githubusercontent.com/93952387/170450601-3350cfe2-208c-477a-a698-90e4a34c02bb.png)


---

11. Создайте `mkfs.ext4` ФС на получившемся LV.

Ответ:  
`sudo mkfs.ext4 /dev/vg1/lvol0`  

![2022-05-25_21-30-32](https://user-images.githubusercontent.com/93952387/170450656-bad8c0a3-68d3-464d-bf91-e650fd855b40.png)


---

12. Смонтируйте этот раздел в любую директорию, например, `/tmp/new`.

Ответ:  

`sudo mkdir /tmp/new`  
`sudo mount /dev/vg1/lvol0/ /tmp/new`  
`sudo mount | grep "lvol0"`  

![2022-05-25_21-36-07](https://user-images.githubusercontent.com/93952387/170450756-30a4461a-b6af-4883-bdd8-3ee4706ada78.png)

---

13. Поместите туда тестовый файл, например `wget https://mirror.yandex.ru/ubuntu/ls-lR.gz -O /tmp/new/test.gz`.

Ответ:  
`sudo wget https://mirror.yandex.ru/ubuntu/ls-lR.gz -O /tmp/new/test.gz`  

![2022-05-25_21-39-30](https://user-images.githubusercontent.com/93952387/170450842-fe092028-cafe-4c80-8645-7fe54d3ffa6d.png)


---

14. Прикрепите вывод `lsblk`.

Ответ:  

![2022-05-25_21-42-02](https://user-images.githubusercontent.com/93952387/170450904-2d4e787e-eba5-4de7-b6be-f295a83ea8de.png)


---

15. Протестируйте целостность файла:

    ```bash
    root@vagrant:~# gzip -t /tmp/new/test.gz
    root@vagrant:~# echo $?
    0
    ```

Ответ:  
Вывод команды echo $? равен 0. Значит команда gzip -t отработала успешно и файл доступен.  

![2022-05-25_21-43-25](https://user-images.githubusercontent.com/93952387/170450945-e45feea2-d1ae-4660-b545-8a261f05f75e.png)


---

16. Используя pvmove, переместите содержимое PV с RAID0 на RAID1.

Ответ:  
`sudo pvmove -n lvol0 /dev/md0 /dev/md1`  

![2022-05-25_21-54-37](https://user-images.githubusercontent.com/93952387/170451020-2903ae2d-f554-4a9f-83f1-78410ca7bb13.png)


---

17. Сделайте `--fail` на устройство в вашем RAID1 md.

Ответ:  
`sudo mdadm --fail /dev/md1 /dev/sdb1`  

![2022-05-25_22-14-40](https://user-images.githubusercontent.com/93952387/170451246-236952a0-c7c1-4984-87fd-be8971e40760.png)


---

18. Подтвердите выводом `dmesg`, что RAID1 работает в деградированном состоянии.

Ответ:  
`dmesg -T`  

![2022-05-25_22-13-27](https://user-images.githubusercontent.com/93952387/170451292-0f1bfe7b-c3aa-47cf-95ba-1ff105d77afa.png)

---

19. Протестируйте целостность файла, несмотря на "сбойный" диск он должен продолжать быть доступен:

    ```bash
    root@vagrant:~# gzip -t /tmp/new/test.gz
    root@vagrant:~# echo $?
    0
    ```

Ответ:  

![2022-05-25_22-17-44](https://user-images.githubusercontent.com/93952387/170451351-e3d299d4-d1a4-46c3-8e0b-e4709d1ca1ab.png)  

Вывод команды echo $? равен 0. Значит команда gzip -t отработала успешно и файл остается доступным.

---

20. Погасите тестовый хост, `vagrant destroy`.

Ответ:  

![2022-05-25_22-22-24](https://user-images.githubusercontent.com/93952387/170451408-1bcfab58-6baa-4948-b69e-4a9756ba65b1.png)


 ---
