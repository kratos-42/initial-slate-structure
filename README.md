# initial-slate-structure

Tracking and documenting an API endpoints can be an hard task, especially when it start growing. This `Go` package aims to create an initial structure for documenting an API using [Slate](#https://github.com/slatedocs/slate) (`.md` format).

This package is supposed to work along with another [script](https://github.com/ReiKratos/bash/tree/master/parsers#route_finder) I built:
  - The `bash` script reads from the project modules, outputting a file with a list of routes and related middlewares;
  - This `Go` package, reads from the previous output file and generates a documentation skeleton for the provided API routes.

## Handling middlewares

Considering each middleware may have different structure and arguments, we need to especially handle one by one. To do so, we just need to had the name of the middleware to the switch case in `storage/write_date` and add the desired template for it in `storage/templates`.

Here is the list of the middlewares names already being handled:

  * `include`
  * `sort`
  * `filter`
  * `paginate`
  * `authorization`

## Input document example

```txt
GET /countries/:id authorization
GET /countries authorization sort filter include
POST /countries
```

## Output page example

[![Page example](https://j.gifs.com/K194vz.gif)](https://j.gifs.com/K194vz.gif)

## Usage

```sh
go run main.go inputFilename [outputFilename]
```

### Arguments

  * `inputFilename:` File containing the list of endpoints, one per line
  * `ouputFilename:` Name of the result file. If one is not provided, the result filename is the current directory.
