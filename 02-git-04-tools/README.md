# Домашнее задание к занятию «2.4. Инструменты Git»

Для выполнения заданий в этом разделе склонировали репозиторий с исходным кодом 
терраформа https://github.com/hashicorp/terraform 

***
## 1. Найдите полный хеш и комментарий коммита, хеш которого начинается на `aefea`.

### Решение:

git show aefea

### Ответ:

aefead2207ef7e2aa5dc81a34aedf0cad4c32545

Update CHANGELOG.md

***

## 2. Какому тегу соответствует коммит `85024d3`?

### Решение:

git show 85024d3

### Ответ:

v0.12.23

commit 85024d3100126de36331c6982bfaac02cdab9e76 (tag: v0.12.23)

***

## 3. Сколько родителей у коммита `b8d720`? Напишите их хеши.

### Решение:

Смотрим информацию по коммиту

git show b8d720

commit b8d720f8340221f2146e4e4870bf2ee0bc48f2d5
Merge: 56cd7859e 9ea88f22f

Так как это мерж коммит у него 2 родителя 
56cd7859e и 9ea88f22f

Смотрим первого рдителя

git show b8d720^

commit 56cd7859e05c36c06b56d013b55a252d0bb7e158

Смотрим второго родителя

git show b8d720^2
commit 9ea88f22fc6269854151c571162c5bcf958bee2b

### Ответ:

У коммита `b8d720` два родителя  
Первый родитель  
56cd7859e05c36c06b56d013b55a252d0bb7e158  
Второй родитель  
9ea88f22fc6269854151c571162c5bcf958bee2b

***

## 4. Перечислите хеши и комментарии всех коммитов которые были сделаны между тегами  v0.12.23 и v0.12.24.

### Решение:  
git log  v0.12.23..v0.12.24  --oneline  
### Ответ:

33ff1c03b (tag: v0.12.24) v0.12.24  
b14b74c49 [Website] vmc provider links  
3f235065b Update CHANGELOG.md  
6ae64e247 registry: Fix panic when server is unreachable  
5c619ca1b website: Remove links to the getting started guide's old location  
06275647e Update CHANGELOG.md  
d5f9411f5 command: Fix bug when using terraform login on Windows  
4b6d06cc5 Update CHANGELOG.md  
dd01a3507 Update CHANGELOG.md  
225466bc3 Cleanup after v0.12.23 release  

***

## 5. Найдите коммит в котором была создана функция `func providerSource`, ее определение в коде выглядит так `func providerSource(...)` (вместо троеточего перечислены аргументы).

### Решение:

git log -S'func providerSource' --oneline

### Ответ:

5af1e6234 main: Honor explicit provider_installation CLI config when present  
8c928e835 main: Consult local directories as potential mirrors of providers

***

## 6. Найдите все коммиты в которых была изменена функция `globalPluginDirs`.

### Решение:

Ищем фунцию

git grep 'func globalPluginDirs'  
plugins.go:func globalPluginDirs() []string {  

Запросим лог где этот файл и эта функция были изменены

git log -L :'func globalPluginDirs':plugins.go --oneline

### Ответ:

78b122055 Remove config.go and update things using its aliases  
52dbf9483 keep .terraform.d/plugins for discovery  
41ab0aef7 Add missing OS_ARCH dir to global plugin paths  
66ebff90c move some more plugin search path logic to command  
8364383c3 Push plugin discovery down into command package  

***

## 7. Кто автор функции `synchronizedWriters`? 

Решение:

git log -S'func synchronizedWriters' --pretty=format:'%h - %an %ae'

Ответ:

git show 5ac311e2a - Martin Atkins mart@degeneration.co.uk

