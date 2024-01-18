# timeSpliceFile
A library that implements the io.Writer interface and splits files by time.

## Install
```
go get github.com/xusenlin/timeSpliceFile@v0.0.2-alpha
```


## Usage

```golang
import (
    "github.com/xusenlin/timeSpliceFile"
)


func main() {
    logFile, err := timeSpliceFile.New(
        "/Users/xusenlin/golangProject/errLog",
        timeSpliceFile.SplitIntervalMinute,
        "log",
    )
    if err != nil {
        fmt.Println(err)
    return
    }
    textHandler := slog.NewTextHandler(io.MultiWriter(logFile, os.Stdout), &slog.HandlerOptions{
        Level:     slog.LevelInfo,
        AddSource: true,
    })
    logger := slog.New(textHandler)

    for range time.Tick(time.Second) {
        logger.Warn("test", slog.Int("int", 2))
    }
}
```
### Output file
- 2024-01-18-10-06.log
- 2024-01-18-10-07.log
- 2024-01-18-10-08.log
- ...

The constant "SplitIntervalMinute" is used to split your files by minutes, as well as by days or months.