package main

import (
    "fmt"
    "time"
    "flag"
    "github.com/kelvins/sunrisesunset"
    "math"
    "os"
)

const DATE_FORMAT = "2006-01-02"
const OUTPUT_FORMAT = "2006-01-02 15:04:05 -0700 MST"

type CLIArgs struct {
    lat, long float64
    date string
    nopause bool
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
    flag.BoolVar(&args.nopause, "nopause", false, "don't wait for user to press Enter")
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
    sunrise = time.Date(date.Year(), date.Month(), date.Day(), sunrise.Hour(), sunrise.Minute(), sunrise.Second(), 0, time.UTC)
    sunset = time.Date(date.Year(), date.Month(), date.Day(), sunset.Hour(), sunset.Minute(), sunset.Second(), 0, time.UTC)

    // If no error has occurred, print the results
    if err == nil {
        fmt.Printf("Sunrise\t: %s\n", sunrise.Local().Format(OUTPUT_FORMAT))
        fmt.Printf("Sunset\t: %s\n", sunset.Local().Format(OUTPUT_FORMAT))
    } else {
        fmt.Println(err)
    }

    var x string
    if !args.nopause {
        fmt.Fprintln(os.Stderr, "\nPress ENTER to continue...")
        fmt.Scanln(&x)
    }
}
