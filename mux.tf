resource "null_resource" "helm_install_http_grpc_mux" {
  provisioner "local-exec" {
    # environment = {
    #   KUBECONFIG = "${path.root}/creds/config"
    # }
    command = "/usr/local/bin/helm install http-grpc-mux --atomic -f ${path.module}/http-grpc-mux/values.yaml ${path.module}/http-grpc-mux/ --namespace transforms-demo"
  }
}
