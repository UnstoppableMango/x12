{
  description = "Path based application framework";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    systems.url = "github:nix-systems/default";
    flake-parts.url = "github:hercules-ci/flake-parts";

    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.flake-utils.inputs.systems.follows = "systems";
    };

    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs =
    inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = import inputs.systems;
      imports = [ inputs.treefmt-nix.flakeModule ];

      perSystem =
        {
          inputs',
          pkgs,
          system,
          ...
        }:
        let
          inherit (inputs'.gomod2nix.legacyPackages) buildGoApplication;

          x12 = buildGoApplication {
            pname = "x12";
            version = "0.0.1";
            src = ./.;
            modules = ./gomod2nix.toml;
          };
        in
        {
          _module.args.pkgs = import inputs.nixpkgs {
            inherit system;
            overlays = [ inputs.gomod2nix.overlays.default ];
          };

          packages.x12 = x12;
          packages.default = x12;

          devShells.default = pkgs.mkShell {
            buildInputs = with pkgs; [
              ginkgo
              git
              gnumake
              go
              gomod2nix
            ];

            GINKGO = "${pkgs.ginkgo}/bin/ginkgo";
            GO = "${pkgs.go}/bin/go";
          };

          treefmt = {
            programs.nixfmt.enable = true;
            programs.gofmt.enable = true;
          };
        };
    };
}
