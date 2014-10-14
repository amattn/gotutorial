[GoTutorial.net][] Link Shortener
=================================


[GoTutorial.net]: http://gotutorial.net


This is a sample project accompanying GoTutorial.net's web app lesson.


Database
--------

Install or update the driver with 

    go get -u github.com/lib/pq

You can install the test db with [macports](http://www.macports.org/install.php)

    # install postgres
    sudo port install postgresql93-server

    # setup db locally
    sudo mkdir -p /gotutorial/db
    sudo chown postgres:postgres /gotutorial/db
    sudo su postgres -c '/opt/local/lib/postgresql93/bin/initdb -D /gotutorial/db' 

In a different terminal window, start like so:

    sudo su postgres -c '/opt/local/lib/postgresql93/bin/pg_ctl -D /gotutorial/db -l /gotutorial/db/postgres.log start'

To uninstall simple remove the root gotutorial directory:

    sudo rm -Rf /gotutorial

To setup the db:

    sudo su postgres -c '/opt/local/lib/postgresql93/bin/createuser --superuser tutorial -U postgres'
    sudo su postgres -c "/opt/local/lib/postgresql93/bin/psql -c \"ALTER USER tutorial WITH PASSWORD 'changeme';\""
    sudo su postgres -c "/opt/local/lib/postgresql93/bin/psql -c \"CREATE TABLE links (code VARCHAR(64) NOT NULL,url TEXT NOT NULL,PRIMARY KEY(code));\""


