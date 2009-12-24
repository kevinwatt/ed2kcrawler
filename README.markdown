
Install/Run Example
-------------------

    $ git clone git://github.com/kevinwatt/ed2kcrawler.git

First, Install mysqlgo

    http://github.com/eden/mysqlgo

Create an config.cnf file

    [default]
    DB = mysql://user:passwd@127.0.0.1:3306/ED2K
    ARS = 10.8.0.1
    ARP = 4712
    ARPS = amulepasswd


Create Database

    $ mysqladmin -u root -p create ED2KDB
    $ mysql -u root -p ED2KDB < ed2kcrawler.sql

Running ed2kcrawler
    
    $ ./ed2kcrawler -l vv

