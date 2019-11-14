# Knative Serving + Knative Eventing + gRPC/HTTP `cmux` + Gloo Transforms

> **NOTE:** This repo and documentation is still under construction so errors are likely to be present

This repo contains sample code intended to be used for debugging and testing Gloo's Proxy + Knative Eventing, more specifically, given the (current) limitations on Eventing not being able to handle/accept/use gRPC relying instead on HTTP.

The present code creates a KNative Service which exposes both HTTP and gRPC endpoints (with the help of [`cmux`][1]) and applies a response transformation which **should** redirect you to Google.

---

## Tech Stack

- [Gloo][2]
- [Knative][3]

## Repo Structure

- `api/`: Contains the proto definitions
- `k8s/`: Contains the `service.yaml` and `virtualservice.yaml` files to deploy onto K8s
- `src/`: Contains the code that is executed on the container
- `templates/`: Contains sample HTML templates to use as HTML output
- `tooling/`:
  - `bazel/`: Contains the bazel dependencies and configurations
  - `docker/`: Contains the files required to build the sandbox Docker images. It also **MUST** contain you Docker Hub credentials if you intend to push images to your repo

---

## Deployment

In order to deploy the service into K8s execute the following command:

<!-- TODO -->
```bash
kubectl apply -f k8s/service.yaml
```

Once the service is deployed, use `glooctl` to get the name of the `Upstream` associated with the Knative Service on your cluster and lookup the `Upstream` on port 81 (HTTP to gRPC conversion) which should be something like `default-transforms-html-xxxx-81`

```bash
glooctl get upstreams
```

Now you can deploy the `VirtualService` replacing the `.spec.VirtualHost.routes.matcher.routeAction.single.upstream.name` with the `Upstream` name you got on the previous step
Alternatively, you can define the `Upstream` spec and name and use a predefined name (the one assigned to the `Upstream`)

```bash
kubectl apply -f k8s/virtualservice.yaml
```

Verify the `VirtualService` was properly created and is in `Accepted` state

```bash
glooctl get virtualservices
```

---

## Testing & Executing

If you use Visual Studio Code and the [REST Client Extension][4], then you can use the `rest-client.http` file to execute the requests to the service, on both gRPC and HTTP endpoints

---

## Development

For your convenience, there is a `Makefile` available that provides a sandboxed build environment (Docker container) complete with Bazel and Gazelle, that is capable of building the required binaries and Docker images to test and deploy the services into a K8s cluster.

In order to be able to use `make` you will have to create a `Makefile.conf` file. Use the provided `Makefile.conf.sample` to get you started. Once done, you can executing `make` commands.

To setup the build environment just run:

```bash
make setup
```

To use the build environment use

```bash
make work
```

## Using Bazel

Bazel is the build system/platform used to build libraries, binaries, container images and packaging everything for deployment. The provided sandbox already has bazel installed so you can just start using it.

### - Building the binary

On the sandbox project root directory (`/workspace`) execute the following command:

```bash
bazel build //:mux_function
```

### - Building the Docker Container

> __NOTE__: In order to be able to push images, there must exist a `config.json` file in the `tooling/docker` directory with your Docker credentials. Use the `tooling/docker/registry.config.json.sample` as a starting point to create your own credentials file.

To use Bazel to build and push a Docker image execute this command:

```bash
bazel run //:mux_image_push --define DOCKER_REGISTRY_IMAGE_NAME=$DOCKER_REGISTRY_IMAGE_NAME
```

[1]: https://github.com/soheilhy/cmux
[2]: https://www.solo.io/glooe
[3]: https://knative.dev
[4]: https://marketplace.visualstudio.com/items?itemName=humao.rest-client
