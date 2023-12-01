day: {buildGoModule}: let
  pname = "sample-go";
  version = "0.0.1";
in
  buildGoModule {
    inherit pname version;
    src = ./.;
    vendorHash = "";

    ldflags = ["-s" "-w"];

    meta = {
      description = "Advent of Code Day ${day}";
      mainProgram = pname;
    };
  }
