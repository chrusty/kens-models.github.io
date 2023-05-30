default:
	@mkdir -p bin
	@go build -o bin/scrape-flickr-albums cmd/scan-collections/main.go

clean:
	@mkdir -p collections/_aircraft collections/_buildings collections/_carsandtrucks collections/_militaryvehicles collections/_rail collections/_ships
	@rm -rf collections/_aircraft/* collections/_buildings/* collections/_carsandtrucks/* collections/_militaryvehicles/* collections/_rail/* collections/_ships/*
