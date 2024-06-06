# Notie
### Markdown sharing service that uses SSH for minimalism.

## Usage:
- Open editor and write a note: ```ssh -p9990 bald.su```, where bald.su is host and 9990 is default editor port
- Read an existing note ```ssh -p9991 bald.su <note_name>```, where bald.su is host and 9991 is default viewer port

I recommend using aliases:
- ```alias nput="ssh -p9990 bald.su"```
- ```alias nget="ssh -p9991 bald.su"```

## Selfhosting:
Read ```doc/SELFHOSTING.md```
