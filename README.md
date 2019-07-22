# jlog
Basic JSON logger

## How to Use
---

## Without Fields

```golang
import log "github.com/alwaysbespoke/jlog"

log.Log(log.INFO, "Insert message", nil)
```

## With Fields

```golang
import log "github.com/alwaysbespoke/jlog"

log.Log(log.INFO, "Insert message", log.Fields{
    "fieldName":"fieldData",
})
```

## Sample Output

```golang
2019/07/18 10:52:57 {"src":"C:/Users/User/go/src/project/main.go","fname":"function name","type":"level","message":"Hello world","data":{}}
```


### Levels
INFO    
ERROR   
WARNING 
DEBUG  
TRACE  
INIT   
FATAL  