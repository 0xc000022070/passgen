# CLI tool to generate simple passwords

Generates a simple password up to 64 characters, using uppercase and lowercase letters, digits and symbols (by default).

## Install

### With Go

`go install github.com/0xc000022070/passgen@latest`

### With Nix (flakes)

Add the passgen flake to your `flake.nix`:

```nix
{
  inputs = {
    passgen = {
      url = "github:0xc000022070/passgen";
      inputs = {
        nixpkgs.follows = "nixpkgs";
        flake-utils.follows = "flake-utils";
      };
    };
  };
}
```

Then install it in your configuration:

```nix
# In your home-manager configuration
home.packages = [
  inputs.passgen.defaultPackage.${system}
];

# Or in your system configuration
environment.systemPackages = [
  inputs.passgen.defaultPackage.${system}
];
```

## Usage

## Examples

Default flags, password length 32 with symbols:

- `passgen`

Outputs `cls6{uUN{&VZs:9i?Qhe(C7)6V(U#-hX`

Password length 30 without symbols:

- `passgen -l=30 -s=false`

Outputs `sDZCw4YO5EUa37FPrbHO82zo0RWjXU`

## CLI Reference

```console
$ passgen --help
Usage: passgen [flags]...
Generates a random password from CLI.

Examples:
  passgen -l=16 -s
  passgen -l=128 --symbols=false

Flags:
  -l, --length  Password length
  -s, --symbols Indicates if the password should contain symbols
```

## Disclaimer

This password generator does **NOT** implement high entropy strategies and should not be used to extreme security requirement applications.
