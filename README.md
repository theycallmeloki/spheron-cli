# spheron-cli

A CLI for Spheron Protocol

## Installation 

Pick a release flavor for a distribution of your choice and install the binary to a location that is in the PATH 

### Windows 

Drag and drop the binary into Windows/System32 folder, accepting permissions if any

### Linux / Mac

chmod +x spheronctl && cp spheronctl /usr/local/bin

## Running

`spheronctl --help` - lists top level commands you can run using spheronctl 

You probably want to configure the CLI to be used locally, 

Collect an API Key from the Spheron Console. 
Profile > User Settings > Tokens > Create Token

Run the following command locally: 

`spheronctl configure`

You should also pick a `project`, you can do so with: 

`spheronctl set project` 

Navigate to the project you're working on currently and select it

You should be all setup! 🎉

## Directory Structure

The `pkg` /spheron folder is a thin wrapper around API contracts exposed by spheron (this can be maintained in a seperate repo, and added to this repo as a submodule, treating it as the Go-SDK)

The `cmd` /spheron folder is for logic that glues the Go-SDK to Spheronctl CLI, think of it as thin stubs that work withs with `pkg` to enable local CLI developer use cases like writing to a config file / interacting with spheron protocol without having to switch to the browser

## TODO

[ ] Distribution channels - QoL / GTM strategy

- [ ] Mac - brew (better to come from `argoapp-live` organization)

- [ ] Linux - `curl -sSL <gist.sh> | sh` installer from releases

- [ ] Windows - scoop / choco / Powershell (?) / `.msi` package

[x] Configure local with secret 

[x] Organization endpoints

[ ] Print a table of reasons for why a organization is overdue

[x] Coupon endpoint

[ ] Coupon - test happy path if one is present (should work, mostly)

[x] Invites endpoints

[x] Project - Environment Variables endpoints

[x] `spheronctl env push --envfile=.env` to automatically push deployment environment variables from `.env` to spheron

[ ] `spheronctl env pull --envfilepath=.env.local` to automatically pull down local environment variables from spheron

[x] Project - Deployment Environment endpoints

[x] Project - Domains endpoints

[x] Project - Deployment - Create (Post API) is ~~not~~ functional ATM, ~~likely needs plumbing from OAuth 2.0 layers which is likely not currently present without a FE context~~

[x] Project - Deployment endpoints

[ ] Find out if the Get Framework Suggestions endpoint is relevant from local development environment - Only after local is in remote would it make much sense to get a framework suggestion, assuming they haven't pushed it to remote, this might not add much value, though if they already connected their repo with Git, likely might help with Project - Domain - Create (Post API) 

[x] Refactor `utils` from `pkg` to `cmd`

[x] Seems every API provides a `error` bool and a `message` string, add support to all the responses to error out at a `pkg` level if these keys exist, panicking and exiting from the CLI

[x] Find out if we can stream logs with `--watch` depending on the topic ID for the deployment, seems likely it might be coming from a rabbitmq / kafka queue 

[ ] Not particularly sure if `set deployment` made sense, since if they made a new commit, that seems to be creating a seperate deployment, so when you redeploy and look for logs, it redeploys the previous commit and shows logs, thereby not making too much sense, moved behind a flag for now