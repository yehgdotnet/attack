## Welcome

This repo hosts utilities/scripts/files to assist emulation of MITRE ATT&CK / PRE-ATT&CK.


## Web tools

* **obfuscator/obfuscator-using-env-vars.htm** Spoof command line characters with environment variables [https://yehgdotnet.github.io/attack/obfuscator/obfuscator-using-env-vars.htm](https://yehgdotnet.github.io/attack/obfuscator/obfuscator-using-env-vars.htm)

## Command-line tools

* **rtl/rtl.go**  - Generate spoofed file extension with right-to-left overide character.
  * Usage: go run rtl.go -src benign.exe -ext PDF (will copy benign.exe to benignexe.pdf)
* **rtl/rtl_avbypass.go**  - Generate spoofed file extension with right-to-left overide character using AV bypass timer approach.
* **masquerader** - Simulate execution of untrusted binary under the disguise of window binary names.
* **file-read** - Simulate execution of file read access operation.
* **http-post** - Simulate execution of a binary posting data to http servers.

## Files

* **r2l-benignexe-pdf.zip** -  benign.exe spoofed as PDF file using right-to-left overide character

## Contact

[@cyberseckb](https://twitter.com/CyberSecKB)
