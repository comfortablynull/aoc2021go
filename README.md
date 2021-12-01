# Advent of Code 2021 Go 1.18

## Building go 1.18 typeparams

1. Clone go repo into goroot `git clone https://go.googlesource.com/go goroot`
2. Checkout `dev.typeparams`
3. Go to the `src` directory and run `./all.bash`
4. Point your ide to the built binary
5. Add `GOEXPERIMENT=unified` to your build environment

## Testing

1. Your build environment for testing has to be set to `GOEXPERIMENT=unified`
2. Your build arg has to turn vet off for testing `-vet=off`