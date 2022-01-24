# Pull base image.
FROM mariadb

VOLUME ["/var/lib/mysql"]

RUN chown -R mysql /var/lib/mysql
RUN chgrp -R mysql /var/lib/mysql


# ====== Docker command ======
# docker build -t apiserver .
# docker images
# docker run --network="host" api

# docker compose