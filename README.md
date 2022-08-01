<div align="center">
<img  width="320"  src="https://user-images.githubusercontent.com/4745789/182104882-ebdc05d2-26e7-4535-ad74-5ec447ff2f3a.png" align="center"  alt="retorrent logo" />

---------------------------------------

[![Release](https://img.shields.io/github/release/relogX/retorrent/all.svg)](https://github.com/relogX/retorrent/releases)
[![Twitter Follow](https://img.shields.io/twitter/follow/relog_x.svg?label=Follow&style=social)](https://twitter.com/intent/follow?screen_name=relog_x)
[![License](https://img.shields.io/github/license/relogX/retorrent.svg)](LICENSE)

</div>

#  What is retorrent?

retorrent is a re-implementation of BitTorrent protocol in Golang.

##  Using retorrent

To use retorrent, download the latest release for your platform and fire the following commands.

 1. Download the latest [retorrent/releases](https://github.com/relogX/retorrent/releases) for your platform.
 2. Extract the binary from the just downloaded compressed artifact
 3. Execute the binary as shown in the following steps

```
$ ./retorrent

        ██████  ███████ ████████  ██████  ██████  ██████  ███████ ███    ██ ████████ 
        ██   ██ ██         ██    ██    ██ ██   ██ ██   ██ ██      ████   ██    ██    
        ██████  █████      ██    ██    ██ ██████  ██████  █████   ██ ██  ██    ██    
        ██   ██ ██         ██    ██    ██ ██   ██ ██   ██ ██      ██  ██ ██    ██    
        ██   ██ ███████    ██     ██████  ██   ██ ██   ██ ███████ ██   ████    ██    
```

##  Developing retorrent

If you are a developer and want to modify retorrent, you will have first to set up a dev environment, and it has the following pre-requisites

- [Go 1.17.1](https://golang.org/)

Once you have set up all the pre-requisites, following the steps to start your development.

- Clone the repository https://github.com/relogX/retorrent
- Fire the command `go run main.go`

Once you start the server, it will download all the necessary packages and execute the root command.

##  Linting

Maintaining coding standards is extremely critical, and retorrent follows the standard [Gofmt](https://pkg.go.dev/cmd/gofmt) to reformat the code. It is customary to fire the following command before you commit.

```
make lint
```

##  Contribution Guidelines

The Code Contribution Guidelines are published at [CONTRIBUTING.md](https://github.com/relogX/retorrent/blob/master/CONTRIBUTING.md); please read them before you start making any changes. This would allow us to have a consistent standard of coding practices and developer experience.

##  The RelogX Umbrella
<div align="center">
<br />
<img  width="240"  src="https://user-images.githubusercontent.com/4745789/133601178-711aa4eb-f836-4e93-a554-22006648f75f.png" align="center"  alt="relog logo" />
<br />
<br />
</div>

[Relog](https://relog.in) is an initiative that aims to transform engineering education and provide high-quality engineering courses, projects, and resources to the community. To better understand all the common systems, we aim to build our own replicated utilities, for example, a load balancer, static file server, API rate limiter, etc. All the projects fall under [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0), and you can find their source code at [github.com/relogX](https://github.com/relogX).

##  License
retorrent is under [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)
