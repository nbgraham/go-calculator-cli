rm -rf ./coverage
rm -rf ./bin

mkdir bin
go build -cover -o ./bin/calculator ./cmd/calculator

mkdir -p ./coverage/int
export GOCOVERDIR=./coverage/int

if ./bin/calculator ; then
  echo "should exit with non-zero exit code when no args provided"
  exit 1
fi
if ./bin/calculator invalid 1 1; then
  echo "should exit with non-zero exit code when invalid operation provided"
  exit 1
fi
if ./bin/calculator add invalid 1 ; then
  echo "should exit with non-zero exit code when non int arg1 provided"
  exit 1
fi
if ./bin/calculator add 1 invalid ; then
  echo "should exit with non-zero exit code when non int arg2 provided"
  exit 1
fi
./bin/calculator add 3 5
./bin/calculator multiply 8 5

go tool covdata percent -i=./coverage/int
go tool covdata textfmt -i=./coverage/int -o coverage/profile
