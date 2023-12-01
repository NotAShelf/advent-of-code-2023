{
  buildGoModule,
  day ? "1",
  ...
}: let
  pname = "aoc" + day;
  version = "0.0.1";
in
  buildGoModule {
    inherit pname version;
    src = ./.;
    vendorHash = null;

    ldflags = ["-s" "-w"];

    meta = {
      description = "Advent of Code Day" + day + " in Go";
      mainProgram = "aoc2023";
    };
  }
