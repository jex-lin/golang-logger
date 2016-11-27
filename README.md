# Logger

It's a lightweight and simple log package.





# Features

* Show package name, func name and line when you log.
* Set level to filter your logs.
* Set trigger when log level reached.




# Examples

### Normal

    var log = logger.New(os.Stdout)
    log.Debug("debug")
    log.Info("info")
    log.Notice("notice")
    log.Warnf("warn %s", "!!")
    log.Errorf("error %s", "!!!")
    log.Critical("critical")

output :

    2016/11/27 16:24:30 [Debug] main main(:12) > debug
    2016/11/27 16:24:30 [Info] main main(:13) > info
    2016/11/27 16:24:30 [Notice] main main(:14) > notice
    2016/11/27 16:24:30 [Warn] main main(:15) > warn !!
    2016/11/27 16:24:30 [Error] main main(:16) > error !!!
    2016/11/27 16:24:30 [Critical] main main(:17) > critical

### Set Level

    var log = logger.New(os.Stdout)
    log.SetLevel("notice")
    log.Debug("debug") // won't print
    log.Info("info")   // won't print
    log.Notice("notice")
    log.Warn("warn")
    log.Error("error")
    log.Critical("critical")

output :

    2016/11/27 16:27:58 [Notice] main main(:15) > notice
    2016/11/27 16:27:58 [Warn] main main(:16) > warn
    2016/11/27 16:27:58 [Error] main main(:17) > error
    2016/11/27 16:27:58 [Critical] main main(:18) > critical


### Set trigger

    func main() {
        var log = logger.New(os.Stdout)
        log.SetTrigger("critical", do)
        log.Warn("warn")
        log.Error("error")
        log.Critical("critical")
    }

    func do() {
        fmt.Println("Critical happended.")
    }

output :

    2016/11/27 16:37:17 [Warn] main main(:13) > warn
    2016/11/27 16:37:17 [Error] main main(:14) > error
    2016/11/27 16:37:17 [Critical] main main(:15) > critical
    Critical happened.




### Alternative output

File

    f, err := os.OpenFile("dev.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
    if err != nil {
      log.Printf("Create log file error: %v", err)
    }
    var log = logger.New(f)

Standard output

    var log = logger.New(os.Stdout)

Buffer

    var buf bytes.Buffer
    var log = logger.New(&buf)




# TODO list

* Custom format
