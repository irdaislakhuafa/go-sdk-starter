{
  description = "A Project Starter for https://github.com/irdaislakhuafa/go-sdk.git.";

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
        minimalPkgs = [
          pkgs.gopls
          pkgs.gotools
        ];
        devEnv = {
          COLORTERM = "truecolor";
        };
      in
      {
        inherit self;
        devShells = {
          ide = pkgs.mkShell {
            buildInputs = minimalPkgs ++ [
              pkgs.go
              pkgs.helix
              inputs.scls.defaultPackage.${system}
            ];
            env = devEnv;
            shellHook = '''';
          };
          ide-minimal = pkgs.mkShell {
            buildInputs = minimalPkgs;
            env = devEnv;
            shellHook = '''';
          };
        };
      }
    );
}
