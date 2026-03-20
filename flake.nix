{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    systems.url = "github:nix-systems/default";
  };

  outputs = {
    self,
    nixpkgs,
    systems,
  }: let
    inherit (nixpkgs) lib;
    eachSystem = lib.genAttrs (import systems);
  in {
    overlays.default = final: prev: {
      passgen = self.packages.${prev.system}.passgen;
    };

    overlay = self.overlays.default;

    packages = eachSystem (system: let
      pkgs = import nixpkgs {
        inherit system;
      };
    in {
      passgen = pkgs.buildGoModule {
        pname = "passgen";
        version = "latest";
        src = builtins.path {
          name = "passgen-src";
          path = ./.;
        };

        buildTarget = ".";

        vendorHash = null;
        doCheck = true;
      };

      default = self.packages.${system}.passgen;
    });

    apps = eachSystem (system: {
      default = {
        type = "app";
        program = "${self.packages.${system}.passgen}/bin/passgen";
      };
    });

    devShells = eachSystem (system: {
      default = let
        pkgs = import nixpkgs {
          inherit system;
        };
      in
        pkgs.mkShell {
          buildInputs = [self.packages.${system}.passgen];
        };
    });
  };
}
