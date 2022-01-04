# Colored

This tool make your console output more colored

## How to use

#### Download package

```bash
$ go get github.com/Jourloy/Colored
```

#### Import package

```golang
import cl "github.com/Jourloy/Colored"
```

#### Print or create 

```golang
fmt.Println(cl.ChangeBG("Change background", "white"))
// or
log := cl.AddOptions("Blink text", "blink")
```