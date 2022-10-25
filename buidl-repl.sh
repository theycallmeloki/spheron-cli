echo '----------------------------------------'
echo 'dropping previous configuration file' 
rm -rf ~/.config/spheron
echo '----------------------------------------'

echo '----------------------------------------'
echo 'creating new test build of spheronctl'
go build -o spheronctl -ldflags="-X 'github.com/theycallmeloki/spheron-cli/cmd/spheron.version=TEST_BUILD'" main.go
echo '----------------------------------------'

echo '----------------------------------------'
echo 'checking for spheron in configuration directory'
ls -1 ~/.config | grep spheron
echo '----------------------------------------'

echo '----------------------------------------'
echo 'CMD: spheronctl --help'
./spheronctl --help
echo '----------------------------------------'

echo '----------------------------------------'
echo 'CMD: spheronctl configure'
cat .token | ./spheronctl configure
echo '----------------------------------------'

echo '----------------------------------------'
echo 'display current configuration file'
cat ~/.config/spheron/spheron.json
echo '----------------------------------------'

echo '----------------------------------------'
echo 'CMD: spheronctl organization --help'
./spheronctl organization --help
echo '----------------------------------------' 

echo '----------------------------------------'
echo 'CMD: spheronctl organization'
./spheronctl organization
echo '----------------------------------------'

echo '----------------------------------------'
echo 'CMD: spheronctl set --help'
./spheronctl set --help
echo '----------------------------------------'

echo '----------------------------------------'
echo 'CMD: spheronctl set project'
./spheronctl set project
echo '----------------------------------------'

echo '----------------------------------------'
echo 'CMD: spheronctl env push -e=.env.sample'
./spheronctl env push -e=.env.sample
echo '----------------------------------------'