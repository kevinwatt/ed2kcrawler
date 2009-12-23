
Install/Run Example
-------------------

    $ git clone git://github.com/kevinwatt/ed2kcrawler.git

First install mysqlgo

    $ http://github.com/eden/mysqlgo

Create an config.inc file

    DB = mysql://useraccount:password@127.0.0.1:3306/ED2KDB

Create Database

    $ mysqladmin -u root -p create ED2KDB
    $ mysql -u root -p ED2KDB < ed2kcrawler.sql

Running ed2kcrawler
    
    $ ./ed2kcrawler -l vv

