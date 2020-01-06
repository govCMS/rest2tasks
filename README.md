#Rest2Tasks

A stripped-back version of Amazeeio's [rest2tasks](https://github.com/amazeeio/lagoon/tree/312dfe719119f93f9293686320d0a083670c2280/services/rest2tasks) service which is now deprecated in [Lagoon](https://github.com/amazeeio/lagoon). It serves as a replacement which performs _nothing_ other than reporting back to the sender and receiver.

The sender will receive a json response which can be used programmatically if needed. The receiver will log a message to the stdout based on the input.

## Availability

A public Docker image is available, you can pull it using the following:

```shell script
$ docker pull quay.io/fubarhouse/rest2tasks
```

## Getting started

There're three ways to run this application at the moment

1. Using docker to build
    ```shell script
    $ docker build -t rest2tasks . && \
      docker run -p 3000:3000 rest2tasks;
    ```
2. Using docker to pull/run
    ```shell script
    $ docker pull quay.io/fubarhouse/rest2tasks && \
      docker run -p 3000:3000 quay.io/fubarhouse/rest2tasks;
    ```
3. Using local compilation
    ```shell script
    $ git clone git@github.com:fubarhouse/rest2tasks.git && \
      cd rest2tasks && \
      go run main.go;
    ```

## Usage

It accepts two possible input values - `branchName` and `projectName` and will simply provide messages back containing those values.

You can use `-X GET` fine however, the application assumes the user uses `-X POST` to interact with this application.

```shell script
# Trigger deploy (based on MR/PR)
$ curl http://localhost:3000/pullrequest/deploy -X POST -d projectName=myproject -d branchName=master

# Trigger deploy (not based on MR/PR)
$ curl http://localhost:3000/deploy -X POST -d projectName=myproject -d branchName=master

# Trigger promote deploy
$ curl http://localhost:3000/promote -X POST -d projectName=myproject -d branchName=master
```

## Limitations

This is intended to be a drop-in replacement for the service provided by Amazee.io which has since been deprecated for the Lagoon API. This could be re-purposed to accommodate the API if you were prepared to make it do so using the [Lagoon CLI API](https://github.com/amazeeio/lagoon-cli).

## License

MIT - This is a recreation of a small subset of features rest2tasks came with and this repository comes with no guarantees or warranties. 