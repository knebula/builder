# KNebula Builder

KNebula repositories use the [mage](https://github.com/magefile/mage) build tool. This repository provides reusable build logic that provide consistent builds and the KNebula Build Tool (kbt) that can be used to manage build configuration files and run build commands.

## KNebula Build Tool (kbt)

The KNebula Build Tool is a command line build tool. **kbt** is simplest to use if repositories follow a consistent layout, however, configuration files can be used if different layouts are required or preferred.

A typical user flow using **kbt** is as follows:

- `kbt init`: initialises the repository to use **kbt**
- `kbt test`: runs all tests defined in the repository
- `kbt build`: build cmds
- `kbt package`: generates docker images for cmds
- `kbt deploy`: deploys cmds to a kubernetes cluster

With the exception of `kbt init`, all commands are proxies for **mage** targets. If you have **mage** installed you can replace `kbt` with `mage` in the above commands.

General help can be printed using `kbt help`. Some command have options that can be provided to modify the default behavior, these can be printed using `kbt help <command>`.