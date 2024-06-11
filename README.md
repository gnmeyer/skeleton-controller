A simple app that creates a service and a ingress when a deployment is created.



Setting up Kind cluster for ingress controller


`extraPortMappings` allow the local host to make requests to the Ingress controller over ports 80/443

`node-labels` only allow the ingress controller to run on a specific node(s) matching the label selector

`cat <<EOF | kind create cluster --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  kubeadmConfigPatches:
  - |
    kind: InitConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        node-labels: "ingress-ready=true"
  extraPortMappings:
  - containerPort: 80
    hostPort: 80
    protocol: TCP
  - containerPort: 443
    hostPort: 443
    protocol: TCP
EOF`

kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml


kubectl port-forward --namespace=ingress-nginx service/ingress-nginx-controller 8080:80
