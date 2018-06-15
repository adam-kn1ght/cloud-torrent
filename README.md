<img src="https://user-images.githubusercontent.com/15851231/41191283-8f4e751e-6c06-11e8-85e3-1ffed5a6fd44.png" alt="screenshot"/>

**Insane torrent** is a a self-hosted remote torrent client, written in Go (golang). You start torrents remotely, which are downloaded as sets of files on the local disk of the server, which are then retrievable or streamable via HTTP.

### Features

* Single binary
* Cross platform
* Embedded torrent search
* Real-time updates
* Mobile-friendly
* Fast [content server](http://golang.org/pkg/net/http/#ServeContent)

See [Future Features here](#future-features)

### Install

<!-- **Binaries**

[![Releases](https://img.shields.io/github/release/admknight/insane-torrent.svg)](https://github.com/admknight/insane-torrent/releases) [![Releases](https://img.shields.io/github/downloads/admknight/insane-torrent/total.svg)](https://github.com/admknight/insane-torrent/releases)

See [the latest release](https://github.com/admknight/insane-torrent/releases/latest) or download and install it now with

```
curl https://i.adamknight.com/insane-torrent! | bash
```

*Tip*: [Auto-run `insane-torrent` on boot](https://github.com/admknight/insane-torrent/wiki/Auto-Run-on-Reboot)

**Docker**

[![Docker Pulls](https://img.shields.io/docker/pulls/admknight/insane-torrent.svg)][dockerhub] [![Image Size](https://images.microbadger.com/badges/image/admknight/insane-torrent.svg)][dockerhub]

[dockerhub]: https://hub.docker.com/r/admknight/insane-torrent/

``` sh
$ docker run -d -p 3000:3000 -v /path/to/my/downloads:/downloads admknight/insane-torrent
```
-->
**Source**

*[Go](https://golang.org/dl/) is required to install from source*

``` sh
$ go get -v github.com/admknight/insane-torrent
```

**VPS**

**[Digital Ocean]**<!--(https://m.do.co/c/011fa87fde07)-->

  <!--1. [Sign up with free $10 credit](https://m.do.co/c/011fa87fde07)-->
  1. "Create Droplet"
  2. "One-Click Apps"
  3. "Docker X.X.X on X.X"
  4. Choose server size ("$5/month" is enough)
  5. Choose server location
  6. **OPTIONAL** Add your SSH key
  7. "Create"
  8. You will be emailed the server details (`IP Address: ..., Username: root, Password: ...`)
  9. SSH into the server using these details (Windows: [Putty](https://the.earth.li/~sgtatham/putty/latest/x86/putty.exe), Mac: Terminal)
  10. Follow the prompts to set a new password
  11. Run `insane-torrent` with:

    docker run --name ct -d -p 63000:63000 \
      --restart always \
      -v /root/downloads:/downloads \
      admknight/insane-torrent --port 63000

  12. Visit `http://<IP Address from email>:63000/`
  13. **OPTIONAL** In addition to `--port` you can specify the options below

**[Vultr]**<!--(http://www.vultr.com/?ref=6947403-3B)-->

<!--* [Sign up with free $10 credit here](http://www.vultr.com/?ref=6947403-3B)-->
* Follow the DO tutorial above, very similar steps ("Applications" instead of "One-Click Apps")
* Offers different server locations

<!--[AWS](https://aws.amazon.com)-->

**Heroku**

<p>Click this button to...</p>
<p>
<a href="https://heroku.com/deploy" rel="nofollow"><img src="https://camo.githubusercontent.com/c0824806f5221ebb7d25e559568582dd39dd1170/68747470733a2f2f7777772e6865726f6b7563646e2e636f6d2f6465706c6f792f627574746f6e2e706e67" alt="Deploy" data-canonical-src="https://www.herokucdn.com/deploy/button.png" style="max-width:100%;"></a>
</p>

### Usage

```
$ insane-torrent --help

  Usage: insane-torrent [options]

  Options:
  --title, -t        Title of this instance (default Insane - Torrent, env TITLE)
  --port, -p         Listening port (default 3000, env PORT)
  --host, -h         Listening interface (default all)
  --auth, -a         Optional basic auth in form 'user:password' (env AUTH)
  --config-path, -c  Configuration file path (default insane-torrent.json)
  --key-path, -k     TLS Key file path
  --cert-path, -r    TLS Certicate file path
  --log, -l          Enable request logging
  --open, -o         Open now with your default browser
  --help
  --version, -v

  Version:
    0.X.Y

  Read more:
    https://github.com/admknight/insane-torrent

```

### Future features

The next set of [core features can be tracked here](https://github.com/admknight/insane-torrent/issues?q=is%3Aopen+is%3Aissue+label%3Acore-feature). This feature set requires large structural changes and therefore requires a complete rewrite for best results. This rewrite is in progress in the `0.9` branch though it will take quite some time.

In summary, the core features will be:

* **Remote backends**

  It's looking like `0.9` will be more of a general purpose cloud transfer engine. It will be capable of transfering files from and source file-system to any destination file-system. A torrent can be viewed a folder with files, just like your local disk, and Dropbox. As long as it has a concept of files and folders, it could potentially be a insane-torrent file-system backend. Track this issue https://github.com/admknight/insane-torrent/issues/24 for the list of proposed backends.

* **File Transforms**

  During a file tranfer, one could apply different transforms against the byte stream for various effect. For example, supported transforms might include: video transcoding (using ffmpeg), encryption and decryption, [media sorting](https://github.com/admknight/insane-torrent/issues/4) (file renaming), and writing multiple files as a single zip file.

* **Automatic updates** Binary will upgrade itself, adding new features as they get released.

* **RSS** Automatically add torrents, with smart episode filter.

Once completed, insane-torrent will no longer be a simple torrent client and most likely project be renamed.

#### Donate

Will be updated later.
<!--
If you'd like to buy me a coffee or more, you can donate via [PayPal](https://www.paypal.com/cgi-bin/webscr?cmd=_xclick&business=dev%40jpillora%2ecom&lc=AU&item_name=Open%20Source%20Donation&button_subtype=services&currency_code=USD&bn=PP%2dBuyNowBF%3abtn_buynowCC_LG%2egif%3aNonHosted) or BitCoin `1AxEWoz121JSC3rV8e9MkaN9GAc5Jxvs4`.

-->

### Notes

This project is the rewrite of the original [Node version](https://github.com/jpillora/node-torrent-cloud).

### Wiki

Please check wiki page.

<!--
![overview](https://docs.google.com/drawings/d/1ekyeGiehwQRyi6YfFA4_tQaaEpUaS8qihwJ-s3FT_VU/pub?w=606&h=305)
-->

Credits to @anacrolix for https://github.com/anacrolix/torrent

Copyright (c) 2018 Adam Knight #Insane-torrent
