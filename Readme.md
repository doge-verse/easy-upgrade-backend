![avatar](https://socialify.git.ci/doge-verse/easy-upgrade-backend/image?description=1&font=KoHo&language=1&owner=1&pattern=Plus&stargazers=1&theme=Dark)

# Upgrade-Doge Backend

![avatar](https://3448297496-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2FUYPGtVjV80SevIasXRO6%2Fuploads%2FGVcmIl7gDsSRibFtwKjW%2F0xnomad_A_doge_head2.png?alt=media&token=c49cee8a-d54f-4a03-8d53-5ad82990767d)

## What is Upgrade-Doge?

Upgradeable smart contract toolkit

Front-end project：[GitHub Repo](contract-tool-web)


### Background

Deployment of upgradeable smart contracts enables rapid iteration of new project features and reduces the impact of fatal bugs on the project.
An increasing number of projects are choosing to implement upgradable contracts with Upgrades Plugins in openzeppelin.
Token rugged caused by the confusion of upgrade rights management often occurs.

### Problems

Web3 project owners are not aware that upgradeable contracts have admin rights independent of business logic. As a result, when project handover, the remaining permissions are in the hands of the departing coders, resulting in security vulnerabilities.
Web3 project members do not have easy access to the rights information of the upgradable contract.
Web3 project managers have no easy way to change contract upgrade permissions.
There is no way to timely inform the management of the change of the upgrade authority of the web3 project contract.
​

### What Upgrade-Doge is

Provide a series of visual upgrade management tools for Openzeppelin-based upgradeable contract projects.In particular, the friendly interactive interface is used to enable non-technical members to view and manage the contract upgrade authority conveniently and timely.

## Local testing and running the backend service

```bash
cp config.example.yaml config.yaml
# Modify the congfig.yaml file according to your infrastructure
make
```

## Learn more

[Upgrade-Doge Documents](https://docs.upgrade-doge.xyz/)
