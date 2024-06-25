{
	inputs = {
		nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
	};

	outputs = { self, nixpkgs, ... }:
	let system = "aarch64-linux";
	pkgs = nixpkgs.legacyPackages.${system};
	in {
		packages."${system}".default = pkgs.buildGoModule {
			name = "notie";
			src = ./.;
			vendorHash = "";
		};
		nixosModules.notie = { config, lib, ... }: {
			# Just don't look here, it's for my server.
			options = {
				server.notie.enable = lib.mkEnableOption "Enable notie, MD share using SSH";
			};
			config = lib.mkIf config.server.notie.enable {
				systemd.services.notie = {
					wantedBy = [ "multi-user.target" ];
					serviceConfig = {
						ExecStart = "${self.packages."${system}".default}/bin/notie";
					};
				};

				services.frp.settings.proxies = [
					{
						name = "notie editor";
						localIP = "127.0.0.1";
						localPort = 9990;
						remotePort = 9990;
					}
					{
						name = "notie viewer";
						localIP = "127.0.0.1";
						localPort = 9991;
						remotePort = 9991;
					}
				];

				networking.firewall.allowedTCPPorts = [ 9991 9990 ];
			};
		};
	};
}
