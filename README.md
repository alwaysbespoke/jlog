# jlog
Basic JSON logger

## How to Use

### Without Fields

```golang
import log "github.com/alwaysbespoke/jlog"

log.Log(log.INFO, "Insert message", nil)
```

### With Fields

```golang
import log "github.com/alwaysbespoke/jlog"

log.Log(log.INFO, "Insert message", log.Fields{
    "fieldName":"fieldData",
})
```

### Sample Output

```golang
2019/07/18 10:52:57 {"src":<source file path string>,"fname":<function name string>,"type":<level string>,"message":<message string>,"data":{}}
```


### Levels

INFO

ERROR

WARNING

DEBUG

TRACE

INIT

FATAL