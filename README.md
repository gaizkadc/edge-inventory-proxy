# edge-inventory-proxy

This component receives data from the [edge-controller](https://github.com/nalej/edge-controller)
and send it to the [inventory-manager](https://github.com/nalej/inventory-manager).
​
## Getting Started
​
### Prerequisites
​
* [inventory-manager](https://github.com/nalej/inventory-manager)
* [network-manager](https://github.com/nalej/network-manager)
* [nalej-bus](https://github.com/nalej/nalej-bus)
* [vpn-server](https://github.com/nalej/vpn-server)
​
### Build and compile
​
In order to build and compile this repository use the provided Makefile:
​
```shell script
make all
```
​
This operation generates the binaries for this repo, download dependencies,
run existing tests and generate ready-to-deploy Kubernetes files.
​
### Run tests
​
Tests are executed using Ginkgo. To run all the available tests:
​
```shell script
make test
```
​
### Update dependencies
​
Dependencies are managed using Godep. For an automatic dependencies download use:
​
```shell script
make dep
```
​
In order to have all dependencies up-to-date run:
​
```shell script
dep ensure -update -v
```
​
## Known Issues​
​
## Contributing
​
Please read [contributing.md](contributing.md) for details on our code of conduct, and the process for submitting pull requests to us.
​
​
## Versioning
​
We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/nalej/edge-inventory-proxy/tags). 
​
## Authors
​
See also the list of [contributors](https://github.com/nalej/edge-inventory-proxy/contributors) who participated in this project.
​
## License
This project is licensed under the Apache 2.0 License - see the [LICENSE-2.0.txt](LICENSE-2.0.txt) file for details.
