
1. Deploy a version of rocketchat
   1. Default to latest
2. Backup the data
3. An easier way to handle API calls (use Go.SDK)
4. Upgrade deployments without damaging existing one
5. Save settings
6. Save states as contexts
7. Save state of running deployments

[//]: # (roctl deploy --docker --version=latest --name=something)
[//]: # (roctl list)
[//]: # (roctl backup <deployment name/id>)

we want to save some data about the deployments

starting with docker