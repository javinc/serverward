curl https://api.ratatxt.com/outbox \
-X POST \
-u "key_Mi5iMzkyMjMwMDFiZTdlYTVlNzM2MTRkZDM0OTYyNDMxNmUwNTRkMDY2:" \
-d '{"text": "'`hostname`' '$1' is very loaded right now. please pay attention","origin": "639751303274", "address": "09353708662"}'

echo $1 `date` >> out.log
