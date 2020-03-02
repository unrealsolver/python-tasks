# Basics
## 01. Leap Year
## 02. Sequence Sum
## 03. Evergreen Tree
## 04. Luhn Algorithm
## 05. ASCII Font

# Text Files
## 06. CSV to JSON
Классическая задача - прочесть CSV, записать JSON. Обратный процесс мало чем отличается.
[file.csv](https://raw.githubusercontent.com/unrealsolver/python-tasks/master/sample.csv)
Рекомендации:
* стандартные модули `csv` и `json`
* Помимо ридера по-умолчанию, можно использоать [csv.DictReader](https://docs.python.org/3/library/csv.html#csv.DictReader)

## 07. Go Type to YAML
Иногда необходимо сохранить информацию о какой-то модели в более удобно формате. 
Возьмите файл [sample.go](https://raw.githubusercontent.com/unrealsolver/python-tasks/master/sample.go) (взято [отсюда](https://github.com/evergreen-ci/evergreen/blob/master/model/task/task.go)) и сохраните его в YAML, по-возможности сохранив максимум информации
Рекомендации:
* Библиотеку `pyyaml` придется поставить отдельно (`pip3 install pyyaml`)
* Регулярки точно нужны

# Binary Files
## BMP to ASCII art
Работать с бинарными файлами в пайтон не сильно сложнее, чем с текстовыми.
Для ознакомления отлично подойтет формат BMP, который нагляден и прост.

Программа должна открывать BMP файл (можете использовать модуль `struct`) и выводить
изображение в терминал в виде ASCII графики.

Рекомендации:
* Использование модуля `struct`
* Создание класса для описания пикселя
* очевидно, вывод будет черно-белым, но можно так же заморочиться с цветным!
(используя управляющие мета-символы)

Пример:

![example](https://raw.githubusercontent.com/unrealsolver/python-tasks/master/bmpascii.bmp)
