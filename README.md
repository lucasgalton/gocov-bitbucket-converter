# GO code coverage format converter for Bitbucket plugin

This is a simple helper tool for generating JSON output required by 
[bitbucket-code-coverage](https://bitbucket.org/atlassian/bitbucket-code-coverage) plugin
from `go test -coverprofile=coverage.out` [output](https://blog.golang.org/cover).

## Installation

```bash
$ go get bitbucket.org/atlassian/gocov-bitbucket-converter
```

## Usage

`gocover-bitbucket` reads from the standard input:

```txt
$ go test -coverprofile=coverage.out ...
$ gocover-bitbucket -prefix=/unnecessary-chunk-of-the-path/ < coverage.out > coverage.json

  Options:

    -prefix          the bit of the file paths to remove, usually between $GOPATH/src and git root of your repo
```

You will then need to
[POST](https://bitbucket.org/atlassian/bitbucket-code-coverage/src/master/code-coverage-plugin/)
resulting JSON to your Bitbucket server.

## Contributors

Pull requests, issues and comments welcome. For pull requests:

* Add tests for new features and bug fixes
* Follow the existing style
* Separate unrelated changes into multiple pull requests

Test your changes locally before creating pull request.

See the existing issues for things to start contributing.

For bigger changes, make sure you start a discussion first by creating
an issue and explaining the intended change.

Atlassian requires contributors to sign a Contributor License Agreement,
known as a CLA. This serves as a record stating that the contributor is
entitled to contribute the code/documentation/translation to the project
and is willing to have it used in distributions and derivative works
(or is willing to transfer ownership).

Prior to accepting your contributions we ask that you please follow the appropriate
link below to digitally sign the CLA. The Corporate CLA is for those who are
contributing as a member of an organization and the individual CLA is for
those contributing as an individual.

* [CLA for corporate contributors](https://na2.docusign.net/Member/PowerFormSigning.aspx?PowerFormId=e1c17c66-ca4d-4aab-a953-2c231af4a20b)
* [CLA for individuals](https://na2.docusign.net/Member/PowerFormSigning.aspx?PowerFormId=3f94fbdc-2fbe-46ac-b14c-5d152700ae5d)

## License

Copyright (c) 2017 Atlassian and others.
Apache 2.0 licensed, see [LICENSE](LICENSE) file.