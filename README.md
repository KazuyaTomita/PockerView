# PockerView

CUI program which supports UPI(Universal Porker Interface) protocol. see (this pdf)[https://cdn.shopify.com/s/files/1/0769/9693/files/UPI-documentation_e9050fc5-e6e6-4f37-8611-04819b636333.pdf]
Now we are thinking PockerView supports three type of uses cases.
+ CUI through standart input
+ Remote server which supports a certain protocol
+ multiple engines games

Note: In these three uses cases, remote server support will have some protocols.

### BUILD
You have to set your engine binary and write some configs in config.toml.

``make -f docker.Makefile binary`` for linux.

``make -f docker.Makefile binary-osx`` for macOS.

``make -f docker.Makefile binary-windows`` for windows.


