# Домашнее задание к занятию "3.1. Работа в терминале, лекция 1"

1. Установите средство виртуализации [Oracle VirtualBox](https://www.virtualbox.org/).
   
   Решение:  
sudo apt install virtualbox

1. Установите средство автоматизации [Hashicorp Vagrant](https://www.vagrantup.com/).
   
    Решение:  
   curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo apt-key add -  
sudo apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main"  
sudo apt-get update && sudo apt-get install vagrant

1. В вашем основном окружении подготовьте удобный для дальнейшей работы терминал. Можно предложить:
    
    Решение:  
    sudo apt install fish  
    fish

1. С помощью базового файла конфигурации запустите Ubuntu 20.04 в VirtualBox посредством Vagrant:

    Решение:

        * Создали директорию

        * Инициализировали ранее созданный каталог как среду Vagrant командой vagrant init при этом создаеться стандартный файл Vagrantfile  

        * Заменинили содержимое Vagrantfile по умолчанию следующим:

		```bash
		Vagrant.configure("2") do |config|
			config.vm.box = "bento/ubuntu-20.04"
		end
		```

      	* Выполнели в этой директории `vagrant up`. Скачался необходимый образ и запустилась виртуальная машина.


1. Ознакомьтесь с графическим интерфейсом VirtualBox, посмотрите как выглядит виртуальная машина, которую создал для вас Vagrant, какие аппаратные ресурсы ей выделены. Какие ресурсы выделены по-умолчанию?

    Выполнено  
        Ресурсы выделенные по умолчанию для созданной машины:  
      *  RAM:1024mb  
      *  CPU:2 cpu  
      *  HDD:64gb
  
1. Ознакомьтесь с возможностями конфигурации VirtualBox через Vagrantfile: [документация](https://www.vagrantup.com/docs/providers/virtualbox/configuration.html). Как добавить оперативной памяти или ресурсов процессора виртуальной машине?

    Решение:  
        Добавление оперативной памяти и ресурсов процессора осуществляеться путем добавление в файл Vagrantfile параметров:  

            v.memory = 
            v.cpus = 


1. Команда `vagrant ssh` из директории, в которой содержится Vagrantfile, позволит вам оказаться внутри виртуальной машины без каких-либо дополнительных настроек. Попрактикуйтесь в выполнении обсуждаемых команд в терминале Ubuntu.

    Выполнено `vagrant ssh`

1. Ознакомиться с разделами `man bash`, почитать о настройках самого bash:
    * какой переменной можно задать длину журнала `history`, и на какой строчке manual это описывается?
    * что делает директива `ignoreboth` в bash?

    Ответ:  
      
    * HISTFILESIZE - максимальное число строк в файле истории для сохранения.  
     строка 698
    * HISTSIZE - число команд для сохранения  
     строка 709

    * Igrnoreboth - это сокращение от ignorespace и ignoredups
    * ignorespace - не сохранять команды начинающиеся с пробела
    * ignoredups - не сохранять команду если такая уже имеется в истории

1. В каких сценариях использования применимы скобки `{}` и на какой строчке `man bash` это описано?

    Решение:  
    ```/\{```  
    Строка 203

        {} - зарезервированные слова.

        {list;}  
        Команда интерпретируется как список команд, разделенных точкой с запятой.

        ${}  
        Подстановка параметра.

        {}  
        Блоки кода.



1. С учётом ответа на предыдущий вопрос, как создать однократным вызовом `touch` 100000 файлов?  
   Решение:  
     touch {1..100000}  
   
   Получится ли аналогичным образом создать 300000?  
    Нет  
    Если нет, то почему?  
    Слишком дилинный список аргументов


1. В man bash поищите по `/\[\[`. Что делает конструкция `[[ -d /tmp ]]`

    Ответ:  
    Конструкция [[ -d /tmp ]] возвращает 1 (True) если директория /tmp существует и 0 (False) если отсутствует

1.  Основываясь на знаниях о просмотре текущих (например, PATH) и установке новых переменных; командах, которые мы рассматривали, добейтесь в выводе type -a bash в виртуальной машине наличия первым пунктом в списке:

	```bash
	bash is /tmp/new_path_directory/bash
	bash is /usr/local/bin/bash
	bash is /bin/bash
	```

	(прочие строки могут отличаться содержимым и порядком)
    В качестве ответа приведите команды, которые позволили вам добиться указанного вывода или соответствующие скриншоты.

    Решение:
    ```
    mkdir /tmp/new_path_directory && cp /usr/bin/bash /tmp/new_path_directory/ && export PATH="/tmp/new_path_directory:$PATH"
    ```
    

1. Чем отличается планирование команд с помощью `batch` и `at`?

    at - выполнение команд в определенное время.

    batch - выполняет команды когда позволяют уровни нагрузки системы; другими словами, когда средняя нагрузка падает ниже 1,5.


1. Завершите работу виртуальной машины чтобы не расходовать ресурсы компьютера и/или батарею ноутбука.

    Выполнено:  
    vagrant halt
 
 ---