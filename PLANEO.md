# planeodev fork of stubbi/paperclip-operator

Branch `planeo/main` = an upstream release tag + Planeo's patches (default branch of this fork). Design:
[ohanalabs-ai/intelligent-engagement-os design 0001 §3.7 / 0002](https://github.com/ohanalabs-ai/intelligent-engagement-os/blob/main/docs/design/0002-repo-and-cluster-topology.md).
Deployed by `planeodev/planeo-infra` `addons/product/paperclip-operator` on EKS `planeo-prod`.

| Patch | Why | Upstream PR |
|---|---|---|
| **Namespaced mode** — `WATCH_NAMESPACES="ns-a,ns-b"` restricts the manager cache to those namespaces; in that mode the operator refuses cluster-scoped RBAC (multi-namespace cloud sandbox, Kubernetes execution mode) and skips the cluster-scoped finalizer cleanup | The chart's ClusterRole (Secrets everywhere, `pods/exec`, namespace create/delete, ClusterRole/Binding writes) bound cluster-wide is cluster-admin-equivalent on a shared platform cluster; namespaced mode runs with RoleBindings per watched namespace + a read-only cluster role | to open |

Base: `v0.19.1` (30d762c). Rebase `planeo/main` onto each upstream release we adopt (kind rehearsal first).
Image: `ghcr.io/planeodev/paperclip-operator:<branch>-<sha7>` (`.github/workflows/planeo-docker-multiarch-cicd.yaml`); upstream's
release/OperatorHub workflows are disabled in this fork.
