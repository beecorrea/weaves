# Weaves

`Weaves` is a monorepo for my projects. It uses [`moon`](https://github.com/moonrepo/moon) for codegen and monorepo management. 

To list all projects and their descriptions, run `moon query projects` at the root of `Weaves`. 

## Architecture
Each directory in this repo is called a `weave`, which extend of a `moon` project.
For now, `weaves` support a resource called `Hack`([\[1\]](https://github.com/kubernetes/kops/issues/444#issuecomment-246913433), [\[2\]](https://github.com/kubernetes/kubernetes/blob/master/hack/README.md)), which are scripts and utils.

