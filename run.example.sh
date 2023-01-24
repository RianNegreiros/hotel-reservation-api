go build -o lodging cmd/web/*go
./lodging -dbname=golodging -dbuser=golodging -cache=false -production=false