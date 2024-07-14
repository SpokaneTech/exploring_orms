Getting started
```shell
go mod tidy
go run . --help
```


Add a new vehicle to your garage
```shell
$ go run . add --manufacturer Mitsubishi --name 'Montero Sport'
Added a Mitsubishi Montero Sport to your gararge
```

List the vehicles in your garage
```shell
$ go run .
Mitsubishi Montero Sport
```

Add a new vehicle part to your garage
```shell
$ go run . add-part --name Muffler --cost '399.99'
Added a new part Muffler to your garage
```

List the vehicle parts in your gararge
```shell
$ go run . list-parts
Muffler ($399.99)
```
