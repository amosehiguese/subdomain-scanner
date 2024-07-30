<br/>

<div align="center">
  Like this project ? Leave us a star ⭐
</div>

<br/>

<div align="center">
  <a href="#" target="_blank">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="assets/subdomain.png">
    <img alt="Subd Logo" src="assets/subdomain.png" width="280"/>
  </picture>
  </a>
</div>


<h3 align="center">
  Subdomain Enumeration Tool 🔥.
</h3>

<br/>

<p align="center">
  <a href="CODE_OF_CONDUCT.md">
    <img src="https://img.shields.io/badge/Contributor%20Covenant-2.0-4baaaa.svg" alt="Code of conduct">
  </a>
  <a href="LICENSE">
    <img src="https://img.shields.io/badge/license-Apache%20-blue" alt="MIT">
  </a>
  <img src="https://img.shields.io/badge/status-experimental-red" alt="Experimental">
</p>


<div>
<span>

Subdomain Enumeration Scanner is a cloud-first microservices tool designed to help you discover all the subdomains associated with a specific domain.

This tool provides a way to gather valuable information that can be used for security testing, or just gaining insights into a target domain's online presence.
</span>
</div>

---




<p align="center">
  <img src="assets/subdomain-scanner.png" alt="Application Banner" width="640" >
</p>

<br/>

## Architecture

<p align="center">
  <img src="assets/withbrute.png" alt="Application Banner" width="640" >
</p>



## How It Works

The user initiates a POST request containing the target domain. Upon receiving the request, the handler parses and deserializes the payload into a local data structure. This data structure is then passed to a `scan` method responsible for identifying subdomains related to the given domain.

### Steps in the `scan` Method:

1. **Subdomain Discovery:**
   - The arguments are passed to `apiQuerySvc`, `aiBruteSvc`, and `bruteSvc`, which both return lists of subdomains. These lists are combined into a single result.

2. **DNS Resolution:**
   - The combined subdomain list is passed to `dnsResolveSvc`, which resolves the subdomains into their corresponding DNS addresse.

3. **Port Scanning:**
   - Here each subdomain are scanned for open ports.

4. **Response Construction:**
   - The final list, including subdomains and their open ports, is sent back as a response to the `frontend`.

| Service                                              | Language      | Description                                                                                                                       |
| ---------------------------------------------------- | ------------- | --------------------------------------------------------------------------------------------------------------------------------- |
| [frontend](/src/frontend)                           | Go            | Exposes an HTTP server to serve the website.|
| [apiqueryservice](/src/apiquerysvc)                     | Typescript            | Queries external api to get associated subdomains                                                           |
| [bruteservice](/src/brutesvc) | Rust           | Uses the brute force methodology of finding subdomains                       |
| [dnsresolveservice](/src/dnsresolvesvc)             | Java       | Responsible for resolving domain names to its ip addresses |
| [portscanservice](/src/portscansvc)               | Rust       | Responsible for scanning subdomains for open ports                                    |
| [aibruteservice](/src/aibrutesvc)             | Python           | Leverages Gen AI for finding subdomains by brute force using semantic understanding from target domain                           |


| Technologies                                           | Uses
| ----------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------
| Kubernetes                         | Container Orchestration System for managing containers
| Docker                         | A tool for packaging your application and its dependecies into an image that can be run as a container
| Helm                       | A dependency management tool for kubernetes environment.
| Skaffold                         | Used for speeding up development processes. It builds, runs and deploys containers into a cluster.
| Github Actions                        | Used for setting up CI/CD to improve developement time.
| Open Telemetry                         | A standard for telemetry data.
| Jaeger                     | Used for handling metrics that comes from Open Telemetry.
| Prometheus                     | Responsible for storing service insights in a time series format
| Google Gemini                         | A cutting-edge LLM to generate subdomains based on sematic understanding.
| gRPC                         | A RPC framework for service-to-service communication used in microservices.
| Protocol Buffer                         | A serialization format used by gRPC to exchange data over HTTP 2.0 protocol.
| Fluent Bit                         | A log and metrics processor which serves as a cluster-level log collector agent.
| Elastic Search                         | A logging backend.
| Kibana                         | A data visualization dashboard for Elastic-search.
|




## License

Copyright 2024 Subdomain Enumeration Tool

Licensed under the Apache License. <br/> See [LICENSE.md](LICENSE) for more information.

## Contributors ✨

<a href="https://github.com/remarkablemark">
  <img src="https://avatars.githubusercontent.com/u/93928881?s=50&u=b468eec8d146b8a918bcae959e3ee7b74ba336c2&v=4&mask=circle">
</a>



## Star History

Truly grateful for your support 💖

Happy Hacking!

