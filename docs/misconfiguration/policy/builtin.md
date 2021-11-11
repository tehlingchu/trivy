# Built-in Policies

## Policy Sources

Built-in policies are mainly written in [Rego][rego].
Those policies are managed under [AppShield repository][appshield].
Terraform policies are currently powered by [tfsec][tfsec] and CloudFormation policies are powered by [cfsec][cfsec].

| Config type    | Source                        |
| ---------------| ----------------------------- |
| Kubernetes     | [AppShield][kubernetes]       |
| Dockerfile     | [AppShield][docker]           |
| Terraform      | [tfsec][tfsec-checks]         |
| CloudFormation | [cfsec][cfsec-checks]         |

For suggestions or issues regarding policy content, please open an issue under [AppShield][appshield], [tfsec][tfsec] or [cfsec][cfsec] repository.

Ansible are coming soon.

## Policy Distribution
AppShield policies are distributed as an OPA bundle on [GitHub Container Registry][ghcr] (GHCR).
When misconfiguration detection is enabled, Trivy pulls the OPA bundle from GHCR as an OCI artifact and stores it in the cache.
Those policies are then loaded into Trivy OPA engine and used for detecting misconfigurations.

## Update Interval
Trivy checks for updates to OPA bundle on GHCR every 24 hours and pulls it if there are any updates.

[rego]: https://www.openpolicyagent.org/docs/latest/policy-language/
[appshield]: https://github.com/aquasecurity/appshield
[kubernetes]: https://github.com/aquasecurity/appshield/tree/master/kubernetes
[docker]: https://github.com/aquasecurity/appshield/tree/master/docker
[tfsec-checks]: https://tfsec.dev/docs/aws/home/
[tfsec]: https://github.com/aquasecurity/tfsec
[cfsec-checks]: https://cfsec.dev/
[cfsec]: https://github.com/aquasecurity/cfsec
[ghcr]: https://github.com/aquasecurity/appshield/pkgs/container/appshield

[dockerfile-bestpractice]: https://docs.docker.com/develop/develop-images/dockerfile_best-practices/
[pss]: https://kubernetes.io/docs/concepts/security/pod-security-standards/
[azure]: https://docs.microsoft.com/en-us/azure/security/fundamentals/network-best-practices
[kics]: https://github.com/Checkmarx/kics/
