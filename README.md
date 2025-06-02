# to-age-git

## en 🇬🇧

`to-age-git` - a tool that makes a project appear older by a desired number of years. It should be used for development anonymization. `to-age-git` will age all commits by N years.

Instructions:

1) First, make a backup copy of your project!

2) Make sure you have git and go installed. If not, download them from their official websites.

3) Clone the tool:

```bash
git clone https://github.com/alexKudryavtsev-web/to-age-git
```

4) Build and run the script:

```bash
go build -o to-age-git main.go
./to-age-git <название папки с git> <кол-во лет>
```

5) Wait for the process to complete and force-push all changes:

```bash
git push --force
```

Recommendation: Remember that the libraries you use have version numbers that can help approximate the real date. Use this tool wisely!



## ru 🇷🇺

`to-age-git` - инструмент, которые сделает проект старше на желаемое количество лет. Его следует использовать для анонимизации разработки. `to-age-git` состарит все коммиты на N-ое количество лет.

Инструкция:

1) Предварительно сделайте копию проекта!

2) Убедитесь, что у вас есть `git` и `go`. Если нет, скачайте с официальных сайтов.

3) Склонируйте инструмент:

```bash
git clone https://github.com/alexKudryavtsev-web/to-age-git
```

4) Соберите и запустите скрипт:

```bash
go build -o to-age-git main.go
./to-age-git <название папки с git> <кол-во лет>
```

5) Дождитесь окончания и сделайте push все изменений:

```bash
git push --force
```

Рекомендация: не забывайте, что используемые библиотеки имеют версии, по которым можно будет приблизительно узнать реальную дату. Используйте этот инструмент с умом!
