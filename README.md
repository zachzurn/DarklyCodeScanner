# Darkly Malicious Code Scanner

Scan web files for malicious code. Support for scanning PHP files at the moment. 

I created this so I can havle a simple command line option to check for malicious code on new projects or for clients who have been hacked and need their website cleaned.

##Roadmap

###Scanner Types

- [X] Scans filesystem. 
- [ ] Scans website.

###PHP Scanner

- [X] Scans for malicious eval code `eval(base64())`. 

- [X] Scans for malicious obscured eval code `<?php $XZKsyG=’as’;$RqoaUO=’e’;`.

- [ ] Scans for extra long lines in PHP code.

###Web Scanner

- [ ] Scans for malicious code inside html and javascript.

###Htaccess Scanner

- [ ] Scans for malicious code inside htaccess looking for possible cloaking scenarios.

###Wordpress Scanner

- [ ] Scans for malicious PHP code using the PHP Scanner. 

- [ ] Scans for malicious front end code using Web Scanner.

- [ ] Scans for malicious code in the database using the credentials from wp-config.

- [ ] Downloads matching version of Wordpress and compares for changed files in wp-admin.
