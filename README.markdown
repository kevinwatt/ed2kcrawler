
Install/Run Example
-------------------

    $ git clone git://github.com/kevinwatt/ed2kcrawler.git

First, Install Go-MySQL-Client-Library

    http://github.com/thoj/Go-MySQL-Client-Library

Create an config.cfg file

    [default]
    DBA = user
    DBP = password
    DBIP = 127.0.0.1:3306
    DBN = ED2KDB

    [amule]
    ARS = 10.8.0.1
    ARP = 4712
    ARPS = amulepasswd


Create Database

    $ mysqladmin -u root -p create ED2KDB
    $ mysql -u root -p ED2KDB < ed2kcrawler.sql

Running ed2kcrawler
    
    $ ./ed2kcrawler -l vv

