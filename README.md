## What

`depoch` is a simple filter to convert epoch timestamps, in [Concourse](http://concourse.ci) logs, to something human readable.

For example, this line:

```
{"timestamp":"1479847503.795126200","source":"atc","message":"atc.baggage-collector.could-not-locate-worker",...}
```

Will be converted to this:

```
{"timestamp":"2016/11/22 20:45:03.795126 (UTC)","source":"atc","message":"atc.baggage-collector.could-not-locate-worker",...}
```

Using the `-z` option you can also pass in a timezone string and the times will be displayed appropriately.

Download binaries [here](https://github.com/jdeppe-pivotal/depoch/releases).

## Build

`go build github.com/jdeppe-pivotal/depoch/...`
