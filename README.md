# ssl-handshake
A command-line debugging tool to check the latency of SSL handshake.

## Usage
```
ssl-handshake v1.3 ( https://tuladhar.github.com/ssl-handshake )

A debugging tool to check the latency of SSL handshake.

Usage:
  ssl-handshake [options] hostname:port

Options:
  -t <timeout>   SSL handshake timeout in milliseconds (default: 5000)
  -c <count>     Stop after <count> SSL handshakes (default: 0)
  -i <interval>  Wait <interval> milliseconds between SSL handshakes (default: 1000)

Examples:
  ssl-handshake tuladhar.github.com:443
  ssl-handshake -c 3 -i 500 tuladhar.github.com:443
  ssl-handshake -t 500 imap.gmail.com:993
```

## Examples

Continously send SSL handshake every 1 second (similar to ping). To cancel, press `Ctrl-C`, and it will display the handshake statistics summary and exit.
```
$ ssl-handshake tuladhar.github.com:443
Starting ssl-handshake v1.3 ( https://tuladhar.github.com/ssl-handshake ) at Fri, 24 Dec 2021 14:53:40 +0545
SSL handshake with tuladhar.github.com:443: time=102 ms
SSL handshake with tuladhar.github.com:443: time=101 ms
SSL handshake with tuladhar.github.com:443: time=105 ms
SSL handshake with tuladhar.github.com:443: time=103 ms
SSL handshake with tuladhar.github.com:443: time=125 ms
^C
--- tuladhar.github.com:443 handshake statistics ---
5 handshakes sent, finished: 5, failed: 0, time: 536 ms
rtt max: 125 ms, avg: 107 ms, min: 101 ms
```

Send SSL handshake every 500ms and stop after 5 successful handshakes. 
```
$ ssl-handshake -c 5 -i 500 tuladhar.github.com:443                           14:48:10
Starting ssl-handshake v1.3 ( https://tuladhar.github.com/ssl-handshake ) at Fri, 24 Dec 2021 14:48:10 +0545
SSL handshake with tuladhar.github.com:443: time=185 ms
SSL handshake with tuladhar.github.com:443: time=94 ms
SSL handshake with tuladhar.github.com:443: time=113 ms
SSL handshake with tuladhar.github.com:443: time=115 ms
SSL handshake with tuladhar.github.com:443: time=109 ms

--- tuladhar.github.com:443 handshake statistics ---
5 handshakes sent, finished: 5, failed: 0, time: 616 ms
rtt max: 185 ms, avg: 123 ms, min: 94 ms
```

Send SSL handshake every 1 second with 110ms timeout and stop after 2 successful handshake.
```
$ ssl-handshake -t 110 -c 2 tuladhar.github.com:443                                                                15:02:20
Starting ssl-handshake v1.3 ( https://tuladhar.github.com/ssl-handshake ) at Fri, 24 Dec 2021 15:02:22 +0545
SSL handshake with tuladhar.github.com:443: time=109 ms
SSL handshake with tuladhar.github.com:443: time=0 ms error=timeout (110 ms) exceeded
SSL handshake with tuladhar.github.com:443: time=108 ms

--- tuladhar.github.com:443 handshake statistics ---
3 handshakes sent, finished: 2, failed: 1, time: 217 ms
rtt max: 109 ms, avg: 108 ms, min: 108 ms
```