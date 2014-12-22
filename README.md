# thingz-server

Thingz understood


## Queries

### All Thingz

List thingz reported over last hour

```
select * from /.*/ where time > now() - 1h limit 1
```

### All Thingz by source

List thingz reported from `demo` source over last `10m` with a limit of `100` records

```
select * from /^demo-*/ where time > now() - 10m limit 100
```

### All Thingz by metric

List CPU metrics reported from all thingz over last `hour` in `5min` groups

```
select max(total) as MaxVal,
       PERCENTILE(total, 80) as High80,
       min(total) as MinVal,
       PERCENTILE(total, 20) as Low20
from /-cpu$/
where time > now() - 1h
group by time(5m)
```
