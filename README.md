# mackerel-agent-checks-plugin

These scripts are used as mackerel-checks-plugins that are defind `[plugin.checks.xxxx]` in mackerel-agent.conf.
You also can use these scripts as nagios plugin and sensu plugin. The specification of these scripts are same as nagios plugin and sensu plugin. (Note: I didn't do testing on Nagios and Sensu)

Details are following links

* English: http://help.mackerel.io/entry/custom-checks
* Japanese: http://help-ja.mackerel.io/entry/custom-checks

# Build and install

```
go get -d github.com/hiroakis/mackerel-agent-checks-plugins
cd $GOPATH/src/github.com/hiroakis/mackerel-agent-checks-plugins
make
sudo make install
# the binaries are installed to /usr/local/bin
```

If you would like to use on other OS, you can edit TARGET_OSARCH in Makefile.

# How to use

## mackerel-check-proc

```
mackerel-check-proc -name=ntpd -critunder=1 -critover=1 -warnunder=1 -warnover=1
```

## mackerel-check-port

```
mackerel-check-port -host=127.0.0.1 -port=11211 -level=warn
```

## mackerel-check-mysql-replication

```
mackerel-check-mysql-replication -host=127.0.0.1 -port=3306 -username=USER -password=PASSWORD -warn=5 -crit=10
```

## mackerel-check-mysql-connection

```
mackerel-check-mysql-connection -host=127.0.0.1 -port=3306 -username=USER -password=PASSWORD -warn=250 -crit=280
```

## mackerel-check-ntpoffset

```
mackerel-check-ntpoffset -warn=50 -crit=100

```
# LICENSE

MIT.
