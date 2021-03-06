<h1>Testing building a REST API in Golang using the net, encoding, and ioutil packages as well as the Gorilla toolkit</h1>

<h2>Firstly, for the car data</h2>
Go to localhost:10000/api/v1 to access the server home page, and visit paths /listCars and /listCars2 to retrieve the data on the server using two different methods with the same outcome. 

Alternatively, a specific make of car can be retrieved by specifying its name in the URL path \listCars\"brand name". 

Car data can also be POSTed to the database on the path /addCar using 
```console
go run client/client.go post
```

Car data can be DELETEd from the database on the path /deleteCar/"brand name" using
```console
go run client/client.go delete "brand name"
```

<h2>Secondly, for the book data </h2>
Go to localhost:10000/api/v2 to access the server home page. Next, list all books by going to path /listBooks
If you wish to query the database by the books' ID, Name, Author, ISBN10, or Language, go to localhost:10000/api/v2/get"enterCategory"/query. 

For example, localhost:10000/api/v2/getLanguage/eng