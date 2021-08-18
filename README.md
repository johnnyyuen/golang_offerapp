# golang_offerapp

## Dev start up file
you can create a startup file for your environment (e.g. dev)

```
#!/bin/bash

export TOKEN_SECRET="secretfortoken"
export DBNAME=offerapp
export DBUSER=postgres
export DBPASSWORD=postgres

echo "starting go server at port 3000"

go run main.go
```

and then make sure it is executable by chmod 755 dev


Issue and trouble shooting

Got error about:
> Error illegal base64 data at input byte 0

Fixed by "Bearer " instead of "Bearer" in AuthMiddleWare function