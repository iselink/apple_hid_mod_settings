# Apple HID Keyboard module settings

Why:
My keyboard (Keychron K4) using Apple HID module and at some point after system start I need to switch FN keys.  
I want to press shift+F10 and run program in IDE rather than muting audio output.  
Normal peoples will set this in config, making this change permanently.  
And because I'm lazy (and/or not normal), I decided to write dedicated program instead.

### Installation

```bash
git clone https://github.com/iselink/apple_hid_mod_settings &&
cd apple_hid_mod_settings &&
go build -o applehid *.go
```

Then you need to copy binary file into `/usr/local/bin`, set owner, group a permissions.  
This must be done under root user or with sudo access.  

```bash
sudo mv -v applehid /usr/local/bin &&
sudo chown -v root:root /usr/local/bin/applehid &&
sudo chmod -v a=rxs /usr/local/bin/applehid
```

For uninstallation just remove binary `sudo rm -fv /usr/local/bin/applehid`

### Usage

Just type into your terminal `appelhid -fn 0` (in my case).  
For list of all parameters use `applehid -h`.  

You can change following aspects:  

 - `-fn` Swapping F keys and multimedia poop.  
 - `-swap-fn-lctrl` Swapping Fn key with ctrl.  
 - `-swap-opt-cmd` Swap option and command key.  
 - `-iso` Enable/disable hardcoded ISO layout.
 
All of these parameters accept number where `0` mean false and `1` mean true.  
Exception is fn parameter which accepts also `2`.  
(info outsourced from https://wiki.archlinux.org/title/Apple_Keyboard)
