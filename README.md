# Epoch

This CLI automates translating from human readable dates to epochs → stored on MongoDB.

## CLI Commands

#### Int epoch to Date

`./epoch translate --epoch 12343123241`

This would just return the corresponding date with a nice understandable format (e.g., "2006-05-12; 14:30h")

#### Date to int epoch

`./epoch translate --date "2012-05-06"`

This command returns the corresponding `.Unix()` value using the time package. For now, the date should be passed in `RFC3339` format


### Additional notes

In order to make a binary file executable from everywhere in your system, you need the following steps:

 - Build the Go project with `go build`;
 - Add the compiled binary to `~/bin`;
 - Make sure that the `~/bin` is added to $PATH → this means adding the following line to `.zshrc`:
 
```sh
export PATH=$PATH:~/bin
```

