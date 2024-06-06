# Notie
### Markdown sharing service that uses SSH for minimalism.

## Usage:
- Open editor and write a note: ```ssh -p3000 bald.su```, where bald.su is host and 3000 is default editor port
- Read an existing note ```ssh -p3001 bald.su <note_name>```, where bald.su is host and 3001 is default viewer port

I recommend using aliases:
- ```alias nput="ssh -p3000 bald.su"```
- ```alias nget="ssh -p3001 bald.su"```

## Selfhosting:
Read ```doc/SELFHOSTING.md```
