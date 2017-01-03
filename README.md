# httpcmd
Execute any command through http

# Usage

```
Usage of httpcmd:
  -port string
        Port number. (default "9002")
  -v    Show version.

Example:
  $ curl localhost:9002 -F cmd="./a.sh hello"
  hello
  error will output here too(if it has error)
  exit status 1
  $ curl localhost:9002 -F cmd="./b.sh hello"                          
  hello
  $
```