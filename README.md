# Knative Serving + gRPC + Gloo Transforms + HTML output

> **NOTE:** This repo and documentation is still under construction so errors are likely to be present

This repo contains sample code intended to be used for debugging and testing Gloo's Proxy + Knative Eventing, more specifically, given the (current) limitations on Eventing not being able to handle/accept/use gRPC relying instead on HTTP.

TODO

## Tech Stack

- [Gloo][1]
- [Knative][2]

## Deployment

In order to deploy the service into K8s execute the following command

<!-- TODO -->
```bash
kubectl apply -f k8s/service.yaml
```

Once the service is deployed, use `glooctl` to get the name of the `Upstream` associated with the Knative Service on your cluster and lookup the `Upstream` on port 81 (HTTP to gRPC conversion) which should be something like `default-transforms-html-xxxx-81`

```bash
glooctl get upstreams
```

Now you can deploy the `VirtualService` replacing the `.spec.VirtualHost.routes.matcher.routeAction.single.upstream.name` with the `Upstream` name you got on the previous step

```bash
kubectl apply -f k8s/virtualservice.yaml
```

Verify the `VirtualService` was properly created and is in `Accepted` state

```bash
glooctl get virtualservices
```

---

## Development

For your convenience, there is a `Makefile` available that provides a sandboxed build environment (Docker container) complete with Bazel and Gazelle, that is capable of building the required binaries and Docker images to test and deploy the services into a K8s cluster.

To setup the build environment just run:

```bash
make setup
```

To use the build environment use

```bash
make work
```

### Repo Structure

- `api/`: Contains the proto definitions
- `k8s/`: Contains the `service.yaml` and `virtualservice.yaml` files to deploy onto K8s
- `src/`: Contains the code that is executed on the container
- `templates/`: Contains sample HTML templates to use as HTML output
- `tooling/`:
  - `bazel/`: Contains the bazel dependencies and configurations
  - `docker/`: Contains the files required to build the sandbox Docker images. It also **MUST** contain you Docker Hub credentials if you intend to push images to your repo



[1]: https://www.solo.io/glooe
[2]: https://knative.dev
