# Network questions

### Какие модели сетевого взаимодействия знаете?
OSI -- 

#### Модель ОСИ?
7 Прикладной
6 Представления
5 Сеансовый
4 Транспортный
3 Сетевой
2 Канальный
1 Физический

#### Что такое DHCP?
Протокол прикладного уровня модели TCP/IP, служит для назначения IP-адреса клиенту. Dynamic Host Configuration Protocol.
Принцип:
1 DHCP Discovery broadcast
2 DHCP Offer от сервера клиенту
3 DHCP Request с принятым IP, broadcast
4 HCP Ack подтверждение клиенту, может быть broadcast
Что если в сети 2 DHCP сервера?
Предпочтение тому, с кем общался ранее. Если первый раз, то первый полученный.
https://selectel.ru/blog/dhcp-protocol/

#### Что такое NAT?
Network Address Translation - это механизм в сетях TCP/IP, позволяющий изменять IP адрес в заголовке пакета, проходящего через устройство маршрутизации трафика.
Принимая пакет от локального компьютера, маршрутизатор смотрит на IP-адрес назначения. Если это локальный адрес, то пакет пересылается другому локальному компьютеру. Если нет, то пакет надо переслать наружу в интернет.

Статический NAT применяют для отображения незарегистрированного IP-адреса на зарегистрированный IP-адрес на основании один к одному. Особенно полезно, когда устройство должно быть доступным снаружи сети.

Динамический NAT отображает незарегистрированный IP-адрес на зарегистрированный адрес из группы зарегистрированных IP-адресов. Отображение может меняться в зависимости от зарегистрированного адреса, доступного в пуле адресов, во время коммуникации.

Перегруженный NAT - форма динамического NAT, который отображает несколько незарегистрированных адресов в единственный зарегистрированный IP-адрес, используя различные порты. При перегрузке каждый компьютер в частной сети транслируется в тот же самый адрес, но с различным номером порта.

#### Что такое ARP протокол?
(канальный уровень) Адресация в сети Internet представляет собой 32-битовую последовательность 0 и 1, называющихся IP-адресами. Но непосредственно связь между двумя устройствами в сети осуществляется по адресам канального уровня (MAC-адресам). Для определения соответствия между логическим адресом сетевого уровня (IP) и физическим адресом устройства (MAC) используется описанный в RFC 826 протокол ARP (Address Resolution Protocol, протокол разрешения адресов).
ARP состоит из двух частей. Первая – определяет физический адрес при посылке пакета, вторая – отвечает на запросы других станций.
Протокол имеет буферную память (ARP-таблицу), в которой хранятся пары адресов (IP-адрес, MAC-адрес) с целью уменьшения количества посылаемых запросов, следовательно, экономии трафика и ресурсов.

Как происходит определение маршрута с участием протокола ARP:
Пусть отправитель A и получатель B имеют свои адреса с указанием маски подсети.
Если адреса находятся в одной подсети, то вызывается протокол ARP и определяется физический адрес получателя, после чего IP-пакет инкапсулируется в кадр канального уровня и отправляется по указанному физическому адресу, соответствующему IP-адресу назначения.
Если нет – начинается просмотр таблицы в поисках прямого маршрута.
Если маршрут найден, то вызывается протокол ARP и определяется физический адрес соответствующего маршрутизатора, после чего пакет инкапсулируется в кадр канального уровня и отправляется по указанному физическому адресу.
В противном случае, вызывается протокол ARP и определяется физический адрес маршрутизатора по умолчанию, после чего пакет инкапсулируется в кадр канального уровня и отправляется по указанному физическому адресу.

Главным достоинством проткола ARP является его простота, что порождает в себе и главный его недостаток – абсолютную незащищенность, так как протокол не проверяет подлинность пакетов, и, в результате, можно осуществить подмену записей в ARP-таблице (материал для отдельной статьи), вклинившись между отправителем и получателем.


#### Что такое DNS?
DNS (Domain Name System) ― это система, которая позволяет браузеру найти запрошенный пользователем сайт по имени домена.

#### Чем отличается UDP от TCP?
Ключевым различием между TCP и UDP является скорость, поскольку TCP сравнительно медленнее UDP. В целом, UDP является быстрым, простым и эффективным протоколом, однако повторная передача потерянных пакетов данных возможна только в TCP. 
Еще одно заметное различие между TCP и UDP заключается в том, что первый обеспечивает упорядоченную доставку данных от пользователя к серверу (и наоборот). UDP, в свою очередь, не проверяет готовность получателя и может доставлять пакеты вразнобой.

Про TCP:
1 Требуется установленное соединение для передачи данных (соединение должно быть закрыто после завершения передачи)
2 Может гарантировать доставку данных получателю
3 Повторная передача нескольких кадров в случае потери одного из них
4 Полная проверка ошибок
5 Используется для передачи сообщений электронной почты, HTML-страниц браузеров

Про UDP:
1 Протокол без соединения, без требований к открытию, поддержанию или прерыванию соединения
2 Не гарантирует доставку данных получателю
3 Отсутствие повторной передачи потерянных пакетов
4 Базовый механизм проверки ошибок. Использует вышестоящие протоколы для проверки целостности
5 UDP-пакеты с определенными границами; отправляются по отдельности и проверяются на целостность по прибытии
6 Видеоконференции, потоковое вещание, DNS, VoIP, IPTV

Использования данных протоколов при VPN-соединениях. К примеру, в OpenVPN существует возможность выбора между TCP- и UDP-протоколами. Условимся, что VPN заворачивает передаваемые данные в еще один протокол (на самом деле все намного сложнее). Если ваш VPN-туннель использует в качестве транспортного протокола TCP, то передача данных по UDP-протоколу теряет свои преимущества. Как минимум на определенном участке пути. Поэтому для VPN-туннеля советуют использовать UDP-протокол, ведь TCP будет штатно работать внутри UDP-туннеля.

#### Какие протоколы прикладного уровня знаете?
#### Чем отличается HTTP от HTTPS?
https://selectel.ru/blog/http-https/
