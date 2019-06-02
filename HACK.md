## Prepping libmodsecurity environment

TODO: compiling to /server/...

TL;DR: Follow compilation recipe: https://github.com/SpiderLabs/ModSecurity/wiki/Compilation-recipes-for-v3.x
./configure --prefix=/server/local

## fix config files being available

% cp -a $ModSecurity/unicode.mapping .
% cp -a $ModSecurity/basic_rules.conf .

## fix lib lookup
You need to have libmodsecurity modules available to lookup from Go

```
% go build -o foo ./example
% ldd foo
        linux-vdso.so.1 (0x00007ffc19fbb000)
        libmodsecurity.so.3 => not found
        libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007ff9436ef000)
        libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007ff9432fe000)
        /lib64/ld-linux-x86-64.so.2 (0x00007ff94390e000)
# HACK to fix the dependency
% ln -s /server/local/lib/libmodsecurity.so.3 /lib/x86_64-linux-gnu/

% ./foo
Rules:
Phase: 0 (0 rules)
Phase: 1 (0 rules)
Phase: 2 (2 rules)
    Rule ID: 200000--0x27bd440
    Rule ID: 200001--0x27bdb70
Phase: 3 (4 rules)
    Rule ID: 200002--0x27be820
    Rule ID: 200003--0x27c2740
    Rule ID: 200004--0x27c2f70
    Rule ID: 200005--0x27c3c20
Phase: 4 (0 rules)
Phase: 5 (0 rules)
Phase: 6 (0 rules)
```
