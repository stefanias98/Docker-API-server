Testing building a REST API in Golang using the net, encoding, and ioutil packages as well as the Gorilla toolkit

Go to localhost:10000 to access the server home page, and visit paths /listCars and /listCars2 to retrieve the data on the server using two different methods with the same outcome. 

Alternatively, a specific make of car can be retrieved by specifying its name in the URL path \listCars\"brand name". 

Car data can also be POSTed to the database on the path /addCar using 
```console
go run client/client.go post
```

Car data can be DELETEd from the database on the path /deleteCar/"brand name" using
```console
go run client/client.go delete "car name"
```