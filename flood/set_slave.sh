#!/bin/bash

MASTER_USER="replication"

mysql --port 3306 -u root -e "
FLUSH TABLES WITH READ LOCK;
SHOW MASTER STATUS;"

MASTER_STATUS=$(mysql --port 3306 -u root -e "SHOW MASTER STATUS" | grep mysql-bin)
MASTER_LOG_FILE=$(echo $MASTER_STATUS | awk '{print $1}')
MASTER_LOG_POS=$(echo $MASTER_STATUS | awk '{print $2}')

mysql --port 3307 -u root -e "
STOP SLAVE;
CHANGE MASTER TO
MASTER_HOST='mysql-master',
MASTER_USER='$MASTER_USER',
MASTER_LOG_FILE='$MASTER_LOG_FILE',
MASTER_LOG_POS=$MASTER_LOG_POS;
START SLAVE;"

mysql --port 3306 -u root -e 'UNLOCK TABLES;'

mysql --port 3307 -u root -e 'SHOW SLAVE STATUS\G'
