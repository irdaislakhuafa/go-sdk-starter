{
  description = "A very basic flake";

  inputs = {
    # nixpkgs input
    nixpkgs.url = "github:NixOS/nixpkgs/9755fc6210088c24e6a0d95d484776c6dcad4e3d";

    # nix flake utils to build each system
    flake-utils = {
      url = "github:numtide/flake-utils";
    };

    # simple completion language server, used for helix
    scls = {
      inputs.nixpkgs.follows = "nixpkgs";
      url = "github:estin/simple-completion-language-server";
    };
  };

  outputs =
    { self, ... }@inputs:
    inputs.flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import inputs.nixpkgs { inherit system; };
      in
      {
        inherit self;
        devShells = {
          ide = pkgs.mkShell {
            buildInputs = [
              pkgs.gopls
              pkgs.go
              pkgs.gotools
              inputs.scls.defaultPackage.${system}
              pkgs.helix
            ];
            env = {
              COLORTERM = "truecolor";
            };
            shellHook = '''';
          };
          ide-minimal = pkgs.mkShell {
            buildInputs = [
              pkgs.gopls
              pkgs.gotools
              pkgs.helix
            ];
            env = {
              COLORTERM = "truecolor";
            };
            shellHook = '''';
          };
        };
      }
    );
}
