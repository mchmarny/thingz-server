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
select * from /^demo.*/ where time > now() - 10m limit 100
```

### All Thingz by metric

List CPU metrics reported from all thingz over last `hour` in `5min` groups

```
select min(value) as MinVal,
       PERCENTILE(value, 25) as LowPercentile,
       mean(value) as MedVal,
       PERCENTILE(value, 75) as HighPercentile,
       max(value) as MaxVal
from /^demo.*/
where time > now() - 1h
group by time(5m)
```

### Speeding things up a bit

Continuous queries let us pre-compute expensive select into another time series in real-time. Here is for example a continuous down-sampling of many series for a single host:

```
select min(value) as MinVal,
       PERCENTILE(value, 25) as LowPercentile,
       mean(value) as MedVal,
       PERCENTILE(value, 75) as HighPercentile,
       max(value) as MaxVal
from /^demo.*/
group by time(5m)
into 5m.:series_name
```

Now we can execute the complex query for each series from that host with an instant response

```
select * from /^5m.demo.*/ limit 1
```