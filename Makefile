car: car.go
	go build car.go
install:
	cp ./car /usr/bin
clean:
	rm car
