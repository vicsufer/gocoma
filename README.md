# Gocoma

Simple command-line tool for managing AWS credentials across different projects and tools.
Based on the Ruby gem [Tacoma](https://github.com/pantulis/tacoma).

## Installation

    $ go get github.com/vicsufer/gocoma
    
## Commands
- [cd](#cd)
- [help](#help)
- [init](#init)
- [ls](#ls)
- [use](#use)

### cd
Change current directory to the one in the`repo` field of an environment. If no environment is specified, the current environment is selected.
```sh
gocoma cd
```
```sh
gocoma cd <environment>
```
### help
Will display help for gocoma or any of its subcommands. Works both with `-h` and `--help` flags 
```sh
gocoma --help
```
```sh
gocoma <command> --help
```
### init
Generate the `.gocoma.yml` file in the user home directory. This command will not override an existing `.gocoma.yml` file.
### ls
```sh
gocoma init
```
### use
Change the current environment, gocoma will add the specified identity file into the SSH agent, and will generate configuration files for the supported tools, which at this time are

- [Boto](https://github.com/boto/boto)
- [s3cmd](https://github.com/s3tools/s3cmd)
- [aws cli](https://github.com/aws/aws-cli)

```sh
gocoma use <environment>
```

## Configuring environments

Gocoma needs a special file `.gocoma.yml` in your home directory. It can create a sample for you with:

     gocoma init

The format of the `.gocoma.yml` file is pretty straighforward:
```yml
currentenv: another_environment
environments
    environment:
      aws_identity_file: "/path/to/pem/file/my_project.pem"
      aws_secret_access_key: "YOURSECRETACCESSKEY"
      aws_access_key_id: "YOURACCESSKEYID"
      region: "REGION"
      repo: "$HOME/projects/my_project"
    another_environment:
      aws_identity_file: "/path/to/another_pem.pem"
      aws_secret_access_key: "ANOTHERECRETACCESSKEY"
      aws_access_key_id: "ANOTHERACCESSKEYID"
      region: "REGION"
      repo: "$HOME/projects/another_project"
```

## Future plans

- Support for bash completion.
- Support for Azure and GCP environments.
- [Hashicorp Vault](https://www.vaultproject.io/) integration to generate temporary keys.
- Modules to perform common tasks such as:
    - Listing resources.
    - Connecting to instances.

