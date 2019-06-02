package pack

import (
	"github.com/forgemods/forge-mod-manager/tools/all_curseforge_mods/command/launcher/config"
)

func init() {
	VersionJSONMap["forge-14.23.5.2823"] = &config.InheritedVersion{
		ID:                 "forge-14.23.5.2823",
		Time:               "2019-03-22T16:53:53+0000",
		ReleaseTime:        "2019-03-22T16:53:53+0000",
		Type:               "release",
		MinecraftArguments: "--username ${auth_player_name} --version ${version_name} --gameDir ${game_directory} --assetsDir ${assets_root} --assetIndex ${assets_index_name} --uuid ${auth_uuid} --accessToken ${auth_access_token} --userType ${user_type} --tweakClass net.minecraftforge.fml.common.launcher.FMLTweaker --versionType Forge",
		MainClass:          "net.minecraft.launchwrapper.Launch",
		InheritsFrom:       "1.12.2",
		Jar:                "1.12.2",
		Logging:            map[string]*config.LoggingConfig{},
		Libraries: []*config.Library{
			{
				Name: "net.minecraftforge:forge:1.12.2-14.23.5.2823",
				Downloads: &config.DownloadInfo{
					Artifact: &config.ArtifactDownloadInfo{
						Path: "net/minecraftforge/forge/1.12.2-14.23.5.2823/forge-1.12.2-14.23.5.2823.jar",
						SHA1: "cec39eddde28eb6f7ac921c8d82d6a5b7916e81b",
						Size: 5012267,
						Url:  "https://files.minecraftforge.net/maven/net/minecraftforge/forge/1.12.2-14.23.5.2823/forge-1.12.2-14.23.5.2823-universal.jar",
					},
				},
			},
			{
				Name:      "net.minecraft:launchwrapper:1.12",
				ServerReq: true,
			},
			{
				Name:      "org.ow2.asm:asm-all:5.2",
				Url:       "http://files.minecraftforge.net/maven/",
				Checksums: []string{"2ea49e08b876bbd33e0a7ce75c8f371d29e1f10a"},
				ServerReq: true,
				ClientReq: true,
			},
			{
				Name:      "org.jline:jline:3.5.1",
				Url:       "http://files.minecraftforge.net/maven/",
				Checksums: []string{"51800e9d7a13608894a5a28eed0f5c7fa2f300fb"},
				ServerReq: true,
				ClientReq: false,
			},
			{
				Name:      "net.java.dev.jna:jna:4.4.0",
				ServerReq: true,
				ClientReq: false,
			},
			{
				Name:      "com.typesafe.akka:akka-actor_2.11:2.3.3",
				Url:       "http://central.maven.org/maven2/",
				Checksums: []string{"ed62e9fc709ca0f2ff1a3220daa8b70a2870078e", "25a86ccfdb6f6dfe08971f4825d0a01be83a6f2e"},
				ServerReq: true,
				ClientReq: true,
			},
			{
				Name:      "com.typesafe:config:1.2.1",
				Url:       "http://central.maven.org/maven2/",
				Checksums: []string{"f771f71fdae3df231bcd54d5ca2d57f0bf93f467", "7d7bc36df0989d72f2d5d057309675777acc528b"},
				ServerReq: true,
				ClientReq: true,
			},
			{
				Name:      "org.scala-lang:scala-actors-migration_2.11:1.1.0",
				Url:       "http://central.maven.org/maven2/",
				Checksums: []string{"dfa8bc42b181d5b9f1a5dd147f8ae308b893eb6f", "8c9aaeeb68487ca519411a14068e1b4d69739207"},
				ServerReq: true,
				ClientReq: true,
			},
			{
				Name:      "org.scala-lang:scala-compiler:2.11.1",
				Url:       "http://central.maven.org/maven2/",
				Checksums: []string{"56ea2e6c025e0821f28d73ca271218b8dd04926a", "1444992390544ba3780867a13ff696a89d7d1639"},
				ServerReq: true,
				ClientReq: true,
			},
			{
				Name:      "org.scala-lang.plugins:scala-continuations-library_2.11:1.0.2",
				Url:       "http://central.maven.org/maven2/",
				Checksums: []string{"87213338cd5a153a7712cb574c0ddd2edfee0386", "0b4c1bf8d48993f138d6e10c0c144e50acfff581"},
				ServerReq: true,
				ClientReq: true,
			},
			{
				Name:      "org.scala-lang.plugins:scala-continuations-plugin_2.11.1:1.0.2",
				Url:       "http://central.maven.org/maven2/",
				Checksums: []string{"1f7371605d4ba42aa26d3443440c0083c587b4e9", "1ea655dda4504ae0a367327e2340cd3beaee6c73"},
				ServerReq: true,
				ClientReq: true,
			},
			{
				Name:      "org.scala-lang:scala-library:2.11.1",
				Url:       "http://central.maven.org/maven2/",
				Checksums: []string{"0e11da23da3eabab9f4777b9220e60d44c1aab6a", "1e4df76e835201c6eabd43adca89ab11f225f134"},
				ServerReq: true,
				ClientReq: true,
			},
			{
				Name:      "org.scala-lang.modules:scala-parser-combinators_2.11:1.0.1",
				Url:       "http://central.maven.org/maven2/",
				Checksums: []string{"f05d7345bf5a58924f2837c6c1f4d73a938e1ff0", "a1cbbcbde1dcc614f4253ed1aa0b320bc78d8f1d"},
				ServerReq: true,
				ClientReq: true,
			},
			{
				Name:      "org.scala-lang:scala-reflect:2.11.1",
				Url:       "http://central.maven.org/maven2/",
				Checksums: []string{"6580347e61cc7f8e802941e7fde40fa83b8badeb", "91ce0f0be20f4a536321724b4b3bbc6530ddcd88"},
				ServerReq: true,
				ClientReq: true,
			},
			{
				Name:      "org.scala-lang.modules:scala-swing_2.11:1.0.1",
				Url:       "http://central.maven.org/maven2/",
				Checksums: []string{"b1cdd92bd47b1e1837139c1c53020e86bb9112ae", "d77152691dcf5bbdb00529af37aa7d3d887b3e63"},
				ServerReq: true,
				ClientReq: true,
			},
			{
				Name:      "org.scala-lang.modules:scala-xml_2.11:1.0.2",
				Url:       "http://central.maven.org/maven2/",
				Checksums: []string{"7a80ec00aec122fba7cd4e0d4cdd87ff7e4cb6d0", "62736b01689d56b6d09a0164b7ef9da2b0b9633d"},
				ServerReq: true,
				ClientReq: true,
			},
			{
				Name:      "lzma:lzma:0.0.1",
				ServerReq: true,
			},
			{
				Name:      "net.sf.jopt-simple:jopt-simple:5.0.3",
				ServerReq: true,
			},
			{
				Name:      "java3d:vecmath:1.5.2",
				ClientReq: true,
				ServerReq: true,
			},
			{
				Name:      "net.sf.trove4j:trove4j:3.0.3",
				ClientReq: true,
				ServerReq: true,
			},
			{
				Name:      "org.apache.maven:maven-artifact:3.5.3",
				Url:       "http://central.maven.org/maven2/",
				ServerReq: true,
				ClientReq: true,
			},
		},
	}
}