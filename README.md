# api.kite

#### install ectd
```
brew install etcd
```

#### start by cloning down 3 repos
```
cd /Users/bhalle/Development/go/src
mkdir github.com
cd github.com
mkdir bhalle
cd bhalle
git clone https://github.com/bhalle/pingpong.kite.git
git clone https://github.com/bhalle/foobar.kite.git
git clone https://github.com/bhalle/api.kite.git
cd api.kite
```

#### install needed kites libs
```
go get github.com/koding/kite
go get github.com/koding/kite/kontrol/kontrol
```

#### generate pem files
```
openssl genrsa -out key.pem 2048
openssl rsa -in key.pem -pubout > key_pub.pem
mkdir certs
mv key.pem certs/
mv key_pub.pem certs/
```

#### place into .bash_profile
```
KONTROL_PORT=6000
KONTROL_USERNAME="kontrol"
KONTROL_STORAGE="etcd"
KONTROL_KONTROLURL="http://127.0.0.1:6000/kite"
```

#### create kite.key file (change to use your path)
```
../../../../bin/kontrol -publickeyfile /Users/bhalle/Development/go/src/github.com/bhalle/api.kite/certs/key_pub.pem -privatekeyfile /Users/bhalle/Development/go/src/github.com/bhalle/api.kite/certs/key.pem -port 6000 -username kontrol -initial
```

#### start etcd (in 1 bash instance)
```
etcd
```


#### start kontrol (in another bash instance ... change to use your path)
```
../../../../bin/kontrol -publickeyfile /Users/bhalle/Development/go/src/github.com/bhalle/api.kite/certs/key_pub.pem -privatekeyfile /Users/bhalle/Development/go/src/github.com/bhalle/api.kite/certs/key.pem -port 6000 -username kontrol
```

#### start up the pingpong kite (in a new bash instance)
```
cd /Users/bhalle/Development/go/src/github.com/bhalle/pingpong.kite/
go run pingpong.go
```

#### start up the foobar kite (in a new bash instance)
```
cd /Users/bhalle/Development/go/src/github.com/bhalle/foobar.kite/
go run foobar.go
```

#### start up each of the kites (in a new bash instance)
```
cd /Users/bhalle/Development/go/src/github.com/bhalle/api.kite/
go run main.go
```

## hit api
[http://localhost:3000/](http://localhost:3000/)

should see: hello world

#### next try pingpong
[http://localhost:3000/ping](http://localhost:3000/ping)

should see: ping >< pong

#### next try foobar
[http://localhost:3000/foo](http://localhost:3000/foo)

should see: foo-bar

## notes & next steps
* this code sucks (ben-wa) ... DON'T use it ... only learn how kites works
* ive only tried this on osx
* you obviously need to change paths used above to fit your environment
* next will try to map out how to use kitectl
* after that will try to map out how to use with docker
* and then ... will try to map out how to use with tutum.co (and digital ocean)
* lastly say hi to your mother for me!
