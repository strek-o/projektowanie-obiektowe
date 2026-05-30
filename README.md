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

## Zadanie 3 `Wzorce kreacyjne`

Spring Boot (Kotlin).

Proszę stworzyć prosty serwis do autoryzacji, który zasymuluje autoryzację użytkownika za pomocą przesłanej nazwy użytkownika oraz hasła. Serwis powinien zostać wstrzyknięty do kontrolera (4.5).
Aplikacja ma oczywiście zawierać jeden kontroler i powinna zostać napisana w języku Kotlin. Oparta powinna zostać na frameworku Spring Boot. Serwis do autoryzacji powinien być singletonem.

- [x] **3.0** Należy stworzyć jeden kontroler wraz z danymi wyświetlanymi z listy na endpoint’cie w formacie JSON - Kotlin + Spring Boot [[commit]](https://github.com/strek-o/projektowanie-obiektowe/tree/31bd054254eb90a2458dc3ffd4bf4c27a17bb470)
- [ ] **3.5** Należy stworzyć klasę do autoryzacji (mock) jako Singleton w formie eager
- [ ] **4.0** Należy obsłużyć dane autoryzacji przekazywane przez użytkownika
- [ ] **4.5** Należy wstrzyknąć singleton do głównej klasy via @Autowired lub konstruktor (constructor injection)
- [ ] **5.0** Obok wersji Eager do wyboru powinna być wersja Singletona w wersji lazy

## Zadanie 4 `Wzorce strukturalne`

Echo (Go).

Należy stworzyć aplikację w Go na frameworku echo. Aplikacja ma mieć jeden endpoint, minimum jedną funkcję proxy, która pobiera dane np. o pogodzie, giełdzie, etc. (do wyboru) z zewnętrznego API. Zapytania do endpointu można wysyłać jako GET lub POST.

- [x] **3.0** Należy stworzyć aplikację we frameworku echo w j. Go, która będzie miała kontroler Pogody, która pozwala na pobieranie danych o pogodzie (lub akcjach giełdowych) [[commit]](https://github.com/strek-o/projektowanie-obiektowe/tree/aad6fc82ee3a137cf4714c0c69e98126b8611f19)
- [x] **3.5** Należy stworzyć model Pogoda (lub Giełda) wykorzystując gorm, a dane załadować z listy przy uruchomieniu [[commit]](https://github.com/strek-o/projektowanie-obiektowe/tree/da212ca92b00f564569a7efa01cde2f6ac7b0e9c)
- [ ] **4.0** Należy stworzyć klasę proxy, która pobierze dane z serwisu zewnętrznego podczas zapytania do naszego kontrolera
- [ ] **4.5** Należy zapisać pobrane dane z zewnątrz do bazy danych
- [ ] **5.0** Należy rozszerzyć endpoint na więcej niż jedną lokalizację (Pogoda), lub akcje (Giełda) zwracając JSONa

## Zadanie 5 `Wzorce behawioralne`

React (JavaScript/Typescript).

- [x] **3.0** W ramach projektu należy stworzyć komponenty Produkty oraz Płatności; komponent Produkty powinien pobierać listę produktów z aplikacji serwerowej, natomiast komponent Płatności powinien wysyłać dane płatności do aplikacji serwerowej. [[commit]](https://github.com/strek-o/projektowanie-obiektowe/tree/5a46aadc02c1a2273699b7b62fe783ece88e8c26)
- [ ] **3.5** Należy dodać komponent Koszyk wraz z osobnym widokiem; aplikacja powinna umożliwiać przechodzenie pomiędzy widokami przy użyciu routingu.
- [ ] **4.0** Dane pomiędzy komponentami, takimi jak Produkty, Koszyk i Płatności, powinny być przekazywane z wykorzystaniem React hooks, np. useState, useEffect lub useContext.
- [ ] **4.5** Należy przygotować konfigurację umożliwiającą uruchomienie aplikacji klienckiej oraz serwerowej w kontenerach Docker za pomocą docker-compose.
- [ ] **5.0** Należy wykorzystać bibliotekę axios do komunikacji z serwerem oraz skonfigurować obsługę CORS, aby frontend mógł poprawnie komunikować się z backendem.

## Zadanie 6 `Zapaszki`

Należy sprawdzić kod projektów JS 3.0, 3.5, 4.0, kotlin, go, js - 4.5, 5.0.

- [x] **3.0** Należy skonfigurować husky + lint-staged uruchamianie lintowania przed commitem [[commit]](https://github.com/strek-o/projektowanie-obiektowe/tree/aacf7ddd3260901ca400425f1b9c1125a12ee463)
- [ ] **3.5** Należy wyeliminować wszystkie bugi w kodzie w Sonarze (kod aplikacji klienckiej)
- [ ] **4.0** Przeskanować oraz naprawić dowolny projekt open source narzędziem [CodeQL](https://codeql.github.com/)
- [ ] **4.5** Należy usunąć problemy typu Code Smell w kodzie w Sonarze (kotlin, go, js). Należy dodać badge z Sonara
- [ ] **5.0** Skonfigurować Github Actions z linterem oraz CodeQL

## Zadanie 7 `Vapor`

Proszę napisać prostą aplikację w Vaporze, wykorzystując Leaf jako silnik szablonów or Fluent jako ORM. Proszę stworzyć trzy modele oraz CRUD dla każdego z nich. Należy stworzyć model z minimum jedną relacją. CRUD powinien mieć odzwierciedlenie w szablonach.

- [x] **3.0** Należy stworzyć kontroler wraz z modele Produktów zgodny z CRUD w ORM Fluent [[commit]](https://github.com/strek-o/projektowanie-obiektowe/tree/a66f7d65c7c40b00d50dc22dc26df6708acac0f8)
- [ ] **3.5** Należy stworzyć szablony w Leaf
- [ ] **4.0** Należy stworzyć drugi model oraz kontroler Kategorii wraz z relacją
- [ ] **4.5** Należy wykorzystać Redis do przechowywania danych
- [ ] **5.0** Wrzucić aplikację na [heroku](https://www.heroku.com/github-students)

## Zadanie 8 `Selenium`

Wykorzystaj do tego zadania Selenium/Webdrivera (Kotlin, Python, Java, JS, Go, Scala).

- [ ] **3.0** Przetestuj formularz rejestracji użytkownika pod kątem walidacji pól obowiązkowych oraz zachowania aplikacji po wprowadzeniu niepoprawnego formatu adresu e-mail.
- [ ] **3.5** Przeprowadź testy bezpieczeństwa typu Cross-Site Scripting (XSS), próbując wstrzyknąć złośliwy kod JavaScript w aplikacji z Reactem.
- [ ] **4.0** Przetestuj działanie koszyka zakupowego przy jednoczesnym otwarciu aplikacji w kilku osobnych kartach tej samej przeglądarki, sprawdzając spójność stanów zamówienia (aplikacja z zadania z React'em).
- [ ] **4.5** Do zadania z React'a należy dodać formularz logowania. Następnie przeprowadź testy podatności na ataki typu Cross-Site Request Forgery (CSRF), próbując wymusić nieautoryzowaną zmianę ustawień konta
      spreparowanym linkiem, podczas gdy użytkownik posiada aktywną sesję w innej karcie.
- [ ] **5.0** Stwórz scenariusz End-to-End w Playwright (minimum 50 asercji).

## Zadanie 9 `Chmura`

- [ ] **3.0** Należy stworzyć odpowiednie instancje po stronie chmury na dockerze
- [ ] **3.5** Stworzyć odpowiedni pipeline w Github Actions do budowania aplikacji
- [ ] **4.0** Dodać notyfikację mailową o zbudowaniu aplikacji
- [ ] **4.5** Dodać krok z deploymentem aplikacji serwerowej oraz klienckiej na chmurę
- [ ] **5.0** Dodać uruchomienie regresyjnych testów automatycznych (funkcjonalnych) jako krok w Actions
