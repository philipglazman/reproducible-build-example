with import <nixpkgs> {}; with buildGo117Module;

buildGo117Module rec {
    name = "example";
    version = "0.1.0";
    src = ./.;
    vendorSha256 = null;

    ldflags = [
      "-s" "-w"
    ];
}