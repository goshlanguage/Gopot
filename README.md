# GooneyPot

A wordpress honeypot written in Golang.
This repo is inspired by DustyFresh's python Wordpress - [honeypress]. 

Status
=
Currently this project serves an index, and can be customized via the data map within handler. It also will serve wp-admin, albeit without the ability to execute load-styles.php. 

We perhaps should/could execute php within a container, and use escaped html templates from this. 


   [honeypress]: <https://github.com/dustyfresh/HoneyPress/>

