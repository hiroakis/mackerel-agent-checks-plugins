# mackerel-agent-checks-plugin

These scripts are used as mackerel-checks-plugins that are defind `[plugin.checks.xxxx]` in mackerel-agent.conf.
You also can use these scripts as nagios plugin and sensu plugin. The specification of these scripts are same as nagios plugin and sensu plugin. (Note: I didn't do testing on Nagios and Sensu)

Details are following links

* English: http://help.mackerel.io/entry/custom-checks
* Japanese: http://help-ja.mackerel.io/entry/custom-checks

# Build

```
git clone git@github.com:hiroakis/mackerel-agent-checks-plugins.git
cd mackerel-agent-checks-plugins
make
```

If you would like to use on other OS, you can edit TARGET_OSARCH in Makefile.

# How to use

## mackerel-check-proc

```
mackerel-check-proc -name ntpd -critunder 1 -critover 1 -warnunder 1 -warnover 1
```

## mackerel-check-port

```
mackerel-check-port -host localhost -port 11211 -level warn
```

## mackerel-check-mysql-replication

```
mackerel-check-mysql-replication -host=127.0.0.1 -port=3306 -username=USER -password=PASSWORD -warn=5 -crit=10
```

# LICENSE

MIT.