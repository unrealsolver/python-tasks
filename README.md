# Basics
## Leap Year
## Sequence Sum
## Evergreen Tree
## Luhn Algorithm
## ASCII Font

# Text Files
## CSV to JSON
Классическая задача - прочесть CSV, записать JSON. Обратный процесс мало чем отличается.
[file.csv](https://raw.githubusercontent.com/unrealsolver/python-tasks/master/sample.csv)
Рекомендации:
* стандартные модули `csv` и `json`
* Помимо ридера по-умолчанию, можно использоать [csv.DictReader](https://docs.python.org/3/library/csv.html#csv.DictReader)

## 07. Go Type to YAML
Иногда необходимо сохранить информацию о какой-то модели в более удобно формате. [sample.go](https://raw.githubusercontent.com/unrealsolver/python-tasks/master/sample.go)
Возьмите файл (sample.go) (взято отсюда https://github.com/evergreen-ci/evergreen/blob/master/model/task/task.go) и сохраните его в YAML, по-возможности сохранив максимум информации
Рекомендации:
* Библиотеку `pyyaml` придется поставить отдельно (`pip3 install pyyaml`)
* Регулярки точно нужны
