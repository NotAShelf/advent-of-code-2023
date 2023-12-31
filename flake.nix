{
  description = "Advent of Code 2023";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs";
  };

  outputs = {
    self,
    nixpkgs,
  }: let
    systems = ["x86_64-linux"];
    forEachSystem = nixpkgs.lib.genAttrs systems;
    pkgsForEach = nixpkgs.legacyPackages;
  in {
    packages = forEachSystem (system: {
      day-1 = pkgsForEach.${system}.callPackage ./day-1 {};
      day-2 = pkgsForEach.${system}.callPackage ./day-2 {};
    });

    devShells = forEachSystem (system: {
      default = pkgsForEach.${system}.mkShell {
        buildInputs = with pkgsForEach.${system}; [
          go
          gopls
        ];
      };
    });

    templates = {
      default = {
        path = ./template;
        description = "Template of daily AoC challenges";
      };
    };
  };
}
