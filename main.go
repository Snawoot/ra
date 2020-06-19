package main

import (
    "fmt"
    "time"
    "flag"
    "github.com/kelvins/sunrisesunset"
    "math"
    "os"
)

const DATE_FORMAT = "2006.01.02"

type CLIArgs struct {
    lat, long float64
    date string
}

func perror(msg string) {
    fmt.Fprintln(os.Stderr, "")
    fmt.Fprintln(os.Stderr, msg)
}

func arg_fail(msg string) {
    perror(msg)
    perror("Usage:")
    flag.PrintDefaults()
    os.Exit(2)
}

func parseArgs() CLIArgs {
    var args CLIArgs
    flag.Float64Var(&args.lat, "lat", math.NaN(), "latitude")
    flag.Float64Var(&args.long, "long", math.NaN(), "longitude")
    flag.StringVar(&args.date, "date", time.Now().Format(DATE_FORMAT),
                   "date in YYYY.MM.DD format. Default is current date.")
    flag.Parse()

    if math.IsNaN(args.lat) {
        arg_fail("Latitude is not specified")
    }
    if math.IsNaN(args.long) {
        arg_fail("Longitude is not specified")
    }
    if _, err := time.Parse(DATE_FORMAT, args.date) ; err != nil {
        arg_fail("Bad date string: " + err.Error())
    }
    return args
}

func main() {
    args := parseArgs()
    date, err := time.Parse(DATE_FORMAT, args.date)
    if err != nil {
        panic(err)
    }

    p := sunrisesunset.Parameters{
        Latitude:  args.lat,
        Longitude: args.long,
        UtcOffset: 0.0,
        Date:      date,
    }

    // Calculate the sunrise and sunset times
    sunrise, sunset, err := p.GetSunriseSunset()

    // If no error has occurred, print the results
    if err == nil {
        fmt.Println("Sunrise:", sunrise.Format("15:04:05")) // Sunrise: 06:11:44
        fmt.Println("Sunset:", sunset.Format("15:04:05")) // Sunset: 18:14:27
    } else {
        fmt.Println(err)
    }
}
