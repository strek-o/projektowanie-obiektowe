# Projektowanie obiektowe

[Docker Hub](https://hub.docker.com/r/streko/projektowanie-obiektowe/tags)  
[WFAIS.IF-XO301.0]

## Zadanie 1 `Paradygmaty`

Sortowanie bąbelkowe.

Proszę napisać program w Pascalu, który zawiera dwie procedury, jedna generuje listę 50 losowych liczb od 0 do 100.  
Druga procedura sortuje liczbę za pomocą sortowania bąbelkowego.

- [x] **3.0** Procedura do generowania 50 losowych liczb od 0 do 100 [[commit]](https://github.com/strek-o/projektowanie-obiektowe/tree/dc12ea6e01e936569613cd0d2228a7b95850ea92)
- [x] **3.5** Procedura do sortowania liczb [[commit]](https://github.com/strek-o/projektowanie-obiektowe/tree/d08f9f3ff3c73dc419557b5da1bbe417b6bc7ae7)
- [x] **4.0** Dodanie parametrów do procedury losującej określającymi zakres losowania: od, do, ile [[commit]](https://github.com/strek-o/projektowanie-obiektowe/tree/bcfa0c3b1b4bbf0df5f32039617fd53a01e7b102)
- [ ] **4.5** 5 testów jednostkowych testujące procedury
- [ ] **5.0** Skrypt w bashu do uruchamiania aplikacji w Pascalu via docker

## Zadanie 2 `Wzorce architektury`

Symfony (PHP).

Należy stworzyć aplikację webową na bazie frameworka Symfony na obrazie [kprzystalski/projobj-php:latest](https://hub.docker.com/r/kprzystalski/projobj-php). Baza danych dowolna, sugeruję SQLite.

- [x] **3.0** Należy stworzyć jeden model z kontrolerem z produktami, zgodnie z CRUD (JSON) [[commit]](https://github.com/strek-o/projektowanie-obiektowe/tree/c80bc210cec53f80916bc97c2d818d135545479e)
- [ ] **3.5** Należy stworzyć skrypty do testów endpointów via curl (JSON)
- [ ] **4.0** Należy stworzyć dwa dodatkowe kontrolery wraz z modelami (JSON)
- [ ] **4.5** Należy stworzyć widoki do wszystkich kontrolerów
- [ ] **5.0** Stworzenie panelu administracyjnego
